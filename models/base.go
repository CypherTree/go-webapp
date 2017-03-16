package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BaseModel - Base model
type BaseModel struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" required:"true"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}

// BeforeSave - Set object id and other attrs before save
func (m *BaseModel) BeforeSave() {
	if !m.ID.Valid() {
		m.ID = bson.NewObjectId()
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	if !(m.CreatedOn > 0) {
		m.CreatedOn = now
	}

	m.UpdatedOn = now
}
