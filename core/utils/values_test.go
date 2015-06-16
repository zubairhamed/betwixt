package utils
import "testing"

func TestDecodeValue (t *testing.T) {
	var data []byte

	// Manufacturer
	data = []byte{ 79, 112, 101, 110, 32, 77, 111, 98, 105, 108, 101, 32, 65, 108, 108, 105, 97, 110, 99, 101 }
	DecodeValue(data)

	// Model Number
	data = []byte{ 76, 105, 103, 104, 116, 119, 101, 105, 103, 104, 116, 32, 77, 50, 77, 32, 67, 108, 105, 101, 110, 116, }
	DecodeValue(data)

	// Serial
	data = []byte{ 51, 52, 53, 48, 48, 48, 49, 50, 51 }
	DecodeValue(data)

	// Firmware
	data = []byte{ 49, 46, 48 }
	DecodeValue(data)

	// Available Power
	data = []byte{ 134, 6, 65, 0, 1, 65, 1, 5 }
	DecodeValue(data)

	// Power Source Voltage
	data = []byte{ 136, 7, 8, 66, 0, 14, 216, 66, 1, 19, 136 }
	DecodeValue(data)

	// Power Source Current
	data = []byte{ 135, 8, 65, 0, 125, 66, 1, 3, 132 }
	DecodeValue(data)

	// Memory Free
	data = []byte{ 49, 53 }
	DecodeValue(data)
}

/*
Serial		 	[51 52 53 48 48 48 49 50 51]
Firmware		[49 46 48]
Avail. Power 	[134 6 65 0 1 65 1 5]
Power Src Volt. [136 7 8 66 0 14 216 66 1 19 136]
Power Src Curr. [135 8 65 0 125 66 1 3 132]
Memory Free		[49 53]
 */