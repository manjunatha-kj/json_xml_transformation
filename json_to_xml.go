package json_xml_transformation

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func (c *converterImpl) JSONToXML(jsonData []byte, opts ...Option) ([]byte, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	var data any
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, ErrInvalidJSON
	}

	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)

	if options.PrettyPrint {
		enc.Indent("", "  ")
	}

	root := xml.StartElement{Name: xml.Name{Local: options.RootName}}
	if err := enc.EncodeToken(root); err != nil {
		return nil, err
	}

	if err := jsonToXMLValue(data, enc); err != nil {
		return nil, err
	}

	if err := enc.EncodeToken(root.End()); err != nil {
		return nil, err
	}

	if err := enc.Flush(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func jsonToXMLValue(v any, enc *xml.Encoder) error {
	switch val := v.(type) {

	case map[string]any:
		for k, v2 := range val {
			elem := xml.StartElement{Name: xml.Name{Local: k}}
			if err := enc.EncodeToken(elem); err != nil {
				return err
			}
			if err := jsonToXMLValue(v2, enc); err != nil {
				return err
			}
			if err := enc.EncodeToken(elem.End()); err != nil {
				return err
			}
		}

	case []any:
		for _, item := range val {
			if err := jsonToXMLValue(item, enc); err != nil {
				return err
			}
		}

	default:
		return enc.EncodeToken(xml.CharData([]byte(fmt.Sprint(val))))
	}

	return nil
}
