# Prompt für Claude Code (Aufgabe 2)

Ein einziger Prompt, der das Beispiel zum Composite-Konstruktionsprinzip
(Folder-Docs) in Go implementieren lässt. Er ist bewusst kurz und enthält nur
das, was das Ergebnis sonst dem Zufall überlässt: Rollen des Patterns, Beispiel
aus der Angabe, Tests, Verifikation.

## Der Prompt

```text
Implement the Composite design pattern in Go in this directory, using the
Folder-Docs example from the lecture.

Create one Go module (module github.com/Tikariki/SE_Aufgabe_2_Katrin_Linz, go 1.21):

1. Package composite:
   - Doc: component interface with Name() string and Size() int64
   - TextDoc: leaf, holds name and size in bytes, implements both methods
   - Folder: composite, holds a name and []Doc, has Add(Doc), and its Size()
     sums the children by calling doc.Size() on each of them, recursively.
     No type switches anywhere: a folder must treat files and subfolders the
     same way.
2. cmd/demo/main.go: build the lecture example - Root with two TextDocs
   (10 and 50 bytes) and one subfolder containing a TextDoc of 100 bytes -
   print the tree and the total size, which must be 160.
3. composite/composite_test.go: tests for leaf size, empty folder, the 160
   example, nesting at least three levels deep, and a []Doc holding a TextDoc
   and a Folder side by side.

Style: idiomatic Go like the gopl.io examples - small exported API, no getter
boilerplate, doc comments on exported identifiers, errors only where they can
actually occur, English comments.

Also add .gitignore (binaries, test binaries, coverage, .idea/, .env) and a
README.md describing the class structure, the example calculation and how to
run demo and tests.

Finally run gofmt -l ., go vet ./... and go test ./... and show me the real
output. Do not create a Git repository and do not push anything.
```

## Warum der Prompt so aussieht

- **Rollen explizit** (Component / Leaf / Composite): ohne diese drei Wörter
  entsteht oft nur eine Klasse mit einer Liste, nicht das Pattern.
- **"No type switches"** ist die eigentliche Prüfungsfrage des Patterns. Wer
  `if _, ok := doc.(*Folder)` schreibt, hat den Nutzen schon verloren.
- **Beispiel mit 160** aus der Angabe: damit ist das Ergebnis nachrechenbar
  statt nur "sieht plausibel aus".
- **Tests und Verifikation im Prompt**: der Agent muss sein Ergebnis belegen,
  nicht behaupten.
- **"Do not create a Git repository"**: Repo und Push sind bei dieser Abgabe
  nicht Teil der Aufgabe.
- Weggelassen: Struktur-Vorgaben wie Ordnernamen, Formatierungsregeln,
  Sprachwahl der Kommentare im Detail. Das ergibt sich aus "idiomatic Go".

## Was der Prompt erzeugt hat

Dieses Repository ist das Ergebnis: `composite/` mit den drei Typen und den
Tests, `cmd/demo` mit dem Beispiel aus der Angabe. Siehe `README.md`.

```text
$ go run ./cmd/demo
Root (160 Bytes)
  notizen.txt (10 Bytes)
  bericht.txt (50 Bytes)
  Unterordner (100 Bytes)
    anhang.txt (100 Bytes)

Gesamtgröße von "Root": 160 Bytes

$ go test ./...
ok  github.com/Tikariki/SE_Aufgabe_2_Katrin_Linz/composite
```
