package actionoverseer

import (
	"reflect"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(action.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("action", Team).
		WithIntroduction("http action workers")
}
