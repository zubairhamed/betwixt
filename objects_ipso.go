package betwixt

import "log"

const (
	IPSO_OBJECT_IPSO_DIGITAL_INPUT            LWM2MObjectType = 3200
	IPSO_OBJECT_IPSO_DIGITAL_OUTPUT           LWM2MObjectType = 3201
	IPSO_OBJECT_IPSO_ANALOG_INPUT             LWM2MObjectType = 3202
	IPSO_OBJECT_IPSO_ANALOG_OUTPUT            LWM2MObjectType = 3203
	IPSO_OBJECT_IPSO_GENERIC_SENSOR           LWM2MObjectType = 3300
	IPSO_OBJECT_IPSO_ILLUMINANCE              LWM2MObjectType = 3301
	IPSO_OBJECT_IPSO_PRESENCE                 LWM2MObjectType = 3302
	IPSO_OBJECT_IPSO_TEMPERATURE              LWM2MObjectType = 3303
	IPSO_OBJECT_IPSO_HUMIDITY                 LWM2MObjectType = 3304
	IPSO_OBJECT_IPSO_POWER_MEASUREMENT        LWM2MObjectType = 3305
	IPSO_OBJECT_IPSO_ACTUATION                LWM2MObjectType = 3306
	IPSO_OBJECT_IPSO_SET_POINT                LWM2MObjectType = 3308
	IPSO_OBJECT_IPSO_LOAD_CONTROL             LWM2MObjectType = 3310
	IPSO_OBJECT_IPSO_LIGHT_CONTROL            LWM2MObjectType = 3311
	IPSO_OBJECT_IPSO_POWER_CONTROL            LWM2MObjectType = 3312
	IPSO_OBJECT_IPSO_ACCELEROMETER            LWM2MObjectType = 3313
	IPSO_OBJECT_IPSO_MAGNETOMETER             LWM2MObjectType = 3314
	IPSO_OBJECT_IPSO_BAROMETER                LWM2MObjectType = 3315
	IPSO_OBJECT_IPSO_VOLTAGE                  LWM2MObjectType = 3316
	IPSO_OBJECT_IPSO_CURRENT                  LWM2MObjectType = 3317
	IPSO_OBJECT_IPSO_FREQUENCY                LWM2MObjectType = 3318
	IPSO_OBJECT_IPSO_DEPTH                    LWM2MObjectType = 3319
	IPSO_OBJECT_IPSO_PERCENTAGE               LWM2MObjectType = 3320
	IPSO_OBJECT_IPSO_ALTITUDE                 LWM2MObjectType = 3321
	IPSO_OBJECT_IPSO_LOAD                     LWM2MObjectType = 3322
	IPSO_OBJECT_IPSO_PRESSURE                 LWM2MObjectType = 3323
	IPSO_OBJECT_IPSO_LOUDNESS                 LWM2MObjectType = 3324
	IPSO_OBJECT_IPSO_CONCENTRATION            LWM2MObjectType = 3325
	IPSO_OBJECT_IPSO_ACIDITY                  LWM2MObjectType = 3326
	IPSO_OBJECT_IPSO_CONDUCTIVITY             LWM2MObjectType = 3327
	IPSO_OBJECT_IPSO_POWER                    LWM2MObjectType = 3328
	IPSO_OBJECT_IPSO_POWER_FACTOR             LWM2MObjectType = 3329
	IPSO_OBJECT_IPSO_DISTANCE                 LWM2MObjectType = 3330
	IPSO_OBJECT_IPSO_ENERGY                   LWM2MObjectType = 3331
	IPSO_OBJECT_IPSO_DIRECTION                LWM2MObjectType = 3332
	IPSO_OBJECT_IPSO_TIME                     LWM2MObjectType = 3333
	IPSO_OBJECT_IPSO_GYROMETER                LWM2MObjectType = 3334
	IPSO_OBJECT_IPSO_COLOR                    LWM2MObjectType = 3335
	IPSO_OBJECT_IPSO_GPS_LOCATION             LWM2MObjectType = 3336
	IPSO_OBJECT_IPSO_POSITIONER               LWM2MObjectType = 3337
	IPSO_OBJECT_IPSO_BUZZER                   LWM2MObjectType = 3338
	IPSO_OBJECT_IPSO_AUDIO_CLIP               LWM2MObjectType = 3339
	IPSO_OBJECT_IPSO_TIMER                    LWM2MObjectType = 3340
	IPSO_OBJECT_IPSO_ADDRESSABLE_TEXT_DISPLAY LWM2MObjectType = 3341
	IPSO_OBJECT_IPSO_ONOFF_SWITCH             LWM2MObjectType = 3342
	IPSO_OBJECT_IPSO_LEVER_CONTROL            LWM2MObjectType = 3343
	IPSO_OBJECT_IPSO_UPDOWN_CONTROL           LWM2MObjectType = 3344
	IPSO_OBJECT_IPSO_MULTIPLE_AXIS_JOYSTICK   LWM2MObjectType = 3345
	IPSO_OBJECT_IPSO_RATE                     LWM2MObjectType = 3346
	IPSO_OBJECT_IPSO_PUSH_BUTTON              LWM2MObjectType = 3347
	IPSO_OBJECT_IPSO_MULTISTATE_SELECTOR      LWM2MObjectType = 3348
)

// This represents a list of LWM2M objects registered to the OMA NA for LWM2M donated by IPSO
type IPSOSmartObjects struct {
	models map[LWM2MObjectType]ObjectDefinition
}

func (o *IPSOSmartObjects) Initialize() {
	o.models = make(map[LWM2MObjectType]ObjectDefinition)

	data, err := Asset("objdefs/ipso.json")
	if err != nil {
		log.Println("Error parsing ipso.json")
	}

	objdefs := ParseObjectDefinitionsJson(data)

	o.AddObjects(objdefs)
}

func (o *IPSOSmartObjects) GetObject(n LWM2MObjectType) ObjectDefinition {
	return o.models[n]
}

func (o *IPSOSmartObjects) GetObjects() map[LWM2MObjectType]ObjectDefinition {
	return o.models
}

func (o *IPSOSmartObjects) AddObject(m ObjectDefinition, res []ResourceDefinition) {
	m.SetResources(res)
	o.models[m.GetType()] = m
}

func (o *IPSOSmartObjects) AddObjects(od []ObjectDefinition) {
	for _, v := range od {
		o.models[v.GetType()] = v
	}
}
