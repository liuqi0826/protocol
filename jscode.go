package protocol

import "fmt"

// jsDecodeBinaryIR ?? JS/TS ???????????little-endian DataView???
func jsDecodeBinaryIR(field *FieldIR, indent string, boundsCheck bool) string {
	if field.Optional {
		var code string
		code += indent + "{\n"
		if boundsCheck {
			code += jsNeed(indent+"\t", 1)
		}
		code += indent + "\tconst present = data[pointer];\n"
		code += indent + "\tpointer += 1;\n"
		code += indent + "\tif (!present) {\n"
		code += indent + "\t\tthis." + field.DisplayName() + " = null;\n"
		code += indent + "\t} else {\n"
		inner := *field
		inner.Optional = false
		code += jsDecodeBinaryIR(&inner, indent+"\t\t", boundsCheck)
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	}

	name := field.DisplayName()
	switch field.Kind {
	case KindInt8:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+" = view.getInt8(pointer);")
	case KindUint8, KindByte:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+" = view.getUint8(pointer);")
	case KindInt16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+" = view.getInt16(pointer, true);")
	case KindUint16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+" = view.getUint16(pointer, true);")
	case KindInt32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+" = view.getInt32(pointer, true);")
	case KindUint32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+" = view.getUint32(pointer, true);")
	case KindInt64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+" = Number(view.getBigInt64(pointer, true));")
	case KindUint64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+" = Number(view.getBigUint64(pointer, true));")
	case KindFloat32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+" = view.getFloat32(pointer, true);")
	case KindFloat64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+" = view.getFloat64(pointer, true);")
	case KindBool:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+" = data[pointer] !== 0;")
	case KindString:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tthis." + name + " = new TextDecoder().decode(data.slice(pointer, pointer + len));\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	case KindSlice:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst count = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedCount(indent+"\t", "count")
		}
		code += indent + "\tthis." + name + " = [];\n"
		code += indent + "\tfor (let i = 0; i < count; i++) {\n"
		code += jsDecodeSliceItem(field, indent+"\t\t", boundsCheck)
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tthis." + name + " = new " + field.ElemType + "();\n"
		code += indent + "\tthis." + name + ".decode(new Uint8Array(data.buffer, data.byteOffset + pointer, len));\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	case KindMap:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst count = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedCount(indent+"\t", "count")
		}
		code += indent + "\tthis." + name + " = {};\n"
		code += indent + "\tfor (let i = 0; i < count; i++) {\n"
		if boundsCheck {
			code += jsNeed(indent+"\t\t", 4)
		}
		code += indent + "\t\tconst keyLen = view.getInt32(pointer, true);\n"
		code += indent + "\t\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t\t", "keyLen")
		}
		code += indent + "\t\tconst key = new TextDecoder().decode(data.slice(pointer, pointer + keyLen));\n"
		code += indent + "\t\tpointer += keyLen;\n"
		code += jsDecodeMapValue(field, indent+"\t\t", boundsCheck)
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("// unsupported field %s\n", name)
	}
}

func jsReadFixed(indent string, boundsCheck bool, size int, assign string) string {
	var code string
	if boundsCheck {
		code += indent + fmt.Sprintf("if (pointer + %d > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);\n", size)
	}
	code += indent + assign + "\n"
	code += indent + fmt.Sprintf("pointer += %d;\n", size)
	return code
}

func jsNeed(indent string, size int) string {
	return indent + fmt.Sprintf("if (pointer + %d > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);\n", size)
}

func jsNeedLen(indent, lenExpr string) string {
	return indent + "if (" + lenExpr + " < 0 || " + lenExpr + " > 16777216) throw new Error('protocol: length out of range: ' + " + lenExpr + ");\n" +
		indent + "if (pointer + " + lenExpr + " > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);\n"
}

func jsNeedCount(indent, countExpr string) string {
	return indent + "if (" + countExpr + " < 0 || " + countExpr + " > 1048576) throw new Error('protocol: count out of range: ' + " + countExpr + ");\n"
}

