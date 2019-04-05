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

func (db *Dashboard) AddTierData(tier int) {
	db.TierDataMap[tier] = NewTierData()
}

func (db *Dashboard) AddTierRef(tier int) {
	db.TierRefMap[tier] = NewTierRef()
}

func (db *Dashboard) AddRef(tier, refMapID int) {
	db.TierRefMap[tier].RefMap[refMapID] = NewRef()
}

func (db *Dashboard) ModRefName(tier, refMapID int, name string) {
	db.TierRefMap[tier].RefMap[refMapID].ModName(name)
}

func (db *Dashboard) AddTarget(tier, refMapID, key,
	targetTier, targetMapID int, selector ...string) {
	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key] =
		NewTarget(targetTier, targetMapID, selector...)
}

// func (db *Dashboard) ModTarget(tier, refMapID, key,
// 	targetTier, targetMapID int, selector ...string) {
// 	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Tier = tier
// 	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].MapID = refMapID
// 	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Selector = selector
// }

func (db *Dashboard) ModTargetTier(tier, refMapID, key,
	targetTier int) {
	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Tier = tier
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].MapID = refMapID
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Selector = selector
}

func (db *Dashboard) ModTargetMapID(tier, refMapID, key,
	targetMapID int) {
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Tier = tier
	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].MapID = refMapID
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Selector = selector
}

func (db *Dashboard) ModTargetSelector(tier, refMapID, key int,
	selector ...string) {
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Tier = tier
	// db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].MapID = refMapID
	db.TierRefMap[tier].RefMap[refMapID].TargetMap[key].Selector = selector
}

func (db *Dashboard) AddData(tier, dataMapID int) {
	db.TierDataMap[tier].DataMap[dataMapID] = NewData()
}

func (db *Dashboard) ModDataName(tier, dataMapID int, name string) {
	db.TierDataMap[tier].DataMap[dataMapID].ModName(name)
}

func (db *Dashboard) ModData(tier, dataMapID, key int, value interface{}) {
	db.TierDataMap[tier].DataMap[dataMapID].IfcMap[key] = value
}

func (db *Dashboard) SaveJSON(fname string) {
	//Save
	fmt.Println("save obj:", db)
	if err := storestruct.Save(fname+".json", db); err != nil {
		log.Println("save json obj err:", err)
	}
}

func (db *Dashboard) LoadJSON(fname string) {
	// load
	fmt.Println("load obj:", db)
	if err := storestruct.Load(fname+".json", db); err != nil {
		log.Println("load json obj err", err)
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
		// fmt.Println()
	}
	fmt.Println("TierDataMap")
	fmt.Println("-----------")
	for k1 := range db.TierDataMap {
		fmt.Printf("T%v:\n", k1)
		db.TierDataMap[k1].Print()
		// fmt.Println()
	}
	fmt.Println()
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
