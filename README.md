# SE Aufgabe 2 – Composite-Konstruktionsprinzip in Go

Umsetzung des Beispiels aus der Angabe (Verzeichnishierarchie "Folder-Docs")
mit dem Composite Design Pattern in Go, plus der Prompt, mit dem die
Implementierung erzeugt wurde.

Die eigentliche Abgabe dieser Aufgabe ist der Prompt: **[PROMPT.md](PROMPT.md)**.
Der Code hier ist das Ergebnis dieses Prompts und dient als Beleg, dass er
funktioniert.

## Klassenstruktur

| Rolle im Pattern | Typ in Go | Aufgabe |
|---|---|---|
| Component | `composite.Doc` (Interface) | gemeinsamer Obertyp mit `Name() string` und `Size() int64` |
| Leaf | `composite.TextDoc` | einzelne Datei; `Size()` liefert die echte Dateigröße |
| Composite | `composite.Folder` | enthält `[]Doc`, `Add(Doc)`, `Size()` summiert rekursiv |

`Folder` implementiert `Doc` selbst (`var _ Doc = (*Folder)(nil)`). Deshalb kann
ein Ordner genauso in einem Ordner stecken wie eine Datei, und der aufrufende
Code muss nie unterscheiden, worum es sich handelt:

```go
func (f *Folder) Size() int64 {
	var total int64
	for _, doc := range f.docs {
		total += doc.Size() // TextDoc oder Folder - egal
	}
	return total
}
```

Kein `type switch`, keine Sonderfälle: die Rekursion entsteht dadurch, dass ein
`Folder` beim Summieren wieder `Size()` aufruft und dabei auf sich selbst
treffen kann.

## Beispielrechnung

```
Root
├── notizen.txt            10 Bytes
├── bericht.txt            50 Bytes
└── Unterordner           100 Bytes
    └── anhang.txt        100 Bytes
```

10 + 50 + 100 = **160 Bytes** für `Root`.

## Projektstruktur

```
.
├── PROMPT.md                  Prompt, der dieses Projekt erzeugt hat
├── go.mod                     Modul github.com/Tikariki/SE_Aufgabe_2_Katrin_Linz
├── composite/
│   ├── doc.go                 Interface Doc (Component)
│   ├── textdoc.go             TextDoc (Leaf)
│   ├── folder.go              Folder (Composite)
│   └── composite_test.go      Tests
└── cmd/
    └── demo/
        └── main.go            Beispiel aus der Angabe
```

## Ausführen

Voraussetzung: Go 1.21 oder neuer.

```bash
go run ./cmd/demo
go test ./...
```

Erwartete Ausgabe des Demos:

```
Root (160 Bytes)
  notizen.txt (10 Bytes)
  bericht.txt (50 Bytes)
  Unterordner (100 Bytes)
    anhang.txt (100 Bytes)

Gesamtgröße von "Root": 160 Bytes
```

## Tests

`composite/composite_test.go` prüft:

- `TestTextDocSize` – Leaf liefert seine eigene Größe (tabellengetrieben)
- `TestEmptyFolderSize` – leerer Ordner ist 0
- `TestFolderSizeExample` – das Beispiel aus der Angabe ergibt 160
- `TestNestedFolders` – Rekursion über mehrere Ebenen
- `TestUniformTreatment` – `TextDoc` und `Folder` liegen gemeinsam in einem
  `[]Doc` und werden ohne Unterscheidung aufsummiert

## Entwurfsentscheidungen

- **Interface statt Basisklasse.** Go kennt keine Vererbung; die gemeinsame
  Schnittstelle `Doc` übernimmt die Rolle der abstrakten Component-Klasse.
- **Zeiger auf `Folder`.** `Add` verändert den Ordner, deshalb `*Folder`.
  `TextDoc` ist unveränderlich und könnte auch als Wert übergeben werden; die
  einheitliche Verwendung von Zeigern hält die Aufrufstellen einfach.
- **Kleine API.** Nur `Name` und `Size` stehen im Interface. Eine Methode wie
  `Children()` gibt es nur am konkreten `Folder`, weil sie für das Pattern
  nicht nötig ist – das Demo nutzt sie lediglich für die Baumausgabe.
- **`int64` für Größen.** Dateigrößen können die Grenzen von `int` auf
  32-Bit-Systemen überschreiten.
