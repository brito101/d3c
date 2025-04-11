package commands

import (
	"os/user"
)

type Whoami struct {
	Command string
}

func (instance Whoami) Exec() (response string, err error) {
	user, err := user.Current()
	response = user.Username
	return response, err
}
