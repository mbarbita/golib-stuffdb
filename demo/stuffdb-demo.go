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

	tier := 0
	db.AddTier(tier)
	db.ModVal(tier, 1, "Dude A")
	db.ModVal(tier, 2, "Dude B")
	db.ModVal(tier, 3, "Dude C")
	db.ModVal(tier, 4, 666)

	tier = 1
	db.AddTier(tier)
	db.ModVal(tier, 1, "Dudes")
	db.ModVal(tier, 2, "Friends")

	listID := 1
	db.AddRef(tier, listID)
	db.ModRef(tier, listID, 1, 0, 1) //add functions
	db.ModRef(tier, listID, 2, 0, 2)
	db.ModRef(tier, listID, 3, 0, 3)

	listID = 2
	db.AddRef(tier, listID)
	db.ModRef(tier, listID, 1, 0, 2)
	db.ModRef(tier, listID, 2, 0, 3)

	db.Print()

	fmt.Println("Save:")
	db.Save("db")
	fmt.Println()

	fmt.Println("Load:")
	var dbs = new(stuffdb.Dashboard)
	dbs.Load("db")

	dbs.Print()

	fmt.Println("Save GOB:")
	db.SaveGob("db")
	var dbsg = new(stuffdb.Dashboard)
	fmt.Println("Load GOB:")
	dbsg.LoadGob("db")
	dbsg.Print()

}
