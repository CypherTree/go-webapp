package facebook

import (
	"encoding/json"
	"fmt"
	"go-webapp/config"

	"github.com/gin-gonic/gin"
)

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
}

// Listen - Facebook changes listener
func Listen(c *gin.Context) {
	c.String(200, "")
	update := fbUpdate{}
	json.NewDecoder(c.Request.Body).Decode(&update)
	fmt.Printf("New update: ", update)
}
