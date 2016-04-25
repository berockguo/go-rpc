package util

import (
	log "github.com/Sirupsen/logrus"
)

type Error struct {
	ErrMsg string
}

func (e *Error) Error() string {
	log.Fatalf(e.ErrMsg)
	return e.ErrMsg
}
