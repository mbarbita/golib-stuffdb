package stuffdb

import "fmt"

type TierRef struct {
	Name   string
	RefMap map[int]map[int]Target
}

func NewTierRef() *TierRef {
	return &TierRef{
		RefMap: make(map[int]map[int]Target),
	}
}

func (tr *TierRef) Print() {
	// fmt.Println("  RefMap:")
	for k1 := range tr.RefMap {
		fmt.Printf("  tier ref name: %v k(ID): %v\n", tr.Name, k1)
		for k2, v2 := range tr.RefMap[k1] {
			fmt.Printf("    ref k: %v v: %v\n", k2, v2)
		}
	}
}

// func (tr *TierRef) AddRef(id int) {
// 	tr.RefMap[id] = make(map[int]Target)
// }

// func (tr *TierRef) Save(fname string) {
// 	//Save
// 	fmt.Println("save obj:", tr)
// 	if err := storestruct.Save("./"+fname+".json", tr); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func (tr *TierRef) Load(fname string) {
// 	// load
// 	fmt.Println("load obj:", tr)
// 	if err := storestruct.Load("./"+fname+".json", tr); err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func (tr *TierRef) PrintList() {
// 	fmt.Printf("tier listmap %+v\n", tr.RefMap)
// 	for k, v := range tr.RefMap {
// 		fmt.Printf("tier listmap k: %+v v: %+v\n", k, v)
// 	}
// }