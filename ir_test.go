package protocol

import (
	"strings"
	"testing"
)

func TestClassifyType(t *testing.T) {
	classes := []*Class{{Name: "Account"}}

	cases := []struct {
		goType   string
		kind     FieldKind
		elemType string
		elemKind FieldKind
		custom   bool
	}{
		{"int8", KindInt8, "", KindUnknown, false},
		{"string", KindString, "", KindUnknown, false},
		{"[]int8", KindSlice, "int8", KindInt8, false},
		{"[]string", KindSlice, "string", KindString, false},
		{"Account", KindStruct, "Account", KindUnknown, true},
		{"[]Account", KindSlice, "Account", KindStruct, true},
		{"map[string]int32", KindMap, "int32", KindUnknown, false},
		{"*Account", KindUnknown, "", KindUnknown, false}, // pointer stripped before classify in Load
	}

	for _, c := range cases {
		kind, elemType, elemKind, custom := classifyType(c.goType, classes)
		if kind != c.kind || elemType != c.elemType || elemKind != c.elemKind || custom != c.custom {
			t.Fatalf("classifyType(%q) = (%v,%q,%v,%v), want (%v,%q,%v,%v)",
				c.goType, kind, elemType, elemKind, custom, c.kind, c.elemType, c.elemKind, c.custom)
		}
	}
}

func TestClassToIR(t *testing.T) {
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("sample.go", []byte(sampleSource())); err != nil {
		t.Fatal(err)
	}
	login := exporter.Classes[0]
	fields := login.ToIR(exporter.Classes)
	if len(fields) != 7 {
		t.Fatalf("fields = %d, want 7", len(fields))
	}
	if fields[0].Kind != KindInt8 || fields[0].JSONName != "a" {
		t.Fatalf("field0 = %+v", fields[0])
	}
	if fields[5].Kind != KindStruct || fields[5].ElemType != "Account" {
		t.Fatalf("field Q = %+v", fields[5])
	}
	if fields[6].Kind != KindSlice || fields[6].ElemKind != KindStruct {
		t.Fatalf("field R = %+v", fields[6])
	}
}

func TestOptionalAndMapIR(t *testing.T) {
	src := `package protocol

type Meta struct {
	Note *string           ` + "`json:\"note\"`" + `
	Tags map[string]string ` + "`json:\"tags\"`" + `
}
`
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("meta.go", []byte(src)); err != nil {
		t.Fatal(err)
	}
	fields := exporter.Classes[0].ToIR(exporter.Classes)
	if len(fields) != 2 {
		t.Fatalf("fields = %d", len(fields))
	}
	if !fields[0].Optional || fields[0].Kind != KindString {
		t.Fatalf("note = %+v", fields[0])
	}
	if fields[1].Kind != KindMap || fields[1].MapKeyKind != KindString || fields[1].MapValKind != KindString {
		t.Fatalf("tags = %+v", fields[1])
	}
	code := exporter.codingToGolang()
	if !strings.Contains(code, "Note *string") {
		t.Fatal("expected *string field in Go export")
	}
	if !strings.Contains(code, "sort.Strings") {
		t.Fatal("expected sorted map keys in Go export")
	}
}

func TestFieldKindBinarySize(t *testing.T) {
	if KindInt16.BinarySize() != 2 || KindFloat64.BinarySize() != 8 || KindString.BinarySize() != 0 {
		t.Fatal("unexpected BinarySize")
	}
}
