package protocol

import (
	"fmt"
	"strings"
)

func (this *ProtocolExporter) codingToCpp() string {
	var code string
	code += "// Auto-generated C++ protocol code\n"
	code += "#ifndef PROTOCOL_H\n"
	code += "#define PROTOCOL_H\n\n"
	code += "#include <cstdint>\n"
	code += "#include <string>\n"
	code += "#include <vector>\n"
	code += "#include <memory>\n"
	code += "#include <cstring>\n"
	code += "#include <iostream>\n"
	code += "#include \"nlohmann/json.hpp\"\n\n"
	code += "using json = nlohmann::json;\n\n"

	// Forward declarations
	for _, class := range this.Classes {
		code += "class " + class.Name + ";\n"
	}

	code += "\n"

	// Class definitions
	for _, class := range this.Classes {
		code += "class " + class.Name + " {\n"
		code += "public:\n"

		// Member variables
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
				code += "std::string"
			default:
				if strings.Contains(attr.Type, "[]") {
					var innerType = attr.Type[2:]
					switch innerType {
					case "int8":
						code += "std::vector<int8_t>"
					case "uint8":
						code += "std::vector<uint8_t>"
					case "int16":
						code += "std::vector<int16_t>"
					case "uint16":
						code += "std::vector<uint16_t>"
					case "int32":
						code += "std::vector<int32_t>"
					case "uint32":
						code += "std::vector<uint32_t>"
					case "int64":
						code += "std::vector<int64_t>"
					case "uint64":
						code += "std::vector<uint64_t>"
					case "float32":
						code += "std::vector<float>"
					case "float64":
						code += "std::vector<double>"
					case "bool":
						code += "std::vector<bool>"
					case "byte":
						code += "std::vector<uint8_t>"
					case "string":
						code += "std::vector<std::string>"
					default:
						code += "std::vector<" + innerType + ">"
					}
				} else {
					code += attr.Type
				}
			}
			code += " " + name + ";\n"
		}

		// Constructor
		code += "\n    " + class.Name + "() {\n"
		code += "        // All members are default initialized\n"
		code += "    }\n"

		// Decode method
		code += "\n    void decode(const uint8_t* data, size_t data_len) {\n"
		code += "        if (data == nullptr || data_len == 0) return;\n"
		code += "        \n"
		code += "        // Check if JSON format\n"
		code += "        if (data[0] == '{') {\n"
		code += "            try {\n"
		code += "                std::string json_str(reinterpret_cast<const char*>(data), data_len);\n"
		code += "                json j = json::parse(json_str);\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			code += cppDecodeJSON(attr, this.Classes)
		}
		code += "            } catch (const std::exception& e) {\n"
		code += "                std::cerr << \"JSON parse error: \" << e.what() << std::endl;\n"
		code += "            }\n"
		code += "            return;\n"
		code += "        }\n"
		code += "        \n"
		code += "        // Binary decoding\n"
		code += "        size_t pointer = 0;\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			code += cppDecodeBinary(attr, this.Classes)
		}
		code += "    }\n"

		// Encode binary method
		code += "\n    std::vector<uint8_t> encodeBinary() const {\n"
		code += "        std::vector<uint8_t> buffer;\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			code += cppEncodeBinary(attr, this.Classes)
		}
		code += "        return buffer;\n"
		code += "    }\n"

		// Encode JSON method
		code += "\n    std::string encodeJson() const {\n"
		code += "        json j;\n"
		for _, field := range class.ToIR(this.Classes) {
			attr := field.AsAttribute()
			code += cppEncodeJSON(attr, this.Classes)
		}
		code += "        return j.dump();\n"
		code += "    }\n"

		code += "};\n\n"
	}

	code += "#endif // PROTOCOL_H\n"

	return code
}

