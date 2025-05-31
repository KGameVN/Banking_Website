package main

import (
	"comb.com/banking/controller"
)

func main() {
	c := controller.GetController()
	c.Start()
}
