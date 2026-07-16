package protocol

import (
	"fmt"
	"strings"
)

func (this *ProtocolExporter) codingToC() string {
	var code string
	code += "// Auto-generated C protocol code\n"
	code += "#ifndef PROTOCOL_H\n"
	code += "#define PROTOCOL_H\n\n"
	code += "#include <stdint.h>\n"
	code += "#include <stdbool.h>\n"
	code += "#include <stdlib.h>\n"
	code += "#include <string.h>\n"
	code += "#include <stdio.h>\n"
	code += "#include \"cjson/cJSON.h\"\n\n"

	// Forward declarations
	for _, class := range this.Classes {
		code += "typedef struct " + class.Name + " " + class.Name + ";\n"
	}

	code += "\n"

	// Structure definitions
	for _, class := range this.Classes {
		code += "struct " + class.Name + " {\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			var name = getValueNameFromLable(attr.Label)
			code += "    "
			switch attr.Type {
			case "int8":
				code += "int8_t"
			case "uint8":
				code += "uint8_t"
			case "int16":
				code += "int16_t"
			case "uint16":
				code += "uint16_t"
			case "int32":
				code += "int32_t"
			case "uint32":
				code += "uint32_t"
			case "int64":
				code += "int64_t"
			case "uint64":
				code += "uint64_t"
			case "float32":
				code += "float"
			case "float64":
				code += "double"
			case "bool":
				code += "bool"
			case "byte":
				code += "uint8_t"
			case "string":
				code += "char*"
			default:
				if strings.Contains(attr.Type, "[]") {
					var innerType = attr.Type[2:]
					switch innerType {
					case "int8":
						code += "int8_t*"
					case "uint8":
						code += "uint8_t*"
					case "int16":
						code += "int16_t*"
					case "uint16":
						code += "uint16_t*"
					case "int32":
						code += "int32_t*"
					case "uint32":
						code += "uint32_t*"
					case "int64":
						code += "int64_t*"
					case "uint64":
						code += "uint64_t*"
					case "float32":
						code += "float*"
					case "float64":
						code += "double*"
					case "bool":
						code += "bool*"
					case "byte":
						code += "uint8_t*"
					case "string":
						code += "char**"
					default:
						code += innerType + "*"
					}
					code += "; int " + name + "_count"
				} else {
					code += attr.Type + "*"
				}
			}
			code += " " + name
			if !strings.Contains(attr.Type, "[]") && attr.Type != "string" && !isCustomType(attr.Type, this.Classes) {
				code += ";\n"
			} else if strings.Contains(attr.Type, "[]") {
				code += ";\n"
			} else if attr.Type == "string" {
				code += ";\n"
			} else {
				code += ";\n"
			}
		}
		code += "};\n\n"

		// Function declarations
		code += "void " + class.Name + "_init(" + class.Name + "* self);\n"
		code += "void " + class.Name + "_decode(" + class.Name + "* self, const uint8_t* data, size_t data_len);\n"
		code += "uint8_t* " + class.Name + "_encode_binary(" + class.Name + "* self, size_t* out_len);\n"
		code += "char* " + class.Name + "_encode_json(" + class.Name + "* self);\n"
		code += "void " + class.Name + "_free(" + class.Name + "* self);\n\n"
	}

	code += "#endif // PROTOCOL_H\n\n"

	// Implementation file content
	var implCode string
	implCode += "// Auto-generated C protocol implementation\n"
	implCode += "#include \"protocol.h\"\n\n"

	for _, class := range this.Classes {
		// Init function
		implCode += "void " + class.Name + "_init(" + class.Name + "* self) {\n"
		implCode += "    if (self == NULL) return;\n"
		implCode += "    memset(self, 0, sizeof(" + class.Name + "));\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			var name = getValueNameFromLable(attr.Label)
			if attr.Type == "string" {
				implCode += "    self->" + name + " = NULL;\n"
			} else if strings.Contains(attr.Type, "[]") {
				implCode += "    self->" + name + " = NULL;\n"
				implCode += "    self->" + name + "_count = 0;\n"
			} else if isCustomType(attr.Type, this.Classes) {
				implCode += "    self->" + name + " = (" + attr.Type + "*)malloc(sizeof(" + attr.Type + "));\n"
				implCode += "    " + attr.Type + "_init(self->" + name + ");\n"
			}
		}
		implCode += "}\n\n"

		// Decode function
		implCode += "void " + class.Name + "_decode(" + class.Name + "* self, const uint8_t* data, size_t data_len) {\n"
		implCode += "    if (self == NULL || data == NULL || data_len == 0) return;\n"
		implCode += "    \n"
		implCode += "    // Check if JSON format\n"
		implCode += "    if (data[0] == '{') {\n"
		implCode += "        cJSON* json = cJSON_Parse((const char*)data);\n"
		implCode += "        if (json != NULL) {\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			implCode += cDecodeJSON(attr, this.Classes)
		}
		implCode += "            cJSON_Delete(json);\n"
		implCode += "        }\n"
		implCode += "        return;\n"
		implCode += "    }\n"
		implCode += "    \n"
		implCode += "    // Binary decoding\n"
		implCode += "    size_t pointer = 0;\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			implCode += cDecodeBinary(attr, this.Classes)
		}
		implCode += "}\n\n"

		// Encode binary function
		implCode += "uint8_t* " + class.Name + "_encode_binary(" + class.Name + "* self, size_t* out_len) {\n"
		implCode += "    if (self == NULL) {\n"
		implCode += "        if (out_len) *out_len = 0;\n"
		implCode += "        return NULL;\n"
		implCode += "    }\n"
		implCode += "    \n"
		implCode += "    size_t buffer_size = 1024;\n"
		implCode += "    uint8_t* buffer = (uint8_t*)malloc(buffer_size);\n"
		implCode += "    size_t offset = 0;\n"
		implCode += "    \n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			implCode += cEncodeBinary(attr, this.Classes)
		}
		implCode += "    \n"
		implCode += "    if (out_len) *out_len = offset;\n"
		implCode += "    return buffer;\n"
		implCode += "}\n\n"

		// Encode JSON function
		implCode += "char* " + class.Name + "_encode_json(" + class.Name + "* self) {\n"
		implCode += "    if (self == NULL) return NULL;\n"
		implCode += "    \n"
		implCode += "    cJSON* json = cJSON_CreateObject();\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			implCode += cEncodeJSON(attr, this.Classes)
		}
		implCode += "    \n"
		implCode += "    char* json_string = cJSON_Print(json);\n"
		implCode += "    cJSON_Delete(json);\n"
		implCode += "    return json_string;\n"
		implCode += "}\n\n"

		// Free function
		implCode += "void " + class.Name + "_free(" + class.Name + "* self) {\n"
		implCode += "    if (self == NULL) return;\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			var name = getValueNameFromLable(attr.Label)
			if attr.Type == "string" {
				implCode += "    if (self->" + name + ") free(self->" + name + ");\n"
			} else if strings.Contains(attr.Type, "[]") {
				var innerType = attr.Type[2:]
				if innerType == "string" {
					implCode += "    if (self->" + name + ") {\n"
					implCode += "        for (int i = 0; i < self->" + name + "_count; i++) {\n"
					implCode += "            if (self->" + name + "[i]) free(self->" + name + "[i]);\n"
					implCode += "        }\n"
					implCode += "        free(self->" + name + ");\n"
					implCode += "    }\n"
				} else if isCustomType(innerType, this.Classes) {
					implCode += "    if (self->" + name + ") {\n"
					implCode += "        for (int i = 0; i < self->" + name + "_count; i++) {\n"
					implCode += "            " + innerType + "_free(&self->" + name + "[i]);\n"
					implCode += "        }\n"
					implCode += "        free(self->" + name + ");\n"
					implCode += "    }\n"
				} else {
					implCode += "    if (self->" + name + ") free(self->" + name + ");\n"
				}
			} else if isCustomType(attr.Type, this.Classes) {
				implCode += "    if (self->" + name + ") {\n"
				implCode += "        " + attr.Type + "_free(self->" + name + ");\n"
				implCode += "        free(self->" + name + ");\n"
				implCode += "    }\n"
			}
		}
		implCode += "}\n\n"
	}

	return code + "\n// Implementation\n" + implCode
}

