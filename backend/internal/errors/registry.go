// Metadata provides extra info for each Code (docs, severity, owner)
package errors

type Metadata struct {
	Code           Code
	Title          string
	Description    string
	Severity       string // e.g., "info", "warning", "critical"
	SuggestedAction string
	Owner          string
}

var registry = map[Code]Metadata{}

func RegisterMeta(m Metadata) {
	registry[m.Code] = m
}

func GetMeta(code Code) (Metadata, bool) {
	v, ok := registry[code]
	return v, ok
}
