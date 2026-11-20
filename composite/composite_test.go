package composite_test

import (
	"testing"

	"github.com/Tikariki/SE_Aufgabe_2_Katrin_Linz/composite"
)

func TestTextDocSize(t *testing.T) {
	tests := []struct {
		name string
		size int64
	}{
		{"leer.txt", 0},
		{"notizen.txt", 10},
		{"bericht.txt", 50},
	}
	for _, tt := range tests {
		doc := composite.NewTextDoc(tt.name, tt.size)
		if got := doc.Size(); got != tt.size {
			t.Errorf("%s: Size() = %d, want %d", tt.name, got, tt.size)
		}
		if got := doc.Name(); got != tt.name {
			t.Errorf("Name() = %q, want %q", got, tt.name)
		}
	}
}

func TestEmptyFolderSize(t *testing.T) {
	if got := composite.NewFolder("leer").Size(); got != 0 {
		t.Errorf("Size() = %d, want 0", got)
	}
}

// TestFolderSizeExample builds the hierarchy from the lecture example:
//
//	Root            -> TextDoc 10, TextDoc 50, Unterordner -> TextDoc 100
//
// and expects a total size of 160.
func TestFolderSizeExample(t *testing.T) {
	root := composite.NewFolder("Root")
	root.Add(composite.NewTextDoc("a.txt", 10))
	root.Add(composite.NewTextDoc("b.txt", 50))

	sub := composite.NewFolder("Unterordner")
	sub.Add(composite.NewTextDoc("c.txt", 100))
	root.Add(sub)

	if got, want := root.Size(), int64(160); got != want {
		t.Errorf("Root.Size() = %d, want %d", got, want)
	}
	if got, want := sub.Size(), int64(100); got != want {
		t.Errorf("Unterordner.Size() = %d, want %d", got, want)
	}
}

// TestNestedFolders checks that the recursion works over several levels.
func TestNestedFolders(t *testing.T) {
	// Root(1) -> Ebene1(2) -> Ebene2(4) -> Datei(8) = 15
	deepest := composite.NewFolder("Ebene2")
	deepest.Add(composite.NewTextDoc("datei.txt", 8))

	middle := composite.NewFolder("Ebene1")
	middle.Add(composite.NewTextDoc("eins.txt", 4))
	middle.Add(deepest)

	root := composite.NewFolder("Root")
	root.Add(composite.NewTextDoc("null.txt", 2))
	root.Add(composite.NewTextDoc("zwei.txt", 1))
	root.Add(middle)

	if got, want := root.Size(), int64(15); got != want {
		t.Errorf("Root.Size() = %d, want %d", got, want)
	}
}

// TestUniformTreatment shows the point of the pattern: a Folder and a TextDoc
// live in the same slice and are summed up without any distinction.
func TestUniformTreatment(t *testing.T) {
	sub := composite.NewFolder("Unterordner")
	sub.Add(composite.NewTextDoc("c.txt", 100))

	docs := []composite.Doc{
		composite.NewTextDoc("a.txt", 10),
		composite.NewTextDoc("b.txt", 50),
		sub,
	}

	var total int64
	for _, doc := range docs {
		total += doc.Size()
	}
	if got, want := total, int64(160); got != want {
		t.Errorf("Summe = %d, want %d", got, want)
	}
}
