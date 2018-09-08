package notifier

import (
	"github.com/caarlos0/env"
	"github.com/leogsouza/running-table/internal/config"
	"github.com/leogsouza/running-table/internal/db"
	"github.com/pusher/pusher-http-go"
)

// Notifier contains a channel to notify the users
type Notifier struct {
	notifyChannel chan<- bool
}

func notifier(database *db.Database, notifyChannel <-chan bool) {

	cfg := config.Variables{}
	err := env.Parse(&cfg)

	if err != nil {
		panic(err)
	}

	client := pusher.Client{
		AppId:   cfg.AppId,
		Key:     cfg.Key,
		Secret:  cfg.Secret,
		Cluster: cfg.Cluster,
		Secure:  true,
	}

	for {
		<-notifyChannel
		data := map[string][]db.Record{"results": database.GetRecords()}
		client.Trigger("results", "results", data)
	}
}

// New creates a new notifier
func New(database *db.Database) Notifier {
	notifyChannel := make(chan bool)
	go notifier(database, notifyChannel)
	return Notifier{
		notifyChannel,
	}
}

// Notify notifies through the channel
func (notifier *Notifier) Notify() {
	notifier.notifyChannel <- true
}
