package composite

// TextDoc is the leaf of the pattern: a single document without children.
type TextDoc struct {
	name string
	size int64
}

// A TextDoc is a Doc.
var _ Doc = (*TextDoc)(nil)

// NewTextDoc returns a TextDoc with the given name and size in bytes.
func NewTextDoc(name string, size int64) *TextDoc {
	return &TextDoc{name: name, size: size}
}

// Name returns the file name of the document.
func (t *TextDoc) Name() string { return t.name }

// Size returns the size of the document itself.
func (t *TextDoc) Size() int64 { return t.size }
