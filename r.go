package main

import (
	"flag"
	"fmt"

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
	fmt.Println(s.Server())
}

func main() {
	c := flag.Bool("c", true, "")
	flag.Parse()

	if *c {
		create()
	} else {
		read()
	}
}
