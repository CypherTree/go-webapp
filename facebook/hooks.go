package facebook

import (
	"encoding/json"
	"go-webapp/config"
	"time"

	"go-webapp/db"
	"go-webapp/dbapi/user"

	"github.com/gin-gonic/gin"
)

var UpdateChannel chan fbUpdate

type fbUpdateEntry struct {
	Time          uint32   `json:"time" bson:"time"`
	ID            string   `json:"id" bson:"id"`
	ChangedFields []string `json:"changed_fields" bson:"changed_fields"`
}

type fbUpdate struct {
	Entry  []fbUpdateEntry `json:"entry" bson:"entry"`
	Object string          `json:"object" bson:"object"`
}

// Subscribe - get endpoint for facebook subscribe
func Subscribe(c *gin.Context) {
	if hookToken, _ := c.GetQuery("hub.verify_token"); hookToken != config.Settings.FbWebHookToken {
		c.JSON(400, "Wrong token")
		return
	}

	hubChallenge, _ := c.GetQuery("hub.challenge")
	c.String(200, hubChallenge)

	// if previously subscribed and channel is running then close it.
	if UpdateChannel != nil {
		close(UpdateChannel)
	}

	// (re)create channel for handling updates from fb
	UpdateChannel = make(chan fbUpdate, 100)
	go ListenUpdates()
}

// Listen - Facebook changes listener
func Listen(c *gin.Context) {
	c.String(200, "")
	update := fbUpdate{}
	json.NewDecoder(c.Request.Body).Decode(&update)
	UpdateChannel <- update
}

func ListenUpdates() {
	redisPrefix := "LastFbUpdate_"
	for update := range UpdateChannel {
		if update.Entry == nil || len(update.Entry) == 0 {
			continue
		}

		since, _ := db.Redis.Get(redisPrefix + update.Entry[0].ID).Result()
		user, err := userapi.GetByFbID(update.Entry[0].ID)

		if err != nil || user == nil {
			continue
		}

		FetchFeed(user, since)

		db.Redis.Set(redisPrefix+update.Entry[0].ID, update.Entry[0].Time, time.Duration(time.Hour*24*10))
	}
}
