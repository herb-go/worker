package worker

import (
	"reflect"
)

type Dummy struct {
}

var DummyTeam = reflect.ValueOf(NewDummyWorker()).Type().String()

func NewDummyWorker() *Dummy {
	return &Dummy{}
}
