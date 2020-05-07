package worker

//Dummy dummy struct
type Dummy struct {
}

//DummyTeam dummy team
var DummyTeam = GetWorkerTeam(NewDummyWorker())

//NewDummyWorker create new dummy worker
func NewDummyWorker() *Dummy {
	return &Dummy{}
}
