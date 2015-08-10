package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func handlerExec(w http.ResponseWriter, r *http.Request) {
	var cmd Command
	var cmdin cmdDTOIN
	var status = http.StatusNotFound

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		sendError(w, status)
		fmt.Println(err)

		return
	}
	if err := r.Body.Close(); err != nil {
		sendError(w, status)
		fmt.Println(err)

		return
	}
	if err := json.Unmarshal(body, &cmdin); err != nil {
		sendError(w, status)
		fmt.Println(err)

		return
	}

	(&cmd).Build(cmdin)
	log.Printf("lsldldld = %s", cmd)
	if pid, err := (&cmd).StartProcess(); err == nil {
		runningProcesses[pid] = cmd
		cmdout := generateCmdOut(cmd)
		if err := json.NewEncoder(w).Encode(cmdout); err != nil {
			sendError(w, status)
		}
		status = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}

func handlerKill(w http.ResponseWriter, r *http.Request) {
	//TODO: Log
	status := http.StatusNotFound
	vars := mux.Vars(r)
	psID := (vars["psId"])
	start := time.Now()
	id, errAtoi := strconv.Atoi(psID)

	if errAtoi != nil {
		sendError(w, status)
	}
	prettyPrintPSList()

	if _, ok := runningProcesses[id]; ok == true {
		runningProcesses[id].StopProcess()
		delete(runningProcesses, id)
		status = http.StatusOK
	} else {
		log.Printf("Process not found (PID = %d)  \t\t%s", id, time.Since(start))
	}

	w.WriteHeader(status)
}

func handlerListPSEntity(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNotFound
	vars := mux.Vars(r)
	entityID := (vars["keyid"])
	start := time.Now()
	id, errAtoi := strconv.Atoi(entityID)

	if errAtoi != nil {
		sendError(w, status)
	}

	if value, ok := runningProcesses[id]; ok != false {
		tmp := generateCmdOut(value)
		if err := json.NewEncoder(w).Encode(tmp); err != nil {
			sendError(w, status)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		log.Printf("Process not found (PID = %d)  \t\t%s", id, time.Since(start))
	}

	w.WriteHeader(status)
}

func handlerListPS(w http.ResponseWriter, r *http.Request) {

	status := http.StatusNotFound

	var temp []cmdDTOOut
	var tmp cmdDTOOut
	for _, value := range runningProcesses {
		tmp = generateCmdOut(value)
		temp = append(temp, tmp)
	}
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		sendError(w, status)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func sendError(w http.ResponseWriter, status int) {
	w.WriteHeader(status) // unprocessable entity
}
