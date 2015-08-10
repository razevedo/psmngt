package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
	"utils"
)

type Command struct {
	CmdName      string
	Arguments    string
	PID          int
	StartingTime time.Time
	running      *exec.Cmd
}

/*
func run() {
	for state := startState; state != nil {
		state = state(lex)
	}
}

*/

/*a String func for a type is used by fmt.print* - !
func (cmd Command) String() string {
	return fmt.Sprintf("ola")
}
*/
func (cmd *Command) Build(cmdin cmdDTOIN) {
	cmd.CmdName = cmdin.CmdName
	cmd.Arguments = cmdin.Arguments
	cmd.StartingTime = time.Now()
}

func (cmd Command) StopProcess() error {
	//TODO: log
	/*if cmd == nil {
		return nil
	}*/
	start := time.Now()
	if err := cmd.running.Process.Kill(); err != nil {
		log.Printf("Error: Process with %d was not terminated \t\t%s", cmd.PID, time.Since(start))
		return err
	}
	log.Printf("Process with %d terminated \t\t%s", cmd.PID, time.Since(start))
	return nil
}

func (cmd *Command) StartProcess() (n int, e error) {
	pid, err := cmd.startProcess()

	if err != nil || pid == 0 {
		return 0, err
	}
	return pid, nil
}

func (cmd *Command) startProcess() (n int, err error) {
	//TODO:log
	start := time.Now()
	args := strings.Split(cmd.Arguments, " ")
	execcmd := exec.Command(cmd.CmdName, args...)

	if err := execcmd.Start(); err != nil {
		fmt.Println("Error running ", err)
		log.Printf("Error: Process not started - %s \t\t%s", args, time.Since(start))
		return 0, err
	}

	cmd.PID = execcmd.Process.Pid
	cmd.StartingTime = time.Now()
	cmd.running = execcmd
	log.Printf("Process started - %s with PID - %d\t\t%s", utils.ConcateStrings(cmd.CmdName, cmd.Arguments), cmd.PID, time.Since(start))
	return cmd.PID, nil

}
