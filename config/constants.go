package config

import (
	"time"
)

// FbAccessTokenURL - Facebook access token url
const FbAccessTokenURL = "https://graph.facebook.com/v2.7/oauth/access_token"

// FbGraphAPIURL - Facebook graph api url
const FbGraphAPIURL = "https://graph.facebook.com/v2.7"

// IgAccessTokenURL - Instagram access token url
const IgAccessTokenURL = "https://api.instagram.com/oauth/access_token"

// IgAPIURL - Instagram API url
const IgAPIURL = "https://api.instagram.com/v1"

// AuthHeaderName - Authentication Header Name
const AuthHeaderName = "X-Auth-Token"

// TokenExpiry - Set token expiration time to 2 days
const TokenExpiry = time.Hour * 48

// RefreshTokenExpiry - Set refresh token expiration time to 15 days
const RefreshTokenExpiry = time.Hour * 24 * 15

// Collection names
const UserColl = "users"
const PostColl = "posts"
