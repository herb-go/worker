package memberoverseer

import (
	"reflect"

	"github.com/herb-go/member"
	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(member.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("member", Team).
		WithIntroduction("member workers")
}
