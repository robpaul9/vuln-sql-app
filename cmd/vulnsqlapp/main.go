package main

import (
	"github.com/joho/godotenv"
	"github.com/robpaul9/vulnsqlapp/adapters/aws"
	"github.com/robpaul9/vulnsqlapp/adapters/db"
	"github.com/robpaul9/vulnsqlapp/adapters/records"
	"github.com/robpaul9/vulnsqlapp/adapters/server"
	"github.com/robpaul9/vulnsqlapp/config"
)

func main() {
	godotenv.Load()

	config := config.NewConfig()

	awsService := aws.New(&aws.Config{
		Logger: config.Logger,
	})

	dbPassword, err := awsService.SSMVC.Param(config.DBPasswordParam, true).GetValue()
	if err != nil {
		config.Logger.Panic(err)
	}

	database, err := db.New(
		&db.Config{
			DatabaseName: config.DBName,
			Host:         config.DBHost,
			Password:     dbPassword,
			User:         config.DBUser,
			Port:         config.DBPort,
		})
	if err != nil {
		config.Logger.Panic(err)
	}
	defer database.Close()

	recordsService := records.New(&records.Config{
		Logger:   config.Logger,
		Database: database,
	})

	s := server.New(&server.Config{
		ServiceName:    config.ServiceName,
		ServicePort:    config.ServicePort,
		Logger:         config.Logger,
		RecordsService: *recordsService,
	})

	s.Start()

}
