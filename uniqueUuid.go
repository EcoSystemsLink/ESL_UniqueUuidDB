package esluniqueuuiddb

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateUniqueUUID(db *gorm.DB, model any) (uuid.UUID, error) {
	var uniqueID uuid.UUID
	for {
		uniqueID = uuid.New()
		if isUnique(db, model, uniqueID.String()) {
			break
		}
	}
	return uniqueID, nil
}

func isUnique(db *gorm.DB, model any, id string) bool {
	var count int64
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	tableName := stmt.Schema.Table
	db.Table(tableName).Where("id = ?", id).Count(&count)
	return count == 0
}
