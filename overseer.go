package worker

import (
	"fmt"
)

//Overseer overseet interface
type Overseer interface {
	//Team overseer team
	Team() string
	//Introduction overseer introduction.
	Introduction() string
	//Train init given workers
	Train(workers []*Worker) error
	//Muted is overseer muted.
	Muted() bool
	//Evaluate evaluate given worker.
	//Return data and any error if raised
	Evaluate(*Worker) (interface{}, error)
	//EvaluationReport create evaluation report by given worker .
	//Return plain report and any error if raised
	EvaluationReport(*Worker) (string, error)
	//Command exec command on given workder.
	//Return data and any error if raised.
	Command(*Worker, []byte) (interface{}, error)
	//Outsource create worker by outsourced config.
	Outsource(*Outsourced) error
	//ID overseer id.
	ID() string
	//Init init overseer
	Init(*OverseerTranning) error
}

func defaultOutsource(o *Outsourced) error {
	return fmt.Errorf("worker: outsourcing for [%s] not supported", o.Team)
}
func defaultInit(*OverseerTranning) error {
	return nil
}
func defaultTrain(workers []*Worker) error {
	return nil
}

func defaultEvaluate(*Worker) (interface{}, error) {
	return nil, nil
}

func defaultevaluationReport(*Worker) (string, error) {
	return "", nil
}
func defaultCommand(*Worker, []byte) (interface{}, error) {
	return nil, ErrUnknownCommand
}

//PlainOverseer plain oversser struct
type PlainOverseer struct {
	id               string
	team             string
	introduction     string
	muted            bool
	outsource        func(o *Outsourced) error
	init             func(*OverseerTranning) error
	train            func(workers []*Worker) error
	evaluate         func(*Worker) (interface{}, error)
	evaluationReport func(*Worker) (string, error)
	command          func(*Worker, []byte) (interface{}, error)
}

//ID overseer id.
func (o *PlainOverseer) ID() string {
	return o.id
}

//Team overseer team
func (o *PlainOverseer) Team() string {
	return o.team
}

//Introduction overseer introduction.
func (o *PlainOverseer) Introduction() string {
	return o.introduction
}

//WithIntroduction set overseer introduction.
//Reutrn overseer itself.
func (o *PlainOverseer) WithIntroduction(intro string) *PlainOverseer {
	o.introduction = intro
	return o
}

//Muted is overseer muted
func (o *PlainOverseer) Muted() bool {
	return o.muted
}

//WithMuted set over muted
//Return overseer itself.
func (o *PlainOverseer) WithMuted(muted bool) *PlainOverseer {
	o.muted = muted
	return o
}

//Init init overseer
func (o *PlainOverseer) Init(tranning *OverseerTranning) error {
	return o.init(tranning)
}

//WithInitFunc set overseer init func.
//return overseer self.
func (o *PlainOverseer) WithInitFunc(f func(*OverseerTranning) error) *PlainOverseer {
	o.init = f
	return o
}

//Train init given workers
func (o *PlainOverseer) Train(workers []*Worker) error {
	return o.train(workers)
}

//WithTrainFunc set overseer train func.
//return overseer self.
func (o *PlainOverseer) WithTrainFunc(f func([]*Worker) error) *PlainOverseer {
	o.train = f
	return o
}

//Evaluate evaluate given worker.
//Return data and any error if raised
func (o *PlainOverseer) Evaluate(worker *Worker) (interface{}, error) {
	return o.evaluate(worker)
}

//WithEvaluateFunc set overseer evalutate function.
//Reutrn overseer self.
func (o *PlainOverseer) WithEvaluateFunc(f func(worker *Worker) (interface{}, error)) *PlainOverseer {
	o.evaluate = f
	return o
}

//EvaluationReport create evaluation report by given worker.
//Return plain report and any error if raised
func (o *PlainOverseer) EvaluationReport(worker *Worker) (string, error) {
	return o.evaluationReport(worker)
}

//WithEvaluationReportFunc set overseer evalution report function.
//Return overseer self.
func (o *PlainOverseer) WithEvaluationReportFunc(f func(worker *Worker) (string, error)) *PlainOverseer {
	o.evaluationReport = f
	return o
}

//Command exec command on given workder.
//Return data and any error if raised.
func (o *PlainOverseer) Command(worker *Worker, cmd []byte) (interface{}, error) {
	return o.command(worker, cmd)
}

//WithCommandFunc set overseer command function.
//Return overseer self
func (o *PlainOverseer) WithCommandFunc(f func(worker *Worker, cmd []byte) (interface{}, error)) *PlainOverseer {
	o.command = f
	return o
}

//Outsource create worker by outsourced config.
func (o *PlainOverseer) Outsource(outsourced *Outsourced) error {
	return o.outsource(outsourced)
}

//WithOutsourceFunc set overseer outsource function.
//Return overseer self
func (o *PlainOverseer) WithOutsourceFunc(f func(outsourced *Outsourced) error) *PlainOverseer {
	o.outsource = f
	return o
}

//NewOrverseer create new overseer with given id and team
func NewOrverseer(id string, v interface{}) *PlainOverseer {
	return &PlainOverseer{
		id:               id,
		team:             GetWorkerTeam(v),
		init:             defaultInit,
		train:            defaultTrain,
		evaluate:         defaultEvaluate,
		evaluationReport: defaultevaluationReport,
		command:          defaultCommand,
		outsource:        defaultOutsource,
	}
}
