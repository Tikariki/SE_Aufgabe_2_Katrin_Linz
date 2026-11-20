// Package composite implements the Composite design pattern for a
// folder/document hierarchy ("Folder-Docs").
//
// A Folder is built from Docs. Because both a single document (TextDoc) and a
// folder (Folder) implement the same interface, a folder may contain files as
// well as other folders, and clients never have to distinguish between them.
package composite

// Doc is the component of the pattern: the common type of single documents and
// folders. Anything that is a Doc can be put into a Folder and its Size can be
// requested.
type Doc interface {
	// Name returns the name of the document or folder.
	Name() string

	// Size returns the size in bytes. For a TextDoc this is the real file
	// size, for a Folder it is the sum of the sizes of all contained Docs,
	// computed recursively.
	Size() int64
}
