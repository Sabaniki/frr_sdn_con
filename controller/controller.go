package main

import (
	controller "github.com/Sabaniki/frr_sdn_con/controller/lib"
)

func main() {
	go controller.Start()
	for {
	}
}
