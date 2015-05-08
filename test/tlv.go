package main
import (
    // "encoding/hex"
    "log"
)

func main() {
    /*
    var hexString = "8606410001410105"
    rawMessage, _ := hex.DecodeString(hexString)

    for _, b := range rawMessage {
        log.Println (b)
    }
    */

    var binval byte = 85
    // 1 2 4 8
    // 00000110
    // 00000111

    /*
    if type == OBJECT_INSTANCE
    if type == RESOURCE_WITH_VALUE
    if type == MULTIPLE_RESOURCE
    if type == RESOURCE_INSTANCE
    */


    log.Println ((binval & 0xC0) >> 6)
}


/*
    for bytes {
        read first byte
        if type == object instance

        if type == resource instance
        if type == multiple resources
        if type == resource with value
    }
*/


