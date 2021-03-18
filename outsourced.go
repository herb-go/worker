package worker

var OutsourcedPrefix = "outsourced-"

type Outsourced struct {
	Name         string
	introduction string
	Team         string
	TranningPlan func(v interface{}) error `config:", lazyload"`
}
