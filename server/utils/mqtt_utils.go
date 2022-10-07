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
		Audience:  