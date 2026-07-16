package protocol

import (
	"fmt"
)

func (this *ProtocolExporter) codingToC3() string {
	var code string
	code += "// Auto-generated C3 protocol code (C3 0.8+)\n"
	code += "// Binary: little-endian, self-contained.\n"
	code += "// JSON: encode via json::marshal; decode via json::parse + Object getters.\n"
	code += "module protocol;\n\n"
	code += "import std::io;\n"
	code += "import std::encoding::json;\n"
	code += "import std::collections::object;\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "struct " + class.Name + "\n{\n"
		for _, field := range fields {
			code += "    " + c3FieldType(field) + " " + field.DisplayName() + ";\n"
		}
		code += "}\n\n"

		code += "fn void " + class.Name + ".init(&self)\n{\n"
		for _, field := range fields {
			code += c3InitFieldIR(field)
		}
		code += "}\n\n"

		code += "fn void " + class.Name + ".free(&self)\n{\n"
		for _, field := range fields {
			code += c3FreeFieldIR(field)
		}
		code += "}\n\n"

		code += "fn void " + class.Name + ".decode_from_object(&self, Object* obj)\n{\n"
		for _, field := range fields {
			code += c3DecodeJSONField(field)
		}
		code += "}\n\n"

		code += "fn void " + class.Name + ".decode(&self, char[] data)\n{\n"
		code += "    if (data.len == 0) return;\n"
		code += "    if (data[0] == '{')\n"
		code += "    {\n"
		code += "        String s = (String)data;\n"
		code += "        Object*? parsed = json::tparse(s);\n"
		code += "        if (catch parsed) return;\n"
		code += "        Object* obj = parsed;\n"
		code += "        defer obj.free();\n"
		code += "        self.decode_from_object(obj);\n"
		code += "        return;\n"
		code += "    }\n"
		code += "    usz pointer = 0;\n"
		for _, field := range fields {
			code += c3DecodeBinaryIR(field)
		}
		code += "}\n\n"

		code += "fn String " + class.Name + ".encode_binary(&self, Allocator allocator)\n{\n"
		code += "    DString buffer = dstring::temp();\n"
		for _, field := range fields {
			code += c3EncodeBinaryIR(field)
		}
		code += "    return buffer.copy_str(allocator);\n"
		code += "}\n\n"

		code += "fn String " + class.Name + ".encode_json(&self, Allocator allocator)\n{\n"
		code += "    return json::marshal(allocator, *self);\n"
		code += "}\n\n"
	}

	code += c3HelperFunctions()
	return code
}

func c3FieldType(field *FieldIR) string {
	base := c3TypeFromKind(field.Kind, field.ElemType, field.ElemKind, field)
	// C3 暂无可空指针表示 optional 复杂类型；标量 optional 仍用值类型 + 二进制 present 标志
	return base
}

func c3TypeFromKind(kind FieldKind, elemType string, elemKind FieldKind, field *FieldIR) string {
	switch kind {
	case KindInt8:
		return "ichar"
	case KindUint8, KindByte:
		return "char"
	case KindInt16:
		return "short"
	case KindUint16:
		return "ushort"
	case KindInt32:
		return "int"
	case KindUint32:
		return "uint"
	case KindInt64:
		return "long"
	case KindUint64:
		return "ulong"
	case KindFloat32:
		return "float"
	case KindFloat64:
		return "double"
	case KindBool:
		return "bool"
	case KindString:
		return "String"
	case KindSlice:
		return c3TypeFromKind(elemKind, elemType, KindUnknown, nil) + "[]"
	case KindStruct:
		return elemType
	case KindMap:
		return "Object*" // map 暂用 Object* 承载；二进制仍按协议编解码到临时结构时需扩展
	default:
		return "void"
	}
}

func c3InitFieldIR(field *FieldIR) string {
	name := field.DisplayName()
	switch field.Kind {
	case KindInt8, KindUint8, KindByte, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64:
		return "    self." + name + " = 0;\n"
	case KindFloat32, KindFloat64:
		return "    self." + name + " = 0.0;\n"
	case KindBool:
		return "    self." + name + " = false;\n"
	case KindString:
		return "    self." + name + " = \"\";\n"
	case KindSlice:
		return "    self." + name + " = {};\n"
	case KindStruct:
		return "    self." + name + ".init();\n"
	case KindMap:
		return "    self." + name + " = null;\n"
	}
	return ""
}

