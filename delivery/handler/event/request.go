package event

type EventRequest struct {
	Name        string `json:"name" form:"name"`
	Date        string `json:"date" form:"date"`
	Location    string `json:"location" form:"location"`
	Images      string `json:"images" form:"images"`
	Description string `json:"description" form:"description"`
	Category    string `json:"category" form:"category"`
}
