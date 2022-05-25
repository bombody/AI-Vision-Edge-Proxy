package models

type MQTTProcessOperation string
type MQTTProcessType string

const (
	DeviceOperationAdd      string = "add"           // add device
	DeviceOperationRemove   string = "remove"        // remove device
	DeviceOperationState    string = "state"   