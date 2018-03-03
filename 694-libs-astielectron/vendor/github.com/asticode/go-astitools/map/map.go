package astimap

// Map represents a bi-directional map
type Map struct {
	defaultA interface{}
	defaultB interface{}
	mapAToB  map[interface{}]interface{}
	mapBToA  map[interface{}]interface{}
}

// NewMap builds a new *Map
func NewMap(defaultA, defaultB interface{}) *Map {
	return &Map{
		defaultA: defaultA,
		defaultB: defaultB,
		mapAToB:  make(map[interface{}]interface{}),
		mapBToA:  make(map[interface{}]interface{}),
	}
}

// A retrieves a based on b
func (m *Map) A(b interface{}) interface{} {
	if a, ok := m.mapBToA[b]; ok {
		return a
	}
	return m.defaultA
}

// B retrieves b based on a
func (m *Map) B(a interface{}) interface{} {
	if b, ok := m.mapAToB[a]; ok {
		return b
	}
	return m.defaultB
}

// InA checks whether a exists
func (m *Map) InA(a interface{}) (ok bool) {
	_, ok = m.mapAToB[a]
	return
}

// InB checks whether b exists
func (m *Map) InB(b interface{}) (ok bool) {
	_, ok = m.mapBToA[b]
	return
}

// Set sets a key/value couple
func (m *Map) Set(a, b interface{}) *Map {
	m.mapAToB[a] = b
	m.mapBToA[b] = a
	return m
}
