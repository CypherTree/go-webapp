package postapi

import (
	"go-webapp/config"
	"go-webapp/db"
	"go-webapp/models"

	"gopkg.in/mgo.v2/bson"
)

// GetBySrcID - Fetch one post by its source id
func GetBySrcID(srcID string, srcType int) (*models.Post, error) {
	post := new(models.Post)

	query := make(bson.M)
	query["src"] = srcType
	query["post_id"] = srcID

	err := db.Conn.GetOne(query, config.PostColl, &post)

	return post, err
}

// Upsert - Insert or update post
func Upsert(post *models.Post) error {
	post.BeforeSave()
	query := bson.M{
		"_id": post.ID,
	}

	err := db.Conn.Upsert(query, config.PostColl, &post)
	return err
}
