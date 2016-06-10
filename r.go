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
	//s, _ := session()
}

func read() {
	s, _ := session()
	res, err := r.Table("2do").Filter(map[string]string{"msg": "2"}, r.FilterOpts{}).Run(s)

	if err != nil {
		log.Fatalf(err.Error())
	}

	var rows []map[string]interface{}

	err = res.All(&rows)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, v := range rows {
		fmt.Println(v["msg"])
	}
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
