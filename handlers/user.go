package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"go-webapp/config"
	"go-webapp/dbapi/user"
	"go-webapp/facebook"
	"go-webapp/models"
	"go-webapp/utils"
)

// UserViewModel - User view model
type UserViewModel struct {
	ID        bson.ObjectId `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Token     string        `json:"token"`
}

// Save user, handle response and return token if necessary
func saveUser(user *models.User, c *gin.Context) {
	session := c.MustGet("UserSession").(*models.Session)

	err := userapi.Upsert(user)

	if err != nil {
		panic(err)
	}

	userVM := &UserViewModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ID:        user.BaseModel.ID,
	}

	if session.UserID == "" {
		userVM.Token = userapi.GenerateToken(*user)
	}

	c.JSON(200, userVM)
}

// Facebook - handler for facebook login
func Facebook(c *gin.Context) {
	var fb models.Facebook

	// get access token
	q, _, err := utils.MakeOauthQparams(c.Request.Body)
	q.Add("grant_type", "authorization_code")
	q.Add("client_secret", config.Settings.FbSecret)
	err = facebook.GetAccessToken(&fb, q)

	if err != nil {
		panic(err)
	}

	// get facebook profile
	err = facebook.GetProfile(&fb)
	if err != nil {
		panic(err)
	}

	user, err := userapi.GetByFbID(fb.ProfileID)
	if err != nil {
		panic(err)
	}

	if user == nil {
		user = &models.User{
			Fb:        fb,
			FirstName: fb.FirstName,
			LastName:  fb.LastName,
			Email:     fb.Email,
		}
	} else {
		user.Fb = fb
	}

	saveUser(user, c)
}
