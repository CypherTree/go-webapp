package instagram

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-webapp/config"
	"go-webapp/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

type InstagramUser struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	ProfilePic string `json:"profile_picture"`
	Username   string `json:"username"`
}

type InstagramResponse struct {
	User        InstagramUser `json:"user"`
	AccessToken string        `json:"access_token"`
}

// GetAccessToken - Get access token
func GetAccessToken(ig *models.Instagram, q url.Values) error {
	resp, _ := http.Post(config.IgAccessTokenURL, "application/x-www-form-urlencoded", strings.NewReader(q.Encode()))

	if resp.StatusCode != http.StatusOK {
		return errors.New("Not able to get access token. Status Code: " + string(resp.StatusCode))
	}

	var igResp InstagramResponse
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&igResp)

	if err != nil {
		return err
	}

	names := strings.Split(igResp.User.FullName, " ")
	ig.ProfileID = igResp.User.ID
	ig.FirstName = names[0]
	ig.LastName = names[1]
	ig.Username = igResp.User.Username
	ig.AccessToken = igResp.AccessToken

	return nil
}

// SaveInitFeed - Fetch followers and initial posts of user from instagram
func SaveInitFeed(ig *models.Instagram) error {
	q := url.Values{}
	q.Set("access_token", ig.AccessToken)
	userResp, _ := http.Get(config.IgAPIURL + "/users/self/followed-by?" + q.Encode())
	fmt.Println("status from ig is ", userResp.StatusCode)
	if userResp.StatusCode != http.StatusOK {
		return errors.New("Not able to fetch instagram followers")
	}

	defer userResp.Body.Close()
	json, err := ioutil.ReadAll(userResp.Body)

	if err != nil {
		return errors.New("Unable to parse response from facebook feed: " + err.Error())
	}

	jsonStr := string(json[:])
	dataJSON := gjson.Get(jsonStr, "data")

	fmt.Println("Got something", dataJSON)
	for _, item := range dataJSON.Array() {
		fmt.Println(item.String())
	}

	return nil
}
