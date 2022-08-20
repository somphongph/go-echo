package book

type Book struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name"`
	Title string `json:"title"`
}
