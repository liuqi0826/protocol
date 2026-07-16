package protocol

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	GO          = 0
	CSHARP      = 1
	C           = 2
	CPLUSPLUS   = 3
	TYPE_SCRIPT = 4
	JAVA_SCRIPT = 5
	GD_SCRIPT   = 6
	RUST        = 7
	ZIG         = 8
	C3          = 9
)

type Class struct {
	Name       string
	Attributes []*Attribute
	Raw        string
	Hash       string
}

type Attribute struct {
	Name     string
	Type     string
	Label    string
	Optional bool // 来自 *T
}

type ProtocolExporter struct {
	PackageName string
	Classes     []*Class
}

func (this *ProtocolExporter) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return this.LoadSource(path, data)
}

// LoadSource 使用 go/parser 解析 Go 源码中的结构体定义。
func (this *ProtocolExporter) LoadSource(filename string, src []byte) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse %s: %w", filename, err)
	}

	this.PackageName = file.Name.Name
	this.Classes = make([]*Class, 0)

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok || st.Fields == nil {
				continue
			}

			class := &Class{Name: ts.Name.Name}
			var rawBuf bytes.Buffer
			if err := printer.Fprint(&rawBuf, fset, ts); err == nil {
				class.Raw = rawBuf.String()
			}

			for _, field := range st.Fields.List {
				if len(field.Names) == 0 {
					continue // 跳过嵌入字段
				}
				typeStr, optional := normalizeTypeString(exprString(fset, field.Type))
				label := ""
				if field.Tag != nil {
					label = field.Tag.Value
				}
				for _, name := range field.Names {
					class.Attributes = append(class.Attributes, &Attribute{
						Name:     name.Name,
						Type:     typeStr,
						Label:    label,
						Optional: optional,
					})
				}
			}
			this.Classes = append(this.Classes, class)
		}
	}

	return nil
}

func exprString(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, expr); err != nil {
		return ""
	}
	return buf.String()
}

func normalizeTypeString(typeStr string) (string, bool) {
	typeStr = strings.ReplaceAll(typeStr, "interface{}", "any")
	optional := false
	for strings.HasPrefix(typeStr, "*") {
		optional = true
		typeStr = strings.TrimPrefix(typeStr, "*")
		typeStr = strings.TrimSpace(typeStr)
	}
	return strings.TrimSpace(typeStr), optional
}

func getLanguageFileExtension(language int) string {
	switch language {
	case GO:
		return ".go"
	case CSHARP:
		return ".cs"
	case C:
		return ".h"
	case CPLUSPLUS:
		return ".h"
	case TYPE_SCRIPT:
		return ".ts"
	case JAVA_SCRIPT:
		return ".js"
	case GD_SCRIPT:
		return ".gd"
	case RUST:
		return ".rs"
	case ZIG:
		return ".zig"
	case C3:
		return ".c3"
	default:
		return ".txt"
	}
}

func ensureDirectory(dirPath string) error {
	info, err := os.Stat(dirPath)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		dirPath = filepath.Dir(dirPath)
	} else if os.IsNotExist(err) {
		if strings.HasSuffix(dirPath, string(os.PathSeparator)) ||
			strings.HasSuffix(dirPath, "/") ||
			strings.HasSuffix(dirPath, "\\") {
			return os.MkdirAll(dirPath, 0755)
		}
		parentDir := filepath.Dir(dirPath)
		if parentDir != "." && parentDir != dirPath {
			return os.MkdirAll(parentDir, 0755)
		}
	}
	return nil
}

func normalizePath(path string) string {
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))
	path = strings.ReplaceAll(path, "\\", string(os.PathSeparator))
	return filepath.Clean(path)
}

func writeFile(filePath string, content []byte) error {
	filePath = normalizePath(filePath)
	if err := ensureDirectory(filePath); err != nil {
		return err
	}
	return os.WriteFile(filePath, content, 0644)
}

func (this *ProtocolExporter) Export(path string, language int) (err error) {
	path = normalizePath(path)

	switch language {
	case GD_SCRIPT:
		return this.exportGDScript(path)
	case RUST:
		return this.exportRust(path)
	}

	var code string
	switch language {
	case GO:
		code, err = this.formatGolangCode(this.codingToGolang())
		if err != nil {
			return err
		}
	case CSHARP:
		code = this.codingToCSharp()
	case C:
		code = this.codingToC()
	case CPLUSPLUS:
		code = this.codingToCpp()
	case TYPE_SCRIPT:
		code = this.codingToTypeScript()
	case JAVA_SCRIPT:
		code = this.codingToJavaScript()
	case ZIG:
		code = this.codingToZig()
	case C3:
		code = this.codingToC3()
	default:
		return os.ErrInvalid
	}

	if !strings.Contains(filepath.Base(path), ".") {
		path = path + getLanguageFileExtension(language)
	}

	return writeFile(path, []byte(code))
}

func (this *ProtocolExporter) formatGolangCode(code string) (string, error) {
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", fmt.Errorf("gofmt generated Go code: %w", err)
	}
	return string(formatted), nil
}

func (this *ProtocolExporter) exportGDScript(path string) error {
	files := this.codingToGDScriptFiles()

	path = strings.TrimSuffix(path, ".gd")
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if err := ensureDirectory(path); err != nil {
		return err
	}

	for fname, content := range files {
		if err := writeFile(filepath.Join(path, fname), []byte(content)); err != nil {
			return err
		}
	}
	return nil
}

func (this *ProtocolExporter) exportRust(path string) error {
	code := this.codingToRust()

	if !strings.Contains(filepath.Base(path), ".") {
		path = path + ".rs"
	}
	if err := writeFile(path, []byte(code)); err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if dir != "." && dir != path {
		cargoTomlPath := filepath.Join(dir, "Cargo.toml")
		if _, err := os.Stat(cargoTomlPath); os.IsNotExist(err) {
			_ = writeFile(cargoTomlPath, []byte(this.generateCargoToml()))
		}
	}
	return nil
}

func (this *ProtocolExporter) generateCargoToml() string {
	var code string
	code += "[package]\n"
	code += "name = \"" + strings.ToLower(this.PackageName) + "\"\n"
	code += "version = \"0.1.0\"\n"
	code += "edition = \"2021\"\n\n"
	code += "[dependencies]\n"
	code += "serde = { version = \"1.0\", features = [\"derive\"] }\n"
	code += "serde_json = \"1.0\"\n"
	return code
}

// getValueNameFromLable 从 struct tag 中提取 json 字段名。
func getValueNameFromLable(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	tag := reflect.StructTag(strings.Trim(s, "`"))
	if jsonTag := tag.Get("json"); jsonTag != "" {
		name := strings.Split(jsonTag, ",")[0]
		if name == "-" {
			return ""
		}
		if name != "" {
			return name
		}
	}
	l := strings.Index(s, ":")
	if l > -1 {
		l += 2
		r := len(s) - 2
		if l < r {
			return s[l:r]
		}
	}
	return s
}
