package form

type House struct {
	ID         int64   `json:"ID"`
	BuildingID int64   `json:"buildingID"`
	Unit       string  `json:"unit"`
	Acreage    float32 `json:"acreage"`
	State      int32   `json:"state"`
	Rental     int32   `json:"rental"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
	OperatorID int64   `json:"operatorID"`
}
