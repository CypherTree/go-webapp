package facebook

import (
	"encoding/json"
	"errors"
	"go-webapp/config"
	"go-webapp/models"
	"io/ioutil"
	"net/http"
	"net/url"

	"fmt"

	"github.com/tidwall/gjson"
)

// GetAccessToken - Get access token
func GetAccessToken(fb *models.Facebook, q url.Values) error {
	resp, _ := http.Get(config.FbAccessTokenURL + "?" + q.Encode())

	if resp.StatusCode != http.StatusOK {
		return errors.New("Not able to get access token. Status Code: " + string(resp.StatusCode))
	}

	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&fb)

	if err != nil {
		return err
	}

	return nil
}

// GetProfile - Get facebook profile
func GetProfile(fb *models.Facebook) error {
	q := url.Values{}
	q.Set("access_token", fb.AccessToken)
	q.Add("fields", "id,email,first_name,last_name,link,name,birthday,gender,picture")
	userResp, _ := http.Get(config.FbGraphAPIURL + "/me?" + q.Encode())

	if userResp.StatusCode != http.StatusOK {
		return errors.New("Not able to fetch facebook profile. Status Code: " + string(userResp.StatusCode))
	}

	err := json.NewDecoder(userResp.Body).Decode(&fb)

	if err != nil {
		return err
	}

	return nil
}

// SaveInitFeed - Fetch first n records from /feed of facebook api
func SaveInitFeed(fb *models.Facebook) error {
	q := url.Values{}
	q.Set("access_token", fb.AccessToken)
	q.Add("fields", "reactions.limit(0).summary(total_count),message,full_picture,type,attachments{url},comments.limit(0).summary(total_count)")
	q.Add("limit", "100")
	userResp, _ := http.Get(config.FbGraphAPIURL + "/me/feed?" + q.Encode())

	if userResp.StatusCode != http.StatusOK {
		return errors.New("Not able to fetch facebook feed")
	}

	defer userResp.Body.Close()
	json, err := ioutil.ReadAll(userResp.Body)

	if err != nil {
		return errors.New("Unable to parse response from facebook feed: " + err.Error())
	}

	jsonStr := string(json[:])
	dataJSON := gjson.Get(jsonStr, "data")

	for _, item := range dataJSON.Array() {
		fmt.Println(item.String())
	}

	return nil
}
