package model

// model bread
type Bread struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Taste string `json:"rasa"`
	Price int    `json:"price"`
}
