package xmljsontransformation

import "errors"

var (
	ErrInvalidJSON = errors.New("invalid json payload")
	ErrInvalidXML  = errors.New("invalid xml payload")
)
