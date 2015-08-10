package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Client struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Rule  string `json:"rule"`
}

type clients []Client

/*
func (a *) UnmarshalJSON(b []byte) (err error) {
	j, s, n := author{}, "", uint64(0)
	if err = json.Unmarshal(b, &j); err == nil {
		*a = Author(j)
		return
	}
	if err = json.Unmarshal(b, &s); err == nil {
		a.Email = s
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		a.ID = n
	}
	return
}*/

var authorizedClients map[string]Client

func isAuthorized(token string) bool {
	if _, ok := authorizedClients[token]; ok {
		return true
	}
	return false
}

//TODO: refactor
func loadAuthorizedClients(filename string) error {
	authorizedClients = make(map[string]Client)

	file, e := ioutil.ReadFile(filename)
	if e != nil {
		log.Printf("File error: %v\n", e)
		return e
	}
	var jsontype clients
	json.Unmarshal(file, &jsontype)
	for _, each := range jsontype {
		authorizedClients[each.Token] = each
	}
	return nil //TODO:
}
