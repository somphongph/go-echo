package models

type FacebookPost struct {
	// Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Message      string `json:"message" bson:"message,omitempty"`
	Created_time string `json:"created_time" bson:"created_time,omitempty"`
	Post_id      string `json:"post_id" bson:"id,omitempty"`
}