func cDecodeJSON(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	code += "            cJSON* " + name + "_item = cJSON_GetObjectItemCaseSensitive(json, \"" + name + "\");\n"
	switch attr.Type {
	case "int8", "int16", "int32":
		code += "            if (cJSON_IsNumber(" + name + "_item)) {\n"
		code += "                self->" + name + " = (int32_t)cJSON_GetNumberValue(" + name + "_item);\n"
		code += "            }\n"
	case "uint8", "uint16", "uint32":
		code += "            if (cJSON_IsNumber(" + name + "_item)) {\n"
		code += "                self->" + name + " = (uint32_t)cJSON_GetNumberValue(" + name + "_item);\n"
		code += "            }\n"
	case "int64", "uint64":
		code += "            if (cJSON_IsNumber(" + name + "_item)) {\n"
		code += "                self->" + name + " = (int64_t)cJSON_GetNumberValue(" + name + "_item);\n"
		code += "            }\n"
	case "float32", "float64":
		code += "            if (cJSON_IsNumber(" + name + "_item)) {\n"
		code += "                self->" + name + " = (double)cJSON_GetNumberValue(" + name + "_item);\n"
		code += "            }\n"
	case "bool":
		code += "            if (cJSON_IsBool(" + name + "_item)) {\n"
		code += "                self->" + name + " = cJSON_IsTrue(" + name + "_item);\n"
		code += "            }\n"
	case "byte":
		code += "            if (cJSON_IsNumber(" + name + "_item)) {\n"
		code += "                self->" + name + " = (uint8_t)cJSON_GetNumberValue(" + name + "_item);\n"
		code += "            }\n"
	case "string":
		code += "            if (cJSON_IsString(" + name + "_item)) {\n"
		code += "                const char* str = cJSON_GetStringValue(" + name + "_item);\n"
		code += "                if (self->" + name + ") free(self->" + name + ");\n"
		code += "                self->" + name + " = (char*)malloc(strlen(str) + 1);\n"
		code += "                strcpy(self->" + name + ", str);\n"
		code += "            }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "            if (cJSON_IsArray(" + name + "_item)) {\n"
			code += "                // Array decoding - simplified\n"
			code += "            }\n"
		} else if isCustomType(attr.Type, classes) {
			code += "            if (cJSON_IsObject(" + name + "_item)) {\n"
			code += "                char* " + name + "_str = cJSON_Print(" + name + "_item);\n"
			code += "                " + attr.Type + "_decode(self->" + name + ", (uint8_t*)" + name + "_str, strlen(" + name + "_str));\n"
			code += "                free(" + name + "_str);\n"
			code += "            }\n"
		}
	}
	return code
}

