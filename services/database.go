package services

import (
	"github.com/jinzhu/gorm"
)

type DatabaseOptions struct {
	Driver string `mapstructure:"driver" yaml:"driver"`
	Dsn    string `mapstructure:"dsn" yaml:"dsn"`
}

// DatabaseService type
type DatabaseService struct {
	*gorm.DB
}

// NewDatabaseService returns instance of database service
func NewDatabaseService(opts DatabaseOptions) (*DatabaseService, error) {
	conn, err := gorm.Open(opts.Driver, opts.Dsn)
	if err != nil {
		return nil, err
	}
	return &DatabaseService{conn}, nil
}
