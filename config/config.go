package config

import (
	"os"
	"project/config/postgres"
	"strconv"

	"gorm.io/gorm"
)

type (
	config struct {
	}

	Config interface {
		ServiceName() string
		ServiceHost() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServiceHost() string {
	return os.Getenv("HOST")
}

func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
