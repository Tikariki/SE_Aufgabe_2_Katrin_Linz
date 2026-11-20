package composite

// Folder is the composite of the pattern: it holds any number of Docs, which
// may be TextDocs or further Folders. That makes every Folder a tree node, so
// hierarchies of arbitrary depth can be built.
type Folder struct {
	name string
	docs []Doc
}

// A Folder is a Doc as well, which is exactly what makes the pattern work.
var _ Doc = (*Folder)(nil)

// NewFolder returns an empty Folder with the given name.
func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

// Add appends doc to the folder. It does not matter whether doc is a TextDoc
// or another Folder.
func (f *Folder) Add(doc Doc) { f.docs = append(f.docs, doc) }

// Name returns the name of the folder.
func (f *Folder) Name() string { return f.name }

// Size returns the total size of the folder, i.e. the sum of the sizes of all
// contained Docs. Nested folders are handled implicitly: Size simply calls
// Size on every child and lets each child decide whether it is a file or a
// folder. No type switch is needed.
func (f *Folder) Size() int64 {
	var total int64
	for _, doc := range f.docs {
		total += doc.Size()
	}
	return total
}

// Children returns the contained Docs.
func (f *Folder) Children() []Doc { return f.docs }
