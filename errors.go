package worker

import (
	"errors"
	"fmt"
)

//ErrUnknownCommand error raised if given command is unknown.
var ErrUnknownCommand = errors.New("worker:unknow command")

var ErrWorkerNotFound = errors.New("worker not found")

func NewWorkerNotFounderError(id string) error {
	return fmt.Errorf("%w [%s]", ErrWorkerNotFound, id)
}
