package tlv
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateIdentifierField(t *testing.T) {
	var id []byte
	id = CreateTlvIdentifierField(1)
	assert.Equal(t, []byte{1}, id)

	id = CreateTlvIdentifierField(10)
	assert.Equal(t, []byte{10}, id)

	id = CreateTlvIdentifierField(100)
	assert.Equal(t, []byte{100}, id)

	id = CreateTlvIdentifierField(1000)
	assert.Equal(t, []byte{232, 3}, id)

	id = CreateTlvIdentifierField(10000)
	assert.Equal(t, []byte{16, 39}, id)
}

func TestCreateTlvTypeField(t *testing.T) {
	/*
	func CreateTlvTypeField(identType byte, value interface{}, ident uint16) byte {
		var typeField byte

		valueTypeLength, _ := typeval.GetValueByteLength(value)

		// Bit 7-6: identifier
		typeField |= identType

		// Bit 5
		if ident > 255 {
			typeField |= 32
		}

		// Bit 4-3
		if valueTypeLength > 7 {
			if valueTypeLength < 256 {
				typeField |= 8
			} else {
				if valueTypeLength < 65535 {
					typeField |= 16
				} else {
					if valueTypeLength > 16777215 {
						// Error, size exceeds allowed (> 16.7MB)
					} else {
						// Size is 16777215 or less
						typeField |= 24
					}
				}
			}
		} else {
			// Set bit 2-0 instead
			b := byte(valueTypeLength)
			typeField |= b
		}

		return typeField
	}
	 */
}

func TestCreateTlvLengthField(t *testing.T) {

}
/*
func CreateTlvLengthField(value interface{}) []byte {
	valueTypeLength, _ := typeval.GetValueByteLength(value)

	if valueTypeLength > 7 {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, valueTypeLength)

		return bytes.Trim(buf.Bytes(), "\x00")
	}
	return []byte{}
}
*/

func TestCreateTlvValueField(t *testing.T) {

}
/*
func CreateTlvValueField(value int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint64(value))
	if value == 0 {
		return []byte{0}
	} else {
		return bytes.Trim(buf.Bytes(), "\x00")
	}
}
*/