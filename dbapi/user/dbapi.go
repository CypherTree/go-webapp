package userapi

import (
	"fmt"
	"go-webapp/config"
	"go-webapp/db"
	"go-webapp/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"gopkg.in/mgo.v2/bson"
)

// CustomClaims - Custom claims struct for jwt
type CustomClaims struct {
	UserRole int `json:"user_role"`
	jwt.StandardClaims
}

// GetByFbID - Fetch one user by facebook id
func GetByFbID(fbID string) (*models.User, error) {
	user := new(models.User)

	query := make(bson.M)
	query["facebook.id"] = fbID

	err := db.Conn.GetOne(query, config.UserColl, &user)

	return user, err
}

// GetByIgID - Fetch one user by instagram id
func GetByIgID(igID string) (*models.User, error) {
	user := new(models.User)

	query := make(bson.M)
	query["instagram.id"] = igID

	err := db.Conn.GetOne(query, config.UserColl, &user)

	return user, err
}

// Upsert - Insert or update user
func Upsert(user *models.User) error {
	user.BeforeSave()
	query := bson.M{
		"_id": user.ID,
	}

	err := db.Conn.Upsert(query, config.UserColl, &user)
	return err
}

// ValidateAuthToken - Validate auth token and return user obj
func ValidateAuthToken(tokenString string, session *models.Session) {
	if tokenString == "" || IsTokenBlacklisted(tokenString) {
		return
	}

	tokenSecretFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unknown signing method: %v", token.Header["alg"])
		}

		return config.Settings.AuthTokenSecret, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, tokenSecretFunc)
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(CustomClaims); ok && token.Valid {
		session = &models.Session{
			UserID:      bson.ObjectId(claims.Id),
			UserRole:    claims.UserRole,
			TokenString: tokenString,
			IssuedAt:    claims.IssuedAt,
		}
	}
}

// GenerateToken - Generate jwt token using user object
func GenerateToken(u models.User) string {
	signingKey := config.Settings.AuthTokenSecret
	claims := &CustomClaims{
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		panic(err)
	}

	return signedToken
}

// IsTokenBlacklisted - check if token is blacklisted in redis
func IsTokenBlacklisted(tokenString string) bool {
	val, _ := db.Redis.Get("Blacklisted:" + tokenString).Result()
	return val != ""
}

// BlacklistToken - Add token to blacklist in redis with expiry
func BlacklistToken(session *models.Session) {

	// expiry should be issued time + refresh expiry. This is to prevent anyone
	// from refreshing any blacklisted token.
	expiry := time.Unix(session.IssuedAt, 0).Add(config.RefreshTokenExpiry).Unix()

	key := "Blacklisted:" + session.TokenString
	err := db.Redis.Set(key, session.TokenString, time.Duration(expiry)).Err()

	if err != nil {
		panic(err)
	}
}
