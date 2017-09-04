package services

import (
	"github.com/fpay/gopress"
	"github.com/jinzhu/gorm"
)

const (
	// DatabaseServiceName is the identity of database service
	DatabaseServiceName = "database"
)

// DatabaseService type
type DatabaseService struct {
	// Uncomment this line if this service has dependence on other services in the container
	// c *gopress.Container

	*gorm.DB
}

// NewDatabaseService returns instance of database service
func NewDatabaseService() *DatabaseService {
	return new(DatabaseService)
}

// ServiceName is used to implements gopress.Service
func (s *DatabaseService) ServiceName() string {
	return DatabaseServiceName
}

// RegisterContainer is used to implements gopress.Service
func (s *DatabaseService) RegisterContainer(c *gopress.Container) {
	// Uncomment this line if this service has dependence on other services in the container
	// s.c = c
}

func (s *DatabaseService) SampleMethod() {
}
