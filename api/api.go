package main

import (
	"fmt"
	"golang-mod/api/config"
	"log"
	"os"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
» Reading variables using the model
fmt.Println("Reading variables using the model..")
fmt.Println("Database is\t", configuration.Database.DBName)
fmt.Println("Port is\t\t", configuration.Server.Port)
fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

» Reading variables without using the model
 fmt.Println("\nReading variables without using the model..")
 fmt.Println("Database is\t", viper.GetString("database.dbname"))
 fmt.Println("Port is\t\t", viper.GetInt("server.port"))
 fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
 fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
*/

func dbConn(username string, userpass string, dbname string, dbhost string, dbport string) {
	// dsn := "dev1:dev_test@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + userpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%v", err)
	}
	sqlDB, err := db.DB()

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("connection to database error !!!")
	} else {
		fmt.Println("connected to database")
	}
}

func main() {

	// test env
	viper.BindEnv("id")
	fmt.Printf("before %s\n", viper.Get("id"))
	os.Setenv("ID", "10")
	fmt.Printf("after %s\n", viper.Get("id"))

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "local")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	config := map[string]string{
		"username": viper.GetString("database.dbuser"),
		"userpass": viper.GetString("database.dbpassword"),
		"dbname":   viper.GetString("database.dbname"),
		"dbport":   viper.GetString("database.dbport"),
		"dbhost":   viper.GetString("database.dbhost"),
	}

	// validation for db connection
	for key, val := range config {
		if len(val) == 0 {
			log.Fatalf("%s required value\n", key)
		}
	}

	dbConn(config["username"],
		config["userpass"],
		config["dbname"],
		config["dbhost"],
		config["dbport"])

}
