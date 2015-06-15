package ipso

import (
	. "github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/core/objects"
	. "github.com/zubairhamed/betwixt/core/resources"
	. "github.com/zubairhamed/go-commons/typeval"
)

const (
	OBJECT_IPSO_DIGITAL_INPUT     LWM2MObjectType = 3200
	OBJECT_IPSO_DIGITAL_OUTPUT    LWM2MObjectType = 3201
	OBJECT_IPSO_ANALOG_INPUT      LWM2MObjectType = 3202
	OBJECT_IPSO_ANALOG_OUTPUT     LWM2MObjectType = 3203
	OBJECT_IPSO_GENERIC_SENSOR    LWM2MObjectType = 3300
	OBJECT_IPSO_ILLUMINANCE       LWM2MObjectType = 3301
	OBJECT_IPSO_PRESENCE          LWM2MObjectType = 3302
	OBJECT_IPSO_TEMPERATURE       LWM2MObjectType = 3303
	OBJECT_IPSO_HUMIDITY          LWM2MObjectType = 3304
	OBJECT_IPSO_POWER_MEASUREMENT LWM2MObjectType = 3305
	OBJECT_IPSO_ACTUATION         LWM2MObjectType = 3306
	OBJECT_IPSO_SET_POINT         LWM2MObjectType = 3308
	OBJECT_IPSO_LOAD_CONTROL      LWM2MObjectType = 3310
	OBJECT_IPSO_LIGHT_CONTROL     LWM2MObjectType = 3311
	OBJECT_IPSO_POWER_CONTROL     LWM2MObjectType = 3312
	OBJECT_IPSO_ACCELEROMETER     LWM2MObjectType = 3313
	OBJECT_IPSO_MAGNETOMETER      LWM2MObjectType = 3314
	OBJECT_IPSO_BAROMETER         LWM2MObjectType = 3315
)

type IPSOSmartObjects struct {
	models map[LWM2MObjectType]ObjectDefinition
}

