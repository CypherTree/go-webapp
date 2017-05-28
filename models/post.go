package models

import "gopkg.in/mgo.v2/bson"

const (
	PostSrcFB = "FB"
	PostSrcIG = "IG"
	PostSrcTW = "TW"
)

// Post - Post model
type Post struct {
	BaseModel     `bson:",inline"`
	PostID        string        `json:"post_id" bson:"post_id"`
	Src           string        `json:"src" bson:"src"`
	ReplyCount    int16         `json:"replay_count" bson:"reply_count"`
	ReactionCount int16         `json:"reaction_count" bson:"reaction_count"`
	UserID        bson.ObjectId `json:"user_id" bson:"user_id"`
}
