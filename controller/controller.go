package main

import (
	controller "github.com/Sabaniki/frr_sdn_con/controller/lib"
)

func main() {
	controller.SetMed("IMPORT_from_vSIX_BB", 15, "permit", 11)
}