func (o *IPSOSmartObjects) Initialize() {
	o.models = make(map[LWM2MObjectType]ObjectDefinition)

	// This IPSO object is a generic object that can be used with any kind of digital input interface.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Digital Input", Id: 3200, Multiple: true, Mandatory: false, Description: "Generic digital input for non-specific sensors"},
		&DefaultResourceDefinition{Id: 5500, Name: "Digital Input State", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "The current state of a digital input."},
		&DefaultResourceDefinition{Id: 5501, Name: "Digital Input Counter", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "", Description: "The cumulative value of active state detected.  "},
		&DefaultResourceDefinition{Id: 5502, Name: "Digital Input Polarity", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "The polarity of the digital input as a Boolean (0 \u003d Normal, 1\u003d Reversed)"},
		&DefaultResourceDefinition{Id: 5503, Name: "Digital Input Debounce", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "ms", Description: "The debounce period in ms. ."},
		&DefaultResourceDefinition{Id: 5504, Name: "Digital Input Edge Selection", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "1-3", Units: "", Description: "The edge selection as an integer (1 \u003d Falling edge, 2 \u003d Rising edge, 3 \u003d Both Rising and Falling edge)."},
		&DefaultResourceDefinition{Id: 5505, Name: "Digital Input Counter Reset", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Counter value."},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The application type of the sensor or actuator as a string, for instance, “Air Pressure”"},
		&DefaultResourceDefinition{Id: 5751, Name: "Sensor Type", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The type of the sensor (for instance PIR type)"},
	)

	// This IPSO object is a generic object that can be used with any kind of digital output interface.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Digital Output", Id: 3201, Multiple: true, Mandatory: false, Description: "Generic digital output for non-specific actuators"},
		&DefaultResourceDefinition{Id: 5550, Name: "Digital Output State", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "The current state of a digital output."},
		&DefaultResourceDefinition{Id: 5551, Name: "Digital Output Polarity", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "The polarity of a digital ouput as a Boolean (0 \u003d Normal, 1\u003d Reversed)."},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The application type of the output as a string, for instance, “LED”"},
	)

	// This IPSO object is a generic object that can be used with any kind of analog input interface.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Analog Input", Id: 3202, Multiple: true, Mandatory: false, Description: "Generic analog input for non-specific sensors "},
		&DefaultResourceDefinition{Id: 5600, Name: "Analog Input Current Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "0-1", Units: "", Description: "The current value of the analog input."},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The application type of the sensor or actuator as a string, for instance, “Air Pressure”"},
		&DefaultResourceDefinition{Id: 5751, Name: "Sensor Type", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The type of the sensor, for instance PIR type"},
	)

	// This IPSO object is a generic object that can be used with any kind of analog output interface.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Analog Output", Id: 3203, Multiple: true, Mandatory: false, Description: "This IPSO object is a generic object that can be used with any kind of analog output interface."},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5650, Name: "Analog Output Current Value", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "0-1", Units: "", Description: "The current state of the analogue output."},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the application type of the actuator as a string, for instance, “Thermostat”"},
	)

	// This IPSO object allows the description of a generic sensor. It is based on the description of a value
	// and measurement units according to the UCUM specification. Thus, any type of value defined within
	// the UCUM specification can be reported using this object.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Generic Sensor", Id: 3300, Multiple: true, Mandatory: false, Description: "This IPSO object allow the description of a generic sensor. It is based on the description of a value and a unit according to the UCUM specification. Thus, any type of value defined within this specification can be reporting using this object. \nSpecific object for a given range of sensors is described later in the document, enabling to identify the type of sensors directly from its Object ID. This object may be used as a generic object if a dedicated one does not exist. \n"},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5700, Name: "Sensor Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "Last or Current Measured Value from the Sensor"},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the application type of the sensor as a string, for instance, “CO2”"},
		&DefaultResourceDefinition{Id: 5751, Name: "Sensor Type", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The type of the sensor (for instance PIR type)"},
	)

	// This IPSO object should be used with an illuminance (light intensity) sensor to report an illuminance
	// measurement. It also provides resources for minimum/maximum measured values and the minimum/maximum
	// range that can be measured by the sensor. An example measurement unit is Lux (ucum:lx).
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Illuminance", Id: 3301, Multiple: true, Mandatory: false, Description: "Illuminance sensor, example units \u003d lx"},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5700, Name: "Sensor Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The current value of the luminosity sensor."},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
	)

	// This IPSO object should be used with a presence sensor to report presence detection. It also provides
	// resources to manage a counter, the type of sensor used (e.g the technology of the probe), and
	// configuration for the delay between busy and clear detection state.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Presence", Id: 3302, Multiple: true, Mandatory: false, Description: "Presence sensor with digital sensing, optional delay parameters"},
		&DefaultResourceDefinition{Id: 5500, Name: "Digital Input State", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "The current state of the presence sensor"},
		&DefaultResourceDefinition{Id: 5501, Name: "Digital Input Counter", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "", Description: "The cumulative value of active state detected."},
		&DefaultResourceDefinition{Id: 5505, Name: "Digital Input Counter Reset", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Counter value"},
		&DefaultResourceDefinition{Id: 5751, Name: "Sensor Type", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The type of the sensor (for instance PIR type)"},
		&DefaultResourceDefinition{Id: 5903, Name: "Busy to Clear delay", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "ms", Description: "Delay from the detection state to the clear state in ms"},
		&DefaultResourceDefinition{Id: 5904, Name: "Clear to Busy delay", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "ms", Description: "Delay from the clear state to the busy state in ms"},
	)

	// This IPSO object should be used with a temperature sensor to report a temperature measurement.
	// It also provides resources for minimum/maximum measured values and the minimum/maximum range that
	// can be measured by the temperature sensor. An example measurement unit is degrees Celsius (ucum:Cel).
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Temperature", Id: 3303, Multiple: true, Mandatory: false, Description: "Description: This IPSO object should be used with a temperature sensor to report a temperature measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the temperature sensor. An example measurement unit is degrees Celsius (ucum:Cel). "},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5700, Name: "Sensor Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "Last or Current Measured Value from the Sensor"},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
	)

	// This IPSO object should be used with a humidity sensor to report a humidity measurement. It also
	// provides resources for minimum/maximum measured values and the minimum/maximum range that can be
	// measured by the humidity sensor. An example measurement unit is relative humidity as a percentage
	// (ucum:%).
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Humidity", Id: 3304, Multiple: true, Mandatory: false, Description: "Description: This IPSO object should be used with a humidity sensor to report a humidity measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the humidity sensor. An example measurement unit is relative humidity as a percentage (ucum:%). "},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5700, Name: "Sensor Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "Last or Current Measured Value from the Sensor"},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
	)

	// This IPSO object should be used with a power measurement sensor to report a power measurement.
	// It also provides resources for minimum/maximum measured values and the minimum/maximum range for
	// both active and reactive power. Il also provides resources for cumulative energy, calibration, and
	// the power factor.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Power Measurement", Id: 3305, Multiple: true, Mandatory: false, Description: "This IPSO object should be used over a power measurement sensor to report a remote power measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range for both active and reactive power. Il also provides resources for cumulative energy, calibration, and the power factor. "},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5800, Name: "Instantaneous active power", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "The current active power"},
		&DefaultResourceDefinition{Id: 5801, Name: "Min Measured active power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "The minimum active power measured by the sensor since it is ON"},
		&DefaultResourceDefinition{Id: 5802, Name: "Max Measured  active power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "The maximum active power measured by the sensor since it is ON"},
		&DefaultResourceDefinition{Id: 5803, Name: "Min  Range  active power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "The minimum active power that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5804, Name: "Max Range active power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "The maximum active power that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5805, Name: "Cumulative active power ", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Wh", Description: "The cumulative active power since the last cumulative energy reset or device start"},
		&DefaultResourceDefinition{Id: 5806, Name: "Active Power Calibration", Operations: OPERATION_W, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "W", Description: "Request an active power calibration by writing the value of a calibrated load. "},
		&DefaultResourceDefinition{Id: 5810, Name: "Instantaneous reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "The current reactive power"},
		&DefaultResourceDefinition{Id: 5811, Name: "Min Measured reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "The minimum reactive power measured by the sensor since it is ON"},
		&DefaultResourceDefinition{Id: 5812, Name: "Max Measured  reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "The maximum reactive power measured by the sensor since it is ON"},
		&DefaultResourceDefinition{Id: 5813, Name: "Min  Range  reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "The minimum active power that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5814, Name: "Max Range reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "The maximum reactive power that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5815, Name: "Cumulative reactive power", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VARh", Description: "The cumulative reactive power since the last cumulative energy reset or device start"},
		&DefaultResourceDefinition{Id: 5816, Name: "Reactive Power Calibration", Operations: OPERATION_W, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "VAR", Description: "Request a reactive power calibration by writing the value of a calibrated load."},
		&DefaultResourceDefinition{Id: 5820, Name: "Power factor", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "If applicable, the power factor of the current consumption."},
		&DefaultResourceDefinition{Id: 5821, Name: "Current Calibration", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "Read or Write the current calibration coefficient"},
		&DefaultResourceDefinition{Id: 5822, Name: "Reset Cumulative energy", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset both cumulative active/reactive power"},
	)

	// This IPSO object is dedicated to remote actuation such as ON/OFF action or dimming. A multi-state
	// output can also be described as a string. This is useful to send pilot wire orders for instance. It
	// also provides a resource to reflect the time that the device has been switched on.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Actuation", Id: 3306, Multiple: true, Mandatory: false, Description: "This IPSO object is dedicated to remote actuation such as ON/OFF action or dimming. A multi-state output can also be described as a string. This is useful to send pilot wire orders for instance. It also provides a resource to reflect the time that the device has been switched on. "},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The Application type of the device, for example “Motion Closure”."},
		&DefaultResourceDefinition{Id: 5850, Name: "On/Off", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "On/Off\n\n"},
		&DefaultResourceDefinition{Id: 5851, Name: "Dimmer", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", Description: "This resource represents a light dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
		&DefaultResourceDefinition{Id: 5852, Name: "On Time", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "sec", Description: "The time in seconds that the device has been on. Writing a value of 0 resets the counter."},
		&DefaultResourceDefinition{Id: 5853, Name: "Muti-state Output", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "A string describing a state for multiple level output such as Pilot Wire"},
	)

	// This IPSO object should be used to set a desired value to a controller, such as a thermostat. This
	// object enables a setpoint to be expressed units defined in the UCUM specification, to match an
	// associated sensor or measurement value. A special resource is added to set the colour of an object.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Set Point", Id: 3308, Multiple: true, Mandatory: false, Description: "Description: This IPSO object should be used to set a desired value to a controller, such as a thermostat. This object enables a setpoint to be expressed units defined in the UCUM specification, to match an associated sensor or measurement value. A special resource is added to set the colour of an object."},
		&DefaultResourceDefinition{Id: 5701, Name: "Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
		&DefaultResourceDefinition{Id: 5706, Name: "Colour", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "A string representing a value in some color space"},
		&DefaultResourceDefinition{Id: 5750, Name: "Application Type", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The Application type of the device, for example “Motion Closure”."},
		&DefaultResourceDefinition{Id: 5900, Name: "Set Point Value", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.\t", Description: "The setpoint value. "},
	)

	// This Object is used for demand-response load control and other load control in automation applications
	// (not limited to power).
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Load Control", Id: 3310, Multiple: true, Mandatory: false, Description: "Description: This Object is used for demand-response load control and other load control in automation application (not limited to power)."},
		&DefaultResourceDefinition{Id: 5823, Name: "Event Identifier", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The event identifier as a string."},
		&DefaultResourceDefinition{Id: 5824, Name: "Start Time", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Time when the load control event will start started."},
		&DefaultResourceDefinition{Id: 5825, Name: "Duration In Min", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "Min", Description: "The duration of the load control event."},
		&DefaultResourceDefinition{Id: 5826, Name: "Criticality Level", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "The criticality of the event.  The device receiving the event will react in an appropriate fashion for the device."},
		&DefaultResourceDefinition{Id: 5827, Name: "Avg Load Adj Pct", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "0-100", Units: "%", Description: "Defines the maximum energy usage of the receivng device, as a percentage of the device\u0027s normal maximum energy usage."},
		&DefaultResourceDefinition{Id: 5828, Name: "Duty Cycle", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "0-100", Units: "%", Description: "Defines the duty cycle for the load control event, i.e, what percentage of time the receiving device is allowed to be on."},
	)

	//  This Object is used to control a light source, such as a LED or other light. It allows a light to
	// be turned on or off and its dimmer setting to be control as a % between 0 and 100. An optional colour
	// setting enables a string to be used to indicate the desired colour.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Light Control", Id: 3311, Multiple: true, Mandatory: false, Description: "Description: This Object is used to control a light source, such as a LED or other light.  It allows a light to be turned on or off and its dimmer setting to be control as a % between 0 and 100. An optional colour setting enables a string to be used to indicate the desired colour."},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
		&DefaultResourceDefinition{Id: 5706, Name: "Colour", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "A string representing a value in some color space"},
		&DefaultResourceDefinition{Id: 5805, Name: "Cumulative active power ", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Wh", Description: "The total power in Wh that the light has used."},
		&DefaultResourceDefinition{Id: 5820, Name: "Power factor", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The power factor of the light."},
		&DefaultResourceDefinition{Id: 5850, Name: "On/Off", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "This resource represents a light, which can be controlled, the setting of which is a Boolean value (1,0) where 1 is on and 0 is off."},
		&DefaultResourceDefinition{Id: 5851, Name: "Dimmer", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", Description: "This resource represents a light dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
		&DefaultResourceDefinition{Id: 5852, Name: "On Time", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "sec", Description: "The time in seconds that the light has been on. Writing a value of 0 resets the counter."},
	)

	// This Object is used to control a power source, such as a Smart Plug. It allows a power relay to be
	// turned on or off and its dimmer setting to be control as a % between 0 and 100.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Power Control", Id: 3312, Multiple: true, Mandatory: false, Description: "Description: This Object is used to control a power source, such as a Smart Plug.  It allows a power relay to be turned on or off and its dimmer setting to be control as a % between 0 and 100."},
		&DefaultResourceDefinition{Id: 5805, Name: "Cumulative active power ", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Wh", Description: "The total power in Wh that has been used by the load."},
		&DefaultResourceDefinition{Id: 5820, Name: "Power factor", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "", Description: "The power factor of the load."},
		&DefaultResourceDefinition{Id: 5850, Name: "On/Off", Operations: OPERATION_RW, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_BOOLEAN, RangeOrEnums: "", Units: "", Description: "This resource represents a power relay, which can be controlled, the setting of which is a Boolean value (1,0) where 1 is on and 0 is off."},
		&DefaultResourceDefinition{Id: 5851, Name: "Dimmer", Operations: OPERATION_W, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", Description: "This resource represents a power dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
		&DefaultResourceDefinition{Id: 5852, Name: "On Time", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_INTEGER, RangeOrEnums: "", Units: "sec", Description: "The time in seconds that the power relay has been on. Writing a value of 0 resets the counter."},
	)

	// This IPSO object can be used to represent a 1-3 axis accelerometer.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Accelerometer", Id: 3313, Multiple: true, Mandatory: false, Description: "Description: This IPSO object can be used to represent a 1-3 axis accelerometer."},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
		&DefaultResourceDefinition{Id: 5702, Name: "X Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the X axis."},
		&DefaultResourceDefinition{Id: 5703, Name: "Y Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the Y axis."},
		&DefaultResourceDefinition{Id: 5704, Name: "Z Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the Z axis."},
	)

	// This IPSO object can be used to represent a 1-3 axis magnetometer with optional compass direction.
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Magnetometer", Id: 3314, Multiple: true, Mandatory: false, Description: "Description: This IPSO object can be used to represent a 1-3 axis magnetometer with optional compass direction."},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
		&DefaultResourceDefinition{Id: 5702, Name: "X Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the X axis."},
		&DefaultResourceDefinition{Id: 5703, Name: "Y Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the Y axis."},
		&DefaultResourceDefinition{Id: 5704, Name: "Z Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The measured value along the Z axis."},
		&DefaultResourceDefinition{Id: 5705, Name: "Compass Direction", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "0-360", Units: "Deg", Description: "The compass direction"},
	)

	// This IPSO object should be used with an air pressure sensor to report a barometer measurement.
	// It also provides resources for minimum/maximum measured values and the minimum/maximum range that
	// can be measured by the barometer sensor. An example measurement unit is kPa (ucum:kPa).
	o.AddObject(
		&DefaultObjectDefinition{Name: "IPSO Barometer", Id: 3315, Multiple: true, Mandatory: false, Description: "Description: This IPSO object should be used with an air pressure sensor to report a barometer measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the barometer sensor. An example measurement unit is kPa (ucum:kPa)."},
		&DefaultResourceDefinition{Id: 5601, Name: "Min Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5602, Name: "Max Measured Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value measured by the sensor since power ON or reset"},
		&DefaultResourceDefinition{Id: 5603, Name: "Min Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The minimum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5604, Name: "Max Range Value", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "The maximum value that can be measured by the sensor"},
		&DefaultResourceDefinition{Id: 5605, Name: "Reset Min and Max Measured Values", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_OPAQUE, RangeOrEnums: "", Units: "", Description: "Reset the Min and Max Measured Values to Current Value"},
		&DefaultResourceDefinition{Id: 5700, Name: "Sensor Value", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: VALUETYPE_FLOAT, RangeOrEnums: "", Units: "Defined by “Units” resource.", Description: "Last or Current Measured Value from the Sensor"},
		&DefaultResourceDefinition{Id: 5701, Name: "Sensor Units", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: VALUETYPE_STRING, RangeOrEnums: "", Units: "", Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
	)
}

func (o *IPSOSmartObjects) GetObject(n LWM2MObjectType) ObjectDefinition {
	return o.models[n]
}

func (o *IPSOSmartObjects) GetObjects() map[LWM2MObjectType]ObjectDefinition {
	return o.models
}

func (o *IPSOSmartObjects) AddObject(m ObjectDefinition, res ...ResourceDefinition) {
	m.SetResources(res)
	o.models[m.GetType()] = m
}
