package stuffdb

import "fmt"

type TierVal struct {
	// Level  int
	ValMap map[int]interface{}
}

func NewTierVal() *TierVal {
	return &TierVal{
		// Level:  level,
		ValMap: make(map[int]interface{}),
		// ListMap: make(map[int]map[int]Target),
	}
}

func (tv *TierVal) PrintVal() {
	// fmt.Printf("tier  level: %+v valmap: %+v\n", tier.Level, tier.ValMap)
	fmt.Printf("valmap: %+v\n", tv.ValMap)
	for k, v := range tv.ValMap {
		fmt.Printf("tier valmap k: %+v v: %+v\n", k, v)
	}
}
