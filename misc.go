package glo

// Color infomration related to a Label's colour
type Color struct {
	R int     `json:"r"`
	G int     `json:"g"`
	B int     `json:"b"`
	A float64 `json:"a"`
}

// Label information related to a Label
type Label struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Color       Color        `json:"color"`
	CreatedDate string       `json:"created_date"`
	CreatedBy   *PartialUser `json:"created_date"`
}

// PartialLabel minmized Label data
type PartialLabel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PartialUser minmized User information
type PartialUser struct {
	ID string `json:"id"`
}
