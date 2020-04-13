package worker

import "reflect"

//Worker worker strut
type Worker struct {
	Name         string
	introduction string
	Interface    interface{}
}

//WithIntroduction set workder introduction.
//Return workder self.
func (u *Worker) WithIntroduction(c string) *Worker {
	u.introduction = c
	return u
}

//Introduction return workder introduction
func (u *Worker) Introduction() string {
	return u.introduction
}

//New create new worker
func New() *Worker {
	return &Worker{}
}

var workers = map[string][]*Worker{}
var overseers = map[string]Overseer{}

//Hire register interface as worker with given name.
//Return workder registered.
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

//Appoint register overseer
func Appoint(t Overseer) Overseer {
	overseers[t.Team()] = t
	return t
}

//TrainWorkers init workders by registered overseers.
func TrainWorkers() error {
	for k := range overseers {
		err := overseers[k].Train(workers[overseers[k].ID()])
		if err != nil {
			return err
		}
	}
	return nil
}

//InitOverseers init overseers
func InitOverseers() error {
	for k := range overseers {
		err := overseers[k].Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//FindWorker find worker by given type and name.
func FindWorker(team string, name string) *Worker {
	t := workers[team]
	if t == nil {
		return nil
	}
	for k := range t {
		if t[k].Name == name {
			return t[k]
		}
	}
	return nil
}

func reset() {
	workers = map[string][]*Worker{}
	overseers = map[string]Overseer{}
}