func c3FreeFieldIR(field *FieldIR) string {
	name := field.DisplayName()
	switch field.Kind {
	case KindSlice:
		var code string
		if field.ElemKind == KindStruct {
			code += "    foreach (&item : self." + name + ")\n"
			code += "    {\n"
			code += "        item.free();\n"
			code += "    }\n"
		}
		code += "    free(self." + name + ");\n"
		code += "    self." + name + " = {};\n"
		return code
	case KindStruct:
		return "    self." + name + ".free();\n"
	case KindMap:
		return "    if (self." + name + ") self." + name + ".free();\n" +
			"    self." + name + " = null;\n"
	}
	return ""
}

func c3DecodeJSONField(field *FieldIR) string {
	name := field.DisplayName()
	key := name
	switch field.Kind {
	case KindInt8:
		return "    if (try v = obj.get_ichar(\"" + key + "\")) self." + name + " = v;\n"
	case KindUint8, KindByte:
		return "    if (try v = obj.get_char(\"" + key + "\")) self." + name + " = v;\n"
	case KindInt16:
		return "    if (try v = obj.get_short(\"" + key + "\")) self." + name + " = v;\n"
	case KindUint16:
		return "    if (try v = obj.get_ushort(\"" + key + "\")) self." + name + " = v;\n"
	case KindInt32:
		return "    if (try v = obj.get_int(\"" + key + "\")) self." + name + " = v;\n"
	case KindUint32:
		return "    if (try v = obj.get_uint(\"" + key + "\")) self." + name + " = v;\n"
	case KindInt64:
		return "    if (try v = obj.get_long(\"" + key + "\")) self." + name + " = v;\n"
	case KindUint64:
		return "    if (try v = obj.get_ulong(\"" + key + "\")) self." + name + " = v;\n"
	case KindFloat32:
		return "    if (try v = obj.get_float(\"" + key + "\")) self." + name + " = (float)v;\n"
	case KindFloat64:
		return "    if (try v = obj.get_float(\"" + key + "\")) self." + name + " = v;\n"
	case KindBool:
		return "    if (try v = obj.get_bool(\"" + key + "\")) self." + name + " = v;\n"
	case KindString:
		return "    if (try v = obj.get_string(\"" + key + "\")) self." + name + " = v;\n"
	case KindSlice:
		return c3DecodeJSONSlice(field)
	case KindStruct:
		return "    if (try nested = obj.get(\"" + key + "\"))\n" +
			"    {\n" +
			"        self." + name + ".init();\n" +
			"        self." + name + ".decode_from_object(nested);\n" +
			"    }\n"
	case KindMap:
		return "    if (try nested = obj.get(\"" + key + "\")) self." + name + " = nested;\n"
	}
	return "    // unsupported json field " + name + "\n"
}

func c3DecodeJSONSlice(field *FieldIR) string {
	name := field.DisplayName()
	var code string
	code += "    if (try arr = obj.get(\"" + name + "\"))\n"
	code += "    {\n"
	code += "        sz count = arr.get_len();\n"
	code += "        self." + name + " = mem::new_array(" + c3TypeFromKind(field.ElemKind, field.ElemType, KindUnknown, nil) + ", count);\n"
	code += "        for (sz i = 0; i < count; i++)\n"
	code += "        {\n"
	switch field.ElemKind {
	case KindInt8:
		code += "            if (try v = arr.get_ichar_at(i)) self." + name + "[i] = v;\n"
	case KindUint8, KindByte:
		code += "            if (try v = arr.get_char_at(i)) self." + name + "[i] = v;\n"
	case KindInt16:
		code += "            if (try v = arr.get_short_at(i)) self." + name + "[i] = v;\n"
	case KindUint16:
		code += "            if (try v = arr.get_ushort_at(i)) self." + name + "[i] = v;\n"
	case KindInt32:
		code += "            if (try v = arr.get_int_at(i)) self." + name + "[i] = v;\n"
	case KindUint32:
		code += "            if (try v = arr.get_uint_at(i)) self." + name + "[i] = v;\n"
	case KindInt64:
		code += "            if (try v = arr.get_long_at(i)) self." + name + "[i] = v;\n"
	case KindUint64:
		code += "            if (try v = arr.get_ulong_at(i)) self." + name + "[i] = v;\n"
	case KindFloat32:
		code += "            if (try v = arr.get_float_at(i)) self." + name + "[i] = (float)v;\n"
	case KindFloat64:
		code += "            if (try v = arr.get_float_at(i)) self." + name + "[i] = v;\n"
	case KindBool:
		code += "            if (try v = arr.get_bool_at(i)) self." + name + "[i] = v;\n"
	case KindString:
		code += "            if (try v = arr.get_string_at(i)) self." + name + "[i] = v;\n"
	case KindStruct:
		code += "            self." + name + "[i].init();\n"
		code += "            self." + name + "[i].decode_from_object(arr.get_at(i));\n"
	}
	code += "        }\n"
	code += "    }\n"
	return code
}

