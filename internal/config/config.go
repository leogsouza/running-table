package config

type Variables struct {
	AppId   string `env:"PUSHER_APP_ID"`
	Key     string `env:"PUSHER_KEY"`
	Secret  string `env:"PUSHER_SECRET"`
	Cluster string `env:"PUSHER_CLUSTER"`
}