func cDecodeBinary(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8":
		code += "    if (pointer + 1 <= data_len) {\n"
		code += "        self->" + name + " = *(int8_t*)(data + pointer);\n"
		code += "        pointer += 1;\n"
		code += "    }\n"
	case "uint8":
		code += "    if (pointer + 1 <= data_len) {\n"
		code += "        self->" + name + " = *(uint8_t*)(data + pointer);\n"
		code += "        pointer += 1;\n"
		code += "    }\n"
	case "int16":
		code += "    if (pointer + 2 <= data_len) {\n"
		code += "        self->" + name + " = *(int16_t*)(data + pointer);\n"
		code += "        pointer += 2;\n"
		code += "    }\n"
	case "uint16":
		code += "    if (pointer + 2 <= data_len) {\n"
		code += "        self->" + name + " = *(uint16_t*)(data + pointer);\n"
		code += "        pointer += 2;\n"
		code += "    }\n"
	case "int32":
		code += "    if (pointer + 4 <= data_len) {\n"
		code += "        self->" + name + " = *(int32_t*)(data + pointer);\n"
		code += "        pointer += 4;\n"
		code += "    }\n"
	case "uint32":
		code += "    if (pointer + 4 <= data_len) {\n"
		code += "        self->" + name + " = *(uint32_t*)(data + pointer);\n"
		code += "        pointer += 4;\n"
		code += "    }\n"
	case "int64":
		code += "    if (pointer + 8 <= data_len) {\n"
		code += "        self->" + name + " = *(int64_t*)(data + pointer);\n"
		code += "        pointer += 8;\n"
		code += "    }\n"
	case "uint64":
		code += "    if (pointer + 8 <= data_len) {\n"
		code += "        self->" + name + " = *(uint64_t*)(data + pointer);\n"
		code += "        pointer += 8;\n"
		code += "    }\n"
	case "float32":
		code += "    if (pointer + 4 <= data_len) {\n"
		code += "        self->" + name + " = *(float*)(data + pointer);\n"
		code += "        pointer += 4;\n"
		code += "    }\n"
	case "float64":
		code += "    if (pointer + 8 <= data_len) {\n"
		code += "        self->" + name + " = *(double*)(data + pointer);\n"
		code += "        pointer += 8;\n"
		code += "    }\n"
	case "bool":
		code += "    if (pointer + 1 <= data_len) {\n"
		code += "        self->" + name + " = data[pointer] != 0;\n"
		code += "        pointer += 1;\n"
		code += "    }\n"
	case "byte":
		code += "    if (pointer + 1 <= data_len) {\n"
		code += "        self->" + name + " = data[pointer];\n"
		code += "        pointer += 1;\n"
		code += "    }\n"
	case "string":
		code += "    if (pointer + 4 <= data_len) {\n"
		code += "        int32_t " + name + "_len = *(int32_t*)(data + pointer);\n"
		code += "        pointer += 4;\n"
		code += "        if (pointer + " + name + "_len <= data_len) {\n"
		code += "            if (self->" + name + ") free(self->" + name + ");\n"
		code += "            self->" + name + " = (char*)malloc(" + name + "_len + 1);\n"
		code += "            memcpy(self->" + name + ", data + pointer, " + name + "_len);\n"
		code += "            self->" + name + "[" + name + "_len] = '\\0';\n"
		code += "            pointer += " + name + "_len;\n"
		code += "        }\n"
		code += "    }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "    if (pointer + 4 <= data_len) {\n"
			code += "        int32_t " + name + "_count = *(int32_t*)(data + pointer);\n"
			code += "        pointer += 4;\n"
			code += "        if (self->" + name + ") free(self->" + name + ");\n"
			code += "        self->" + name + "_count = " + name + "_count;\n"
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "        self->" + name + " = (char**)malloc(sizeof(char*) * " + name + "_count);\n"
				code += "        for (int i = 0; i < " + name + "_count; i++) {\n"
				code += "            if (pointer + 4 <= data_len) {\n"
				code += "                int32_t len = *(int32_t*)(data + pointer);\n"
				code += "                pointer += 4;\n"
				code += "                if (pointer + len <= data_len) {\n"
				code += "                    self->" + name + "[i] = (char*)malloc(len + 1);\n"
				code += "                    memcpy(self->" + name + "[i], data + pointer, len);\n"
				code += "                    self->" + name + "[i][len] = '\\0';\n"
				code += "                    pointer += len;\n"
				code += "                }\n"
				code += "            }\n"
				code += "        }\n"
			} else if isCustomType(innerType, classes) {
				code += "        self->" + name + " = (" + innerType + "*)malloc(sizeof(" + innerType + ") * " + name + "_count);\n"
				code += "        for (int i = 0; i < " + name + "_count; i++) {\n"
				code += "            " + innerType + "_init(&self->" + name + "[i]);\n"
				code += "            if (pointer + 4 <= data_len) {\n"
				code += "                int32_t len = *(int32_t*)(data + pointer);\n"
				code += "                pointer += 4;\n"
				code += "                if (pointer + len <= data_len) {\n"
				code += "                    " + innerType + "_decode(&self->" + name + "[i], data + pointer, len);\n"
				code += "                    pointer += len;\n"
				code += "                }\n"
				code += "            }\n"
				code += "        }\n"
			} else {
				code += "        self->" + name + " = (" + cTypeName(innerType) + "*)malloc(sizeof(" + cTypeName(innerType) + ") * " + name + "_count);\n"
				code += "        for (int i = 0; i < " + name + "_count; i++) {\n"
				code += cDecodeArrayItem(innerType, name)
				code += "        }\n"
			}
			code += "    }\n"
		} else if isCustomType(attr.Type, classes) {
			code += "    if (pointer + 4 <= data_len) {\n"
			code += "        int32_t " + name + "_len = *(int32_t*)(data + pointer);\n"
			code += "        pointer += 4;\n"
			code += "        if (pointer + " + name + "_len <= data_len) {\n"
			code += "            if (self->" + name + " == NULL) {\n"
			code += "                self->" + name + " = (" + attr.Type + "*)malloc(sizeof(" + attr.Type + "));\n"
			code += "                " + attr.Type + "_init(self->" + name + ");\n"
			code += "            }\n"
			code += "            " + attr.Type + "_decode(self->" + name + ", data + pointer, " + name + "_len);\n"
			code += "            pointer += " + name + "_len;\n"
			code += "        }\n"
			code += "    }\n"
		}
	}
	return code
}

