package models

type EdgeCommandPayload struct {
	Type           string   `json:"t"`              // type of the config payload (rtsp)
	Operation      string   `json:"op"`             // a (add), r(remove), u(update)
	Name           string   `json:"n"`              // name of the devi