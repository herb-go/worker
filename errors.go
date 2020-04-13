package worker

import "errors"

//ErrUnknownCommand error raised if given command is unknown.
var ErrUnknownCommand = errors.New("worker:unknow command")