func cppDecodeJSON(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	code += "                if (j.contains(\"" + name + "\")) {\n"
	switch attr.Type {
	case "int8", "int16", "int32", "int64":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<int64_t>();\n"
	case "uint8", "uint16", "uint32", "uint64":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<uint64_t>();\n"
	case "float32":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<float>();\n"
	case "float64":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<double>();\n"
	case "bool":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<bool>();\n"
	case "byte":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<uint8_t>();\n"
	case "string":
		code += "                    this->" + name + " = j[\"" + name + "\"].get<std::string>();\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "                    this->" + name + " = j[\"" + name + "\"].get<std::vector<std::string>>();\n"
			} else if innerType == "bool" {
				code += "                    this->" + name + " = j[\"" + name + "\"].get<std::vector<bool>>();\n"
			} else {
				code += "                    auto arr = j[\"" + name + "\"];\n"
				code += "                    this->" + name + ".clear();\n"
				code += "                    for (auto& item : arr) {\n"
				if isCustomType(innerType, classes) {
					code += "                        " + innerType + " obj;\n"
					code += "                        std::string item_str = item.dump();\n"
					code += "                        obj.decode(reinterpret_cast<const uint8_t*>(item_str.c_str()), item_str.length());\n"
					code += "                        this->" + name + ".push_back(obj);\n"
				} else {
					code += "                        this->" + name + ".push_back(item.get<" + cppTypeName(innerType) + ">());\n"
				}
				code += "                    }\n"
			}
		} else if isCustomType(attr.Type, classes) {
			code += "                    std::string " + name + "_str = j[\"" + name + "\"].dump();\n"
			code += "                    this->" + name + ".decode(reinterpret_cast<const uint8_t*>(" + name + "_str.c_str()), " + name + "_str.length());\n"
		}
	}
	code += "                }\n"
	return code
}

func cppDecodeBinary(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8":
		code += "        if (pointer + 1 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const int8_t*>(data + pointer);\n"
		code += "            pointer += 1;\n"
		code += "        }\n"
	case "uint8":
		code += "        if (pointer + 1 <= data_len) {\n"
		code += "            " + name + " = data[pointer];\n"
		code += "            pointer += 1;\n"
		code += "        }\n"
	case "int16":
		code += "        if (pointer + 2 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const int16_t*>(data + pointer);\n"
		code += "            pointer += 2;\n"
		code += "        }\n"
	case "uint16":
		code += "        if (pointer + 2 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const uint16_t*>(data + pointer);\n"
		code += "            pointer += 2;\n"
		code += "        }\n"
	case "int32":
		code += "        if (pointer + 4 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const int32_t*>(data + pointer);\n"
		code += "            pointer += 4;\n"
		code += "        }\n"
	case "uint32":
		code += "        if (pointer + 4 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const uint32_t*>(data + pointer);\n"
		code += "            pointer += 4;\n"
		code += "        }\n"
	case "int64":
		code += "        if (pointer + 8 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const int64_t*>(data + pointer);\n"
		code += "            pointer += 8;\n"
		code += "        }\n"
	case "uint64":
		code += "        if (pointer + 8 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const uint64_t*>(data + pointer);\n"
		code += "            pointer += 8;\n"
		code += "        }\n"
	case "float32":
		code += "        if (pointer + 4 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const float*>(data + pointer);\n"
		code += "            pointer += 4;\n"
		code += "        }\n"
	case "float64":
		code += "        if (pointer + 8 <= data_len) {\n"
		code += "            " + name + " = *reinterpret_cast<const double*>(data + pointer);\n"
		code += "            pointer += 8;\n"
		code += "        }\n"
	case "bool":
		code += "        if (pointer + 1 <= data_len) {\n"
		code += "            " + name + " = data[pointer] != 0;\n"
		code += "            pointer += 1;\n"
		code += "        }\n"
	case "byte":
		code += "        if (pointer + 1 <= data_len) {\n"
		code += "            " + name + " = data[pointer];\n"
		code += "            pointer += 1;\n"
		code += "        }\n"
	case "string":
		code += "        if (pointer + 4 <= data_len) {\n"
		code += "            int32_t " + name + "_len = *reinterpret_cast<const int32_t*>(data + pointer);\n"
		code += "            pointer += 4;\n"
		code += "            if (pointer + " + name + "_len <= data_len) {\n"
		code += "                " + name + " = std::string(reinterpret_cast<const char*>(data + pointer), " + name + "_len);\n"
		code += "                pointer += " + name + "_len;\n"
		code += "            }\n"
		code += "        }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "        if (pointer + 4 <= data_len) {\n"
			code += "            int32_t " + name + "_count = *reinterpret_cast<const int32_t*>(data + pointer);\n"
			code += "            pointer += 4;\n"
			code += "            " + name + ".clear();\n"
			code += "            " + name + ".reserve(" + name + "_count);\n"
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "            for (int i = 0; i < " + name + "_count; i++) {\n"
				code += "                if (pointer + 4 <= data_len) {\n"
				code += "                    int32_t len = *reinterpret_cast<const int32_t*>(data + pointer);\n"
				code += "                    pointer += 4;\n"
				code += "                    if (pointer + len <= data_len) {\n"
				code += "                        " + name + ".push_back(std::string(reinterpret_cast<const char*>(data + pointer), len));\n"
				code += "                        pointer += len;\n"
				code += "                    }\n"
				code += "                }\n"
				code += "            }\n"
			} else if isCustomType(innerType, classes) {
				code += "            for (int i = 0; i < " + name + "_count; i++) {\n"
				code += "                if (pointer + 4 <= data_len) {\n"
				code += "                    int32_t len = *reinterpret_cast<const int32_t*>(data + pointer);\n"
				code += "                    pointer += 4;\n"
				code += "                    if (pointer + len <= data_len) {\n"
				code += "                        " + innerType + " obj;\n"
				code += "                        obj.decode(data + pointer, len);\n"
				code += "                        " + name + ".push_back(obj);\n"
				code += "                        pointer += len;\n"
				code += "                    }\n"
				code += "                }\n"
				code += "            }\n"
			} else {
				code += "            for (int i = 0; i < " + name + "_count; i++) {\n"
				code += cppDecodeArrayItem(innerType, name)
				code += "            }\n"
			}
			code += "        }\n"
		} else if isCustomType(attr.Type, classes) {
			code += "        if (pointer + 4 <= data_len) {\n"
			code += "            int32_t " + name + "_len = *reinterpret_cast<const int32_t*>(data + pointer);\n"
			code += "            pointer += 4;\n"
			code += "            if (pointer + " + name + "_len <= data_len) {\n"
			code += "                " + name + ".decode(data + pointer, " + name + "_len);\n"
			code += "                pointer += " + name + "_len;\n"
			code += "            }\n"
			code += "        }\n"
		}
	}
	return code
}

