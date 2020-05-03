package worker

type OverseerConfig struct {
	ID       string
	Tranning func(v interface{}) error
}

type WorkerConfig struct {
	ID       string
	Tranning func(v interface{}) error
}

type Config struct {
	Overseers []*OverseerConfig
	Workers   []*WorkerConfig
}

func (c *Config) Apply() error {
	for _, v := range c.Overseers {
		if v.ID != "" {
			overseerLoaders[v.ID] = v.Tranning
		}
	}
	for _, v := range c.Workers {
		if v.ID != "" {
			workerLoaders[v.ID] = v.Tranning
		}
	}
	return nil
}
