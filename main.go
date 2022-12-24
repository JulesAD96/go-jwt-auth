package main

import (
	"github.com/JulesAD96/go-jwt-auth/database"
)

func main() {
	database.Connect("root:root@tcp(localhost:3306)/jwtauth?parseTime=true")
	database.Migrate()
}
