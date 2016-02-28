package betwixt

import (
	"log"
	"encoding/json"
)

type YamlObjectDefinition struct {
	Id       int
	Name     string `json:"Name"`
	Multiple bool
}

func ParseObjectDefinitionsJson(data []byte) []ObjectDefinition {

	var n map[string]interface{}

	err := json.Unmarshal(data, &n)
	if err != nil {
		log.Println(err, n)
	}

	obj := []ObjectDefinition{}
	l := n["Objects"].([]interface{})

	for _, _v := range l {
		v := _v.(map[string]interface{})

		m := &DefaultObjectDefinition{}

		id := int(v["Id"].(float64))
		m.Id = LWM2MObjectType(id)

		if v["Name"] != nil {
			m.Name = v["Name"].(string)
		}

		if v["Description"] != nil {
			m.Description = v["Description"].(string)
		}


		m.Multiple, _ = v["Multiple"].(bool)
		m.Mandatory, _ = v["Mandatory"].(bool)
		if v["Resources"] != nil {
			resources := v["Resources"].([]interface{})

			res := []ResourceDefinition{}

			for _, v2 := range resources {
				r := v2.(map[string]interface{})

				rd := &DefaultResourceDefinition{}

				rd.Id = LWM2MResourceType(int(r["Id"].(float64)))

				if r["Name"] != nil {
					rd.Name = r["Name"].(string)
				}

				if r["Multiple"] != nil {
					rd.Multiple = r["Multiple"].(bool)
				}

				if r["Mandatory"] != nil {
					rd.Mandatory = r["Mandatory"].(bool)
				}

				op := r["Operations"]
				rt := r["ResourceType"]

				if r["RangeOrEnums"] != nil {
					rd.RangeOrEnums = r["RangeOrEnums"].(string)
				}

				// vv := r["ValueValidator"].(yaml.Scalar).String()

				switch rt {
				case "multiple":
					rd.ResourceType = VALUETYPE_MULTIPLE

				case "string":
					rd.ResourceType = VALUETYPE_STRING

				case "byte":
					rd.ResourceType = VALUETYPE_BYTE

				case "int":
					rd.ResourceType = VALUETYPE_INTEGER

				case "int32":
					rd.ResourceType = VALUETYPE_INTEGER32

				case "int64":
					rd.ResourceType = VALUETYPE_INTEGER64

				case "float":
					rd.ResourceType = VALUETYPE_FLOAT

				case "float64":
					rd.ResourceType = VALUETYPE_FLOAT64

				case "bool":
					rd.ResourceType = VALUETYPE_BOOLEAN

				case "opaque":
					rd.ResourceType = VALUETYPE_OPAQUE

				case "time":
					rd.ResourceType = VALUETYPE_TIME

				case "objectlink":
					rd.ResourceType = VALUETYPE_OBJECTLINK

				case "object":
					rd.ResourceType = VALUETYPE_OBJECT

				case "resource":
					rd.ResourceType = VALUETYPE_RESOURCE

				case "multiresource":
					rd.ResourceType = VALUETYPE_MULTIRESOURCE
				}

				switch op {
				case "N":
					rd.Operations = OPERATION_NONE

				case "R":
					rd.Operations = OPERATION_R

				case "W":
					rd.Operations = OPERATION_W

				case "RW":
					rd.Operations = OPERATION_RW

				case "E":
					rd.Operations = OPERATION_E

				case "RE":
					rd.Operations = OPERATION_RE

				case "WE":
					rd.Operations = OPERATION_WE

				case "RWE":
					rd.Operations = OPERATION_RWE
				}

				res = append(res, rd)
			}

			m.SetResources(res)
		}

		obj = append(obj, m)
	}

	return obj
}
