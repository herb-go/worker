package worker

import "reflect"

//Debug debug mode
var Debug bool

//Worker worker strut
type Worker struct {
	Name         string
	introduction string
	Interface    interface{}
}

//GetWorkerTeam get team of given worker
func GetWorkerTeam(v interface{}) string {
	if v == nil {
		return ""
	}
	return reflect.ValueOf(v).Type().String()
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

var allworkers = map[string]*Worker{}
var workersByTeam = map[string][]*Worker{}
var overseers = map[string]Overseer{}

//Hire register interface as worker with given name.
//Return workder registered.
func Hire(name string, v interface{}) *Worker {
	c := New()
	c.Name = name
	c.Interface = v
	allworkers[name] = c
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
		err := overseers[k].Train(workersByTeam[overseers[k].Team()])
		if err != nil {
			return err
		}
	}
	return nil
}

func groupWorkersByTeam() {
	for _, v := range allworkers {
		t := GetWorkerTeam(v)
		workersByTeam[t] = append(workersByTeam[t], v)
	}
}

//InitOverseers init overseers
func InitOverseers() error {
	groupWorkersByTeam()
	for k := range overseers {
		err := overseers[k].Init(overseerTrannings[overseers[k].ID()])
		if err != nil {
			return err
		}
	}
	return nil
}

//FindWorkerInTeam find worker by given type and name.
func FindWorkerInTeam(team string, name string) *Worker {
	t := workersByTeam[team]
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
func FindWorker(name string) *Worker {
	return allworkers[name]
}

//Reset reset workers and overseers
func Reset() {
	workersByTeam = map[string][]*Worker{}
	allworkers = map[string]*Worker{}
	overseers = map[string]Overseer{}
	ResetTranning()
}

var trannings = map[string]*Tranning{}
var overseerTrannings = map[string]*OverseerTranning{}

//GetTranning get worker tranning by given worker id.
func GetTranning(workerid string) *Tranning {
	if workerid == "" || trannings[workerid] == nil {
		return nil
	}
	return trannings[workerid]
}

//ResetTranning reset worker trannings and overseer trannings
func ResetTranning() {
	trannings = map[string]*Tranning{}
	overseerTrannings = map[string]*OverseerTranning{}
}
