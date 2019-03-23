package stuffdb

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	storestruct "github.com/mbarbita/golib-storestruct"
)

type Dashboard struct {
	TierRefMap  map[int]*TierRef
	TierDataMap map[int]*TierData
}

func NewDashboard() *Dashboard {
	return &Dashboard{
		TierRefMap:  make(map[int]*TierRef),
		TierDataMap: make(map[int]*TierData),
	}
}

func (db *Dashboard) AddTier(tier int) {
	if (tier % 2) == 0 {
		db.TierDataMap[tier] = NewTierData()
	}
	if (tier % 2) == 1 {
		db.TierRefMap[tier] = NewTierRef()
	}
}

func (db *Dashboard) AddRef(tier, refMapID int) {
	if (tier % 2) == 0 {
		fmt.Println("cant add ref map on even tier", tier)
		return
	}
	db.TierRefMap[tier].RefMap[refMapID] = make(map[int]Target)
}

func (db *Dashboard) ModRef(tier, refMapID, key, targetTier, targetList int) {
	if (tier % 2) == 0 {
		fmt.Println("cant mod ref map on even tier", tier)
		return
	}
	db.TierRefMap[tier].RefMap[refMapID][key] =
		Target{Tier: targetTier, List: targetList}
}

func (db *Dashboard) AddData(tier, dataMapID int) {
	if (tier % 2) == 1 {
		fmt.Println("cant add data map on odd tier", tier)
		return
	}
	// db.TierDataMap[tier].DataMap[dataMapID] = make(map[int]Data)
	db.TierDataMap[tier].DataMap[dataMapID] = NewData()
}

func (db *Dashboard) ModDataName(tier, dataMapID int, name string) {
	if (tier % 2) == 1 {
		fmt.Println("cant add data map on odd tier", tier)
		return
	}
	t := db.TierDataMap[tier].DataMap[dataMapID]
	t.Name = name
	db.TierDataMap[tier].DataMap[dataMapID] = t
}

func (db *Dashboard) ModData(tier, dataMapID, key int, value interface{}) {
	if (tier % 2) == 1 {
		fmt.Println("cant mod data map on odd tier", tier)
		return
	}
	db.TierDataMap[tier].DataMap[dataMapID].IfcMap[key] = value
}

func (db *Dashboard) Save(fname string) {
	//Save
	fmt.Println("save obj:", db)
	if err := storestruct.Save(fname+".json", db); err != nil {
		log.Println("save json obj err:", err)
	}
}

func (db *Dashboard) Load(fname string) {
	// load
	fmt.Println("load obj:", db)
	if err := storestruct.Load(fname+".json", db); err != nil {
		log.Panicln("load json obj err", err)
	}
}

func (db *Dashboard) Print() {
	fmt.Printf("dashboard %+v\n", db)
	fmt.Println()
	fmt.Println("TierRefMap")
	fmt.Println("----------")
	for k1 := range db.TierRefMap {
		fmt.Printf("T%v:\n", k1)
		db.TierRefMap[k1].Print()
		fmt.Println()
	}
	fmt.Println("TierDataMap")
	fmt.Println("----------")
	for k1 := range db.TierDataMap {
		fmt.Printf("T%v:\n", k1)
		db.TierDataMap[k1].Print()
		fmt.Println()
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
			log.Println("encode gob error:", err)
		}
	}
	if err != nil {
		log.Println("save gob error:", err)
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
			log.Println("decode gob error:", err)
		}
	}
	if err != nil {
		log.Println("load gob error:", err)
	}
}
