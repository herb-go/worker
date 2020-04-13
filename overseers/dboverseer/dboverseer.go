package databaseoverseer

import (
	"reflect"

	"github.com/herb-go/herb/model/sql/db"

	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(db.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("db", Team).
		WithIntroduction("db workers")
}
