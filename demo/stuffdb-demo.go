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
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case map[int]int:
		fmt.Printf("map[int]int: %v\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

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

	// test int
	db.ModVal(tier, 4, 666.2)

	//test map
	tm := make(map[int]int)
	for i := 1; i <= 10; i++ {
		tm[i] = i + 10
	}
	db.ModVal(tier, 5, tm)

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

	//reflect
	// fmt.Println("reflect:", reflect.TypeOf(dbs.TierValMap[0].ValMap[4]))
	// rfv := reflect.ValueOf(dbs.TierValMap[0].ValMap[4])
	// fmt.Println("reflect val:", rfv.Float())
	// fmt.Printf("int: %v\n", int(rfv.Float()))
	// do(int(rfv.Float()))

	dbs.Print()

	fmt.Println("Save GOB:")

	gob.Register(stuffdb.Dashboard{})
	gob.Register(tm)

	db.SaveGob("db")
	var dbsg = new(stuffdb.Dashboard)
	fmt.Println("Load GOB:")
	dbsg.LoadGob("db")
	dbsg.Print()

	do((dbsg.TierValMap[0].ValMap[5]).(map[int]int))

}
