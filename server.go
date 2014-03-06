package main

import "github.com/codegangsta/martini"

func Server() martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	return *m
}

func main() {
	m := Server()
	m.Run()
}
