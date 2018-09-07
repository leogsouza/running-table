package main

import (
	"github.com/leogsouza/running-table/internal/db"
	"github.com/leogsouza/running-table/internal/webapp"
)

func main() {
	database := db.New()
	webapp.StartServer(&database)
}
