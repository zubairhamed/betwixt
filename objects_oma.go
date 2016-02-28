package betwixt

import "log"

const (
	POWERSOURCE_DC            = 0
	POWERSOURCE_INTERNAL      = 1
	POWERSOURCE_EXTERNAL      = 2
	POWERSOURCE_OVER_ETHERNET = 4
	POWERSOURCE_USB           = 5
	POWERSOURCE_AC            = 6
	POWERSOURCE_SOLAR         = 7
)

const (
	SERVER_EXEC_DISABLE            = 4
	SERVER_EXEC_REG_UPDATE_TRIGGER = 8

	DEVICE_EXEC_REBOOT           = 4
	DEVICE_EXEC_FACTORY_RESET    = 5
	DEVICE_EXEC_RESET_ERROR_CODE = 12

	FIRMWARE_EXEC_UPDATE = 2

	CONNECTIVITY_STATS_EXEC_STARTRESET = 6
)

const (
	SECURITYMODE_PRESHAREDKEY = 0
	SECURITYMODE_RAWPK        = 1
	SECURITYMODE_CERTIFICATE  = 2
	SECURITYMODE_NOSEC        = 3
)

const (
	BATTERYSTATUS_NORMAL          = 0
	BATTERYSTATUS_CHARGING        = 1
	BATTERYSTATUS_CHARGE_COMPLETE = 2
	BATTERYSTATUS_DAMAGED         = 3
	BATTERYSTATUS_LOW_BATTERY     = 4
	BATTERYSTATUS_NOT_INSTALLED   = 5
	BATTERYSTATUS_UNKNOWN         = 6
)

const (
	ERRORCODE_NO_ERROR                     = 0
	ERRORCODE_LOW_BATTERY_POWER            = 1
	ERRORCODE_EXTERNAL_POWER_SUPPLY_OFF    = 2
	ERRORCODE_GPS_MODULE_FAILURE           = 3
	ERRORCODE_LOW_RECEIVED_SIGNAL_STRENGTH = 4
	ERRORCODE_OUT_OF_MEMORY                = 5
	ERRORCODE_SMS_FAILURE                  = 6
	ERRORCODE_IP_CONNECTIVITY_FAILURE      = 7
	ERRORCODE_PERIPHERAL_MALFUNCTION       = 8
)

const (
	FWUPDATE_STATE_IDLE        = 1
	FWUPDATE_STATE_DOWNLOADING = 2
	FWUPDATE_STATE_DOWNLOADED  = 3

	FWUPDATE_RESULT_DEFAULT                  = 0
	FWUPDATE_RESULT_SUCCESSFUL               = 1
	FWUPDATE_RESULT_NOT_ENOUGH_STORAGE       = 2
	FWUPDATE_RESULT_OUT_OF_MEMORY            = 3
	FWUPDATE_RESULT_CONNECTION_LOST          = 4
	FWUPDATE_RESULT_CRC_CHECK                = 5
	FWUPDATE_RESULT_UNSUPPORTED_PACKAGE_TYPE = 6
	FWUPDATE_RESULT_INVALID_URI              = 7
)

const (
	OMA_OBJECT_LWM2M_SECURITY                LWM2MObjectType = 0
	OMA_OBJECT_LWM2M_SERVER                  LWM2MObjectType = 1
	OMA_OBJECT_LWM2M_ACCESS_CONTROL          LWM2MObjectType = 2
	OMA_OBJECT_LWM2M_DEVICE                  LWM2MObjectType = 3
	OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING LWM2MObjectType = 4
	OMA_OBJECT_LWM2M_FIRMWARE_UPDATE         LWM2MObjectType = 5
	OMA_OBJECT_LWM2M_LOCATION                LWM2MObjectType = 6
	OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS LWM2MObjectType = 7
)

// This represents a list of LWM2M Starter Pack objects registered to the OMA NA
type LWM2MCoreObjects struct {
	models map[LWM2MObjectType]ObjectDefinition
}

func (o *LWM2MCoreObjects) Initialize() {
	o.models = make(map[LWM2MObjectType]ObjectDefinition)

	data, err := Asset("objdefs/oma.json")
	if err != nil {
		log.Println("Error parsing oma.json")
	}

	objdefs := ParseObjectDefinitionsJson(data)

	o.AddObjects(objdefs)

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

func (o *LWM2MCoreObjects) AddObject(m ObjectDefinition, res []ResourceDefinition) {
	m.SetResources(res)
	o.models[m.GetType()] = m
}

func (o *LWM2MCoreObjects) AddObjects(od []ObjectDefinition) {
	for _, v := range od {
		o.models[v.GetType()] = v
	}
}
