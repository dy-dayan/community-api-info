package form

type Asset struct {
	ID           int64  `json:"ID"`
	SerialNumber string `json:"serialNumber"`
	Category     int32  `json:"category"`
	Loc          string `json:"loc"`
	State        int32  `json:"state"`
	CommunityID  int64  `json:"communityID"`
	Brand        string `json:"brand"`
	Desc         string `json:"desc"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	OperatorID   int64  `json:"operatorID"`
}
