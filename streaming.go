package xmljsontransformation

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

func JSONToXMLStream(r io.Reader, w io.Writer, opts ...Option) error {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	var data any
	dec := json.NewDecoder(r)
	if err := dec.Decode(&data); err != nil {
		return ErrInvalidJSON
	}

	enc := xml.NewEncoder(w)
	if options.PrettyPrint {
		enc.Indent("", "  ")
	}

	root := xml.StartElement{Name: xml.Name{Local: options.RootName}}
	if err := enc.EncodeToken(root); err != nil {
		return err
	}

	if err := jsonToXMLValue(data, enc); err != nil {
		return err
	}

	if err := enc.EncodeToken(root.End()); err != nil {
		return err
	}

	return enc.Flush()
}

func XMLToJSONStream(r io.Reader, w io.Writer, opts ...Option) error {
	xmlData, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	c := New()
	jsonData, err := c.XMLToJSON(xmlData, opts...)
	if err != nil {
		return err
	}

	_, err = w.Write(jsonData)
	return err
}
