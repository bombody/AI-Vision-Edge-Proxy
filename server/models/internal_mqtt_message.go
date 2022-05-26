package models

type MQTTProcessOperation string
type MQTTProcessType string

const (
	DeviceOperationAdd      string = "add"           // add device
	DeviceOperationRemove   string = "remove"        // remove device
	DeviceOperationState    string = "state"         // the current state of device
	DeviceOperationStats    string = "stats"         // device stats (host system and each device)
	GatewayOperationCheckIn string = "gwcheckin"     // gateway checkin
	DeviceInternalTesting   string = "internal_test" // internal development event. Not used or required for regular operations (TODO: movw to unit tests)

	DeviceOper