func cppDecodeArrayItem(innerType, arrayName string) string {
	var code string
	switch innerType {
	case "int8":
		code += "                if (pointer + 1 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const int8_t*>(data + pointer));\n"
		code += "                    pointer += 1;\n"
		code += "                }\n"
	case "uint8":
		code += "                if (pointer + 1 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(data[pointer]);\n"
		code += "                    pointer += 1;\n"
		code += "                }\n"
	case "int16":
		code += "                if (pointer + 2 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const int16_t*>(data + pointer));\n"
		code += "                    pointer += 2;\n"
		code += "                }\n"
	case "uint16":
		code += "                if (pointer + 2 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const uint16_t*>(data + pointer));\n"
		code += "                    pointer += 2;\n"
		code += "                }\n"
	case "int32":
		code += "                if (pointer + 4 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const int32_t*>(data + pointer));\n"
		code += "                    pointer += 4;\n"
		code += "                }\n"
	case "uint32":
		code += "                if (pointer + 4 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const uint32_t*>(data + pointer));\n"
		code += "                    pointer += 4;\n"
		code += "                }\n"
	case "int64":
		code += "                if (pointer + 8 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const int64_t*>(data + pointer));\n"
		code += "                    pointer += 8;\n"
		code += "                }\n"
	case "uint64":
		code += "                if (pointer + 8 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const uint64_t*>(data + pointer));\n"
		code += "                    pointer += 8;\n"
		code += "                }\n"
	case "float32":
		code += "                if (pointer + 4 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const float*>(data + pointer));\n"
		code += "                    pointer += 4;\n"
		code += "                }\n"
	case "float64":
		code += "                if (pointer + 8 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(*reinterpret_cast<const double*>(data + pointer));\n"
		code += "                    pointer += 8;\n"
		code += "                }\n"
	case "bool":
		code += "                if (pointer + 1 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(data[pointer] != 0);\n"
		code += "                    pointer += 1;\n"
		code += "                }\n"
	case "byte":
		code += "                if (pointer + 1 <= data_len) {\n"
		code += "                    " + arrayName + ".push_back(data[pointer]);\n"
		code += "                    pointer += 1;\n"
		code += "                }\n"
	}
	return code
}

