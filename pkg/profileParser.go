package JAA

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func ParseProfile(data []byte) Profile {
	t := Profile{}
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalln(err.Error())
	}
	AskandParseInputs(t)
	return t
}
func Ask(key string, value string) string {
	var val string
	fmt.Println("[User-Input][" + key + "] Description/Usage:" + value)
	fmt.Scanln(&val)
	return val
}

func AskandParseInputs(t Profile) []Job {
	// var Xcommands []Command
	var Jobs []Job
	userInputs := make(map[string]string)
	for key, value := range t.UserInputs {
		val := Ask(key, value)
		userInputs[key] = val
	}
	var cmds []string
	for _, commands := range t.Commands {
		for _, cmd := range commands.Commands {
			temp := cmd
			if strings.Contains(cmd, "%{") {
				for key, val := range userInputs {
					temp = strings.ReplaceAll(temp, "%{"+key+"}", val)
				}
			}
			cmds = append(cmds, temp)
		}
		// var cmds []string
		// fmt.Println(cmds)
		Jobs = append(Jobs, Job{Name: commands.Name, Async: commands.Concurrent, Commands: cmds, Status: "Started"})
		cmds = []string{}
	}
	if t.Concurrent {
		// fmt.Println(Jobs)
		AsyncProc(Jobs)
	} else {
		// fmt.Println()
		Proc(Jobs)
	}
	return Jobs
}
