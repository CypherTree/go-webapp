package facebook

import (
	"encoding/json"
	"errors"
	"go-webapp/config"
	"go-webapp/models"
	"net/http"
	"net/url"

	"strings"

	postapi "go-webapp/dbapi/post"
)

type FbPostVM struct {
	Message string            `json:"message" bson:"message"`
	Place   models.Cordinates `json:"place" bson:"place"`
	ID      string            `json:"id" bson:"id"`
	Link    string            `json:"link" bson:"link"`
	Picture string            `json:"full_picture" bson:"full_picture"`
}

type FbPostCollectionVM struct {
	Posts []FbPostVM `json:"data" bson:"data"`
}

// FbAPIFields Graph API fields to be fetched
var FbAPIFields = []string{
	"link",
	"story",
	"message",
	"full_picture",
	"place{location}",
}

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

// FetchFeed - Fetch first n records from /feed of facebook api
func FetchFeed(user *models.User, since uint32) error {

	q := url.Values{}
	q.Set("access_token", user.Fb.AccessToken)
	q.Add("fields", strings.Join(FbAPIFields, ","))

	if since != 0 {
		q.Add("since", string(since))
	}

	reponse, _ := http.Get(config.FbGraphAPIURL + "/me/feed?" + q.Encode())

	if reponse.StatusCode != http.StatusOK {
		return errors.New("Not able to fetch facebook feed")
	}

	posts := FbPostCollectionVM{}
	json.NewDecoder(reponse.Body).Decode(&posts)

	for _, post := range posts.Posts {
		existing, _ := postapi.GetBySrcID(post.ID, models.PostSrcFB)

		if existing == nil {
			existing = &models.Post{}
			existing.UserID = user.ID
		}

		existing.Text = post.Message
		existing.PostID = post.ID
		existing.Location = post.Place
		existing.Src = models.PostSrcFB

		postapi.Upsert(existing)
	}

	return nil
}
