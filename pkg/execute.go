package JAA

import (
	"fmt"
	"os/exec"
	"strings"
)

func AsyncProc(jobs []Job) {
	// glassy := false
	// fmt.Println("Async Proc")
	fmt.Println(jobs)
	for _, job := range jobs {
		go func(job Job) {
			if job.Async {
				for _, command := range job.Commands {
					Wg.Add(1)
					go func(command string, name string) {
						ExecuteCommand(command, name)
						Wg.Done()
					}(command, job.Name)
				}
			} else {
				for _, command := range job.Commands {
					ExecuteCommand(command, job.Name)
				}
			}
		}(job)
	}
	Wg.Wait()
	fmt.Println("Done.")
}

func Proc(jobs []Job) {
	// fmt.Println("Proc")
	// fmt.Println(jobs)
	for _, job := range jobs {
		if job.Async {
			for _, command := range job.Commands {
				Wg.Add(1)
				go func(command string, name string) {
					// fmt.Println(command + " Async-Async")
					ExecuteCommand(command, name)
					Wg.Done()
				}(command, job.Name)
			}
		} else {
			for _, command := range job.Commands {
				fmt.Println(command)
				ExecuteCommand(command, job.Name)
			}
		}
	}
}

func ExecuteCommand(command string, name string) bool {
	// fmt.Println(command + " being executed .. ")
	// fmt.Println("Executing.")
	args := strings.Fields(command)
	out, err := exec.Command(args[0], args[1:]...).Output()
	// out, err := exec.Command(command).Output()
	if err != nil {
		fmt.Println("[" + name + "] Error: " + err.Error())
		// ResultsChannel <- Result{Name: name, StdOutput: "nil", StdError: err.Error()}
		return false
	} else {
		fmt.Println("[" + name + "] Output: " + string(out))
		// ResultsChannel <- Result{Name: name, StdOutput: string(out), StdError: "nil"}
		return true
	}
}

// func BatchProc(job Job) {

// }

//
// func ExecuteCommand(command Command) {
// 	if false {
// 		fmt.Println("Can't Execute " + command.Name + " on a " + runtime.GOOS + " machine")
// 	} else {
// 		if command.Concurrent {
// 			Sem <- 1
// 			Wg.Add(1)
// 			go func(cmd string) {
// 				out, err := exec.Command(cmd).Output()
// 				if err != nil {
// 					res := Result{Name: command.Name, StdOutput: "nil", StdError: err.Error()}
// 					ResultsChannel <- res
// 					<-Sem
// 					return
// 				}
// 				res := Result{Name: command.Name, StdOutput: string(out), StdError: "nil"}
// 				ResultsChannel <- res
// 				<-Sem
// 			}(command.Command)
// 		} else {
// 			out, err := exec.Command(command.Command).Output()
// 			if err != nil {
// 				res := Result{Name: command.Name, StdOutput: "nil", StdError: err.Error()}
// 				ResultsChannel <- res
// 				return
// 			}
// 			res := Result{Name: command.Name, StdOutput: string(out), StdError: "nil"}
// 			ResultsChannel <- res
// 		}

// 	}
// }
