package worker

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

//Debug debug mode
var Debug bool

//Worker worker strut
type Worker struct {
	Name         string
	introduction string
	Team         string
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
var outsourced = []*Outsourced{}

var onStop = []func(){}
var onStart = []func(){}
var locker = sync.Mutex{}

func OnStop(f func()) {
	locker.Lock()
	defer locker.Unlock()
	onStop = append(onStop, f)
}

func Stop() {
	locker.Lock()
	defer locker.Unlock()
	defer func() { onStop = []func(){} }()
	for _, v := range onStop {
		v()
	}
}
func OnStart(f func()) {
	locker.Lock()
	defer locker.Unlock()
	onStart = append(onStart, f)
}

func Start() {
	locker.Lock()
	defer locker.Unlock()
	defer func() { onStart = []func(){} }()
	for _, v := range onStart {
		v()
	}
}

//Hire register interface as worker with given name.
//Return workder registered.
func Hire(name string, v interface{}) *Worker {
	c := New()
	c.Name = name
	c.Interface = v
	c.Team = GetWorkerTeam(v)
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
		workersByTeam[v.Team] = append(workersByTeam[v.Team], v)
		if Debug {
			fmt.Printf("Worker %s <%s> hired.\n", v.Name, removeStarFromTeam(v.Team))
			if v.introduction != "" {
				fmt.Printf("  Introduction :%s\n", v.introduction)
			}
		}
	}
}
func removeStarFromTeam(team string) string {
	if team[0] == '*' {
		return team[1:]
	}
	return team
}
func addStarToTeam(team string) string {
	return "*" + team
}

//InitOverseers init overseers
func InitOverseers() error {
	if Debug {
		time.Sleep(time.Millisecond)
		fmt.Println("Hiring workers and overseers.")
		fmt.Println("-----------------------------")
		fmt.Println()
	}
	for k := range overseers {
		err := overseers[k].Init(overseerTrannings[overseers[k].ID()])
		if err != nil {
			return err
		}
		if Debug {
			fmt.Printf("Overseer %s appointed to worker team <%s>.\n", overseers[k].ID(), removeStarFromTeam(overseers[k].Team()))
			intro := overseers[k].Introduction()
			if intro != "" {
				fmt.Printf("  Introduction :%s\n", intro)
			}
		}
	}
	for _, v := range outsourced {
		if !strings.HasPrefix(v.Name, OutsourcedPrefix) {
			return fmt.Errorf("worker: outsourced worker name '%s' not start with prefix '%s'", v.Name, OutsourcedPrefix)
		}
		o, ok := overseers[addStarToTeam(v.Team)]
		if !ok {
			return fmt.Errorf("worker: overseer [%s] not found.", v.Team)
		}
		err := o.Outsource(v)
		if err != nil {
			return err
		}
		if Debug {
			fmt.Printf("Hiring outsourced worker %s <%s>.\n", v.Name, v.Team)
		}
	}
	groupWorkersByTeam()

	if Debug {
		fmt.Println()
		fmt.Println("-----------------------------")
		fmt.Println("All workers and overseers hired.")
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
	outsourced = []*Outsourced{}
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
