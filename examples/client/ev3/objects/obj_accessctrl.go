package ev3

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/go-commons/typeval"
)

type AccessControlObject struct {
	Model ObjectDefinition
}

func (o *AccessControlObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *AccessControlObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *AccessControlObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *AccessControlObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val typeval.Value

		// resource := o.Model.GetResource(resourceId)
		switch instanceId {
		case 0:
			switch resourceId {
			case 0:
				val = typeval.Integer(1)
				break

			case 1:
				val = typeval.Integer(0)
				break

			case 2:
				// "/0/2/101", []byte{0, 15}
				break

			case 3:
				break
			}
			break

		case 1:
			switch resourceId {
			case 0:
				break

			case 1:
				break

			case 2:
				break

			case 3:
				break
			}
			break

		case 2:
			switch resourceId {
			case 0:
				break

			case 1:
				break

			case 2:
				break

			case 3:
				break
			}
			break

		case 3:
			switch resourceId {
			case 0:
				break

			case 1:
				break

			case 2:
				break

			case 3:
				break
			}
			break

		case 4:
			switch resourceId {
			case 0:
				break

			case 1:
				break

			case 2:
				break

			case 3:
				break
			}
			break
		}
		if val == nil {
			return NotFound()
		} else {
			return Content(val)
		}
	}
	return NotFound()
}

func (o *AccessControlObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleAccessControlObject(reg Registry) *AccessControlObject {
	return &AccessControlObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_SECURITY),
	}
}