func cDecodeArrayItem(innerType, arrayName string) string {
	var code string
	switch innerType {
	case "int8":
		code += "            if (pointer + 1 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(int8_t*)(data + pointer);\n"
		code += "                pointer += 1;\n"
		code += "            }\n"
	case "uint8":
		code += "            if (pointer + 1 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(uint8_t*)(data + pointer);\n"
		code += "                pointer += 1;\n"
		code += "            }\n"
	case "int16":
		code += "            if (pointer + 2 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(int16_t*)(data + pointer);\n"
		code += "                pointer += 2;\n"
		code += "            }\n"
	case "uint16":
		code += "            if (pointer + 2 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(uint16_t*)(data + pointer);\n"
		code += "                pointer += 2;\n"
		code += "            }\n"
	case "int32":
		code += "            if (pointer + 4 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(int32_t*)(data + pointer);\n"
		code += "                pointer += 4;\n"
		code += "            }\n"
	case "uint32":
		code += "            if (pointer + 4 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(uint32_t*)(data + pointer);\n"
		code += "                pointer += 4;\n"
		code += "            }\n"
	case "int64":
		code += "            if (pointer + 8 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(int64_t*)(data + pointer);\n"
		code += "                pointer += 8;\n"
		code += "            }\n"
	case "uint64":
		code += "            if (pointer + 8 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(uint64_t*)(data + pointer);\n"
		code += "                pointer += 8;\n"
		code += "            }\n"
	case "float32":
		code += "            if (pointer + 4 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(float*)(data + pointer);\n"
		code += "                pointer += 4;\n"
		code += "            }\n"
	case "float64":
		code += "            if (pointer + 8 <= data_len) {\n"
		code += "                " + arrayName + "[i] = *(double*)(data + pointer);\n"
		code += "                pointer += 8;\n"
		code += "            }\n"
	case "bool":
		code += "            if (pointer + 1 <= data_len) {\n"
		code += "                " + arrayName + "[i] = data[pointer] != 0;\n"
		code += "                pointer += 1;\n"
		code += "            }\n"
	case "byte":
		code += "            if (pointer + 1 <= data_len) {\n"
		code += "                " + arrayName + "[i] = data[pointer];\n"
		code += "                pointer += 1;\n"
		code += "            }\n"
	}
	return code
}

