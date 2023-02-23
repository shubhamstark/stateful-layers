package main

import "github.com/shubhamstark/stateful-layers.git/factories/ab"

func main() {

	c := ab.Build("user-1")
	c.Demo()
}
