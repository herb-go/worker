package cacheoverseer

import (
	"reflect"

	"github.com/herb-go/herb/cache"

	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(cache.New()).Type().String()

//Overseer overseer struct
type Overseer struct {
}

//Team overseer team
func (o *Overseer) Team() string {
	return Team
}

//Train init given workers
func (o *Overseer) Train(workers []*worker.Worker) error {
	return nil
}

//Evaluate evaluate given worker.
//Return data and any error if raised
func (o *Overseer) Evaluate(w *worker.Worker) (interface{}, error) {
	return nil, nil
}

//EvaluationReport create evaluation report by given worker .
//Return plain report and any error if raised
func (o *Overseer) EvaluationReport(w *worker.Worker) (string, error) {
	return "", nil
}

//ID overseer id.
func (o *Overseer) ID() string {
	return "cache"
}

//Introduction overseer introduction.
func (o *Overseer) Introduction() string {
	return "cache workers"
}

//New create new overseer.
func New() *Overseer {
	return &Overseer{}
}
