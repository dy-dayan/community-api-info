package form

type Building struct {
	ID          int64     `json:"ID"`
	Name        string    `json:"name"`
	Loc         []float32 `json:"loc"`
	ElevatorIDs []int64   `json:"elevatorIDs"`
	CommunityID int64     `json:"communityID"`
	Period      int32     `json:"period"`
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
	OperatorID  int64     `json:"operatorID"`
}
