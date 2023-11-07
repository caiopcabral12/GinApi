package main

import (
	db "GinGoApi/database"
	rt "GinGoApi/routes"
)

func main() {
	db.DbConnect()

	rt.HandleRoutes()
}
