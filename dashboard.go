package stuffdb

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	storestruct "github.com/mbarbita/golib-storestruct"
)

type Dashboard struct {
	TierRefMap map[int]*TierRef
	TierValMap map[int]*TierVal
}

func NewDashboard() *Dashboard {
	return &Dashboard{
		TierRefMap: make(map[int]*TierRef),
		TierValMap: make(map[int]*TierVal),
	}
}

func (db *Dashboard) AddTierRef(level int) {
	db.TierRefMap[level] = NewTierRef()
}

func (db *Dashboard) AddTierVal(level int) {
	db.TierValMap[level] = NewTierVal()
}

func (db *Dashboard) AddRef(tier, listID int) {
	db.TierRefMap[tier].RefMap[listID] = make(map[int]Target)
}

func (db *Dashboard) ModRef(tier, listID, key int, value Target) {
	db.TierRefMap[tier].RefMap[listID][key] = value
}

func (db *Dashboard) ModVal(tier, key int, value interface{}) {
	db.TierValMap[tier].ValMap[key] = value
}

func (db *Dashboard) Save(fname string) {
	//Save
	fmt.Println("save obj:", db)
	if err := storestruct.Save(fname+".json", db); err != nil {
		log.Fatalln(err)
	}
}

func (db *Dashboard) Load(fname string) {
	// load
	fmt.Println("load obj:", db)
	// var dbs = new(Dashboard)
	if err := storestruct.Load(fname+".json", db); err != nil {
		log.Fatalln(err)
	}
}

func (db *Dashboard) Print() {
	fmt.Printf("dashboard %v\n", db)
	fmt.Println()
	fmt.Println("TierRefMap")
	for k := range db.TierRefMap {
		fmt.Printf("T%v:\n", k)
		// fmt.Println()
		fmt.Println("  RefMap")
		for k := range db.TierRefMap[k].RefMap {
			fmt.Printf("    list id: %v\n", k)
			for k, v := range db.TierRefMap[k].RefMap[k] {
				fmt.Printf("      k: %v v: %v\n", k, v)
			}
		}
		fmt.Println()
	}
	fmt.Println("TierValMap")
	for k := range db.TierValMap {
		fmt.Printf("T%v:\n", k)
		fmt.Println("  ValMap")
		for k, v := range db.TierValMap[k].ValMap {
			fmt.Printf("    k: %v v: %v type: %T\n", k, v, v)
		}
	}
}

func (db *Dashboard) SaveGob(filePath string) {
	gob.Register(Dashboard{})
	// gob.Register(Stuff{})
	file, err := os.Create(filePath + ".gob")
	defer file.Close()
	if err == nil {
		encoder := gob.NewEncoder(file)
		err := encoder.Encode(db)
		if err != nil {
			fmt.Println("gob encode error:", err)
		}
	}
	if err != nil {
		fmt.Println("gob save error:", err)
	}
}

func (db *Dashboard) LoadGob(filePath string) {
	gob.Register(Dashboard{})
	file, err := os.Open(filePath + ".gob")
	defer file.Close()
	if err == nil {
		decoder := gob.NewDecoder(file)
		err := decoder.Decode(db)
		if err != nil {
			fmt.Println("gob decode error:", err)
		}
	}
	if err != nil {
		fmt.Println("gob load error:", err)
	}
}
