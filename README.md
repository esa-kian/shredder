# Shredder

Shredder is a Golang library that automatically generates RESTful CRUD APIs for any defined entity in your application. With Shredder, developers can create APIs by simply defining models and running a single command, making API development faster and more efficient.

## Features

- **Command-Line CRUD Generation**: Generate Create, Read, Update, and Delete endpoints for each entity by running a single command.
- **Model-Based Table Creation**: Reads model files from the `migration` directory, checks for existing tables, and creates them if they do not already exist.
- **Easy Configuration**: Configure connection details and other settings easily using a `.env` file.
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
  "EntityName": "Item",
  "Fields": [
    { "Name": "id", "DataType": "INT", "IsPrimaryKey": true },
    { "Name": "name", "DataType": "VARCHAR(255)", "IsRequired": true },
    { "Name": "description", "DataType": "TEXT" },
    { "Name": "price", "DataType": "DECIMAL(10, 2)", "IsRequired": true }
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
3. **Set Up a CLI to Run Shredder Commands:** In your project, create a `main.go` file in the `cmd` directory or another preferred location. Here’s how to structure it:

```go
package main

import (
    "log"
    "os"
    "github.com/esa-kian/shredder"
)

func main() {
    // Get the entity name from command-line arguments
    if len(os.Args) < 2 {
        log.Fatal("Please provide an entity name using -entity <entity_name>")
    }
    entityName := os.Args[1]

    // Run Shredder to generate CRUD APIs for the specified entity
    err := shredder.GenerateCRUD(entityName)
    if err != nil {
        log.Fatalf("Error generating CRUD for %s: %v", entityName, err)
    }

    log.Printf("CRUD for entity %s generated successfully!", entityName)
}
```
4. **Generate CRUD APIs for an Entity:** Run the following command in your terminal to generate the CRUD operations:
```bash
go run cmd/main.go -entity items
```

5. **Calling the APIs:** After running the command, you can access the generated API endpoints. For example, if your server is running on `localhost:8080`, you can call the APIs as follows:
- **Create a new item** (POST request):
```bash
curl -X POST http://localhost:8080/api/items \
-H "Content-Type: application/json" \
-d '{"name": "Sample Item", "description": "A description of the item.", "price": 10.99}'
```

- **List all items** (GET request):
```bash
curl -X GET http://localhost:8080/api/items
```

- **Get an item by ID** (GET request):
```bash
curl -X GET http://localhost:8080/api/items/1
```

- **Update an item** (PUT request):
```bash
curl -X PUT http://localhost:8080/api/items/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Item", "description": "Updated description.", "price": 12.99}'
```

- **Delete an item** (DELETE request):
```bash
curl -X DELETE http://localhost:8080/api/items/1
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