func cEncodeBinary(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	code += "    // Encode " + name + "\n"
	switch attr.Type {
	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
		"float32", "float64", "bool", "byte":
		var typeSize = cTypeSize(attr.Type)
		code += "    while (offset + " + fmt.Sprintf("%d", typeSize) + " > buffer_size) {\n"
		code += "        buffer_size *= 2;\n"
		code += "        buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
		code += "    }\n"
		code += "    memcpy(buffer + offset, &self->" + name + ", " + fmt.Sprintf("%d", typeSize) + ");\n"
		code += "    offset += " + fmt.Sprintf("%d", typeSize) + ";\n"
	case "string":
		code += "    if (self->" + name + ") {\n"
		code += "        size_t " + name + "_len = strlen(self->" + name + ");\n"
		code += "        while (offset + 4 + " + name + "_len > buffer_size) {\n"
		code += "            buffer_size *= 2;\n"
		code += "            buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
		code += "        }\n"
		code += "        int32_t len = (int32_t)" + name + "_len;\n"
		code += "        memcpy(buffer + offset, &len, 4);\n"
		code += "        offset += 4;\n"
		code += "        memcpy(buffer + offset, self->" + name + ", " + name + "_len);\n"
		code += "        offset += " + name + "_len;\n"
		code += "    } else {\n"
		code += "        int32_t len = 0;\n"
		code += "        memcpy(buffer + offset, &len, 4);\n"
		code += "        offset += 4;\n"
		code += "    }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "    while (offset + 4 > buffer_size) {\n"
			code += "        buffer_size *= 2;\n"
			code += "        buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
			code += "    }\n"
			code += "    int32_t " + name + "_count = self->" + name + "_count;\n"
			code += "    memcpy(buffer + offset, &" + name + "_count, 4);\n"
			code += "    offset += 4;\n"
			code += "    for (int i = 0; i < " + name + "_count; i++) {\n"
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "        if (self->" + name + "[i]) {\n"
				code += "            size_t len = strlen(self->" + name + "[i]);\n"
				code += "            while (offset + 4 + len > buffer_size) {\n"
				code += "                buffer_size *= 2;\n"
				code += "                buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
				code += "            }\n"
				code += "            int32_t str_len = (int32_t)len;\n"
				code += "            memcpy(buffer + offset, &str_len, 4);\n"
				code += "            offset += 4;\n"
				code += "            memcpy(buffer + offset, self->" + name + "[i], len);\n"
				code += "            offset += len;\n"
				code += "        } else {\n"
				code += "            int32_t len = 0;\n"
				code += "            memcpy(buffer + offset, &len, 4);\n"
				code += "            offset += 4;\n"
				code += "        }\n"
			} else if isCustomType(innerType, classes) {
				code += "        size_t item_len = 0;\n"
				code += "        uint8_t* item_data = " + innerType + "_encode_binary(&self->" + name + "[i], &item_len);\n"
				code += "        if (item_data) {\n"
				code += "            while (offset + 4 + item_len > buffer_size) {\n"
				code += "                buffer_size *= 2;\n"
				code += "                buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
				code += "            }\n"
				code += "            int32_t len = (int32_t)item_len;\n"
				code += "            memcpy(buffer + offset, &len, 4);\n"
				code += "            offset += 4;\n"
				code += "            memcpy(buffer + offset, item_data, item_len);\n"
				code += "            offset += item_len;\n"
				code += "            free(item_data);\n"
				code += "        }\n"
			} else {
				var typeSize = cTypeSize(innerType)
				code += "        while (offset + " + fmt.Sprintf("%d", typeSize) + " > buffer_size) {\n"
				code += "            buffer_size *= 2;\n"
				code += "            buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
				code += "        }\n"
				code += "        memcpy(buffer + offset, &self->" + name + "[i], " + fmt.Sprintf("%d", typeSize) + ");\n"
				code += "        offset += " + fmt.Sprintf("%d", typeSize) + ";\n"
			}
			code += "    }\n"
		} else if isCustomType(attr.Type, classes) {
			code += "    if (self->" + name + ") {\n"
			code += "        size_t " + name + "_len = 0;\n"
			code += "        uint8_t* " + name + "_data = " + attr.Type + "_encode_binary(self->" + name + ", &" + name + "_len);\n"
			code += "        if (" + name + "_data) {\n"
			code += "            while (offset + 4 + " + name + "_len > buffer_size) {\n"
			code += "                buffer_size *= 2;\n"
			code += "                buffer = (uint8_t*)realloc(buffer, buffer_size);\n"
			code += "            }\n"
			code += "            int32_t len = (int32_t)" + name + "_len;\n"
			code += "            memcpy(buffer + offset, &len, 4);\n"
			code += "            offset += 4;\n"
			code += "            memcpy(buffer + offset, " + name + "_data, " + name + "_len);\n"
			code += "            offset += " + name + "_len;\n"
			code += "            free(" + name + "_data);\n"
			code += "        }\n"
			code += "    } else {\n"
			code += "        int32_t len = 0;\n"
			code += "        memcpy(buffer + offset, &len, 4);\n"
			code += "        offset += 4;\n"
			code += "    }\n"
		}
	}
	return code
}

