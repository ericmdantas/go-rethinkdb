package main

import (
	"flag"
	"fmt"
	"log"

	r "gopkg.in/dancannon/gorethink.v2"
)

func session() (*r.Session, error) {
	return r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
}

func create() {
	s, _ := session()

	r.DBCreate("ex").Run(s, r.RunOpts{})
}

func read() {
	s, _ := session()
	
	t, err := r.Table("todo").Run(s)
	
	if err != nil {
		log.Fatalf(err.Error())
	}
	
	fmt.Println(t)
}

func main() {
	c := flag.Bool("c", false, "")
	flag.Parse()

	if *c {
		create()
	} else {
		read()
	}
}
