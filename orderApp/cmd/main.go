package main

import (
	"orderapp/api"
	"orderapp/api/db"
	"orderapp/config"
)

func main() {
	//Initalize configs
	config.InitConfig()

	//Initalize DB
	db.InitializeConn()

	//Start the API
	api.StartAPI()
}
