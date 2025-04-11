package commands

import (
	"fmt"

	"github.com/mitchellh/go-ps"
)

type Ps struct {
	Command string
}

func (instance Ps) Exec() (response string, err error) {
	process, _ := ps.Processes()
	for _, v := range process {
		response += fmt.Sprintf("%d -> %d -> %s \n", v.PPid(), v.Pid(), v.Executable())
	}
	return response, err
}
