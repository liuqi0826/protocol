package protocol

func (this *ProtocolExporter) codingToTypeScript() string {
	var code string
	code += "// Auto-generated TypeScript protocol code\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "export class " + class.Name + " {\n"

		for _, field := range fields {
			code += "    " + field.DisplayName() + ": " + field.TSType() + ";\n"
		}

		code += "\n    constructor() {\n"
		for _, field := range fields {
			code += "        this." + field.DisplayName() + " = " + field.TSDefault() + ";\n"
		}
		code += "    }\n"

		code += "\n    decode(data: Uint8Array): void {\n"
		code += "        if (!data || data.length === 0) return;\n"
		code += "        if (data[0] === 123) { // '{'\n"
		code += "            const jsonStr = new TextDecoder().decode(data);\n"
		code += "            const obj = JSON.parse(jsonStr);\n"
		for _, field := range fields {
			code += "            this." + field.DisplayName() + " = obj." + field.DisplayName() + ";\n"
		}
		code += "            return;\n"
		code += "        }\n"
		code += "        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);\n"
		code += "        let pointer = 0;\n"
		for _, field := range fields {
			code += jsDecodeBinaryIR(field, "        ", true)
		}
		code += "    }\n"

		code += "\n    encodeBinary(): Uint8Array {\n"
		code += "        const buffer: number[] = [];\n"
		for _, field := range fields {
			code += jsEncodeBinaryIR(field, "        ")
		}
		code += "        return new Uint8Array(buffer);\n"
		code += "    }\n"

		code += "\n    encodeJson(): string {\n"
		code += "        return JSON.stringify(this);\n"
		code += "    }\n"

		code += "}\n\n"
	}

	return code
}
