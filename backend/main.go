package main

import (
	"comb.com/banking/controller"
	// "comb.com/banking/utils/logger"
	// "comb.com/banking/config"
)

func main() {
	// l := logger.NewLogger()
	

	c := controller.GetController()
	c.Start()
	


}
