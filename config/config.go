package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBHost     string
	DBName     string
	DBPort     int
	DBUser     string
	DBPassword string
}

var Configuration Configurations

func InitConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../etc/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.db_name", "test_db")

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", Configuration.Database.DBName)
	fmt.Println("Database port is\t", Configuration.Database.DBPort)
	fmt.Println("Server port is\t\t", Configuration.Server.Port)

	// 	// Reading variables without using the model
	// 	fmt.Println("\nReading variables without using the model..")
	// 	fmt.Println("Database is\t", viper.GetString("database.db_name"))
	// 	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	// 	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	// 	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
}
