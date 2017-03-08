package models

import (
	"go-webapp/config"
	"go-webapp/db"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BaseModel - Base model
type BaseModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" required:"true"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}

// Save - Save model to db
func (m *BaseModel) Save(tableName string) error {
	if !m.ID.Valid() {
		m.ID = bson.NewObjectId()
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	if !(m.CreatedOn > 0) {
		m.CreatedOn = now
	}

	m.UpdatedOn = now

	query := bson.M{
		"_id": m.ID,
	}

	err := db.Conn.Upsert(query, config.UserColl, &m)

	return err
}
