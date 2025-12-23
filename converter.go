package xmljsontransformation

type Converter interface {
	JSONToXML(jsonData []byte, opts ...Option) ([]byte, error)
	XMLToJSON(xmlData []byte, opts ...Option) ([]byte, error)
}

type converterImpl struct{}

func New() Converter {
	return &converterImpl{}
}
