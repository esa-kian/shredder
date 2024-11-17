# Shredder

Shredder is a Golang library that automatically generates RESTful CRUD APIs for any defined entity in your application. With Shredder, developers can create APIs by simply defining models and running a single command, making API development faster and more efficient.

## Features

- **Automatic CRUD API Generation**: Generates Create, Read, Update, and Delete endpoints for each entity.
- **Model-Based Table Creation**: Reads model files from the `migration` directory, checks for existing tables, and creates them if they do not already exist.
- **Easy Configuration**: Set up connection details and other configurations easily using a `.env` file.
- **Modular Architecture**: Easily extend and customize generated controllers and routes.

## Installation

First, get the Shredder package:

```bash
go get github.com/esa-kian/shredder
```

## Usage
1. **Define Models:** Add your entity model files as JSON files in the `migration` directory. Each JSON file should contain fields like the following example:
```json
{
  "EntityName": "User",
  "Fields": [
    { "Name": "id", "DataType": "INT", "IsPrimaryKey": true },
    { "Name": "name", "DataType": "VARCHAR(255)", "IsRequired": true },
    { "Name": "email", "DataType": "VARCHAR(255)", "IsRequired": true },
    { "Name": "age", "DataType": "INT", "IsRequired": true }
  ]
}
```

2. **Set up Database Configuration:** Create a `.env` file in the root of your project with your database connection details:
```dotenv
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=your_db
NAMESPACE=""
```
3. **Initialize Shredder:** Run your application to load models, generate tables, and create CRUD APIs for each entity:

```go
package main

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/esa-kian/shredder"
)

func main() {
    // Initialize the database connection
    db, err := sql.Open("mysql", "<dsn>")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations
    if err := shredder.RunMigrations(db); err != nil {
        log.Fatal("Migration error:", err)
    }

    // Start Shredder API server
    shredder.StartServer(db)
}
```
4. **Access APIs:** Once the server is running, you can access endpoints at:
```bash
POST    /api/{entity}     // Create a new entity
GET     /api/{entity}     // List entities
GET     /api/{entity}/{id} // Get an entity by ID
PUT     /api/{entity}/{id} // Update an entity
DELETE  /api/{entity}/{id} // Delete an entity
```

## Contributing
Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m "Add a feature"`.
4. Push to the branch: `git push origin feature-name`.
5. Open a pull request.

## License
This project is licensed under the MIT License.


