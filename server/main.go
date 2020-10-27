package main

import (
	_ "database/sql"

	_ "github.com/lib/pq"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/router"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8081
	r.Run(":8081")
}
