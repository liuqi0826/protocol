package protocol

import "fmt"

func (this *ProtocolExporter) codingToGolang() string {
	needSort := false
	needMath := false
	for _, class := range this.Classes {
		for _, field := range class.ToIR(this.Classes) {
			if field.Kind == KindMap {
				needSort = true
			}
			if field.Kind == KindFloat32 || field.Kind == KindFloat64 ||
				field.ElemKind == KindFloat32 || field.ElemKind == KindFloat64 ||
				field.MapValKind == KindFloat32 || field.MapValKind == KindFloat64 {
				needMath = true
			}
		}
	}

	var code string
	code += "package " + this.PackageName + "\n\n"
	code += "import (\n"
	code += "\t\"bytes\"\n"
	code += "\t\"encoding/binary\"\n"
	code += "\t\"encoding/json\"\n"
	code += "\t\"fmt\"\n"
	if needMath {
		code += "\t\"math\"\n"
	}
	if needSort {
		code += "\t\"sort\"\n"
	}
	code += ")\n\n"
	code += "const (\n"
	code += "\tprotocolMaxByteLen = 16 << 20 // 单段字符串/嵌套二进制最大 16MiB\n"
	code += "\tprotocolMaxCount   = 1 << 20  // 数组/map 最大元素数\n"
	code += ")\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)

		code += "type " + class.Name + " struct {\n"
		for _, field := range fields {
			goType := field.GoType
			if field.Optional {
				goType = "*" + goType
			}
			label := ""
			for _, attr := range class.Attributes {
				if attr.Name == field.Name {
					label = attr.Label
					break
				}
			}
			code += "\t" + field.Name + " " + goType + " " + label + "\n"
		}
		code += "}\n\n"

		code += "func (this *" + class.Name + ") Decode(data []byte) error {\n"
		code += "\tif len(data) == 0 {\n"
		code += "\t\treturn nil\n"
		code += "\t}\n"
		code += "\tif data[0] == '{' {\n"
		code += "\t\tif err := json.Unmarshal(data, this); err != nil {\n"
		code += "\t\t\treturn fmt.Errorf(\"json decode: %w\", err)\n"
		code += "\t\t}\n"
		code += "\t\treturn nil\n"
		code += "\t}\n"
		code += "\tvar pointer int\n"
		for _, field := range fields {
			code += golangDecodeField(field, "\t")
		}
		code += "\treturn nil\n"
		code += "}\n\n"

		code += "func (this *" + class.Name + ") EncodeBinary() (data []byte) {\n"
		code += "\tvar buffer = bytes.NewBuffer([]byte{})\n"
		for _, field := range fields {
			code += golangEncodeField(field, "\t")
		}
		code += "\tdata = buffer.Bytes()\n"
		code += "\treturn\n"
		code += "}\n\n"

		code += "func (this *" + class.Name + ") EncodeJson() (data []byte) {\n"
		code += "\tdata, _ = json.Marshal(this)\n"
		code += "\treturn\n"
		code += "}\n\n"
	}

	return code
}

func golangNeed(size int, indent string) string {
	return indent + fmt.Sprintf("if pointer+%d > len(data) {\n", size) +
		indent + "\treturn fmt.Errorf(\"protocol: unexpected end of data at offset %d\", pointer)\n" +
		indent + "}\n"
}

func golangNeedLen(lenExpr string, indent string) string {
	return indent + "if " + lenExpr + " < 0 {\n" +
		indent + "\treturn fmt.Errorf(\"protocol: negative length\")\n" +
		indent + "}\n" +
		indent + "if int(" + lenExpr + ") > protocolMaxByteLen {\n" +
		indent + "\treturn fmt.Errorf(\"protocol: length %d exceeds limit %d\", " + lenExpr + ", protocolMaxByteLen)\n" +
		indent + "}\n" +
		indent + "if pointer+int(" + lenExpr + ") > len(data) {\n" +
		indent + "\treturn fmt.Errorf(\"protocol: unexpected end of data at offset %d\", pointer)\n" +
		indent + "}\n"
}

