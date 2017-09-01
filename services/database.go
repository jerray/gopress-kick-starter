package services

import (
	"github.com/fpay/gopress"
	"github.com/jinzhu/gorm"
)

const (
	DatabaseServiceName = "database"
)

type DatabaseService struct {
	*gorm.DB
}

func NewDatabaseService() *DatabaseService {
	return new(DatabaseService)
}

func (d *DatabaseService) ServiceName() string {
	return DatabaseServiceName
}

func (d *DatabaseService) RegisterContainer(c *gopress.Container) {
}
