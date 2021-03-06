package stuffdb

import "fmt"

type Data struct {
	Name   string
	IfcMap map[int]interface{}
}

type TierData struct {
	DataMap map[int]*Data
}

func NewData() *Data {
	return &Data{
		Name:   "",
		IfcMap: make(map[int]interface{}),
	}
}

func (d *Data) ModName(name string) {
	d.Name = name
}

func NewTierData() *TierData {
	return &TierData{
		DataMap: make(map[int]*Data),
	}
}

func (td *TierData) Print() {
	for k1, v1 := range td.DataMap {
		fmt.Printf("  data map k(ID): %v v(name): %v\n", k1, v1.Name)
		for k2, v2 := range td.DataMap[k1].IfcMap {
			fmt.Printf("    data k: %+v v: %+v type: %T\n", k2, v2, v2)
		}
		fmt.Println()
	}
}
