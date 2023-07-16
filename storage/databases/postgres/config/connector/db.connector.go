package connector

import (
	"gorm.io/gorm"
)

type DBConnector interface {
	Connect() (*gorm.DB, error)
}
