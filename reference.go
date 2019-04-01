package stuffdb

import "fmt"

type Ref struct {
	Name      string
	TargetMap map[int]Target
}

type TierRef struct {
	RefMap map[int]*Ref
}

func NewRef() *Ref {
	return &Ref{
		Name:      "",
		TargetMap: make(map[int]Target),
	}
}

func (r *Ref) ModName(name string) {
	r.Name = name
}

func NewTierRef() *TierRef {
	return &TierRef{
		RefMap: make(map[int]*Ref),
	}
}

func (tr *TierRef) Print() {
	// fmt.Println("  RefMap:")
	for k1, v1 := range tr.RefMap {
		fmt.Printf("  tier ref k(ID): %v v(Name): %v\n", k1, v1.Name)
		for k2, v2 := range tr.RefMap[k1].TargetMap {
			fmt.Printf("    ref k: %v v: %v\n", k2, v2)
		}
	}
}
