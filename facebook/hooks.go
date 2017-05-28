package facebook

import (
	"fmt"
	"go-webapp/config"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

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
	fmt.Println("I received something")
	defer c.Request.Body.Close()
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s", string(x))

	c.String(200, "")
}
