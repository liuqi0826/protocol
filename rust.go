package protocol

import "fmt"

func (this *ProtocolExporter) codingToRust() string {
	var code string
	code += "// Auto-generated Rust protocol code\n"
	code += "use serde::{Deserialize, Serialize};\n"
	code += "use std::collections::BTreeMap;\n"
	code += "use std::io::{Read, Write};\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "#[derive(Debug, Clone, Serialize, Deserialize)]\n"
		code += "pub struct " + class.Name + " {\n"
		for _, field := range fields {
			code += "    pub " + field.DisplayName() + ": " + rustType(field) + ",\n"
		}
		code += "}\n\n"

		code += "impl " + class.Name + " {\n"
		code += "    pub fn new() -> Self {\n"
		code += "        " + class.Name + " {\n"
		for _, field := range fields {
			code += "            " + field.DisplayName() + ": " + rustDefault(field) + ",\n"
		}
		code += "        }\n"
		code += "    }\n\n"

		code += "    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {\n"
		code += "        if data.is_empty() {\n"
		code += "            return Err(\"Empty data\".to_string());\n"
		code += "        }\n"
		code += "        if data[0] == b'{' {\n"
		code += "            match serde_json::from_slice::<" + class.Name + ">(data) {\n"
		code += "                Ok(obj) => {\n"
		code += "                    *self = obj;\n"
		code += "                    Ok(())\n"
		code += "                },\n"
		code += "                Err(e) => Err(format!(\"JSON parse error: {}\", e)),\n"
		code += "            }\n"
		code += "        } else {\n"
		code += "            let mut cursor = std::io::Cursor::new(data);\n"
		for _, field := range fields {
			code += rustDecodeField(field)
		}
		code += "            Ok(())\n"
		code += "        }\n"
		code += "    }\n\n"

		code += "    pub fn encode_binary(&self) -> Vec<u8> {\n"
		code += "        let mut buffer = Vec::new();\n"
		for _, field := range fields {
			code += rustEncodeField(field)
		}
		code += "        buffer\n"
		code += "    }\n\n"

		code += "    pub fn encode_json(&self) -> Result<String, String> {\n"
		code += "        match serde_json::to_string(self) {\n"
		code += "            Ok(s) => Ok(s),\n"
		code += "            Err(e) => Err(format!(\"JSON encode error: {}\", e)),\n"
		code += "        }\n"
		code += "    }\n"
		code += "}\n\n"

		code += "impl Default for " + class.Name + " {\n"
		code += "    fn default() -> Self {\n"
		code += "        Self::new()\n"
		code += "    }\n"
		code += "}\n\n"
	}
	return code
}

func rustType(field *FieldIR) string {
	base := rustKindType(field.Kind, field.ElemType, field.ElemKind, field)
	if field.Optional {
		return "Option<" + base + ">"
	}
	return base
}

func rustKindType(kind FieldKind, elemType string, elemKind FieldKind, field *FieldIR) string {
	switch kind {
	case KindInt8:
		return "i8"
	case KindUint8, KindByte:
		return "u8"
	case KindInt16:
		return "i16"
	case KindUint16:
		return "u16"
	case KindInt32:
		return "i32"
	case KindUint32:
		return "u32"
	case KindInt64:
		return "i64"
	case KindUint64:
		return "u64"
	case KindFloat32:
		return "f32"
	case KindFloat64:
		return "f64"
	case KindBool:
		return "bool"
	case KindString:
		return "String"
	case KindSlice:
		inner := rustKindType(elemKind, elemType, KindUnknown, nil)
		if elemKind == KindStruct {
			inner = elemType
		}
		return "Vec<" + inner + ">"
	case KindStruct:
		return elemType
	case KindMap:
		vt := rustKindType(field.MapValKind, field.MapValType, KindUnknown, nil)
		if field.MapValKind == KindStruct {
			vt = field.MapValType
		}
		return "BTreeMap<String, " + vt + ">"
	default:
		return "()"
	}
}

func rustDefault(field *FieldIR) string {
	if field.Optional {
		return "None"
	}
	switch field.Kind {
	case KindBool:
		return "false"
	case KindString:
		return "String::new()"
	case KindSlice:
		return "Vec::new()"
	case KindMap:
		return "BTreeMap::new()"
	case KindStruct:
		return field.ElemType + "::new()"
	default:
		if field.Kind.IsFloat() {
			return "0.0"
		}
		return "0"
	}
}