func jsDecodeSliceItem(field *FieldIR, indent string, boundsCheck bool) string {
	name := field.DisplayName()
	switch field.ElemKind {
	case KindInt8:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+".push(view.getInt8(pointer));")
	case KindUint8, KindByte:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+".push(view.getUint8(pointer));")
	case KindInt16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+".push(view.getInt16(pointer, true));")
	case KindUint16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+".push(view.getUint16(pointer, true));")
	case KindInt32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+".push(view.getInt32(pointer, true));")
	case KindUint32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+".push(view.getUint32(pointer, true));")
	case KindInt64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+".push(Number(view.getBigInt64(pointer, true)));")
	case KindUint64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+".push(Number(view.getBigUint64(pointer, true)));")
	case KindFloat32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+".push(view.getFloat32(pointer, true));")
	case KindFloat64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+".push(view.getFloat64(pointer, true));")
	case KindBool:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+".push(data[pointer] !== 0);")
	case KindString:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tthis." + name + ".push(new TextDecoder().decode(data.slice(pointer, pointer + len)));\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tconst value = new " + field.ElemType + "();\n"
		code += indent + "\tvalue.decode(new Uint8Array(data.buffer, data.byteOffset + pointer, len));\n"
		code += indent + "\tthis." + name + ".push(value);\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("// unsupported slice elem %s\n", field.ElemType)
	}
}

func jsDecodeMapValue(field *FieldIR, indent string, boundsCheck bool) string {
	name := field.DisplayName()
	switch field.MapValKind {
	case KindInt8:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+"[key] = view.getInt8(pointer);")
	case KindUint8, KindByte:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+"[key] = view.getUint8(pointer);")
	case KindInt16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+"[key] = view.getInt16(pointer, true);")
	case KindUint16:
		return jsReadFixed(indent, boundsCheck, 2, "this."+name+"[key] = view.getUint16(pointer, true);")
	case KindInt32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+"[key] = view.getInt32(pointer, true);")
	case KindUint32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+"[key] = view.getUint32(pointer, true);")
	case KindInt64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+"[key] = Number(view.getBigInt64(pointer, true));")
	case KindUint64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+"[key] = Number(view.getBigUint64(pointer, true));")
	case KindFloat32:
		return jsReadFixed(indent, boundsCheck, 4, "this."+name+"[key] = view.getFloat32(pointer, true);")
	case KindFloat64:
		return jsReadFixed(indent, boundsCheck, 8, "this."+name+"[key] = view.getFloat64(pointer, true);")
	case KindBool:
		return jsReadFixed(indent, boundsCheck, 1, "this."+name+"[key] = data[pointer] !== 0;")
	case KindString:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tthis." + name + "[key] = new TextDecoder().decode(data.slice(pointer, pointer + len));\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		var code string
		if boundsCheck {
			code += jsNeed(indent, 4)
		}
		code += indent + "{\n"
		code += indent + "\tconst len = view.getInt32(pointer, true);\n"
		code += indent + "\tpointer += 4;\n"
		if boundsCheck {
			code += jsNeedLen(indent+"\t", "len")
		}
		code += indent + "\tconst value = new " + field.MapValType + "();\n"
		code += indent + "\tvalue.decode(new Uint8Array(data.buffer, data.byteOffset + pointer, len));\n"
		code += indent + "\tthis." + name + "[key] = value;\n"
		code += indent + "\tpointer += len;\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("// unsupported map value %s\n", field.MapValType)
	}
}

func jsEncodeBinaryIR(field *FieldIR, indent string) string {
	name := field.DisplayName()
	if field.Optional {
		var code string
		code += indent + "if (this." + name + " == null) {\n"
		code += indent + "\tbuffer.push(0);\n"
		code += indent + "} else {\n"
		code += indent + "\tbuffer.push(1);\n"
		inner := *field
		inner.Optional = false
		code += jsEncodeExpr(inner, indent+"\t", "this."+name)
		code += indent + "}\n"
		return code
	}
	return jsEncodeExpr(*field, indent, "this."+name)
}