func cEncodeJSON(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8", "int16", "int32", "int64":
		code += "    cJSON_AddNumberToObject(json, \"" + name + "\", (double)self->" + name + ");\n"
	case "uint8", "uint16", "uint32", "uint64":
		code += "    cJSON_AddNumberToObject(json, \"" + name + "\", (double)self->" + name + ");\n"
	case "float32", "float64":
		code += "    cJSON_AddNumberToObject(json, \"" + name + "\", self->" + name + ");\n"
	case "bool":
		code += "    cJSON_AddBoolToObject(json, \"" + name + "\", self->" + name + ");\n"
	case "byte":
		code += "    cJSON_AddNumberToObject(json, \"" + name + "\", (double)self->" + name + ");\n"
	case "string":
		code += "    if (self->" + name + ") {\n"
		code += "        cJSON_AddStringToObject(json, \"" + name + "\", self->" + name + ");\n"
		code += "    } else {\n"
		code += "        cJSON_AddStringToObject(json, \"" + name + "\", \"\");\n"
		code += "    }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "    cJSON* " + name + "_array = cJSON_CreateArray();\n"
			code += "    for (int i = 0; i < self->" + name + "_count; i++) {\n"
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "        if (self->" + name + "[i]) {\n"
				code += "            cJSON_AddItemToArray(" + name + "_array, cJSON_CreateString(self->" + name + "[i]));\n"
				code += "        }\n"
			} else {
				code += "        cJSON_AddItemToArray(" + name + "_array, cJSON_CreateNumber((double)self->" + name + "[i]));\n"
			}
			code += "    }\n"
			code += "    cJSON_AddItemToObject(json, \"" + name + "\", " + name + "_array);\n"
		} else if isCustomType(attr.Type, classes) {
			code += "    if (self->" + name + ") {\n"
			code += "        char* " + name + "_json = " + attr.Type + "_encode_json(self->" + name + ");\n"
			code += "        if (" + name + "_json) {\n"
			code += "            cJSON* " + name + "_obj = cJSON_Parse(" + name + "_json);\n"
			code += "            if (" + name + "_obj) {\n"
			code += "                cJSON_AddItemToObject(json, \"" + name + "\", " + name + "_obj);\n"
			code += "            }\n"
			code += "            free(" + name + "_json);\n"
			code += "        }\n"
			code += "    }\n"
		}
	}
	return code
}

func cTypeName(goType string) string {
	switch goType {
	case "int8":
		return "int8_t"
	case "uint8":
		return "uint8_t"
	case "int16":
		return "int16_t"
	case "uint16":
		return "uint16_t"
	case "int32":
		return "int32_t"
	case "uint32":
		return "uint32_t"
	case "int64":
		return "int64_t"
	case "uint64":
		return "uint64_t"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "bool":
		return "bool"
	case "byte":
		return "uint8_t"
	default:
		return goType
	}
}

func cTypeSize(goType string) int {
	switch goType {
	case "int8", "uint8", "bool", "byte":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "uint32", "float32":
		return 4
	case "int64", "uint64", "float64":
		return 8
	default:
		return 0
	}
}

