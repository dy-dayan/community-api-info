package form

type Community struct {
	ID           int64     `json:"ID"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serialNumber"`
	Provinces    string    `json:"provinces"`
	City         string    `json:"city"`
	Region       string    `json:"region"`
	Street       string    `json:"street"`
	OrgID        int64     `json:"orgID"`
	HouseCount   int32     `json:"houseCount"`
	CheckInCount int32     `json:"checkInCount"`
	BuildingArea float32   `json:"buildingArea"`
	GreeningArea float32   `json:"greeningArea"`
	SealedState  int32     `json:"SealedState"`
	Loc          []float32 `json:"loc"`
	State        int32     `json:"state"`
	OperatorID   int64     `json:"operatorID"`
	CreatedAt    int64     `json:"createdAt"`
	UpdatedAt    int64     `json:"updatedAt"`
}

type QueryCommunity struct {
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
	Loc      []float32 `json:"loc"`
	Distance float32   `json:"distance"`
}
