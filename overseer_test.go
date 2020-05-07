package worker

import (
	"errors"
	"testing"
)

var testingOverseer = NewOrverseer("testoverseer", newTestingWorker("")).WithIntroduction("test intro")

var errTesting = errors.New("testing")

var testingInit = func(*OverseerTranning) error {
	return errTesting
}
var testingTrain = func(workers []*Worker) error {
	return errTesting
}
var testingEvaluate = func(worker *Worker) (interface{}, error) {
	return nil, errTesting
}
var testingEvaluationReport = func(worker *Worker) (string, error) {
	return "", errTesting
}
var testingCommand = func(worker *Worker, cmd []byte) (interface{}, error) {
	return nil, errTesting
}

func TestOverseer(t *testing.T) {
	defer func() {
		Reset()
	}()
	Appoint(testingOverseer)
	if testingOverseer.Team() != testingWorkerTeam {
		t.Fatal(testingOverseer.Team())
	}
	if testingOverseer.ID() != "testoverseer" {
		t.Fatal(testingOverseer.ID())
	}
	if testingOverseer.Introduction() != "test intro" {
		t.Fatal(testingOverseer.Introduction())
	}
	if testingOverseer.Muted() != false {
		t.Fatal(testingOverseer.Muted())
	}
	if testingOverseer.Init(nil) != nil {
		t.Fatal()
	}
	if testingOverseer.Train(nil) != nil {
		t.Fatal()
	}
	if v, err := testingOverseer.Evaluate(nil); v != nil || err != nil {
		t.Fatal()
	}
	if v, err := testingOverseer.EvaluationReport(nil); v != "" || err != nil {
		t.Fatal()
	}
	if v, err := testingOverseer.Command(nil, nil); v != nil || err != ErrUnknownCommand {
		t.Fatal()
	}
	o := testingOverseer.
		WithMuted(true).
		WithInitFunc(testingInit).
		WithTrainFunc(testingTrain).
		WithEvaluateFunc(testingEvaluate).
		WithEvaluationReportFunc(testingEvaluationReport).
		WithCommandFunc(testingCommand)
	if o != testingOverseer {
		t.Fatal(o)
	}
	if testingOverseer.Muted() != true {
		t.Fatal(testingOverseer.Muted())
	}
	if testingOverseer.Init(nil) != errTesting {
		t.Fatal()
	}
	if testingOverseer.Train(nil) != errTesting {
		t.Fatal()
	}
	if v, err := testingOverseer.Evaluate(nil); v != nil || err != errTesting {
		t.Fatal()
	}
	if v, err := testingOverseer.EvaluationReport(nil); v != "" || err != errTesting {
		t.Fatal()
	}
	if v, err := testingOverseer.Command(nil, nil); v != nil || err != errTesting {
		t.Fatal()
	}

}