func cppEncodeBinary(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	code += "        // Encode " + name + "\n"
	switch attr.Type {
	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
		"float32", "float64", "bool", "byte":
		var typeSize = cppTypeSize(attr.Type)
		code += "        {\n"
		code += "            const " + cppTypeName(attr.Type) + "* ptr = &" + name + ";\n"
		code += "            const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);\n"
		code += "            buffer.insert(buffer.end(), bytes, bytes + " + fmt.Sprintf("%d", typeSize) + ");\n"
		code += "        }\n"
	case "string":
		code += "        {\n"
		code += "            int32_t len = static_cast<int32_t>(" + name + ".length());\n"
		code += "            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);\n"
		code += "            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);\n"
		code += "            const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(" + name + ".c_str());\n"
		code += "            buffer.insert(buffer.end(), str_bytes, str_bytes + len);\n"
		code += "        }\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			code += "        {\n"
			code += "            int32_t count = static_cast<int32_t>(" + name + ".size());\n"
			code += "            const uint8_t* count_bytes = reinterpret_cast<const uint8_t*>(&count);\n"
			code += "            buffer.insert(buffer.end(), count_bytes, count_bytes + 4);\n"
			var innerType = attr.Type[2:]
			if innerType == "string" {
				code += "            for (const auto& item : " + name + ") {\n"
				code += "                int32_t len = static_cast<int32_t>(item.length());\n"
				code += "                const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);\n"
				code += "                buffer.insert(buffer.end(), len_bytes, len_bytes + 4);\n"
				code += "                const uint8_t* str_bytes = reinterpret_cast<const uint8_t*>(item.c_str());\n"
				code += "                buffer.insert(buffer.end(), str_bytes, str_bytes + len);\n"
				code += "            }\n"
			} else if isCustomType(innerType, classes) {
				code += "            for (const auto& item : " + name + ") {\n"
				code += "                std::vector<uint8_t> item_data = item.encodeBinary();\n"
				code += "                int32_t len = static_cast<int32_t>(item_data.size());\n"
				code += "                const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);\n"
				code += "                buffer.insert(buffer.end(), len_bytes, len_bytes + 4);\n"
				code += "                buffer.insert(buffer.end(), item_data.begin(), item_data.end());\n"
				code += "            }\n"
			} else {
				var typeSize = cppTypeSize(innerType)
				code += "            for (const auto& item : " + name + ") {\n"
				code += "                const " + cppTypeName(innerType) + "* ptr = &item;\n"
				code += "                const uint8_t* bytes = reinterpret_cast<const uint8_t*>(ptr);\n"
				code += "                buffer.insert(buffer.end(), bytes, bytes + " + fmt.Sprintf("%d", typeSize) + ");\n"
				code += "            }\n"
			}
			code += "        }\n"
		} else if isCustomType(attr.Type, classes) {
			code += "        {\n"
			code += "            std::vector<uint8_t> " + name + "_data = " + name + ".encodeBinary();\n"
			code += "            int32_t len = static_cast<int32_t>(" + name + "_data.size());\n"
			code += "            const uint8_t* len_bytes = reinterpret_cast<const uint8_t*>(&len);\n"
			code += "            buffer.insert(buffer.end(), len_bytes, len_bytes + 4);\n"
			code += "            buffer.insert(buffer.end(), " + name + "_data.begin(), " + name + "_data.end());\n"
			code += "        }\n"
		}
	}
	return code
}

func cppEncodeJSON(attr *Attribute, classes []*Class) string {
	var code string
	var name = getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8", "int16", "int32", "int64":
		code += "        j[\"" + name + "\"] = static_cast<int64_t>(" + name + ");\n"
	case "uint8", "uint16", "uint32", "uint64":
		code += "        j[\"" + name + "\"] = static_cast<uint64_t>(" + name + ");\n"
	case "float32", "float64":
		code += "        j[\"" + name + "\"] = " + name + ";\n"
	case "bool":
		code += "        j[\"" + name + "\"] = " + name + ";\n"
	case "byte":
		code += "        j[\"" + name + "\"] = static_cast<uint8_t>(" + name + ");\n"
	case "string":
		code += "        j[\"" + name + "\"] = " + name + ";\n"
	default:
		if strings.Contains(attr.Type, "[]") {
			var innerType = attr.Type[2:]
			if innerType == "string" || innerType == "bool" {
				code += "        j[\"" + name + "\"] = " + name + ";\n"
			} else if isCustomType(innerType, classes) {
				code += "        j[\"" + name + "\"] = json::array();\n"
				code += "        for (const auto& item : " + name + ") {\n"
				code += "            std::string item_json = item.encodeJson();\n"
				code += "            j[\"" + name + "\"].push_back(json::parse(item_json));\n"
				code += "        }\n"
			} else {
				code += "        j[\"" + name + "\"] = " + name + ";\n"
			}
		} else if isCustomType(attr.Type, classes) {
			code += "        std::string " + name + "_json = " + name + ".encodeJson();\n"
			code += "        j[\"" + name + "\"] = json::parse(" + name + "_json);\n"
		}
	}
	return code
}

func cppTypeName(goType string) string {
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

func cppTypeSize(goType string) int {
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

