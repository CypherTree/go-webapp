package models

import "gopkg.in/mgo.v2/bson"

const (
	PostSrcFB = 1
	PostSrcIG = 2
	PostSrcTW = 3
)

// Cordinates - Location model
type Cordinates struct {
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}

// Post - Post model
type Post struct {
	BaseModel `bson:",inline"`
	PostID    string        `json:"post_id" bson:"post_id"`
	Src       int           `json:"src" bson:"src"`
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`
	Location  Cordinates    `json:"location" bson:"location"`
	Text      string        `json:"text" bson:"text"`
}