func jsEncodeExpr(field FieldIR, indent, expr string) string {
	switch field.Kind {
	case KindInt8:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(1);\n" +
			indent + "\tnew DataView(buf).setInt8(0, " + expr + ");\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(" + expr + " & 0xFF);\n"
	case KindInt16:
		return jsPushView(indent, 2, "setInt16", expr)
	case KindUint16:
		return jsPushView(indent, 2, "setUint16", expr)
	case KindInt32:
		return jsPushView(indent, 4, "setInt32", expr)
	case KindUint32:
		return jsPushView(indent, 4, "setUint32", expr)
	case KindInt64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigInt64(0, BigInt(" + expr + "), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigUint64(0, BigInt(" + expr + "), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindFloat32:
		return jsPushView(indent, 4, "setFloat32", expr)
	case KindFloat64:
		return jsPushView(indent, 8, "setFloat64", expr)
	case KindBool:
		return indent + "buffer.push(" + expr + " ? 1 : 0);\n"
	case KindString:
		return indent + "{\n" +
			indent + "\tconst bytes = new TextEncoder().encode(" + expr + ");\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	case KindSlice:
		var code string
		code += indent + "{\n"
		code += indent + "\tconst countBuf = new ArrayBuffer(4);\n"
		code += indent + "\tnew DataView(countBuf).setInt32(0, " + expr + ".length, true);\n"
		code += indent + "\tbuffer.push(...new Uint8Array(countBuf));\n"
		code += indent + "\tfor (const item of " + expr + ") {\n"
		code += jsEncodeSliceItem(&field, indent+"\t\t")
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	case KindStruct:
		return indent + "{\n" +
			indent + "\tconst bytes = " + expr + ".encodeBinary();\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	case KindMap:
		var code string
		code += indent + "{\n"
		code += indent + "\tconst keys = Object.keys(" + expr + ").sort();\n"
		code += indent + "\tconst countBuf = new ArrayBuffer(4);\n"
		code += indent + "\tnew DataView(countBuf).setInt32(0, keys.length, true);\n"
		code += indent + "\tbuffer.push(...new Uint8Array(countBuf));\n"
		code += indent + "\tfor (const key of keys) {\n"
		code += indent + "\t\tconst keyBytes = new TextEncoder().encode(key);\n"
		code += indent + "\t\tconst keyLenBuf = new ArrayBuffer(4);\n"
		code += indent + "\t\tnew DataView(keyLenBuf).setInt32(0, keyBytes.length, true);\n"
		code += indent + "\t\tbuffer.push(...new Uint8Array(keyLenBuf));\n"
		code += indent + "\t\tbuffer.push(...keyBytes);\n"
		code += jsEncodeMapValue(&field, indent+"\t\t", expr+"[key]")
		code += indent + "\t}\n"
		code += indent + "}\n"
		return code
	default:
		return indent + fmt.Sprintf("// unsupported field kind\n")
	}
}

func jsPushView(indent string, size int, setter, expr string) string {
	return indent + "{\n" +
		indent + fmt.Sprintf("\tconst buf = new ArrayBuffer(%d);\n", size) +
		indent + "\tnew DataView(buf)." + setter + "(0, " + expr + ", true);\n" +
		indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
		indent + "}\n"
}

func jsEncodeSliceItem(field *FieldIR, indent string) string {
	switch field.ElemKind {
	case KindInt8:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(1);\n" +
			indent + "\tnew DataView(buf).setInt8(0, item);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(item & 0xFF);\n"
	case KindInt16:
		return jsPushView(indent, 2, "setInt16", "item")
	case KindUint16:
		return jsPushView(indent, 2, "setUint16", "item")
	case KindInt32:
		return jsPushView(indent, 4, "setInt32", "item")
	case KindUint32:
		return jsPushView(indent, 4, "setUint32", "item")
	case KindInt64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigInt64(0, BigInt(item), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigUint64(0, BigInt(item), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindFloat32:
		return jsPushView(indent, 4, "setFloat32", "item")
	case KindFloat64:
		return jsPushView(indent, 8, "setFloat64", "item")
	case KindBool:
		return indent + "buffer.push(item ? 1 : 0);\n"
	case KindString:
		return indent + "{\n" +
			indent + "\tconst bytes = new TextEncoder().encode(item);\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	case KindStruct:
		return indent + "{\n" +
			indent + "\tconst bytes = item.encodeBinary();\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	default:
		return indent + fmt.Sprintf("// unsupported slice elem %s\n", field.ElemType)
	}
}

func jsEncodeMapValue(field *FieldIR, indent, expr string) string {
	switch field.MapValKind {
	case KindInt8:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(1);\n" +
			indent + "\tnew DataView(buf).setInt8(0, " + expr + ");\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint8, KindByte:
		return indent + "buffer.push(" + expr + " & 0xFF);\n"
	case KindInt16:
		return jsPushView(indent, 2, "setInt16", expr)
	case KindUint16:
		return jsPushView(indent, 2, "setUint16", expr)
	case KindInt32:
		return jsPushView(indent, 4, "setInt32", expr)
	case KindUint32:
		return jsPushView(indent, 4, "setUint32", expr)
	case KindInt64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigInt64(0, BigInt(" + expr + "), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindUint64:
		return indent + "{\n" +
			indent + "\tconst buf = new ArrayBuffer(8);\n" +
			indent + "\tnew DataView(buf).setBigUint64(0, BigInt(" + expr + "), true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(buf));\n" +
			indent + "}\n"
	case KindFloat32:
		return jsPushView(indent, 4, "setFloat32", expr)
	case KindFloat64:
		return jsPushView(indent, 8, "setFloat64", expr)
	case KindBool:
		return indent + "buffer.push(" + expr + " ? 1 : 0);\n"
	case KindString:
		return indent + "{\n" +
			indent + "\tconst bytes = new TextEncoder().encode(" + expr + ");\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	case KindStruct:
		return indent + "{\n" +
			indent + "\tconst bytes = " + expr + ".encodeBinary();\n" +
			indent + "\tconst lenBuf = new ArrayBuffer(4);\n" +
			indent + "\tnew DataView(lenBuf).setInt32(0, bytes.length, true);\n" +
			indent + "\tbuffer.push(...new Uint8Array(lenBuf));\n" +
			indent + "\tbuffer.push(...bytes);\n" +
			indent + "}\n"
	default:
		return indent + fmt.Sprintf("// unsupported map value %s\n", field.MapValType)
	}
}

func jsDecodeJSONIR(field *FieldIR, classes []*Class, indent string) string {
	name := field.DisplayName()
	var code string
	code += indent + "if (Object.prototype.hasOwnProperty.call(obj, '" + name + "')) {\n"
	if field.Optional {
		code += indent + "\tif (obj." + name + " == null) { this." + name + " = null; }\n"
		code += indent + "\telse {\n"
		inner := *field
		inner.Optional = false
		code += jsDecodeJSONAssign(&inner, classes, indent+"\t\t")
		code += indent + "\t}\n"
	} else {
		code += jsDecodeJSONAssign(field, classes, indent+"\t")
	}
	code += indent + "}\n"
	return code
}

func jsDecodeJSONAssign(field *FieldIR, classes []*Class, indent string) string {
	name := field.DisplayName()
	switch field.Kind {
	case KindBool:
		return indent + "this." + name + " = Boolean(obj." + name + ");\n"
	case KindString:
		return indent + "this." + name + " = String(obj." + name + ");\n"
	case KindSlice:
		if field.ElemKind == KindString {
			return indent + "this." + name + " = Array.isArray(obj." + name + ") ? obj." + name + ".map(s => String(s)) : [];\n"
		}
		if field.ElemKind == KindStruct {
			var code string
			code += indent + "this." + name + " = [];\n"
			code += indent + "if (Array.isArray(obj." + name + ")) {\n"
			code += indent + "\tfor (const item of obj." + name + ") {\n"
			code += indent + "\t\tconst itemObj = new " + field.ElemType + "();\n"
			code += indent + "\t\titemObj.decode(new TextEncoder().encode(JSON.stringify(item)));\n"
			code += indent + "\t\tthis." + name + ".push(itemObj);\n"
			code += indent + "\t}\n"
			code += indent + "}\n"
			return code
		}
		return indent + "this." + name + " = Array.isArray(obj." + name + ") ? obj." + name + ".map(n => Number(n)) : [];\n"
	case KindStruct:
		return indent + "this." + name + ".decode(new TextEncoder().encode(JSON.stringify(obj." + name + ")));\n"
	case KindMap:
		if field.MapValKind == KindStruct {
			var code string
			code += indent + "this." + name + " = {};\n"
			code += indent + "if (obj." + name + " && typeof obj." + name + " === 'object') {\n"
			code += indent + "\tfor (const [k, v] of Object.entries(obj." + name + ")) {\n"
			code += indent + "\t\tconst itemObj = new " + field.MapValType + "();\n"
			code += indent + "\t\titemObj.decode(new TextEncoder().encode(JSON.stringify(v)));\n"
			code += indent + "\t\tthis." + name + "[k] = itemObj;\n"
			code += indent + "\t}\n"
			code += indent + "}\n"
			return code
		}
		return indent + "this." + name + " = obj." + name + " && typeof obj." + name + " === 'object' ? obj." + name + " : {};\n"
	default:
		if field.Kind.IsNumber() || field.Kind == KindByte {
			return indent + "this." + name + " = Number(obj." + name + ");\n"
		}
		_ = classes
		return indent + "this." + name + " = obj." + name + ";\n"
	}
}