func rustDecodeField(field *FieldIR) string {
	name := field.DisplayName()
	indent := "            "
	if field.Optional {
		var code string
		code += indent + "self." + name + " = {\n"
		code += indent + "    let mut buf = [0u8; 1];\n"
		code += indent + "    if cursor.read_exact(&mut buf).is_err() { return Err(\"Read optional flag failed\".to_string()); }\n"
		code += indent + "    if buf[0] == 0 { None } else {\n"
		code += indent + "        Some({\n"
		inner := *field
		inner.Optional = false
		code += rustDecodeValueExpr(&inner, indent+"            ")
		code += indent + "        })\n"
		code += indent + "    }\n"
		code += indent + "};\n"
		return code
	}
	return indent + "self." + name + " = {\n" + rustDecodeValueExpr(field, indent+"    ") + indent + "};\n"
}

func rustDecodeValueExpr(field *FieldIR, indent string) string {
	switch field.Kind {
	case KindInt8:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read int8 failed\".to_string()); }\n" +
			indent + "buf[0] as i8\n"
	case KindUint8, KindByte:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read uint8 failed\".to_string()); }\n" +
			indent + "buf[0]\n"
	case KindInt16:
		return rustReadFixed(indent, 2, "i16")
	case KindUint16:
		return rustReadFixed(indent, 2, "u16")
	case KindInt32:
		return rustReadFixed(indent, 4, "i32")
	case KindUint32:
		return rustReadFixed(indent, 4, "u32")
	case KindInt64:
		return rustReadFixed(indent, 8, "i64")
	case KindUint64:
		return rustReadFixed(indent, 8, "u64")
	case KindFloat32:
		return rustReadFixed(indent, 4, "f32")
	case KindFloat64:
		return rustReadFixed(indent, 8, "f64")
	case KindBool:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read bool failed\".to_string()); }\n" +
			indent + "buf[0] != 0\n"
	case KindString:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read string length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut str_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut str_buf).is_err() { return Err(\"Read string data failed\".to_string()); }\n" +
			indent + "String::from_utf8(str_buf).map_err(|e| format!(\"Invalid UTF-8: {}\", e))?\n"
	case KindSlice:
		var code string
		code += indent + "let mut count_buf = [0u8; 4];\n"
		code += indent + "if cursor.read_exact(&mut count_buf).is_err() { return Err(\"Read array count failed\".to_string()); }\n"
		code += indent + "let count = i32::from_le_bytes(count_buf) as usize;\n"
		code += indent + "let mut vec = Vec::with_capacity(count);\n"
		code += indent + "for _ in 0..count {\n"
		code += rustDecodeSliceItem(field, indent+"    ")
		code += indent + "}\n"
		code += indent + "vec\n"
		return code
	case KindStruct:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read object length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut obj_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut obj_buf).is_err() { return Err(\"Read object data failed\".to_string()); }\n" +
			indent + "let mut obj = " + field.ElemType + "::new();\n" +
			indent + "obj.decode(&obj_buf)?;\n" +
			indent + "obj\n"
	case KindMap:
		var code string
		code += indent + "let mut count_buf = [0u8; 4];\n"
		code += indent + "if cursor.read_exact(&mut count_buf).is_err() { return Err(\"Read map count failed\".to_string()); }\n"
		code += indent + "let count = i32::from_le_bytes(count_buf) as usize;\n"
		code += indent + "let mut map = BTreeMap::new();\n"
		code += indent + "for _ in 0..count {\n"
		code += indent + "    let mut len_buf = [0u8; 4];\n"
		code += indent + "    if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read map key length failed\".to_string()); }\n"
		code += indent + "    let len = i32::from_le_bytes(len_buf) as usize;\n"
		code += indent + "    let mut key_buf = vec![0u8; len];\n"
		code += indent + "    if cursor.read_exact(&mut key_buf).is_err() { return Err(\"Read map key failed\".to_string()); }\n"
		code += indent + "    let key = String::from_utf8(key_buf).map_err(|e| format!(\"Invalid UTF-8: {}\", e))?;\n"
		code += rustDecodeMapValue(field, indent+"    ")
		code += indent + "}\n"
		code += indent + "map\n"
		return code
	default:
		return indent + "return Err(\"unsupported\".to_string());\n"
	}
}

func rustReadFixed(indent string, size int, typ string) string {
	return indent + fmt.Sprintf("let mut buf = [0u8; %d];\n", size) +
		indent + fmt.Sprintf("if cursor.read_exact(&mut buf).is_err() { return Err(\"Read %s failed\".to_string()); }\n", typ) +
		indent + typ + "::from_le_bytes(buf)\n"
}

func rustDecodeSliceItem(field *FieldIR, indent string) string {
	switch field.ElemKind {
	case KindInt8:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read int8 failed\".to_string()); }\n" +
			indent + "vec.push(buf[0] as i8);\n"
	case KindUint8, KindByte:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read uint8 failed\".to_string()); }\n" +
			indent + "vec.push(buf[0]);\n"
	case KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindFloat32, KindFloat64:
		typ := rustKindType(field.ElemKind, "", KindUnknown, nil)
		size := field.ElemKind.BinarySize()
		return indent + fmt.Sprintf("let mut buf = [0u8; %d];\n", size) +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read failed\".to_string()); }\n" +
			indent + "vec.push(" + typ + "::from_le_bytes(buf));\n"
	case KindBool:
		return indent + "let mut buf = [0u8; 1];\n" +
			indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read bool failed\".to_string()); }\n" +
			indent + "vec.push(buf[0] != 0);\n"
	case KindString:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read string length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut str_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut str_buf).is_err() { return Err(\"Read string data failed\".to_string()); }\n" +
			indent + "vec.push(String::from_utf8(str_buf).map_err(|e| format!(\"Invalid UTF-8: {}\", e))?);\n"
	case KindStruct:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read object length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut obj_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut obj_buf).is_err() { return Err(\"Read object data failed\".to_string()); }\n" +
			indent + "let mut obj = " + field.ElemType + "::new();\n" +
			indent + "obj.decode(&obj_buf)?;\n" +
			indent + "vec.push(obj);\n"
	default:
		return indent + "// unsupported\n"
	}
}

func rustDecodeMapValue(field *FieldIR, indent string) string {
	switch field.MapValKind {
	case KindString:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read map value length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut str_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut str_buf).is_err() { return Err(\"Read map value failed\".to_string()); }\n" +
			indent + "map.insert(key, String::from_utf8(str_buf).map_err(|e| format!(\"Invalid UTF-8: {}\", e))?);\n"
	case KindStruct:
		return indent + "let mut len_buf = [0u8; 4];\n" +
			indent + "if cursor.read_exact(&mut len_buf).is_err() { return Err(\"Read map value length failed\".to_string()); }\n" +
			indent + "let len = i32::from_le_bytes(len_buf) as usize;\n" +
			indent + "let mut obj_buf = vec![0u8; len];\n" +
			indent + "if cursor.read_exact(&mut obj_buf).is_err() { return Err(\"Read map value failed\".to_string()); }\n" +
			indent + "let mut obj = " + field.MapValType + "::new();\n" +
			indent + "obj.decode(&obj_buf)?;\n" +
			indent + "map.insert(key, obj);\n"
	default:
		if field.MapValKind.IsNumber() || field.MapValKind == KindByte || field.MapValKind == KindBool {
			typ := rustKindType(field.MapValKind, "", KindUnknown, nil)
			if field.MapValKind == KindBool {
				return indent + "let mut buf = [0u8; 1];\n" +
					indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read failed\".to_string()); }\n" +
					indent + "map.insert(key, buf[0] != 0);\n"
			}
			if field.MapValKind == KindInt8 {
				return indent + "let mut buf = [0u8; 1];\n" +
					indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read failed\".to_string()); }\n" +
					indent + "map.insert(key, buf[0] as i8);\n"
			}
			if field.MapValKind == KindUint8 || field.MapValKind == KindByte {
				return indent + "let mut buf = [0u8; 1];\n" +
					indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read failed\".to_string()); }\n" +
					indent + "map.insert(key, buf[0]);\n"
			}
			size := field.MapValKind.BinarySize()
			return indent + fmt.Sprintf("let mut buf = [0u8; %d];\n", size) +
				indent + "if cursor.read_exact(&mut buf).is_err() { return Err(\"Read failed\".to_string()); }\n" +
				indent + "map.insert(key, " + typ + "::from_le_bytes(buf));\n"
		}
		return indent + "// unsupported map value\n"
	}
}

func rustEncodeField(field *FieldIR) string {
	name := field.DisplayName()
	indent := "        "
	if field.Optional {
		var code string
		code += indent + "match &self." + name + " {\n"
		code += indent + "    None => buffer.push(0),\n"
		code += indent + "    Some(v) => {\n"
		code += indent + "        buffer.push(1);\n"
		inner := *field
		inner.Optional = false
		code += rustEncodeExpr(&inner, indent+"        ", "v")
		code += indent + "    }\n"
		code += indent + "}\n"
		return code
	}
	return rustEncodeExpr(field, indent, "&self."+name)
}

func rustEncodeExpr(field *FieldIR, indent, expr string) string {
	switch field.Kind {
	case KindInt8:
		return indent + "buffer.push(*" + expr + " as u8);\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(*" + expr + ");\n"
	case KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindFloat32, KindFloat64:
		return indent + "buffer.extend_from_slice(&(*" + expr + ").to_le_bytes());\n"
	case KindBool:
		return indent + "buffer.push(if *" + expr + " { 1 } else { 0 });\n"
	case KindString:
		return indent + "{\n" +
			indent + "    let len = " + expr + ".len() as i32;\n" +
			indent + "    buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "    buffer.extend_from_slice(" + expr + ".as_bytes());\n" +
			indent + "}\n"
	case KindSlice:
		var code string
		code += indent + "{\n"
		code += indent + "    let count = " + expr + ".len() as i32;\n"
		code += indent + "    buffer.extend_from_slice(&count.to_le_bytes());\n"
		code += indent + "    for item in " + expr + " {\n"
		code += rustEncodeSliceItem(field, indent+"        ")
		code += indent + "    }\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		return indent + "{\n" +
			indent + "    let data = " + expr + ".encode_binary();\n" +
			indent + "    let len = data.len() as i32;\n" +
			indent + "    buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "    buffer.extend_from_slice(&data);\n" +
			indent + "}\n"
	case KindMap:
		var code string
		code += indent + "{\n"
		code += indent + "    let count = " + expr + ".len() as i32;\n"
		code += indent + "    buffer.extend_from_slice(&count.to_le_bytes());\n"
		code += indent + "    for (key, value) in " + expr + " {\n"
		code += indent + "        let len = key.len() as i32;\n"
		code += indent + "        buffer.extend_from_slice(&len.to_le_bytes());\n"
		code += indent + "        buffer.extend_from_slice(key.as_bytes());\n"
		code += rustEncodeMapValue(field, indent+"        ")
		code += indent + "    }\n"
		code += indent + "}\n"
		return code
	default:
		return indent + "// unsupported\n"
	}
}

func rustEncodeSliceItem(field *FieldIR, indent string) string {
	switch field.ElemKind {
	case KindInt8:
		return indent + "buffer.push(*item as u8);\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(*item);\n"
	case KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindFloat32, KindFloat64:
		return indent + "buffer.extend_from_slice(&item.to_le_bytes());\n"
	case KindBool:
		return indent + "buffer.push(if *item { 1 } else { 0 });\n"
	case KindString:
		return indent + "let len = item.len() as i32;\n" +
			indent + "buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "buffer.extend_from_slice(item.as_bytes());\n"
	case KindStruct:
		return indent + "let data = item.encode_binary();\n" +
			indent + "let len = data.len() as i32;\n" +
			indent + "buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "buffer.extend_from_slice(&data);\n"
	default:
		return indent + "// unsupported\n"
	}
}

func rustEncodeMapValue(field *FieldIR, indent string) string {
	switch field.MapValKind {
	case KindString:
		return indent + "let len = value.len() as i32;\n" +
			indent + "buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "buffer.extend_from_slice(value.as_bytes());\n"
	case KindStruct:
		return indent + "let data = value.encode_binary();\n" +
			indent + "let len = data.len() as i32;\n" +
			indent + "buffer.extend_from_slice(&len.to_le_bytes());\n" +
			indent + "buffer.extend_from_slice(&data);\n"
	case KindBool:
		return indent + "buffer.push(if *value { 1 } else { 0 });\n"
	case KindInt8:
		return indent + "buffer.push(*value as u8);\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(*value);\n"
	default:
		if field.MapValKind.IsNumber() {
			return indent + "buffer.extend_from_slice(&value.to_le_bytes());\n"
		}
		return indent + "// unsupported\n"
	}
}
