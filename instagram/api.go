package instagram

import (
	"encoding/json"
	"errors"
	"go-webapp/config"
	"go-webapp/models"
	"net/http"
	"net/url"
	"strings"
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
