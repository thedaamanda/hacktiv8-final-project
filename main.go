package main

import (
	"project/config"
	"project/src"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// @title Hacktiv8 Final Project API Documentation
// @description This is the documentation for the final project of Hacktiv8
// @version 1.0
//
// @contact.name API Support
// @contact.email thedaamandaa@gmail.com
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}

	log.Infof("read .env from file")

	config := config.NewConfig()
	server := src.InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
