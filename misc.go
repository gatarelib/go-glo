package glo

// Color information related to a Label's colour
type Color struct {
	R int     `json:"r"`
	G int     `json:"g"`
	B int     `json:"b"`
	A float64 `json:"a"`
}

// Label contains information related ot a label
type Label struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Color       Color        `json:"color"`
	CreatedDate string       `json:"created_date"`
	CreatedBy   *PartialUser `json:"created_by"`
}

// PartialLabel minimized Label data
type PartialLabel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PartialUser minimized User information
type PartialUser struct {
	ID string `json:"id"`
}
