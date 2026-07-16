package protocol

func (this *ProtocolExporter) codingToJavaScript() string {
	var code string
	code += "// Auto-generated JavaScript protocol code\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "class " + class.Name + " {\n"

		code += "    constructor() {\n"
		for _, field := range fields {
			code += "        this." + field.DisplayName() + " = " + field.TSDefault() + ";\n"
		}
		code += "    }\n"

		code += "\n    decode(data) {\n"
		code += "        if (!data || data.length === 0) return;\n"
		code += "        if (data.length > 0 && data[0] === 123) { // '{'\n"
		code += "            try {\n"
		code += "                const jsonStr = new TextDecoder().decode(data);\n"
		code += "                const obj = JSON.parse(jsonStr);\n"
		for _, field := range fields {
			code += jsDecodeJSONIR(field, this.Classes, "                ")
		}
		code += "            } catch (e) {\n"
		code += "                console.error('JSON parse error:', e);\n"
		code += "            }\n"
		code += "            return;\n"
		code += "        }\n"
		code += "        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);\n"
		code += "        let pointer = 0;\n"
		for _, field := range fields {
			code += jsDecodeBinaryIR(field, "        ", true)
		}
		code += "    }\n"

		code += "\n    encodeBinary() {\n"
		code += "        const buffer = [];\n"
		for _, field := range fields {
			code += jsEncodeBinaryIR(field, "        ")
		}
		code += "        return new Uint8Array(buffer);\n"
		code += "    }\n"

		code += "\n    encodeJson() {\n"
		code += "        return JSON.stringify(this);\n"
		code += "    }\n"

		code += "}\n\n"
	}

	code += "// Export for Node.js/ES6 modules\n"
	code += "if (typeof module !== 'undefined' && module.exports) {\n"
	code += "    module.exports = {\n"
	for i, class := range this.Classes {
		if i > 0 {
			code += ",\n"
		}
		code += "        " + class.Name + ": " + class.Name
	}
	code += "\n    };\n"
	code += "}\n\n"

	code += "// Export for browser/global scope\n"
	code += "if (typeof window !== 'undefined') {\n"
	for _, class := range this.Classes {
		code += "    window." + class.Name + " = " + class.Name + ";\n"
	}
	code += "}\n"

	return code
}
