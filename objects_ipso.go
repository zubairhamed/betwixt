package lwm2m

const (
    OBJECT_IPSO_DIGITAL_INPUT       LWM2MObjectType = 3200
    OBJECT_IPSO_DIGITAL_OUTPUT      LWM2MObjectType = 3201
    OBJECT_IPSO_ANALOG_INPUT        LWM2MObjectType = 3202
    OBJECT_IPSO_ANALOG_OUTPUT       LWM2MObjectType = 3203
    OBJECT_IPSO_GENERIC_SENSOR      LWM2MObjectType = 3300
    OBJECT_IPSO_ILLUMINANCE         LWM2MObjectType = 3301
    OBJECT_IPSO_PRESENCE            LWM2MObjectType = 3302
    OBJECT_IPSO_TEMPERATURE         LWM2MObjectType = 3303
    OBJECT_IPSO_HUMIDITY            LWM2MObjectType = 3304
    OBJECT_IPSO_POWER_MEASUREMENT   LWM2MObjectType = 3305
    OBJECT_IPSO_ACTUATION           LWM2MObjectType = 3306
    OBJECT_IPSO_SET_POINT           LWM2MObjectType = 3308
    OBJECT_IPSO_LOAD_CONTROL        LWM2MObjectType = 3310
    OBJECT_IPSO_LIGHT_CONTROL       LWM2MObjectType = 3311
    OBJECT_IPSO_POWER_CONTROL       LWM2MObjectType = 3312
    OBJECT_IPSO_ACCELEROMETER       LWM2MObjectType = 3313
    OBJECT_IPSO_MAGNETOMETER        LWM2MObjectType = 3314
    OBJECT_IPSO_BAROMETER           LWM2MObjectType = 3315
)

type IPSOSmartObjects struct {
    models  map[LWM2MObjectType]*ObjectModel
}

