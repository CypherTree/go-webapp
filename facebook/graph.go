package facebook

import (
	"encoding/json"
	"errors"
	"go-webapp/config"
	"go-webapp/models"
	"net/http"
	"net/url"
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
	userResp, _ := http.Get(config.FbGraphAPIURL + "?" + q.Encode())

	if userResp.StatusCode != http.StatusOK {
		return errors.New("Not able to fetch facebook profile. Status Code: " + string(userResp.StatusCode))
	}

	err := json.NewDecoder(userResp.Body).Decode(&fb)

	if err != nil {
		return err
	}

	return nil
}
