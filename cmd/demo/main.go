// Command demo builds the Folder-Docs example from the lecture and prints the
// resulting tree together with the total size of the root folder.
package main

import (
	"fmt"
	"strings"

	"github.com/Tikariki/SE_Aufgabe_2_Katrin_Linz/composite"
)

func main() {
	root := composite.NewFolder("Root")
	root.Add(composite.NewTextDoc("notizen.txt", 10))
	root.Add(composite.NewTextDoc("bericht.txt", 50))

	sub := composite.NewFolder("Unterordner")
	sub.Add(composite.NewTextDoc("anhang.txt", 100))
	root.Add(sub)

	printTree(root, 0)
	fmt.Println()
	fmt.Printf("Gesamtgröße von %q: %d Bytes\n", root.Name(), root.Size())
}

// printTree prints doc and, if it has children, its subtree.
func printTree(doc composite.Doc, depth int) {
	fmt.Printf("%s%s (%d Bytes)\n", strings.Repeat("  ", depth), doc.Name(), doc.Size())

	// Only the tree output needs to look at the children; the pattern itself
	// does not. The anonymous interface keeps the demo independent of the
	// concrete type Folder.
	if parent, ok := doc.(interface{ Children() []composite.Doc }); ok {
		for _, child := range parent.Children() {
			printTree(child, depth+1)
		}
	}
}
