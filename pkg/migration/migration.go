package migration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/esa-kian/shredder/pkg/models"
)

func LoadModelsFromMigrationDir() ([]models.Model, error) {
	var modelsList []models.Model

	files, err := os.ReadDir("./pkg/migration/")
	if err != nil {
		return nil, fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := os.ReadFile(filepath.Join("./pkg/migration", file.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to read model file %s: %w", file.Name(), err)
			}

			var model models.Model
			if err := json.Unmarshal(data, &model); err != nil {
				return nil, fmt.Errorf("failed to parse model file %s: %w", file.Name(), err)
			}

			modelsList = append(modelsList, model)
		}
	}

	return modelsList, nil
}
