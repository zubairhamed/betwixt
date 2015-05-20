package core
import (
    "errors"
    "time"
    . "github.com/zubairhamed/lwm2m/api"
    "bytes"
    "fmt"
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
    } else
    if _, ok := val.(bool); ok {
        return 1, nil
    } else
    if _, ok := val.(string); ok {
        v := val.(string)

        return uint32(len(v)), nil
    } else
    if _, ok := val.(float64); ok {
        v := val.(float64)

        if v > +3.4E+38 || v < -3.4E+38 {
            return 8, nil
        } else {
            return 4, nil
        }
    } else
    if _, ok := val.(time.Time); ok {
        return 8, nil
    } else
    if _, ok := val.([]byte); ok {
        v := val.([]byte)
        return uint32(len(v)), nil
    } else {
        return 0, errors.New("Unknown type")
    }
}

func BuildModelResourceStringPayload(instances LWM2MObjectInstances) (string) {
    var buf bytes.Buffer

    for k, v := range instances {
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
