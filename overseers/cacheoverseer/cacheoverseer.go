package cacheoverseer

import (
	"reflect"

	"github.com/herb-go/herb/cache"

	"github.com/herb-go/worker"
)

//Team overseer team
var Team = reflect.ValueOf(cache.New()).Type().String()

//New create new overseer.
func New() *worker.PlainOverseer {
	return worker.NewOrverseer("cache", Team).
		WithIntroduction("cache workers")
}
