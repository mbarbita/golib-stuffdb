package stuffdb

type Target struct {
	Tier     int
	MapID    int
	Selector []string
}

func NewTarget(tier, mapID int, selector ...string) *Target {
	return &Target{
		Tier:     tier,
		MapID:    mapID,
		Selector: selector,
	}
}
