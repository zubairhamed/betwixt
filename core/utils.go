package core

import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/zubairhamed/go-lwm2m/api"
	"time"
	"sort"
)

func GetValueByteLength(val interface{}) (uint32, error) {
	if _, ok := val.(int); ok {
		v := val.(int)
		if v > 127 || v < -128 {
			if v > 32767 || v < -32768 {
				if v > 2147483647 || v < -2147483648 {
					return 8, nil
				} else {
					return 4, nil
				}
			} else {
				return 2, nil
			}
		} else {
			return 1, nil
		}
	} else if _, ok := val.(bool); ok {
		return 1, nil
	} else if _, ok := val.(string); ok {
		v := val.(string)

		return uint32(len(v)), nil
	} else if _, ok := val.(float64); ok {
		v := val.(float64)

		if v > +3.4E+38 || v < -3.4E+38 {
			return 8, nil
		} else {
			return 4, nil
		}
	} else if _, ok := val.(time.Time); ok {
		return 8, nil
	} else if _, ok := val.([]byte); ok {
		v := val.([]byte)
		return uint32(len(v)), nil
	} else {
		return 0, errors.New("Unknown type")
	}
}

/*
 // To create a map as input
    m := make(map[int]string)
    m[1] = "a"
    m[2] = "c"
    m[0] = "b"

    // To store the keys in slice in sorted order
    var keys []int
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    // To perform the opertion you want
    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", m[k])
    }
*/

func BuildModelResourceStringPayload(instances LWM2MObjectInstances) string {
	var buf bytes.Buffer

	var keys []int
	for k := range instances {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := instances[LWM2MObjectType(k)]
		inst := v.GetObjectInstances()
		if len(inst) > 0 {
			for _, j := range inst {
				buf.WriteString(fmt.Sprintf("</%d/%d>,", k, j.GetId()))
			}
		} else {
			buf.WriteString(fmt.Sprintf("</%d>,", k))
		}
	}
	return buf.String()
}

func IsExecutableResource(m ResourceModel) bool {
	op := m.GetOperations()
	return (op == OPERATION_E || op == OPERATION_RE || op == OPERATION_RWE || op == OPERATION_WE)
}

func IsReadableResource(m ResourceModel) bool {
	op := m.GetOperations()
	return (op == OPERATION_RE || op == OPERATION_R || op == OPERATION_RWE || op == OPERATION_RW)
}

func IsWritableResource(m ResourceModel) bool {
	op := m.GetOperations()
	return (op == OPERATION_RW || op == OPERATION_RWE || op == OPERATION_WE || op == OPERATION_W)
}

//////////////////////////////////////////////////////
type ObjectsData struct {
	Data map[string]interface{}
}

func (o *ObjectsData) Put(path string, value interface{}) {
	o.Data[path] = value
}

func (o *ObjectsData) Get(path string) interface{} {
	return o.Data[path]
}

func (o *ObjectsData) Length() int {
	return len(o.Data)
}


func (o *ObjectsData) Clear() {
	o.Data = make(map[string]interface{})
}
