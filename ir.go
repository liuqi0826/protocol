package protocol

import (
	"strings"
)

// FieldKind 描述协议字段的中间类型分类，供各语言生成器共用。
type FieldKind int

const (
	KindUnknown FieldKind = iota
	KindInt8
	KindUint8
	KindInt16
	KindUint16
	KindInt32
	KindUint32
	KindInt64
	KindUint64
	KindFloat32
	KindFloat64
	KindBool
	KindByte
	KindString
	KindSlice
	KindStruct
	KindMap
)

// FieldIR 是属性的规范化中间表示。
type FieldIR struct {
	Name     string
	JSONName string
	GoType   string
	Kind     FieldKind
	Optional bool

	ElemKind FieldKind
	ElemType string
	IsCustom bool

	MapKeyKind FieldKind
	MapKeyType string
	MapValKind FieldKind
	MapValType string
}

func (c *Class) ToIR(classes []*Class) []*FieldIR {
	out := make([]*FieldIR, 0, len(c.Attributes))
	for _, attr := range c.Attributes {
		out = append(out, attributeToIR(attr, classes))
	}
	return out
}

func attributeToIR(attr *Attribute, classes []*Class) *FieldIR {
	ir := &FieldIR{
		Name:     attr.Name,
		JSONName: getValueNameFromLable(attr.Label),
		GoType:   attr.Type,
		Optional: attr.Optional,
	}
	ir.Kind, ir.ElemType, ir.ElemKind, ir.IsCustom = classifyType(attr.Type, classes)
	if ir.Kind == KindMap {
		ir.MapKeyType, ir.MapValType = parseMapTypes(attr.Type)
		ir.MapKeyKind, _, _, _ = classifyType(ir.MapKeyType, classes)
		ir.MapValKind, _, _, ir.IsCustom = classifyType(ir.MapValType, classes)
	}
	return ir
}

func classifyType(goType string, classes []*Class) (kind FieldKind, elemType string, elemKind FieldKind, isCustom bool) {
	goType = strings.TrimSpace(goType)
	switch goType {
	case "int8":
		return KindInt8, "", KindUnknown, false
	case "uint8":
		return KindUint8, "", KindUnknown, false
	case "int16":
		return KindInt16, "", KindUnknown, false
	case "uint16":
		return KindUint16, "", KindUnknown, false
	case "int32":
		return KindInt32, "", KindUnknown, false
	case "uint32":
		return KindUint32, "", KindUnknown, false
	case "int64":
		return KindInt64, "", KindUnknown, false
	case "uint64":
		return KindUint64, "", KindUnknown, false
	case "float32":
		return KindFloat32, "", KindUnknown, false
	case "float64":
		return KindFloat64, "", KindUnknown, false
	case "bool":
		return KindBool, "", KindUnknown, false
	case "byte":
		return KindByte, "", KindUnknown, false
	case "string":
		return KindString, "", KindUnknown, false
	}

	if strings.HasPrefix(goType, "map[") {
		_, val := parseMapTypes(goType)
		return KindMap, val, KindUnknown, isCustomType(val, classes)
	}

	if len(goType) > 2 && goType[:2] == "[]" {
		inner := goType[2:]
		ek, _, _, custom := classifyType(inner, classes)
		return KindSlice, inner, ek, custom
	}

	if isCustomType(goType, classes) {
		return KindStruct, goType, KindUnknown, true
	}
	return KindUnknown, "", KindUnknown, false
}

func parseMapTypes(goType string) (key, val string) {
	if !strings.HasPrefix(goType, "map[") {
		return "", ""
	}
	rest := goType[4:]
	depth := 1
	for i := 0; i < len(rest); i++ {
		switch rest[i] {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return strings.TrimSpace(rest[:i]), strings.TrimSpace(rest[i+1:])
			}
		}
	}
	return "", ""
}

func (k FieldKind) IsInteger() bool {
	switch k {
	case KindInt8, KindUint8, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindByte:
		return true
	}
	return false
}

func (k FieldKind) IsFloat() bool {
	return k == KindFloat32 || k == KindFloat64
}

func (k FieldKind) IsNumber() bool {
	return k.IsInteger() || k.IsFloat()
}

func (k FieldKind) BinarySize() int {
	switch k {
	case KindInt8, KindUint8, KindByte, KindBool:
		return 1
	case KindInt16, KindUint16:
		return 2
	case KindInt32, KindUint32, KindFloat32:
		return 4
	case KindInt64, KindUint64, KindFloat64:
		return 8
	default:
		return 0
	}
}

func (f *FieldIR) TSType() string {
	if f.Kind == KindMap {
		vt := tsTypeName(f.MapValKind, f.MapValType)
		base := "{ [key: string]: " + vt + " }"
		if f.Optional {
			return base + " | null"
		}
		return base
	}
	base := tsTypeName(f.Kind, f.ElemType)
	if f.Kind == KindSlice {
		base = tsTypeName(f.ElemKind, f.ElemType) + "[]"
	}
	if f.Optional {
		return base + " | null"
	}
	return base
}

func tsTypeName(kind FieldKind, elemOrStruct string) string {
	switch kind {
	case KindBool:
		return "boolean"
	case KindString:
		return "string"
	case KindStruct:
		return elemOrStruct
	default:
		if kind.IsNumber() || kind == KindByte {
			return "number"
		}
		return "any"
	}
}

func (f *FieldIR) TSDefault() string {
	if f.Optional {
		return "null"
	}
	switch f.Kind {
	case KindBool:
		return "false"
	case KindString:
		return "\"\""
	case KindSlice:
		return "[]"
	case KindMap:
		return "{}"
	case KindStruct:
		return "new " + f.ElemType + "()"
	default:
		if f.Kind.IsFloat() {
			return "0.0"
		}
		return "0"
	}
}

// DisplayName 返回跨语言字段名（优先 json tag）。
func (f *FieldIR) DisplayName() string {
	if f.JSONName != "" {
		return f.JSONName
	}
	return f.Name
}

// AsAttribute 将 FieldIR 转回 Attribute，供尚未完全按 Kind 分支的生成器过渡使用。
func (f *FieldIR) AsAttribute() *Attribute {
	label := ""
	if f.JSONName != "" {
		label = "`json:\"" + f.JSONName + "\"`"
	}
	return &Attribute{
		Name:     f.Name,
		Type:     f.GoType,
		Label:    label,
		Optional: f.Optional,
	}
}
