package main

import (
	router "./routers"
	db "./database"
)

func main() {
		router := router.InitRouter()
		
		router.Static("/static","./static")
		
		defer db.Db.Close()

		router.Run(":8080")	
}