func c3DecodeBinaryIR(field *FieldIR) string {
	if field.Optional {
		return "    {\n" +
			"        if (pointer + 1 > data.len) return;\n" +
			"        char present = data[pointer];\n" +
			"        pointer += 1;\n" +
			"        if (present == 0) { /* optional absent */ }\n" +
			"        else {\n" +
			c3DecodeBinaryValue(field, "            ") +
			"        }\n" +
			"    }\n"
	}
	return c3DecodeBinaryValue(field, "    ")
}

func c3DecodeBinaryValue(field *FieldIR, indent string) string {
	name := field.DisplayName()
	switch field.Kind {
	case KindInt8:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + " = (ichar)data[pointer];\n" +
			indent + "pointer += 1;\n"
	case KindUint8, KindByte:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + " = data[pointer];\n" +
			indent + "pointer += 1;\n"
	case KindInt16:
		return indent + "if (pointer + 2 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_short(data, pointer);\n" +
			indent + "pointer += 2;\n"
	case KindUint16:
		return indent + "if (pointer + 2 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_ushort(data, pointer);\n" +
			indent + "pointer += 2;\n"
	case KindInt32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_int(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindUint32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_uint(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindInt64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_long(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindUint64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_ulong(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindFloat32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_float(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindFloat64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + " = protocol_read_double(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindBool:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + " = data[pointer] != 0;\n" +
			indent + "pointer += 1;\n"
	case KindString:
		return indent + "{\n" +
			indent + "    if (pointer + 4 > data.len) return;\n" +
			indent + "    int len = protocol_read_int(data, pointer);\n" +
			indent + "    pointer += 4;\n" +
			indent + "    if (len < 0 || pointer + (usz)len > data.len) return;\n" +
			indent + "    self." + name + " = (String)data[pointer:(usz)len];\n" +
			indent + "    pointer += (usz)len;\n" +
			indent + "}\n"
	case KindSlice:
		var code string
		code += indent + "{\n"
		code += indent + "    if (pointer + 4 > data.len) return;\n"
		code += indent + "    int count = protocol_read_int(data, pointer);\n"
		code += indent + "    pointer += 4;\n"
		code += indent + "    if (count < 0) return;\n"
		code += indent + "    self." + name + " = mem::new_array(" + c3TypeFromKind(field.ElemKind, field.ElemType, KindUnknown, nil) + ", (usz)count);\n"
		code += indent + "    for (usz i = 0; i < (usz)count; i++)\n"
		code += indent + "    {\n"
		code += c3DecodeArrayItemIR(field, indent+"        ")
		code += indent + "    }\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		return indent + "{\n" +
			indent + "    if (pointer + 4 > data.len) return;\n" +
			indent + "    int len = protocol_read_int(data, pointer);\n" +
			indent + "    pointer += 4;\n" +
			indent + "    if (len < 0 || pointer + (usz)len > data.len) return;\n" +
			indent + "    self." + name + ".init();\n" +
			indent + "    self." + name + ".decode(data[pointer:(usz)len]);\n" +
			indent + "    pointer += (usz)len;\n" +
			indent + "}\n"
	case KindMap:
		return indent + "// map binary decode not fully supported in C3 export\n"
	}
	return indent + fmt.Sprintf("// unsupported field %s\n", name)
}

func c3DecodeArrayItemIR(field *FieldIR, indent string) string {
	name := field.DisplayName()
	switch field.ElemKind {
	case KindInt8:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + "[i] = (ichar)data[pointer];\n" +
			indent + "pointer += 1;\n"
	case KindUint8, KindByte:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + "[i] = data[pointer];\n" +
			indent + "pointer += 1;\n"
	case KindInt16:
		return indent + "if (pointer + 2 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_short(data, pointer);\n" +
			indent + "pointer += 2;\n"
	case KindUint16:
		return indent + "if (pointer + 2 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_ushort(data, pointer);\n" +
			indent + "pointer += 2;\n"
	case KindInt32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_int(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindUint32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_uint(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindInt64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_long(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindUint64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_ulong(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindFloat32:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_float(data, pointer);\n" +
			indent + "pointer += 4;\n"
	case KindFloat64:
		return indent + "if (pointer + 8 > data.len) return;\n" +
			indent + "self." + name + "[i] = protocol_read_double(data, pointer);\n" +
			indent + "pointer += 8;\n"
	case KindBool:
		return indent + "if (pointer + 1 > data.len) return;\n" +
			indent + "self." + name + "[i] = data[pointer] != 0;\n" +
			indent + "pointer += 1;\n"
	case KindString:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "int len = protocol_read_int(data, pointer);\n" +
			indent + "pointer += 4;\n" +
			indent + "if (len < 0 || pointer + (usz)len > data.len) return;\n" +
			indent + "self." + name + "[i] = (String)data[pointer:(usz)len];\n" +
			indent + "pointer += (usz)len;\n"
	case KindStruct:
		return indent + "if (pointer + 4 > data.len) return;\n" +
			indent + "int len = protocol_read_int(data, pointer);\n" +
			indent + "pointer += 4;\n" +
			indent + "if (len < 0 || pointer + (usz)len > data.len) return;\n" +
			indent + "self." + name + "[i].init();\n" +
			indent + "self." + name + "[i].decode(data[pointer:(usz)len]);\n" +
			indent + "pointer += (usz)len;\n"
	}
	return indent + "// unsupported\n"
}

func c3EncodeBinaryIR(field *FieldIR) string {
	if field.Optional {
		return "    protocol_write_char(&buffer, 1);\n" + c3EncodeBinaryValue(field)
	}
	return c3EncodeBinaryValue(field)
}

func c3EncodeBinaryValue(field *FieldIR) string {
	name := field.DisplayName()
	switch field.Kind {
	case KindInt8:
		return "    protocol_write_char(&buffer, (char)self." + name + ");\n"
	case KindUint8, KindByte:
		return "    protocol_write_char(&buffer, self." + name + ");\n"
	case KindInt16:
		return "    protocol_write_short(&buffer, self." + name + ");\n"
	case KindUint16:
		return "    protocol_write_ushort(&buffer, self." + name + ");\n"
	case KindInt32:
		return "    protocol_write_int(&buffer, self." + name + ");\n"
	case KindUint32:
		return "    protocol_write_uint(&buffer, self." + name + ");\n"
	case KindInt64:
		return "    protocol_write_long(&buffer, self." + name + ");\n"
	case KindUint64:
		return "    protocol_write_ulong(&buffer, self." + name + ");\n"
	case KindFloat32:
		return "    protocol_write_float(&buffer, self." + name + ");\n"
	case KindFloat64:
		return "    protocol_write_double(&buffer, self." + name + ");\n"
	case KindBool:
		return "    protocol_write_char(&buffer, self." + name + " ? 1 : 0);\n"
	case KindString:
		return "    protocol_write_int(&buffer, (int)self." + name + ".len);\n" +
			"    buffer.append(self." + name + ");\n"
	case KindSlice:
		var code string
		code += "    protocol_write_int(&buffer, (int)self." + name + ".len);\n"
		code += "    foreach (item : self." + name + ")\n"
		code += "    {\n"
		code += c3EncodeArrayItemIR(field)
		code += "    }\n"
		return code
	case KindStruct:
		return "    {\n" +
			"        String nested = self." + name + ".encode_binary(allocator);\n" +
			"        protocol_write_int(&buffer, (int)nested.len);\n" +
			"        buffer.append(nested);\n" +
			"        free(nested);\n" +
			"    }\n"
	case KindMap:
		return "    // map binary encode not fully supported in C3 export\n"
	}
	return fmt.Sprintf("    // unsupported field %s\n", name)
}

func c3EncodeArrayItemIR(field *FieldIR) string {
	switch field.ElemKind {
	case KindInt8:
		return "        protocol_write_char(&buffer, (char)item);\n"
	case KindUint8, KindByte:
		return "        protocol_write_char(&buffer, item);\n"
	case KindInt16:
		return "        protocol_write_short(&buffer, item);\n"
	case KindUint16:
		return "        protocol_write_ushort(&buffer, item);\n"
	case KindInt32:
		return "        protocol_write_int(&buffer, item);\n"
	case KindUint32:
		return "        protocol_write_uint(&buffer, item);\n"
	case KindInt64:
		return "        protocol_write_long(&buffer, item);\n"
	case KindUint64:
		return "        protocol_write_ulong(&buffer, item);\n"
	case KindFloat32:
		return "        protocol_write_float(&buffer, item);\n"
	case KindFloat64:
		return "        protocol_write_double(&buffer, item);\n"
	case KindBool:
		return "        protocol_write_char(&buffer, item ? 1 : 0);\n"
	case KindString:
		return "        protocol_write_int(&buffer, (int)item.len);\n" +
			"        buffer.append(item);\n"
	case KindStruct:
		return "        {\n" +
			"            String nested = item.encode_binary(allocator);\n" +
			"            protocol_write_int(&buffer, (int)nested.len);\n" +
			"            buffer.append(nested);\n" +
			"            free(nested);\n" +
			"        }\n"
	}
	return "        // unsupported\n"
}

func c3HelperFunctions() string {
	return `
fn void protocol_write_char(DString* buffer, char value)
{
    buffer.append_char(value);
}

fn void protocol_write_short(DString* buffer, short value)
{
    char[2] bytes = { (char)(value & 0xff), (char)((value >> 8) & 0xff) };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_ushort(DString* buffer, ushort value)
{
    protocol_write_short(buffer, (short)value);
}

fn void protocol_write_int(DString* buffer, int value)
{
    char[4] bytes = {
        (char)(value & 0xff),
        (char)((value >> 8) & 0xff),
        (char)((value >> 16) & 0xff),
        (char)((value >> 24) & 0xff)
    };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_uint(DString* buffer, uint value)
{
    protocol_write_int(buffer, (int)value);
}

fn void protocol_write_long(DString* buffer, long value)
{
    char[8] bytes = {
        (char)(value & 0xff),
        (char)((value >> 8) & 0xff),
        (char)((value >> 16) & 0xff),
        (char)((value >> 24) & 0xff),
        (char)((value >> 32) & 0xff),
        (char)((value >> 40) & 0xff),
        (char)((value >> 48) & 0xff),
        (char)((value >> 56) & 0xff)
    };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_ulong(DString* buffer, ulong value)
{
    protocol_write_long(buffer, (long)value);
}

fn void protocol_write_float(DString* buffer, float value)
{
    uint bits = bitcast(value, uint);
    protocol_write_uint(buffer, bits);
}

fn void protocol_write_double(DString* buffer, double value)
{
    ulong bits = bitcast(value, ulong);
    protocol_write_ulong(buffer, bits);
}

fn short protocol_read_short(char[] data, usz pointer)
{
    return (short)((ushort)data[pointer] | ((ushort)data[pointer + 1] << 8));
}

fn ushort protocol_read_ushort(char[] data, usz pointer)
{
    return (ushort)protocol_read_short(data, pointer);
}

fn int protocol_read_int(char[] data, usz pointer)
{
    return (int)((uint)data[pointer]
        | ((uint)data[pointer + 1] << 8)
        | ((uint)data[pointer + 2] << 16)
        | ((uint)data[pointer + 3] << 24));
}

fn uint protocol_read_uint(char[] data, usz pointer)
{
    return (uint)protocol_read_int(data, pointer);
}

fn long protocol_read_long(char[] data, usz pointer)
{
    return (long)((ulong)data[pointer]
        | ((ulong)data[pointer + 1] << 8)
        | ((ulong)data[pointer + 2] << 16)
        | ((ulong)data[pointer + 3] << 24)
        | ((ulong)data[pointer + 4] << 32)
        | ((ulong)data[pointer + 5] << 40)
        | ((ulong)data[pointer + 6] << 48)
        | ((ulong)data[pointer + 7] << 56));
}

fn ulong protocol_read_ulong(char[] data, usz pointer)
{
    return (ulong)protocol_read_long(data, pointer);
}

fn float protocol_read_float(char[] data, usz pointer)
{
    return bitcast(protocol_read_uint(data, pointer), float);
}

fn double protocol_read_double(char[] data, usz pointer)
{
    return bitcast(protocol_read_ulong(data, pointer), double);
}
`
}
