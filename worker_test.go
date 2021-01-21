package worker

import (
	"testing"
)

type TestingWorker struct {
	Value string
}

func newTestingWorker(v string) *TestingWorker {
	return &TestingWorker{
		Value: v,
	}
}

var testingWorkerTeam = GetWorkerTeam(newTestingWorker(""))

func TestWorker(t *testing.T) {
	defer func() {
		Reset()
	}()
	w := Hire("testname", newTestingWorker("test")).WithIntroduction("test intro")
	if w.Name != "testname" || w.Introduction() != "test intro" {
		t.Fatal(w)
	}
	worker := FindWorker("testname")
	if worker == nil {
		t.Fatal(worker)
	}
	worker = FindWorker("notexist")
	if worker != nil {
		t.Fatal(worker)
	}
	team := GetWorkerTeam(nil)
	if team != "" {
		t.Fatal()
	}
	team = GetWorkerTeam(newTestingWorker)
	if team == "" {
		t.Fatal()
	}
	team2 := GetWorkerTeam([]byte{})
	if team2 == "" || team == team2 {
		t.Fatal()
	}
}
