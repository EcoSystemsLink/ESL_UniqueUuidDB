package esluniqueuuiddb

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUniqueUUID(db *gorm.DB, model interface{}) (string, error) {
	var id string
	for {
		id = uuid.New().String()
		if isUnique(db, model, id) {
			break
		}
	}
	return id, nil
}

func isUnique(db *gorm.DB, model interface{}, id string) bool {
	var count int64
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	tableName := stmt.Schema.Table
	db.Table(tableName).Where("id = ?", id).Count(&count)
	return count == 0
}
