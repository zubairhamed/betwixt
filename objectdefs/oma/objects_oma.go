package oma

import (
	. "github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/core/objects"
	. "github.com/zubairhamed/betwixt/core/resources"
	"github.com/zubairhamed/betwixt/core/values/validators"
	. "github.com/zubairhamed/go-commons/typeval"
)

const (
	OBJECT_LWM2M_SECURITY                LWM2MObjectType = 0
	OBJECT_LWM2M_SERVER                  LWM2MObjectType = 1
	OBJECT_LWM2M_ACCESS_CONTROL          LWM2MObjectType = 2
	OBJECT_LWM2M_DEVICE                  LWM2MObjectType = 3
	OBJECT_LWM2M_CONNECTIVITY_MONITORING LWM2MObjectType = 4
	OBJECT_LWM2M_FIRMWARE_UPDATE         LWM2MObjectType = 5
	OBJECT_LWM2M_LOCATION                LWM2MObjectType = 6
	OBJECT_LWM2M_CONNECTIVITY_STATISTICS LWM2MObjectType = 7
)

type LWM2MCoreObjects struct {
	models map[LWM2MObjectType]ObjectDefinition
}

func (o *LWM2MCoreObjects) Initialize() {
	o.models = make(map[LWM2MObjectType]ObjectDefinition)

	o.AddObject(
		&DefaultObjectDefinition{Name: "LWM2M Security", Id: 0, Multiple: true, Mandatory: true,},
		&DefaultResourceDefinition{Id: 0, Name: "LWM2M  Server URI", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "0-255 bytes", ValueValidator: validators.NewRangeValidator(0, 255)},
		&DefaultResourceDefinition{Id: 1, Name: "Bootstrap Server", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN},
		&DefaultResourceDefinition{Id: 2, Name: "Security Mode", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-3", ValueValidator: validators.NewRangeValidator(0, 3)},
		&DefaultResourceDefinition{Id: 3, Name: "Public Key or Identity", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE},
		&DefaultResourceDefinition{Id: 4, Name: "Server Public Key or Identity", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE},
		&DefaultResourceDefinition{Id: 5, Name: "Secret Key", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE},
		&DefaultResourceDefinition{Id: 6, Name: "SMS Security Mode", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-255", ValueValidator: validators.NewRangeValidator(0, 255)},
		&DefaultResourceDefinition{Id: 7, Name: "SMS Binding Key Parameters", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "6 bytes", ValueValidator: validators.NewLengthValidator(6)},
		&DefaultResourceDefinition{Id: 8, Name: "SMS Binding Secret Keys", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "32-48 bytes", ValueValidator: validators.NewRangeValidator(32, 48)},
		&DefaultResourceDefinition{Id: 9, Name: "LWM2M Server SMS Number", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 10, Name: "Short Server ID", Operations: OPERATION_NONE, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "1-65535", ValueValidator: validators.NewRangeValidator(1, 65535)},
		&DefaultResourceDefinition{Id: 11, Name: "Client Hold Off Time", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, Units: "s",},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "LWM2M Server", Id: 1, Multiple: true, Mandatory: true,},
		&DefaultResourceDefinition{Id: 0, Name: "Short Server ID", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "1-65535", ValueValidator: validators.NewRangeValidator(1, 65535)},
		&DefaultResourceDefinition{Id: 1, Name: "Lifetime", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, Units: "s",},
		&DefaultResourceDefinition{Id: 2, Name: "Default Minimum Period", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "s"},
		&DefaultResourceDefinition{Id: 3, Name: "Default Maximum Period", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "s"},
		&DefaultResourceDefinition{Id: 4, Name: "Disable", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 5, Name: "Disable Timeout", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "s"},
		&DefaultResourceDefinition{Id: 6, Name: "Notification Storing When Disabled or Offline", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN},
		&DefaultResourceDefinition{Id: 7, Name: "Binding", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "The possible values of Resource are listed in 5.2.1.1"},
		&DefaultResourceDefinition{Id: 8, Name: "Registration Update Trigger", Operations: OPERATION_E, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "LWM2M Access Control", Id: 2, Multiple: true, Mandatory: false,},
		&DefaultResourceDefinition{Id: 0, Name: "Object ID", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "1-65534", ValueValidator: validators.NewRangeValidator(1, 65534)},
		&DefaultResourceDefinition{Id: 1, Name: "Object Instance ID", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-65535", ValueValidator: validators.NewRangeValidator(0, 65535)},
		&DefaultResourceDefinition{Id: 2, Name: "ACL", Operations: OPERATION_RW, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "16-bit", ValueValidator: validators.NewRangeValidator(-32768, 32767)},
		&DefaultResourceDefinition{Id: 3, Name: "Access Control Owner", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-65535", ValueValidator: validators.NewRangeValidator(0, 65535)},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "Device", Id: 3, Multiple: false, Mandatory: true,},
		&DefaultResourceDefinition{Id: 0, Name: "Manufacturer", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 1, Name: "Model Number", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 2, Name: "Serial Number", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 3, Name: "Firmware Version", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 4, Name: "Reboot", Operations: OPERATION_E, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 5, Name: "Factory Reset", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 6, Name: "Available Power Sources", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-7", ValueValidator: validators.NewRangeValidator(0, 7)},
		&DefaultResourceDefinition{Id: 7, Name: "Power Source Voltage", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "mV"},
		&DefaultResourceDefinition{Id: 8, Name: "Power Source Current", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "mA"},
		&DefaultResourceDefinition{Id: 9, Name: "Battery Level", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", ValueValidator: validators.NewRangeValidator(0, 100)},
		&DefaultResourceDefinition{Id: 10, Name: "Memory Free", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "KB"},
		&DefaultResourceDefinition{Id: 11, Name: "Error Code", Operations: OPERATION_R, Multiple: true, Mandatory: true, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 12, Name: "Reset Error Code", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 13, Name: "Current Time", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_TIME},
		&DefaultResourceDefinition{Id: 14, Name: "UTC Offset", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 15, Name: "Timezone", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 16, Name: "Supported Binding and Modes", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "Connectivity Monitoring", Id: 4, Multiple: false, Mandatory: false,},
		&DefaultResourceDefinition{Id: 0, Name: "Network Bearer", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 1, Name: "Available Network Bearer", Operations: OPERATION_R, Multiple: true, Mandatory: true, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 2, Name: "Radio Signal Strength", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, Units: "dBm"},
		&DefaultResourceDefinition{Id: 3, Name: "Link Quality", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 4, Name: "IP Addresses", Operations: OPERATION_R, Multiple: true, Mandatory: true, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 5, Name: "Router IP Addresse", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 6, Name: "Link Utilization", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", ValueValidator: validators.NewRangeValidator(0, 100)},
		&DefaultResourceDefinition{Id: 7, Name: "APN", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 8, Name: "Cell ID", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 9, Name: "SMNC", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "%"},
		&DefaultResourceDefinition{Id: 10, Name: "SMCC", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "Firmware Update", Id: 5, Multiple: false, Mandatory: false,},
		&DefaultResourceDefinition{Id: 0, Name: "Package", Operations: OPERATION_W, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_OPAQUE},
		&DefaultResourceDefinition{Id: 1, Name: "Package URI", Operations: OPERATION_W, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "0-255 bytes", ValueValidator: validators.NewRangeValidator(0, 255)},
		&DefaultResourceDefinition{Id: 2, Name: "Update", Operations: OPERATION_E, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING},
		&DefaultResourceDefinition{Id: 3, Name: "State", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "1-3", ValueValidator: validators.NewRangeValidator(1, 3)},
		&DefaultResourceDefinition{Id: 4, Name: "Update Supported Objects", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_BOOLEAN},
		&DefaultResourceDefinition{Id: 5, Name: "Update Result", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-6", ValueValidator: validators.NewRangeValidator(0, 6)},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "Location", Id: 6, Multiple: false, Mandatory: false,},
		&DefaultResourceDefinition{Id: 0, Name: "Latitude", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, Units: "Deg"},
		&DefaultResourceDefinition{Id: 1, Name: "Longitude", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, Units: "Deg"},
		&DefaultResourceDefinition{Id: 2, Name: "Altitude", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, Units: "m"},
		&DefaultResourceDefinition{Id: 3, Name: "Uncertainty", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, Units: "m"},
		&DefaultResourceDefinition{Id: 4, Name: "Velocity", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, Units: "Refers to 3GPP GAD specs"},
		&DefaultResourceDefinition{Id: 5, Name: "Timestamp", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_TIME, RangeOrEnums: "0-6", ValueValidator: validators.NewRangeValidator(0, 6)},
	)

	o.AddObject(
		&DefaultObjectDefinition{Name: "Connectivity Statistics", Id: 7, Multiple: false, Mandatory: false},
		&DefaultResourceDefinition{Id: 0, Name: "SMS Tx Counter", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 1, Name: "SMS Rx Counter", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER},
		&DefaultResourceDefinition{Id: 2, Name: "Tx Data", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "Kilo-Bytes"},
		&DefaultResourceDefinition{Id: 3, Name: "Rx Data", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "Kilo-Bytes"},
		&DefaultResourceDefinition{Id: 4, Name: "Max Message Size", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "Byte"},
		&DefaultResourceDefinition{Id: 5, Name: "Average Message Size", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, Units: "Byte"},
		&DefaultResourceDefinition{Id: 6, Name: "StartOrReset", Operations: OPERATION_E, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING},
	)

	/*
		Lock and Wipe

		Software Update

		Cellular connectivity

		APN connection profile

		WLAN connectivity

		Bearer selection

		DevCapMgmt
	*/
}

func (o *LWM2MCoreObjects) GetObject(n LWM2MObjectType) ObjectDefinition {
	return o.models[n]
}

func (o *LWM2MCoreObjects) GetObjects() map[LWM2MObjectType]ObjectDefinition {
	return o.models
}

func (o *LWM2MCoreObjects) AddObject(m ObjectDefinition, res ...ResourceDefinition) {
	m.SetResources(res)
	o.models[m.GetType()] = m
}
