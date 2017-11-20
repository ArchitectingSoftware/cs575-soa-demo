package server

type Pub struct {
	Id int		 	`json:"id"`
	Title string 	`json:"title"`
	Cite string		`json:"cite"`
	Link string		`json:"link"`
	Slides string	`json:"slides"`
	Abstract string `json:"abstract"`

}
