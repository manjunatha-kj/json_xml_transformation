# Converter

A schema-agnostic Go package for real-time conversion between JSON and XML.

## Features
- JSON → XML
- XML → JSON
- Streaming support
- No structs required
- Production-ready

## Usage

```go
conv := converter.New()

xmlData, _ := conv.JSONToXML(
    []byte(`{"name":"Alice","age":30}`),
    converter.WithRoot("Person"),
    converter.WithPrettyPrint(true),
)

jsonData, _ := conv.XMLToJSON(xmlData, converter.WithPrettyPrint(true))