func golangNeedCount(countExpr string, indent string) string {
	return indent + "if " + countExpr + " < 0 {\n" +
		indent + "\treturn fmt.Errorf(\"protocol: negative count\")\n" +
		indent + "}\n" +
		indent + "if int(" + countExpr + ") > protocolMaxCount {\n" +
		indent + "\treturn fmt.Errorf(\"protocol: count %d exceeds limit %d\", " + countExpr + ", protocolMaxCount)\n" +
		indent + "}\n"
}

func golangDecodeField(field *FieldIR, indent string) string {
	if field.Optional {
		var code string
		code += indent + "{\n"
		code += golangNeed(1, indent+"\t")
		code += indent + "\tpresent := data[pointer]\n"
		code += indent + "\tpointer += 1\n"
		code += indent + "\tif present == 0 {\n"
		code += indent + "\t\tthis." + field.Name + " = nil\n"
		code += indent + "\t} else {\n"
		inner := *field
		inner.Optional = false
		code += golangDecodeInto(inner, indent+"\t\t", "tmp_"+field.Name, true)
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	}
	return golangDecodeInto(*field, indent, "this."+field.Name, false)
}

func golangDecodeInto(field FieldIR, indent, dest string, asPtr bool) string {
	name := field.Name
	switch field.Kind {
	case KindInt8:
		return golangDecodeFixed(indent, 1, asPtr, name, field.GoType, dest,
			"int8(data[pointer])")
	case KindUint8, KindByte:
		return golangDecodeFixed(indent, 1, asPtr, name, field.GoType, dest,
			field.GoType+"(data[pointer])")
	case KindBool:
		return golangDecodeFixed(indent, 1, asPtr, name, field.GoType, dest,
			"data[pointer] != 0")
	case KindInt16:
		return golangDecodeFixed(indent, 2, asPtr, name, field.GoType, dest,
			"int16(binary.LittleEndian.Uint16(data[pointer:pointer+2]))")
	case KindUint16:
		return golangDecodeFixed(indent, 2, asPtr, name, field.GoType, dest,
			"binary.LittleEndian.Uint16(data[pointer:pointer+2])")
	case KindInt32:
		return golangDecodeFixed(indent, 4, asPtr, name, field.GoType, dest,
			"int32(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindUint32:
		return golangDecodeFixed(indent, 4, asPtr, name, field.GoType, dest,
			"binary.LittleEndian.Uint32(data[pointer:pointer+4])")
	case KindFloat32:
		return golangDecodeFixed(indent, 4, asPtr, name, field.GoType, dest,
			"math.Float32frombits(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindInt64:
		return golangDecodeFixed(indent, 8, asPtr, name, field.GoType, dest,
			"int64(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindUint64:
		return golangDecodeFixed(indent, 8, asPtr, name, field.GoType, dest,
			"binary.LittleEndian.Uint64(data[pointer:pointer+8])")
	case KindFloat64:
		return golangDecodeFixed(indent, 8, asPtr, name, field.GoType, dest,
			"math.Float64frombits(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindString:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		if asPtr {
			code += indent + "\tv := string(data[pointer : pointer+int(length)])\n"
			code += indent + "\tpointer += int(length)\n"
			code += indent + "\tthis." + name + " = &v\n"
		} else {
			code += indent + "\t" + dest + " = string(data[pointer : pointer+int(length)])\n"
			code += indent + "\tpointer += int(length)\n"
		}
		code += indent + "}\n"
		return code
	case KindSlice:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tcount := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedCount("count", indent+"\t")
		if asPtr {
			code += indent + "\ttmp := make(" + field.GoType + ", 0)\n"
			code += indent + "\tfor i := 0; i < int(count); i++ {\n"
			code += golangDecodeSliceItemAppend(&field, indent+"\t\t", "tmp")
			code += indent + "\t}\n"
			code += indent + "\tthis." + name + " = &tmp\n"
		} else {
			code += indent + "\t" + dest + " = " + dest + "[:0]\n"
			code += indent + "\tfor i := 0; i < int(count); i++ {\n"
			code += golangDecodeSliceItemAppend(&field, indent+"\t\t", dest)
			code += indent + "\t}\n"
		}
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		if asPtr {
			code += indent + "\tv := " + field.ElemType + "{}\n"
			code += indent + "\tif err := v.Decode(data[pointer : pointer+int(length)]); err != nil {\n"
			code += indent + "\t\treturn err\n"
			code += indent + "\t}\n"
			code += indent + "\tpointer += int(length)\n"
			code += indent + "\tthis." + name + " = &v\n"
		} else {
			code += indent + "\t" + dest + " = " + field.ElemType + "{}\n"
			code += indent + "\tif err := " + dest + ".Decode(data[pointer : pointer+int(length)]); err != nil {\n"
			code += indent + "\t\treturn err\n"
			code += indent + "\t}\n"
			code += indent + "\tpointer += int(length)\n"
		}
		code += indent + "}\n"
		return code
	case KindMap:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tcount := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedCount("count", indent+"\t")
		mapType := field.GoType
		if asPtr {
			code += indent + "\ttmp := make(" + mapType + ")\n"
		} else {
			code += indent + "\tthis." + name + " = make(" + mapType + ")\n"
		}
		code += indent + "\tfor i := 0; i < int(count); i++ {\n"
		code += golangNeed(4, indent+"\t\t")
		code += indent + "\t\tkeyLen := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\t\tpointer += 4\n"
		code += golangNeedLen("keyLen", indent+"\t\t")
		code += indent + "\t\tkey := string(data[pointer : pointer+int(keyLen)])\n"
		code += indent + "\t\tpointer += int(keyLen)\n"
		code += golangDecodeMapValue(&field, indent+"\t\t", asPtr)
		code += indent + "\t}\n"
		if asPtr {
			code += indent + "\tthis." + name + " = &tmp\n"
		}
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("return fmt.Errorf(\"protocol: unsupported field %s %s\")\n", name, field.GoType)
	}
}

func golangDecodeFixed(indent string, size int, asPtr bool, name, goType, dest, expr string) string {
	var code string
	code += golangNeed(size, indent)
	if asPtr {
		code += indent + "{\n"
		code += indent + "\tv := " + expr + "\n"
		code += indent + fmt.Sprintf("\tpointer += %d\n", size)
		code += indent + "\tthis." + name + " = &v\n"
		code += indent + "}\n"
		_ = goType
		_ = dest
		return code
	}
	code += indent + dest + " = " + expr + "\n"
	code += indent + fmt.Sprintf("pointer += %d\n", size)
	return code
}

func golangDecodeMapValue(field *FieldIR, indent string, asPtr bool) string {
	target := "this." + field.Name
	if asPtr {
		target = "tmp"
	}
	switch field.MapValKind {
	case KindInt8:
		return golangDecodeMapFixed(indent, target, 1, "int8(data[pointer])")
	case KindUint8, KindByte:
		return golangDecodeMapFixed(indent, target, 1, field.MapValType+"(data[pointer])")
	case KindBool:
		return golangDecodeMapFixed(indent, target, 1, "data[pointer] != 0")
	case KindInt16:
		return golangDecodeMapFixed(indent, target, 2, "int16(binary.LittleEndian.Uint16(data[pointer:pointer+2]))")
	case KindUint16:
		return golangDecodeMapFixed(indent, target, 2, "binary.LittleEndian.Uint16(data[pointer:pointer+2])")
	case KindInt32:
		return golangDecodeMapFixed(indent, target, 4, "int32(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindUint32:
		return golangDecodeMapFixed(indent, target, 4, "binary.LittleEndian.Uint32(data[pointer:pointer+4])")
	case KindFloat32:
		return golangDecodeMapFixed(indent, target, 4, "math.Float32frombits(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindInt64:
		return golangDecodeMapFixed(indent, target, 8, "int64(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindUint64:
		return golangDecodeMapFixed(indent, target, 8, "binary.LittleEndian.Uint64(data[pointer:pointer+8])")
	case KindFloat64:
		return golangDecodeMapFixed(indent, target, 8, "math.Float64frombits(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindString:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		code += indent + "\t" + target + "[key] = string(data[pointer : pointer+int(length)])\n"
		code += indent + "\tpointer += int(length)\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		code += indent + "\tvar value = " + field.MapValType + "{}\n"
		code += indent + "\tif err := value.Decode(data[pointer : pointer+int(length)]); err != nil {\n"
		code += indent + "\t\treturn err\n"
		code += indent + "\t}\n"
		code += indent + "\t" + target + "[key] = value\n"
		code += indent + "\tpointer += int(length)\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("return fmt.Errorf(\"protocol: unsupported map value %s\")\n", field.MapValType)
	}
}

func golangDecodeMapFixed(indent, target string, size int, expr string) string {
	return golangNeed(size, indent) +
		indent + target + "[key] = " + expr + "\n" +
		indent + fmt.Sprintf("pointer += %d\n", size)
}

func golangDecodeSliceItemAppend(field *FieldIR, indent, dest string) string {
	switch field.ElemKind {
	case KindInt8:
		return golangDecodeSliceFixed(indent, dest, 1, "int8(data[pointer])")
	case KindUint8, KindByte:
		return golangDecodeSliceFixed(indent, dest, 1, field.ElemType+"(data[pointer])")
	case KindBool:
		return golangDecodeSliceFixed(indent, dest, 1, "data[pointer] != 0")
	case KindInt16:
		return golangDecodeSliceFixed(indent, dest, 2, "int16(binary.LittleEndian.Uint16(data[pointer:pointer+2]))")
	case KindUint16:
		return golangDecodeSliceFixed(indent, dest, 2, "binary.LittleEndian.Uint16(data[pointer:pointer+2])")
	case KindInt32:
		return golangDecodeSliceFixed(indent, dest, 4, "int32(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindUint32:
		return golangDecodeSliceFixed(indent, dest, 4, "binary.LittleEndian.Uint32(data[pointer:pointer+4])")
	case KindFloat32:
		return golangDecodeSliceFixed(indent, dest, 4, "math.Float32frombits(binary.LittleEndian.Uint32(data[pointer:pointer+4]))")
	case KindInt64:
		return golangDecodeSliceFixed(indent, dest, 8, "int64(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindUint64:
		return golangDecodeSliceFixed(indent, dest, 8, "binary.LittleEndian.Uint64(data[pointer:pointer+8])")
	case KindFloat64:
		return golangDecodeSliceFixed(indent, dest, 8, "math.Float64frombits(binary.LittleEndian.Uint64(data[pointer:pointer+8]))")
	case KindString:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		code += indent + "\t" + dest + " = append(" + dest + ", string(data[pointer : pointer+int(length)]))\n"
		code += indent + "\tpointer += int(length)\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		code += golangNeed(4, indent)
		code += indent + "{\n"
		code += indent + "\tlength := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))\n"
		code += indent + "\tpointer += 4\n"
		code += golangNeedLen("length", indent+"\t")
		code += indent + "\tvar value = " + field.ElemType + "{}\n"
		code += indent + "\tif err := value.Decode(data[pointer : pointer+int(length)]); err != nil {\n"
		code += indent + "\t\treturn err\n"
		code += indent + "\t}\n"
		code += indent + "\t" + dest + " = append(" + dest + ", value)\n"
		code += indent + "\tpointer += int(length)\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("return fmt.Errorf(\"protocol: unsupported slice elem %s\")\n", field.ElemType)
	}
}

func golangDecodeSliceFixed(indent, dest string, size int, expr string) string {
	return golangNeed(size, indent) +
		indent + dest + " = append(" + dest + ", " + expr + ")\n" +
		indent + fmt.Sprintf("pointer += %d\n", size)
}

func golangEncodeField(field *FieldIR, indent string) string {
	name := field.Name
	if field.Optional {
		var code string
		code += indent + "if this." + name + " == nil {\n"
		code += indent + "\tbinary.Write(buffer, binary.LittleEndian, uint8(0))\n"
		code += indent + "} else {\n"
		code += indent + "\tbinary.Write(buffer, binary.LittleEndian, uint8(1))\n"
		inner := *field
		inner.Optional = false
		code += golangEncodeValue(inner, indent+"\t", "(*this."+name+")")
		code += indent + "}\n"
		return code
	}
	return golangEncodeValue(*field, indent, "this."+name)
}

func golangEncodeValue(field FieldIR, indent, expr string) string {
	name := field.Name
	switch field.Kind {
	case KindInt8, KindUint8, KindByte, KindBool, KindInt16, KindUint16,
		KindInt32, KindUint32, KindFloat32, KindInt64, KindUint64, KindFloat64:
		return indent + "binary.Write(buffer, binary.LittleEndian, " + expr + ")\n"
	case KindString:
		return indent + "{\n" +
			indent + "\tvar b = []byte(" + expr + ")\n" +
			indent + "\tvar l = int32(len(b))\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, b)\n" +
			indent + "}\n"
	case KindSlice:
		var code string
		code += indent + "{\n"
		code += indent + "\tvar count = int32(len(" + expr + "))\n"
		code += indent + "\tbinary.Write(buffer, binary.LittleEndian, count)\n"
		code += indent + "\tfor _, item := range " + expr + " {\n"
		code += golangEncodeSliceItem(&field, indent+"\t\t")
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		return indent + "{\n" +
			indent + "\tvar b = (" + expr + ").EncodeBinary()\n" +
			indent + "\tvar l = int32(len(b))\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, b)\n" +
			indent + "}\n"
	case KindMap:
		var code string
		code += indent + "{\n"
		code += indent + "\tvar keys = make([]string, 0, len(" + expr + "))\n"
		code += indent + "\tfor k := range " + expr + " {\n"
		code += indent + "\t\tkeys = append(keys, k)\n"
		code += indent + "\t}\n"
		code += indent + "\tsort.Strings(keys)\n"
		code += indent + "\tvar count = int32(len(keys))\n"
		code += indent + "\tbinary.Write(buffer, binary.LittleEndian, count)\n"
		code += indent + "\tfor _, key := range keys {\n"
		code += indent + "\t\tvar kb = []byte(key)\n" +
			indent + "\t\tvar kl = int32(len(kb))\n" +
			indent + "\t\tbinary.Write(buffer, binary.LittleEndian, kl)\n" +
			indent + "\t\tbinary.Write(buffer, binary.LittleEndian, kb)\n"
		code += golangEncodeMapValue(&field, indent+"\t\t", expr+"[key]")
		code += indent + "\t}\n"
		code += indent + "}\n"
		_ = name
		return code
	default:
		return indent + fmt.Sprintf("// unsupported field %s %s\n", name, field.GoType)
	}
}

func golangEncodeMapValue(field *FieldIR, indent, expr string) string {
	switch field.MapValKind {
	case KindInt8, KindUint8, KindByte, KindBool, KindInt16, KindUint16,
		KindInt32, KindUint32, KindFloat32, KindInt64, KindUint64, KindFloat64:
		return indent + "binary.Write(buffer, binary.LittleEndian, " + expr + ")\n"
	case KindString:
		return indent + "{\n" +
			indent + "\tvar b = []byte(" + expr + ")\n" +
			indent + "\tvar l = int32(len(b))\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, b)\n" +
			indent + "}\n"
	case KindStruct:
		return indent + "{\n" +
			indent + "\tvar b = (" + expr + ").EncodeBinary()\n" +
			indent + "\tvar l = int32(len(b))\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "\tbinary.Write(buffer, binary.LittleEndian, b)\n" +
			indent + "}\n"
	default:
		return indent + fmt.Sprintf("// unsupported map value %s\n", field.MapValType)
	}
}

func golangEncodeSliceItem(field *FieldIR, indent string) string {
	switch field.ElemKind {
	case KindInt8, KindUint8, KindByte, KindBool, KindInt16, KindUint16,
		KindInt32, KindUint32, KindFloat32, KindInt64, KindUint64, KindFloat64:
		return indent + "binary.Write(buffer, binary.LittleEndian, item)\n"
	case KindString:
		return indent + "var b = []byte(item)\n" +
			indent + "var l = int32(len(b))\n" +
			indent + "binary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "binary.Write(buffer, binary.LittleEndian, b)\n"
	case KindStruct:
		return indent + "var b = item.EncodeBinary()\n" +
			indent + "var l = int32(len(b))\n" +
			indent + "binary.Write(buffer, binary.LittleEndian, l)\n" +
			indent + "binary.Write(buffer, binary.LittleEndian, b)\n"
	default:
		return indent + fmt.Sprintf("// unsupported slice elem %s\n", field.ElemType)
	}
}
