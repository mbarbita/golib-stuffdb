package stuffdb

import "fmt"

type TierData struct {
	DataMap map[int]map[int]interface{}
}

func NewTierData() *TierData {
	return &TierData{
		DataMap: make(map[int]map[int]interface{}),
	}
}

func (td *TierData) Print() {
	// fmt.Println("  DataMap:")
	for k1 := range td.DataMap {
		fmt.Printf("  tier data k(ID): %v\n", k1)
		for k2, v2 := range td.DataMap[k1] {
			fmt.Printf("    data k: %+v v: %+v\n", k2, v2)
		}
	}
}
