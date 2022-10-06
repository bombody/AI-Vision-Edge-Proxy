package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	microCrypto "github.com/chryscloud/go-microkit-plugins/crypto"
	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/go-resty/resty/v2"
)

func CallAPIWithBody(apiClient *resty.Client, method string, fullEndpoint string, body interface{}, edgeKey, edgeSecret string) ([]byte, error) {

	payload, err := json.Marshal(body)
	if err != nil {
		g.Log.Error("failed to marshal body", err)
		return nil, err
	}

	h := md5.New()
	h.Write(payload)
	contentMD5 := hex.EncodeToString(h.Sum(nil))
	current_ts := strconv.FormatInt(time.Now().Unix()*1000, 10)
	signPayload := current_ts + contentMD5
	mac := microCrypto.ComputeHmac(sha256.New, signPayload, edgeSecret)

	req := apiClient.R().SetHeader("X-ChrysEdge-Auth", edgeKey+":"+mac).
		SetHeader("X-Chrys-Date", current_ts).
		SetHeader("Content-MD5", contentMD5).SetBody(body)
	resp, sndErr := req.Execute(method, fullEndpoint)

	if sndErr != nil {
		g.Log.Error("failed to send annotations to remote 