package worker

//OverseerTranning overseer config struct
type OverseerTranning struct {
	ID           string
	TranningPlan func(v interface{}) error `config:", lazyload"`
}

//Tranning worker config struct
type Tranning struct {
	ID           string
	TranningPlan func(v interface{}) error `config:", lazyload"`
}

//Config config struct
type Config struct {
	Overseers  []*OverseerTranning
	Workers    []*Tranning
	Outsourced []*Outsourced
}

//Apply apply config
//Return any error if raised.
func (c *Config) Apply() error {
	for _, v := range c.Overseers {
		if v.ID != "" {
			overseerTrannings[v.ID] = v
		}
	}
	outsourced = append(outsourced, c.Outsourced...)
	for _, v := range c.Workers {
		if v.ID != "" {
			trannings[v.ID] = v
		}
	}
	return nil
}
