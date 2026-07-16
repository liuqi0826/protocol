package protocol

import (
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func sampleSource() string {
	return `package protocol

type ProtocolLogin struct {
	A int8      ` + "`json:\"a\"`" + `
	B uint8     ` + "`json:\"b\"`" + `
	M string    ` + "`json:\"m\"`" + `
	N []int8    ` + "`json:\"n\"`" + `
	O []string  ` + "`json:\"o\"`" + `
	Q Account   ` + "`json:\"q\"`" + `
	R []Account ` + "`json:\"r\"`" + `
}

type Account struct {
	Nickname string ` + "`json:\"nickname\"`" + `
	Password string ` + "`json:\"password\"`" + `
}
`
}

func TestLoadSourceParsesStructs(t *testing.T) {
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("sample.go", []byte(sampleSource())); err != nil {
		t.Fatalf("LoadSource: %v", err)
	}

	if exporter.PackageName != "protocol" {
		t.Fatalf("PackageName = %q, want protocol", exporter.PackageName)
	}
	if len(exporter.Classes) != 2 {
		t.Fatalf("Classes len = %d, want 2", len(exporter.Classes))
	}

	login := exporter.Classes[0]
	if login.Name != "ProtocolLogin" {
		t.Fatalf("first class = %q, want ProtocolLogin", login.Name)
	}
	if len(login.Attributes) != 7 {
		t.Fatalf("ProtocolLogin attrs = %d, want 7", len(login.Attributes))
	}

	want := []struct {
		name, typ, json string
	}{
		{"A", "int8", "a"},
		{"B", "uint8", "b"},
		{"M", "string", "m"},
		{"N", "[]int8", "n"},
		{"O", "[]string", "o"},
		{"Q", "Account", "q"},
		{"R", "[]Account", "r"},
	}
	for i, w := range want {
		attr := login.Attributes[i]
		if attr.Name != w.name || attr.Type != w.typ {
			t.Fatalf("attr[%d] = %s %s, want %s %s", i, attr.Name, attr.Type, w.name, w.typ)
		}
		if got := getValueNameFromLable(attr.Label); got != w.json {
			t.Fatalf("json name[%d] = %q, want %q (tag=%s)", i, got, w.json, attr.Label)
		}
	}
}

func TestLoadSourceIgnoresCommentsAndPointers(t *testing.T) {
	src := `package protocol

// type Fake struct { should not parse from comment }
type Account struct {
	Nickname string ` + "`json:\"nickname\"`" + `
	Nested   *Account ` + "`json:\"nested\"`" + `
}
`
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("ptr.go", []byte(src)); err != nil {
		t.Fatalf("LoadSource: %v", err)
	}
	if len(exporter.Classes) != 1 {
		t.Fatalf("Classes len = %d, want 1", len(exporter.Classes))
	}
	attrs := exporter.Classes[0].Attributes
	if len(attrs) != 2 {
		t.Fatalf("attrs = %d, want 2", len(attrs))
	}
	if attrs[1].Type != "Account" {
		t.Fatalf("pointer field type = %q, want Account", attrs[1].Type)
	}
	if !attrs[1].Optional {
		t.Fatal("pointer field should be Optional=true")
	}
}

func TestLoadExampleSourceFile(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	sourcePath := filepath.Join(filepath.Dir(thisFile), "example", "source", "source.go")

	exporter := &ProtocolExporter{}
	if err := exporter.Load(sourcePath); err != nil {
		t.Fatalf("Load(%s): %v", sourcePath, err)
	}
	if len(exporter.Classes) != 5 {
		t.Fatalf("Classes len = %d, want 5", len(exporter.Classes))
	}
	names := make([]string, 0, len(exporter.Classes))
	for _, c := range exporter.Classes {
		names = append(names, c.Name)
	}
	want := []string{"ProtocolLogin", "Account", "ProtocolServerLogin", "ProtocolServerState", "ProtocolServerCommand"}
	for i, name := range want {
		if names[i] != name {
			t.Fatalf("class[%d] = %q, want %q (all=%v)", i, names[i], name, names)
		}
	}
}

func TestCodingToGolangIsGofmtClean(t *testing.T) {
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("sample.go", []byte(sampleSource())); err != nil {
		t.Fatalf("LoadSource: %v", err)
	}

	raw := exporter.codingToGolang()
	formatted, err := format.Source([]byte(raw))
	if err != nil {
		t.Fatalf("format.Source(raw): %v\n----\n%s", err, raw)
	}

	got, err := exporter.formatGolangCode(raw)
	if err != nil {
		t.Fatalf("formatGolangCode: %v", err)
	}
	if got != string(formatted) {
		t.Fatal("formatGolangCode output mismatch")
	}

	// 生成代码必须可被 parser 解析
	if _, err := parser.ParseFile(token.NewFileSet(), "protocol.go", got, 0); err != nil {
		t.Fatalf("generated Go not parseable: %v", err)
	}

	if !strings.Contains(got, "func (this *ProtocolLogin) Decode") {
		t.Fatal("missing ProtocolLogin.Decode")
	}
	if !strings.Contains(got, "Decode(data []byte) error") {
		t.Fatal("Decode should return error")
	}
	if !strings.Contains(got, "unexpected end of data") {
		t.Fatal("Decode should bound-check binary input")
	}
	if !strings.Contains(got, "protocolMaxByteLen") || !strings.Contains(got, "protocolMaxCount") {
		t.Fatal("Decode should define size limits")
	}
	if !strings.Contains(got, "func (this *Account) EncodeBinary") {
		t.Fatal("missing Account.EncodeBinary")
	}
}

func TestExportGoSnapshotFragments(t *testing.T) {
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("sample.go", []byte(sampleSource())); err != nil {
		t.Fatalf("LoadSource: %v", err)
	}

	code, err := exporter.formatGolangCode(exporter.codingToGolang())
	if err != nil {
		t.Fatal(err)
	}

	required := []string{
		"type ProtocolLogin struct {",
		"\tA int8",
		"\tR []Account",
		"if len(data) == 0 {",
		"EncodeBinary()",
		"EncodeJson()",
	}
	for _, frag := range required {
		if !strings.Contains(code, frag) {
			t.Fatalf("generated code missing fragment %q", frag)
		}
	}
}

func TestExportGoRoundTripViaTempDir(t *testing.T) {
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("sample.go", []byte(sampleSource())); err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	out := filepath.Join(dir, "protocol.go")
	if err := exporter.Export(out, GO); err != nil {
		t.Fatalf("Export: %v", err)
	}

	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	formatted, err := format.Source(data)
	if err != nil {
		t.Fatalf("exported file not gofmt-clean: %v", err)
	}
	if string(data) != string(formatted) {
		t.Fatal("exported Go file differs from gofmt output")
	}
}

func TestExportGoRejectsOversizedLength(t *testing.T) {
	src := `package protocol

type Packet struct {
	Msg string ` + "`json:\"msg\"`" + `
}
`
	exporter := &ProtocolExporter{}
	if err := exporter.LoadSource("p.go", []byte(src)); err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	out := filepath.Join(dir, "protocol.go")
	if err := exporter.Export(out, GO); err != nil {
		t.Fatal(err)
	}

	generated, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(generated), "protocolMaxByteLen") {
		t.Fatal("missing max byte len constant")
	}
	if !strings.Contains(string(generated), "exceeds limit") {
		t.Fatal("missing exceeds limit check")
	}
}

func TestGetValueNameFromLable(t *testing.T) {
	cases := map[string]string{
		"`json:\"nickname\"`":        "nickname",
		"`json:\"id,omitempty\"`":    "id",
		"`json:\"-\"`":               "",
		"`xml:\"x\" json:\"token\"`": "token",
	}
	for in, want := range cases {
		if got := getValueNameFromLable(in); got != want {
			t.Fatalf("getValueNameFromLable(%s) = %q, want %q", in, got, want)
		}
	}
}
