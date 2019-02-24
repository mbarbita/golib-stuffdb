package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mbarbita/golib-stuffdb"
)

func main() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()

	db := stuffdb.NewDashboard()

	db.AddTier(0)
	db.ModVal(0, 1, "Dude A")
	db.ModVal(0, 2, "Dude B")
	db.ModVal(0, 3, "Dude C")

	db.AddTier(1)
	db.ModVal(1, 1, "Dudes")
	db.ModVal(1, 2, "Friends")

	db.AddRef(1, 1)
	db.ModRef(1, 1, 1, stuffdb.Target{Tier: 0, List: 1}) //add functions
	db.ModRef(1, 1, 2, stuffdb.Target{Tier: 0, List: 2})
	db.ModRef(1, 1, 3, stuffdb.Target{Tier: 0, List: 3})

	db.AddRef(1, 2)
	db.ModRef(1, 2, 1, stuffdb.Target{Tier: 0, List: 2})
	db.ModRef(1, 2, 2, stuffdb.Target{Tier: 0, List: 3})

	db.Print()

	fmt.Println("Save:")
	db.Save("db")
	fmt.Println()

	fmt.Println("Load:")
	var dbs = new(stuffdb.Dashboard)
	dbs.Load("db")

	dbs.Print()

}
