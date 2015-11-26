package main

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// Third party packages
	"github.com/julienschmidt/httprouter"
)

type Data struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var Resp map[int]string

func updateKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body

	fmt.Println("putting data")

	id := p.ByName("id")
	value := p.ByName("value")

	i, _ := strconv.Atoi(id)

	fmt.Println(i)
	Resp[i] = value

	w.WriteHeader(200)
	//fmt.Fprintf(w, "%s", uj)
}

func readKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Simply write some test data for now
	fmt.Println("reading values")

	id := p.ByName("id")
	i, _ := strconv.Atoi(id)

	u := Data{}

	if val, ok := Resp[i]; ok {
		fmt.Println(i, val)
		u.Key = i
		u.Value = val
	}
	// Marshal provided interface into JSON structure

	fmt.Println("before marshalling map :", u)

	uj, _ := json.Marshal(u)

	fmt.Println("after marshalling map :", u)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func readAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Simply write some test data for now
	fmt.Println("reading values")

	l := len(Resp)

	dataArr := make([]Data, l)
	i := 0

	for k, v := range Resp {
		dataArr[i].Key = k
		dataArr[i].Value = v
		i++
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(dataArr)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	Resp = make(map[int]string)

	// Add a handler on /test
	r.GET("/keys/:id", readKey)
	r.GET("/keys", readAll)
	r.PUT("/keys/:id/:value", updateKey)
	//	r.DELETE("/user/:id", deleteUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
