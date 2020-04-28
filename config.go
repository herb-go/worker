package config

type OverseerConfig struct {
	ID string
}

type WorkerConfig struct {
	ID       string
	Tranning func(v interface{}) error
}

type Config struct {
	Overseers []*OverseerConfig
	Workers   []*WorkerConfig
}
