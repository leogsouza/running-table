package main

import (
	"github.com/leogsouza/running-table/internal/db"
	"github.com/leogsouza/running-table/internal/notifier"
	"github.com/leogsouza/running-table/internal/webapp"
)

func main() {
	database := db.New()
	notifierClient := notifier.New(&database)
	webapp.StartServer(&database, &notifierClient)
}
