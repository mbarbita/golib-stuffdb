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

func (db *Dashboard) AddTier(tier int) {
	db.TierValMap[tier] = NewTierVal()
	if (tier % 2) == 1 {
		db.TierRefMap[tier] = NewTierRef()
	}
}

// func (db *Dashboard) AddTierRef(tier int) {
// 	db.TierRefMap[tier] = NewTierRef()
// }
//
// func (db *Dashboard) AddTierVal(tier int) {
// 	db.TierValMap[tier] = NewTierVal()
// }

func (db *Dashboard) AddRef(tier, listID int) {
	if (tier % 2) == 0 {
		fmt.Println("Cant add ref map on even tier", tier)
		return
	}
	db.TierRefMap[tier].RefMap[listID] = make(map[int]Target)
}

func (db *Dashboard) ModRef(tier, listID, key, targetTier, targetList int) {
	if (tier % 2) == 0 {
		fmt.Println("Cant mod ref map on even tier", tier)
		return
	}
	db.TierRefMap[tier].RefMap[listID][key] =
		Target{Tier: targetTier, List: targetList}
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
	fmt.Printf("dashboard %+v\n", db)
	fmt.Println()
	fmt.Println("TierRefMap")
	fmt.Println("----------")
	for k1 := range db.TierRefMap {
		fmt.Printf("T%v:\n", k1)
		// fmt.Println()
		fmt.Println("  RefMap")
		for k2 := range db.TierRefMap[k1].RefMap {
			fmt.Printf("    list id: %v\n", k2)
			for k3, v3 := range db.TierRefMap[k1].RefMap[k2] {
				fmt.Printf("      k: %v v: %v\n", k3, v3)
			}
		}
		fmt.Println()
	}
	fmt.Println("TierValMap")
	fmt.Println("----------")
	for k := range db.TierValMap {
		fmt.Printf("T%v:\n", k)
		fmt.Println("  ValMap")
		for k, v := range db.TierValMap[k].ValMap {
			fmt.Printf("    k: %v v: %v type: %T\n", k, v, v)
		}
	}
	fmt.Println()
}

func (db *Dashboard) SaveGob(filePath string) {
	// gob.Register(Dashboard{})
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
	// gob.Register(Dashboard{})
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
