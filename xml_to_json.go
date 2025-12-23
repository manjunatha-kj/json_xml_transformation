package json_xml_transformation

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"strings"
)

type xmlNode struct {
	Name       string
	Attributes map[string]string
	Children   map[string][]any
	Text       string
}

func (c *converterImpl) XMLToJSON(xmlData []byte, opts ...Option) ([]byte, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	dec := xml.NewDecoder(bytes.NewReader(xmlData))

	var stack []*xmlNode
	var root *xmlNode

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, ErrInvalidXML
		}

		switch t := tok.(type) {

		case xml.StartElement:
			node := &xmlNode{
				Name:       t.Name.Local,
				Attributes: map[string]string{},
				Children:   map[string][]any{},
			}

			for _, attr := range t.Attr {
				node.Attributes[options.AttrPrefix+attr.Name.Local] = attr.Value
			}

			if len(stack) == 0 {
				root = node
			} else {
				parent := stack[len(stack)-1]
				parent.Children[node.Name] = append(parent.Children[node.Name], node)
			}

			stack = append(stack, node)

		case xml.CharData:
			if len(stack) > 0 {
				text := strings.TrimSpace(string(t))
				if text != "" {
					stack[len(stack)-1].Text += text
				}
			}

		case xml.EndElement:
			stack = stack[:len(stack)-1]
		}
	}

	result := map[string]any{
		root.Name: xmlNodeToMap(root, options),
	}

	if options.PrettyPrint {
		return json.MarshalIndent(result, "", "  ")
	}
	return json.Marshal(result)
}

func xmlNodeToMap(n *xmlNode, opts *Options) any {
	m := map[string]any{}

	for k, v := range n.Attributes {
		m[k] = v
	}

	if len(n.Children) == 0 && n.Text != "" {
		if len(n.Attributes) == 0 {
			return n.Text
		}
		m[opts.TextKey] = n.Text
		return m
	}

	for k, children := range n.Children {
		if len(children) == 1 {
			m[k] = xmlNodeToMap(children[0].(*xmlNode), opts)
		} else {
			arr := make([]any, 0, len(children))
			for _, c := range children {
				arr = append(arr, xmlNodeToMap(c.(*xmlNode), opts))
			}
			m[k] = arr
		}
	}

	if n.Text != "" {
		m[opts.TextKey] = n.Text
	}

	return m
}
