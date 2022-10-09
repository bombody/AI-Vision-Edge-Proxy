package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/dgrijalva/jwt-go"
	qtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v7"
)

// CreateJWT creates RS265 JWT signed token
func CreateJWT(projectID string, privateKeyBytes []byte, expiration time.Duration) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  projectID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	algorithm := "RS256"

	switch algorithm {
	case "RS256":
		privKey, pErr := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
		if pErr != nil {
			g.Log.Error("invalid private key", pErr)
			return "", pErr
		}
		return token.SignedString(privKey)
	case "ES256":
		privKey, _ := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
		return token.SignedString(privKey)
	}

	return "", errors.New("Cannot find JWT algorithm. Specify 'ES256' or 'RS256'")
}

// ParseJWTTokenExpirationTime (no validation parsing of the jwt token in string format)
func ParseJWTTokenExpirationTime(jwtToken string) (time.Time, error) {
	claims := jwt.MapClaims{}
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return time.Time{}, errors.New("Can't convert token's claims to standard claims")
	}
	var tm time.Time
	switch exp :