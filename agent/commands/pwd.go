package commands

import (
	"os"
)

type Pwd struct {
	Command string
}

func (instance Pwd) Exec() (response string, err error) {
	response, err = os.Getwd()
	return response, err
}