func (o *IPSOSmartObjects) Initialize() {
    o.models = make(map[LWM2MObjectType]*ObjectModel)

    o.Add(
        &ObjectModel{ Name: "IPSO Digital Input", Id: 3200, Multiple: true, Mandatory: false, Description: "Generic digital input for non-specific sensors", },
        &ResourceModel{ Id: 5500,  Name: "Digital Input State",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "The current state of a digital input."},
        &ResourceModel{ Id: 5501,  Name: "Digital Input Counter",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "",  Description: "The cumulative value of active state detected.  "},
        &ResourceModel{ Id: 5502,  Name: "Digital Input Polarity",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "The polarity of the digital input as a Boolean (0 \u003d Normal, 1\u003d Reversed)"},
        &ResourceModel{ Id: 5503,  Name: "Digital Input Debounce",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "ms",  Description: "The debounce period in ms. ."},
        &ResourceModel{ Id: 5504,  Name: "Digital Input Edge Selection",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "1-3",  Units: "",  Description: "The edge selection as an integer (1 \u003d Falling edge, 2 \u003d Rising edge, 3 \u003d Both Rising and Falling edge)."},
        &ResourceModel{ Id: 5505,  Name: "Digital Input Counter Reset",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Counter value."},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The application type of the sensor or actuator as a string, for instance, “Air Pressure”"},
        &ResourceModel{ Id: 5751,  Name: "Sensor Type",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The type of the sensor (for instance PIR type)"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Digital Output", Id: 3201, Multiple: true, Mandatory: false, Description: "Generic digital output for non-specific actuators", },
        &ResourceModel{ Id: 5550,  Name: "Digital Output State",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "The current state of a digital output."},
        &ResourceModel{ Id: 5551,  Name: "Digital Output Polarity",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "The polarity of a digital ouput as a Boolean (0 \u003d Normal, 1\u003d Reversed)."},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The application type of the output as a string, for instance, “LED”"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Analog Input", Id: 3202, Multiple: true, Mandatory: false, Description: "Generic analog input for non-specific sensors ", },
        &ResourceModel{ Id: 5600,  Name: "Analog Input Current Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "0-1",  Units: "",  Description: "The current value of the analog input."},
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The application type of the sensor or actuator as a string, for instance, “Air Pressure”"},
        &ResourceModel{ Id: 5751,  Name: "Sensor Type",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The type of the sensor, for instance PIR type"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Analog Output", Id: 3203, Multiple: true, Mandatory: false, Description: "This IPSO object is a generic object that can be used with any kind of analog output interface.", },
        &ResourceModel{ Id: 5603,Name: "Min Range Value",Operations: OPERATION_R,Multiple: false,Mandatory: false,ResourceType: TYPE_FLOAT,RangeOrEnums: "",Units: "",Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,Name: "Max Range Value",Operations: OPERATION_R,Multiple: false,Mandatory: false,ResourceType: TYPE_FLOAT,RangeOrEnums: "",Units: "",Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5650,Name: "Analog Output Current Value",Operations: OPERATION_RW,Multiple: false,Mandatory: true,ResourceType: TYPE_FLOAT,RangeOrEnums: "0-1",Units: "",Description: "The current state of the analogue output."},
        &ResourceModel{ Id: 5750,Name: "Application Type",Operations: OPERATION_R,Multiple: false,Mandatory: false,ResourceType: TYPE_STRING,RangeOrEnums: "",Units: "",Description: "If present, the application type of the actuator as a string, for instance, “Thermostat”"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Generic Sensor", Id: 3300, Multiple: true, Mandatory: false, Description: "This IPSO object allow the description of a generic sensor. It is based on the description of a value and a unit according to the UCUM specification. Thus, any type of value defined within this specification can be reporting using this object. \nSpecific object for a given range of sensors is described later in the document, enabling to identify the type of sensors directly from its Object ID. This object may be used as a generic object if a dedicated one does not exist. \n", },
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5700,  Name: "Sensor Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "Last or Current Measured Value from the Sensor"},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "If present, the application type of the sensor as a string, for instance, “CO2”"},
        &ResourceModel{ Id: 5751,  Name: "Sensor Type",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The type of the sensor (for instance PIR type)"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Illuminance", Id: 3301, Multiple: true, Mandatory: false, Description: "Illuminance sensor, example units \u003d lx", },
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5700,  Name: "Sensor Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The current value of the luminosity sensor."},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Presence", Id: 3302, Multiple: true, Mandatory: false, Description: "Presence sensor with digital sensing, optional delay parameters", },
        &ResourceModel{ Id: 5500,  Name: "Digital Input State",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "The current state of the presence sensor"},
        &ResourceModel{ Id: 5501,  Name: "Digital Input Counter",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "",  Description: "The cumulative value of active state detected."},
        &ResourceModel{ Id: 5505,  Name: "Digital Input Counter Reset",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Counter value"},
        &ResourceModel{ Id: 5751,  Name: "Sensor Type",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The type of the sensor (for instance PIR type)"},
        &ResourceModel{ Id: 5903,  Name: "Busy to Clear delay",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "ms",  Description: "Delay from the detection state to the clear state in ms"},
        &ResourceModel{ Id: 5904,  Name: "Clear to Busy delay",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "ms",  Description: "Delay from the clear state to the busy state in ms"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Temperature",    Id: 3303,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object should be used with a temperature sensor to report a temperature measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the temperature sensor. An example measurement unit is degrees Celsius (ucum:Cel). "},
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5700,  Name: "Sensor Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "Last or Current Measured Value from the Sensor"},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Humidity",    Id: 3304,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object should be used with a humidity sensor to report a humidity measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the humidity sensor. An example measurement unit is relative humidity as a percentage (ucum:%). "},
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5700,  Name: "Sensor Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "Last or Current Measured Value from the Sensor"},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Power Measurement",    Id: 3305,    Multiple: true,    Mandatory: false,    Description: "This IPSO object should be used over a power measurement sensor to report a remote power measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range for both active and reactive power. Il also provides resources for cumulative energy, calibration, and the power factor. "},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5800,  Name: "Instantaneous active power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "The current active power"},
        &ResourceModel{ Id: 5801,  Name: "Min Measured active power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "The minimum active power measured by the sensor since it is ON"},
        &ResourceModel{ Id: 5802,  Name: "Max Measured  active power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "The maximum active power measured by the sensor since it is ON"},
        &ResourceModel{ Id: 5803,  Name: "Min  Range  active power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "The minimum active power that can be measured by the sensor"},
        &ResourceModel{ Id: 5804,  Name: "Max Range active power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "The maximum active power that can be measured by the sensor"},
        &ResourceModel{ Id: 5805,  Name: "Cumulative active power ",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Wh",  Description: "The cumulative active power since the last cumulative energy reset or device start"},
        &ResourceModel{ Id: 5806,  Name: "Active Power Calibration",  Operations: OPERATION_W,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "W",  Description: "Request an active power calibration by writing the value of a calibrated load. "},
        &ResourceModel{ Id: 5810,  Name: "Instantaneous reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "The current reactive power"},
        &ResourceModel{ Id: 5811,  Name: "Min Measured reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "The minimum reactive power measured by the sensor since it is ON"},
        &ResourceModel{ Id: 5812,  Name: "Max Measured  reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "The maximum reactive power measured by the sensor since it is ON"},
        &ResourceModel{ Id: 5813,  Name: "Min  Range  reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "The minimum active power that can be measured by the sensor"},
        &ResourceModel{ Id: 5814,  Name: "Max Range reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "The maximum reactive power that can be measured by the sensor"},
        &ResourceModel{ Id: 5815,  Name: "Cumulative reactive power",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VARh",  Description: "The cumulative reactive power since the last cumulative energy reset or device start"},
        &ResourceModel{ Id: 5816,  Name: "Reactive Power Calibration",  Operations: OPERATION_W,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "VAR",  Description: "Request a reactive power calibration by writing the value of a calibrated load."},
        &ResourceModel{ Id: 5820,  Name: "Power factor",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "If applicable, the power factor of the current consumption."},
        &ResourceModel{ Id: 5821,  Name: "Current Calibration",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "Read or Write the current calibration coefficient"},
        &ResourceModel{ Id: 5822,  Name: "Reset Cumulative energy",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset both cumulative active/reactive power"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Actuation",    Id: 3306,    Multiple: true,    Mandatory: false,    Description: "This IPSO object is dedicated to remote actuation such as ON/OFF action or dimming. A multi-state output can also be described as a string. This is useful to send pilot wire orders for instance. It also provides a resource to reflect the time that the device has been switched on. "},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The Application type of the device, for example “Motion Closure”."},
        &ResourceModel{ Id: 5850,  Name: "On/Off",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "On/Off\n\n"},
        &ResourceModel{ Id: 5851,  Name: "Dimmer",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "0-100",  Units: "%",  Description: "This resource represents a light dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
        &ResourceModel{ Id: 5852,  Name: "On Time",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "sec",  Description: "The time in seconds that the device has been on. Writing a value of 0 resets the counter."},
        &ResourceModel{ Id: 5853,  Name: "Muti-state Output",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "A string describing a state for multiple level output such as Pilot Wire"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Set Point",    Id: 3308,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object should be used to set a desired value to a controller, such as a thermostat. This object enables a setpoint to be expressed units defined in the UCUM specification, to match an associated sensor or measurement value. A special resource is added to set the colour of an object."},
        &ResourceModel{ Id: 5701,  Name: "Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
        &ResourceModel{ Id: 5706,  Name: "Colour",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "A string representing a value in some color space"},
        &ResourceModel{ Id: 5750,  Name: "Application Type",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The Application type of the device, for example “Motion Closure”."},
        &ResourceModel{ Id: 5900,  Name: "Set Point Value",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.\t",  Description: "The setpoint value. "},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Load Control",    Id: 3310,    Multiple: true,    Mandatory: false,    Description: "Description: This Object is used for demand-response load control and other load control in automation application (not limited to power)."},
        &ResourceModel{ Id: 5823,  Name: "Event Identifier",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The event identifier as a string."},
        &ResourceModel{ Id: 5824,  Name: "Start Time",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Time when the load control event will start started."},
        &ResourceModel{ Id: 5825,  Name: "Duration In Min",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "Min",  Description: "The duration of the load control event."},
        &ResourceModel{ Id: 5826,  Name: "Criticality Level",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "The criticality of the event.  The device receiving the event will react in an appropriate fashion for the device."},
        &ResourceModel{ Id: 5827,  Name: "Avg Load Adj Pct",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "0-100",  Units: "%",  Description: "Defines the maximum energy usage of the receivng device, as a percentage of the device\u0027s normal maximum energy usage."},
        &ResourceModel{ Id: 5828,  Name: "Duty Cycle",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "0-100",  Units: "%",  Description: "Defines the duty cycle for the load control event, i.e, what percentage of time the receiving device is allowed to be on."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Light Control",    Id: 3311,    Multiple: true,    Mandatory: false,    Description: "Description: This Object is used to control a light source, such as a LED or other light.  It allows a light to be turned on or off and its dimmer setting to be control as a % between 0 and 100. An optional colour setting enables a string to be used to indicate the desired colour."},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
        &ResourceModel{ Id: 5706,  Name: "Colour",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "A string representing a value in some color space"},
        &ResourceModel{ Id: 5805,  Name: "Cumulative active power ",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Wh",  Description: "The total power in Wh that the light has used."},
        &ResourceModel{ Id: 5820,  Name: "Power factor",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The power factor of the light."},
        &ResourceModel{ Id: 5850,  Name: "On/Off",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "This resource represents a light, which can be controlled, the setting of which is a Boolean value (1,0) where 1 is on and 0 is off."},
        &ResourceModel{ Id: 5851,  Name: "Dimmer",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "0-100",  Units: "%",  Description: "This resource represents a light dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
        &ResourceModel{ Id: 5852,  Name: "On Time",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "sec",  Description: "The time in seconds that the light has been on. Writing a value of 0 resets the counter."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Power Control",    Id: 3312,    Multiple: true,    Mandatory: false,    Description: "Description: This Object is used to control a power source, such as a Smart Plug.  It allows a power relay to be turned on or off and its dimmer setting to be control as a % between 0 and 100."},
        &ResourceModel{ Id: 5805,  Name: "Cumulative active power ",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Wh",  Description: "The total power in Wh that has been used by the load."},
        &ResourceModel{ Id: 5820,  Name: "Power factor",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "",  Description: "The power factor of the load."},
        &ResourceModel{ Id: 5850,  Name: "On/Off",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_BOOLEAN,  RangeOrEnums: "",  Units: "",  Description: "This resource represents a power relay, which can be controlled, the setting of which is a Boolean value (1,0) where 1 is on and 0 is off."},
        &ResourceModel{ Id: 5851,  Name: "Dimmer",  Operations: OPERATION_W,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "0-100",  Units: "%",  Description: "This resource represents a power dimmer setting, which has an Integer value between 0 and 100 as a percentage."},
        &ResourceModel{ Id: 5852,  Name: "On Time",  Operations: OPERATION_RW,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_INTEGER,  RangeOrEnums: "",  Units: "sec",  Description: "The time in seconds that the power relay has been on. Writing a value of 0 resets the counter."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Accelerometer",    Id: 3313,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object can be used to represent a 1-3 axis accelerometer."},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
        &ResourceModel{ Id: 5702,  Name: "X Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the X axis."},
        &ResourceModel{ Id: 5703,  Name: "Y Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the Y axis."},
        &ResourceModel{ Id: 5704,  Name: "Z Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the Z axis."},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Magnetometer",    Id: 3314,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object can be used to represent a 1-3 axis magnetometer with optional compass direction."},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "Measurement Units Definition e.g. “Cel” for Temperature in Celsius."},
        &ResourceModel{ Id: 5702,  Name: "X Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the X axis."},
        &ResourceModel{ Id: 5703,  Name: "Y Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the Y axis."},
        &ResourceModel{ Id: 5704,  Name: "Z Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The measured value along the Z axis."},
        &ResourceModel{ Id: 5705,  Name: "Compass Direction",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "0-360",  Units: "Deg",  Description: "The compass direction"},
    )

    o.Add(
        &ObjectModel{ Name: "IPSO Barometer",    Id: 3315,    Multiple: true,    Mandatory: false,    Description: "Description: This IPSO object should be used with an air pressure sensor to report a barometer measurement.  It also provides resources for minimum/maximum measured values and the minimum/maximum range that can be measured by the barometer sensor. An example measurement unit is kPa (ucum:kPa)."},
        &ResourceModel{ Id: 5601,  Name: "Min Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5602,  Name: "Max Measured Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value measured by the sensor since power ON or reset"},
        &ResourceModel{ Id: 5603,  Name: "Min Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The minimum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5604,  Name: "Max Range Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "The maximum value that can be measured by the sensor"},
        &ResourceModel{ Id: 5605,  Name: "Reset Min and Max Measured Values",  Operations: OPERATION_E,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_OPAQUE,  RangeOrEnums: "",  Units: "",  Description: "Reset the Min and Max Measured Values to Current Value"},
        &ResourceModel{ Id: 5700,  Name: "Sensor Value",  Operations: OPERATION_R,  Multiple: false,  Mandatory: true,  ResourceType: TYPE_FLOAT,  RangeOrEnums: "",  Units: "Defined by “Units” resource.",  Description: "Last or Current Measured Value from the Sensor"},
        &ResourceModel{ Id: 5701,  Name: "Sensor Units",  Operations: OPERATION_R,  Multiple: false,  Mandatory: false,  ResourceType: TYPE_STRING,  RangeOrEnums: "",  Units: "",  Description: "If present, the type of sensor defined as the UCUM Unit Definition e.g. “Cel” for Temperature in Celcius."},
    )
}

func (o *IPSOSmartObjects) Get(n LWM2MObjectType) (*ObjectModel) {
    return o.models[n]
}

func (mr *IPSOSmartObjects) Add(m *ObjectModel, res ...*ResourceModel) {

}