# ESL_UniqueUuidDB

This project provides a function to generate a unique UUID using GORM to verify its uniqueness in a database. 

## Installation

To use this module, you must first import the package into your Go project. Make sure you have Go and GORM installed.

```bash
go get github.com/EcoSystemsLink/ESL_UniqueUuidDB
```

## Usage

Below is an example of how to use the `GenerateUniqueUUID` function in your project:

```go
package main

import (
    "fmt"
    "github.com/EcoSystemsLink/ESL_UniqueUuidDB"
    "gorm.io/gorm"
)

func main() {
    var db *gorm.DB // Initialize your database connection here

    uuid, err := ESL_UniqueUuidDB.GenerateUniqueUUID(db, &YourModel{})
    if err != nil {
        fmt.Println("Error generating UUID:", err)
        return
    }

    fmt.Println("Unique UUID generated:", uuid)
}
```

## Functions

- `GenerateUniqueUUID(db *gorm.DB, model interface{}) (string, error)`: Generates a unique UUID by verifying its uniqueness in the database.
- `isUnique(db *gorm.DB, model interface{}, id string) bool`: Checks if the UUID already exists in the specified table.

## Contributions

Contributions are welcome. If you want to improve this project, please open an issue or a pull request.

## License

This project is under the MIT License.