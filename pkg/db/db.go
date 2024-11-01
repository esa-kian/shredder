package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/esa-kian/shredder/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DBConfig holds configuration for the database connection
type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewConnection establishes a database connection using the provided config
func NewConnection(cfg DBConfig) (*sql.DB, error) {
	var dsn string
	if cfg.Driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	} else if cfg.Driver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	} else {
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GenerateCreateTableSQL(model models.Model) string {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", model.EntityName)
	fields := []string{}

	for _, field := range model.Fields {
		fieldSQL := fmt.Sprintf("%s %s", field.Name, field.DataType)
		if field.IsPrimaryKey {
			fieldSQL += " PRIMARY KEY"
		}
		if field.IsRequired {
			fieldSQL += " NOT NULL"
		}
		fields = append(fields, fieldSQL)
	}

	query += strings.Join(fields, ", ") + ");"
	return query
}

func CreateTableFromModel(db *sql.DB, model models.Model) error {
	query := GenerateCreateTableSQL(model)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table for %s: %w", model.EntityName, err)
	}
	return nil
}

func TableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = '%s'", tableName)
	if err := db.QueryRow(query).Scan(&exists); err != nil {
		return false, fmt.Errorf("error checking table existence for %s: %w", tableName, err)
	}
	return exists, nil
}
