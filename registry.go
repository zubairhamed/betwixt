package betwixt

// NewDefaultObjectRegistry instantiates a default rgistry containing the Starter Pack
// objects and IPSO smart objects
func NewDefaultObjectRegistry() Registry {
	reg := NewObjectRegistry(&LWM2MCoreObjects{}, &IPSOSmartObjects{})

	return reg
}

func NewObjectRegistry(s ...ObjectSource) Registry {
	reg := &ObjectRegistry{}
	reg.sources = []ObjectSource{}

	for _, o := range s {
		reg.Register(o)
	}
	return reg
}

// ObjectRegistry is a registry containing known LWM2M objects registered to te OMA NA as well
// as custom objects
// It contains multiple ObjectSources which in turns each contains multiple Object definitions
type ObjectRegistry struct {
	sources []ObjectSource
}

func (m *ObjectRegistry) GetDefinition(n LWM2MObjectType) ObjectDefinition {
	for _, s := range m.sources {
		if s != nil {
			o := s.GetObject(n)
			if o != nil {
				return o
			}
		}
	}
	return nil
}

func (m *ObjectRegistry) GetDefinitions() []ObjectDefinition {
	defs := []ObjectDefinition{}

	for _, s := range m.sources {
		if s != nil {
			for _, v := range s.GetObjects() {
				defs = append(defs, v)
			}
		}
	}
	return defs
}

// Registers a new ObjectSource to the tegistry
func (m *ObjectRegistry) Register(s ObjectSource) {
	s.Initialize()
	m.sources = append(m.sources, s)
}

// Get all object definitions which are mandatory to be registered by
// a client (such as Firmware, Device etc)
func (m *ObjectRegistry) GetMandatory() []ObjectDefinition {
	mandatory := []ObjectDefinition{}

	for _, s := range m.sources {
		objs := s.GetObjects()
		for _, o := range objs {
			if o.IsMandatory() {
				mandatory = append(mandatory, o)
			}
		}
	}
	return mandatory
}
