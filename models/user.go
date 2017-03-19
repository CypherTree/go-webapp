package models

import "gopkg.in/mgo.v2/bson"

const (
	UserRoleUnknown  = 0
	UserRoleExpert   = 1
	UserRoleCustomer = 2
)

// BaseThirdParty - Base type for all third party login models
type BaseThirdParty struct {
	ProfileID   string `json:"id,_omitempty" bson:"id" required:"true"`
	AccessToken string `json:"access_token" required:"true" bson:"access_token"`
	ExpiresIn   int64  `json:"expires_in" bson:"expires_in"`
	Email       string `json:"email" bson:"email"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
}

// Facebook - Facebook login model
type Facebook struct {
	BaseThirdParty `bson:",inline"`
	Gender         string `json:"gender" bson:"gender"`
	ProfileLink    string `json:"link" bson:"link"`
}

// Instagram - Instagram login model
type Instagram struct {
	BaseThirdParty `bson:",inline"`
	Username       string `json:"username" bson:"username"`
}

// Twitter - Twitter login model
type Twitter struct {
	BaseThirdParty `bson:",inline"`
}

// User - User model
type User struct {
	BaseModel `bson:",inline"`
	Role      int       `json:"role" bson:"role"`
	Pic       string    `json:"picture" bson:"picture"`
	Email     string    `json:"email" bson:"email"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Fb        Facebook  `json:"facebook" bson:"facebook"`
	Ig        Instagram `json:"instagram" bson:"instagram"`
	Tw        Instagram `json:"twitter" bson:"twitter"`
}

// Session - User session model
type Session struct {
	UserID      bson.ObjectId
	UserRole    int
	TokenString string
	IssuedAt    int64
}
