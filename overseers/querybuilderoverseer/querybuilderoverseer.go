package databaseoverseer

import (
	"reflect"

	"github.com/herb-go/herb/model/sql/querybuilder"

	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(querybuilder.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("querybuilder", Team).
		WithIntroduction("querybuilder workers")
}
