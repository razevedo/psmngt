package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"handlerExec",
		"POST",
		"/exec",
		handlerExec,
	},
	Route{
		"handlerListPSEntity",
		"GET",
		"/ps/{keyid}",
		handlerListPSEntity,
	},
	Route{
		"handlerListPS",
		"GET",
		"/ps",
		handlerListPS,
	},
	Route{
		"handlerKill",
		"DELETE",
		"/kill/{psId}",
		handlerKill,
	},
}
