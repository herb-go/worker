package sessionoverseer

import (
	"reflect"

	"github.com/herb-go/herb/cache/session"

	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(session.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("session", Team).
		WithIntroduction("session workers")
}
