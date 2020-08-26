package events

type getEventResponse struct{
	name string `json:"name"`
	date string `json:due""`
	priority int8 `json:"priority"`
}