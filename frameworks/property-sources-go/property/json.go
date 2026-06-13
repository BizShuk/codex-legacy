package property

import (
	"encoding/json"
	"io"
	"log"
	propertyio "propertychain/io"
)

// JsonProperties load properties from a flat JSON object: {"key":"value", ...}
type JsonProperties map[string]string

func (p JsonProperties) ReadFrom(reader io.Reader) (n int64, err error) {
	dec := json.NewDecoder(reader)
	// accept one top-level JSON object
	if err := dec.Decode(&p); err != nil {
		return 0, err
	}
	// n is unknown for streamed JSON decoding; report 0 (caller does not rely on it)
	return 0, nil
}

// Get property with key
func (p JsonProperties) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p JsonProperties) Set(key string, val string) {
	p[key] = val
}

// NewJsonProperties constructor
func NewJsonProperties(uri string) (Properties, error) {

	reader := propertyio.GetProtocolHandler(uri)()
	defer reader.Close()

	p := JsonProperties(make(map[string]string))

	_, err := p.ReadFrom(reader)
	if err != nil {
		log.Fatal("Load JsonProperties failed")
		return nil, err
	}
	return p, nil
}
