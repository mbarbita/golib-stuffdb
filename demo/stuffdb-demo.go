package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"os/exec"

	stuffdb "github.com/mbarbita/golib-stuffdb"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %v is %T\n", v, v)
	case string:
		fmt.Printf("string: %q is %T\n", v, v)
	case map[int]int:
		fmt.Printf("map[int]int: %v is %T\n", v, v)
	default:
		fmt.Printf("unhandeled type: %v is %T\n", v, v)
	}
}

func main() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()

	db := stuffdb.NewDashboard()

	tier := 0
	dataMapID := 0
	db.AddTierData(tier)
	db.AddData(tier, dataMapID)
	db.ModDataName(tier, dataMapID, "Ppl")
	db.ModData(tier, dataMapID, 1, "Dude A")
	db.ModData(tier, dataMapID, 2, "Dude B")
	db.ModData(tier, dataMapID, 3, "Dude C")
	db.ModData(tier, dataMapID, 4, "Dude D")

	// test int
	dataMapID = 1
	db.AddData(tier, dataMapID)
	db.ModDataName(tier, dataMapID, "A number")
	db.ModData(tier, dataMapID, 1, 666.2)

	//test map
	tm := make(map[int]int)
	for i := 1; i <= 10; i++ {
		tm[i] = i + 10
	}
	dataMapID = 2
	db.AddData(tier, dataMapID)
	db.ModDataName(tier, dataMapID, "Some map")
	db.ModData(tier, dataMapID, 1, tm)

	// dataMapID = 3
	// db.AddData(tier, dataMapID)
	// db.ModDataName(tier, dataMapID, "Some cat")
	// db.ModData(tier, dataMapID, 1, "Dudes")
	// db.ModData(tier, dataMapID, 2, "Friends")

	dataMapID = 4
	db.AddData(tier, dataMapID)
	db.ModDataName(tier, dataMapID, "Stuff")
	db.ModData(tier, dataMapID, 1, []string{"slice",
		"of", "strings", "and stuff", "\\ \"", "世界"})

	tier = 1
	refMapID := 1
	db.AddTierRef(tier)
	db.AddRef(tier, refMapID)
	db.ModRefName(tier, refMapID, "Dudes")
	db.ModRef(tier, refMapID, 1, 0, 1)
	db.ModRef(tier, refMapID, 2, 0, 2)
	db.ModRef(tier, refMapID, 3, 0, 3)

	refMapID = 2
	db.AddRef(tier, refMapID)
	db.ModRefName(tier, refMapID, "Friends")
	db.ModRef(tier, refMapID, 1, 0, 2)
	db.ModRef(tier, refMapID, 2, 0, 3)

	db.Print()

	fmt.Println("Save:")
	db.Save("db")
	fmt.Println()

	fmt.Println("Load:")
	var dbs = new(stuffdb.Dashboard)
	dbs.Load("db")

	dbs.Print()

	fmt.Println("Save GOB:")

	gob.Register(stuffdb.Dashboard{})
	gob.Register(tm)

	db.SaveGob("db")
	var dbsg = new(stuffdb.Dashboard)
	fmt.Println("Load GOB:")
	dbsg.LoadGob("db")
	dbsg.Print()

	fmt.Println("DO:------------------")
	do(666)
	do(666.6)
	do("666.666")
	do((dbsg.TierDataMap[0].DataMap[2].IfcMap[1]).(map[int]int))

}
