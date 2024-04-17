package auth

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ary82/urlStash/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

// Custom jwt claims for urlstash
type CustomClaims struct {
	UserId int `json:"uid"`
	jwt.RegisteredClaims
}

// Useful data from Google's JWT
type GoogleData struct {
	Username string
	Name     string
	Email    string
	Picture  string
}

// Context key/val pair
type ContextKey string
type ContextVal struct {
	UserId int
	Email  string
}

// Accepts Google identity's JWT, parse and upsert user into db,
// Return a new generated Jwt for urlstash
func Login(tokenStr []byte, db database.DB) (*string, error) {

	// Get the token Payload from Google idToken JWT
	payload, err := GetPayload(tokenStr)
	if err != nil {
		return nil, err
	}

	// Extract useful data form payload
	data := GetData(payload)

	// Insert or Update User from Google's information
	err = db.UpsertUser(
		data.Username,
		data.Name,
		data.Email,
		data.Picture,
	)
	if err != nil {
		return nil, err
	}

	// Get the upserted User
	user, err := db.GetUserByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	// Generate the urlStash jwt with claims set as userId and email
	jwt, err := GenerateJWT(user.ID, data.Email)
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

// Generate Jwt to be used for urlstash
func GenerateJWT(UserId int, email string) (string, error) {
	claims := &CustomClaims{
		UserId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   email,
		},
	}

	key := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	return ss, err
}

func ParseJWT(tokenStr string) (*ContextVal, error) {

	// Parse claims to get the jwt token
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	// Check validity of claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("claims corrupted")
	}

	// Get email from claims
	email, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	// User struct for use in context
	userInfo := &ContextVal{
		UserId: claims.UserId,
		Email:  email,
	}
	return userInfo, nil
}

// Function for extracting payload from google identity's jwt
func GetPayload(tokenStr []byte) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(
		context.Background(),
		string(tokenStr),
		os.Getenv("AUDIENCE"),
	)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Extract data from google jwt's payload
func GetData(p *idtoken.Payload) *GoogleData {
	return &GoogleData{
		// Generate username from email for now
		// TODO: Let users pick
		Username: strings.Split(p.Claims["email"].(string), "@")[0],
		Email:    p.Claims["email"].(string),
		Name:     p.Claims["name"].(string),
		Picture:  p.Claims["picture"].(string),
	}
}
