package main

import (
	"time"
)

type cmdDTOIN struct {
	CmdName   string `json:"CmdName"`
	Arguments string `json:"Arguments"`
}

type cmdDTOOut struct {
	CmdName   string    `json:"CmdName"`
	Arguments string    `json:"Arguments"`
	PID       int       `json:"PID"`
	StartTime time.Time `json:"StartTime"`
}

func generateCmdOut(cmd Command) cmdDTOOut {
	var out cmdDTOOut
	out.CmdName = cmd.CmdName
	out.Arguments = cmd.Arguments
	out.PID = cmd.PID
	out.StartTime = cmd.StartingTime
	return out
}
