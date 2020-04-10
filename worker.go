package worker

import "reflect"

type Worker struct {
	Name         string
	introduction string
	location     string
	Interface    interface{}
}

func (u *Worker) WithIntroduction(c string) *Worker {
	u.introduction = c
	return u
}
func (u *Worker) Introduction() string {
	return u.introduction
}

func (u *Worker) WithLocation(c string) *Worker {
	u.location = c
	return u
}
func (u *Worker) Location() string {
	return u.location
}
func New() *Worker {
	return &Worker{}
}

//Overseer overseet interface
type Overseer interface {
	//Team overseer team
	Team() string
	//Introduction overseer introduction.
	Introduction() string
	//Train init given workers
	Train(workers []*Worker) error
	//Evaluate evaluate given worker.
	//Return data and any error if raised
	Evaluate(*Worker) (interface{}, error)
	//EvaluationReport create evaluation report by given worker .
	//Return plain report and any error if raised
	EvaluationReport(*Worker) (string, error)
	//ID overseer id.
	ID() string
}

var workers = map[string][]*Worker{}
var overseers = map[string]Overseer{}

func Hire(name string, v interface{}) *Worker {
	ct := reflect.ValueOf(v).Elem().Type().String()
	if workers[ct] == nil {
		workers[ct] = []*Worker{}
	}
	c := New()
	c.Name = name
	c.Interface = v
	workers[ct] = append(workers[ct], c)
	return c
}
func Appoint(t Overseer) {
	overseers[t.Team()] = t
}

func TrainWorkers() error {
	for k := range overseers {
		err := overseers[k].Train(workers[overseers[k].ID()])
		if err != nil {
			return err
		}
	}
	return nil
}

func reset() {
	workers = map[string][]*Worker{}
	overseers = map[string]Overseer{}
}
