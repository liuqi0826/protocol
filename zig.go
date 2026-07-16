package protocol

import (
	"strings"
)

func (this *ProtocolExporter) codingToZig() string {
	var code string
	code += "// Auto-generated Zig protocol code\n"
	code += "const std = @import(\"std\");\n\n"

	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "pub const " + class.Name + " = struct {\n"
		for _, field := range fields {
			attr := field.AsAttribute()
			var name = getValueNameFromLable(attr.Label)
			code += "    " + name + ": " + zigType(attr.Type, this.Classes) + " = " + zigDefault(attr.Type, this.Classes) + ",\n"
		}
		code += "\n"

		// deinit
		code += "    pub fn deinit(self: *" + class.Name + ", allocator: std.mem.Allocator) void {\n"
		for _, field := range fields {
			code += zigDeinitField(field.AsAttribute(), this.Classes)
		}
		code += "        self.* = .{};\n"
		code += "    }\n\n"

		// decode
		code += "    pub fn decode(self: *" + class.Name + ", data: []const u8, allocator: std.mem.Allocator) !void {\n"
		code += "        if (data.len == 0) return;\n"
		code += "        if (data[0] == '{') {\n"
		code += "            const parsed = try std.json.parseFromSlice(" + class.Name + ", allocator, data, .{ .allocate = .alloc_always });\n"
		code += "            defer parsed.deinit();\n"
		code += "            self.deinit(allocator);\n"
		code += "            self.* = try clone" + class.Name + "(parsed.value, allocator);\n"
		code += "            return;\n"
		code += "        }\n"
		code += "        var pointer: usize = 0;\n"
		for _, field := range fields {
			code += zigDecodeBinary(field.AsAttribute(), this.Classes)
		}
		code += "    }\n\n"

		// encodeBinary
		code += "    pub fn encodeBinary(self: *const " + class.Name + ", allocator: std.mem.Allocator) ![]u8 {\n"
		code += "        var list: std.ArrayList(u8) = .empty;\n"
		code += "        errdefer list.deinit(allocator);\n"
		for _, field := range fields {
			code += zigEncodeBinary(field.AsAttribute(), this.Classes)
		}
		code += "        return try list.toOwnedSlice(allocator);\n"
		code += "    }\n\n"

		// encodeJson
		code += "    pub fn encodeJson(self: *const " + class.Name + ", allocator: std.mem.Allocator) ![]u8 {\n"
		code += "        return try std.json.stringifyAlloc(allocator, self.*, .{});\n"
		code += "    }\n"
		code += "};\n\n"
	}

	// clone helpers for JSON ownership
	for _, class := range this.Classes {
		fields := class.ToIR(this.Classes)
		code += "fn clone" + class.Name + "(src: " + class.Name + ", allocator: std.mem.Allocator) !" + class.Name + " {\n"
		code += "    var dst: " + class.Name + " = .{};\n"
		for _, field := range fields {
			code += zigCloneField(field.AsAttribute(), this.Classes)
		}
		code += "    return dst;\n"
		code += "}\n\n"
	}

	return code
}

func zigType(goType string, classes []*Class) string {
	switch goType {
	case "int8":
		return "i8"
	case "uint8", "byte":
		return "u8"
	case "int16":
		return "i16"
	case "uint16":
		return "u16"
	case "int32":
		return "i32"
	case "uint32":
		return "u32"
	case "int64":
		return "i64"
	case "uint64":
		return "u64"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "bool":
		return "bool"
	case "string":
		return "[]const u8"
	default:
		if strings.HasPrefix(goType, "[]") {
			inner := goType[2:]
			return "[]" + zigType(inner, classes)
		}
		return goType
	}
}

func zigDefault(goType string, classes []*Class) string {
	switch goType {
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "uint32", "int64", "uint64":
		return "0"
	case "float32", "float64":
		return "0"
	case "bool":
		return "false"
	case "string":
		return "\"\""
	default:
		if strings.HasPrefix(goType, "[]") {
			return "&.{}"
		}
		if isCustomType(goType, classes) {
			return ".{}"
		}
		return "undefined"
	}
}

func zigDeinitField(attr *Attribute, classes []*Class) string {
	name := getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "string":
		return "        if (self." + name + ".len > 0) allocator.free(self." + name + ");\n"
	default:
		if strings.HasPrefix(attr.Type, "[]") {
			inner := attr.Type[2:]
			var code string
			code += "        {\n"
			if inner == "string" {
				code += "            for (self." + name + ") |item| {\n"
				code += "                if (item.len > 0) allocator.free(item);\n"
				code += "            }\n"
			} else if isCustomType(inner, classes) {
				code += "            for (self." + name + ") |*item| {\n"
				code += "                item.deinit(allocator);\n"
				code += "            }\n"
			}
			code += "            if (self." + name + ".len > 0) allocator.free(self." + name + ");\n"
			code += "        }\n"
			return code
		}
		if isCustomType(attr.Type, classes) {
			return "        self." + name + ".deinit(allocator);\n"
		}
	}
	return ""
}

func zigCloneField(attr *Attribute, classes []*Class) string {
	name := getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "string":
		return "    dst." + name + " = try allocator.dupe(u8, src." + name + ");\n"
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "uint32", "int64", "uint64", "float32", "float64", "bool":
		return "    dst." + name + " = src." + name + ";\n"
	default:
		if strings.HasPrefix(attr.Type, "[]") {
			inner := attr.Type[2:]
			var code string
			code += "    {\n"
			code += "        const items = try allocator.alloc(" + zigType(inner, classes) + ", src." + name + ".len);\n"
			code += "        for (src." + name + ", 0..) |item, i| {\n"
			if inner == "string" {
				code += "            items[i] = try allocator.dupe(u8, item);\n"
			} else if isCustomType(inner, classes) {
				code += "            items[i] = try clone" + inner + "(item, allocator);\n"
			} else {
				code += "            items[i] = item;\n"
			}
			code += "        }\n"
			code += "        dst." + name + " = items;\n"
			code += "    }\n"
			return code
		}
		if isCustomType(attr.Type, classes) {
			return "    dst." + name + " = try clone" + attr.Type + "(src." + name + ", allocator);\n"
		}
	}
	return ""
}

func zigDecodeBinary(attr *Attribute, classes []*Class) string {
	name := getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8":
		return "        if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = @bitCast(data[pointer]);\n" +
			"        pointer += 1;\n"
	case "uint8", "byte":
		return "        if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = data[pointer];\n" +
			"        pointer += 1;\n"
	case "int16":
		return "        if (pointer + 2 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(i16, data[pointer..][0..2], .little);\n" +
			"        pointer += 2;\n"
	case "uint16":
		return "        if (pointer + 2 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(u16, data[pointer..][0..2], .little);\n" +
			"        pointer += 2;\n"
	case "int32":
		return "        if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(i32, data[pointer..][0..4], .little);\n" +
			"        pointer += 4;\n"
	case "uint32":
		return "        if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(u32, data[pointer..][0..4], .little);\n" +
			"        pointer += 4;\n"
	case "int64":
		return "        if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(i64, data[pointer..][0..8], .little);\n" +
			"        pointer += 8;\n"
	case "uint64":
		return "        if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = std.mem.readInt(u64, data[pointer..][0..8], .little);\n" +
			"        pointer += 8;\n"
	case "float32":
		return "        if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = @bitCast(std.mem.readInt(u32, data[pointer..][0..4], .little));\n" +
			"        pointer += 4;\n"
	case "float64":
		return "        if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = @bitCast(std.mem.readInt(u64, data[pointer..][0..8], .little));\n" +
			"        pointer += 8;\n"
	case "bool":
		return "        if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"        self." + name + " = data[pointer] != 0;\n" +
			"        pointer += 1;\n"
	case "string":
		return "        {\n" +
			"            if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));\n" +
			"            pointer += 4;\n" +
			"            if (pointer + len > data.len) return error.UnexpectedEnd;\n" +
			"            self." + name + " = try allocator.dupe(u8, data[pointer .. pointer + len]);\n" +
			"            pointer += len;\n" +
			"        }\n"
	default:
		if strings.HasPrefix(attr.Type, "[]") {
			inner := attr.Type[2:]
			var code string
			code += "        {\n"
			code += "            if (pointer + 4 > data.len) return error.UnexpectedEnd;\n"
			code += "            const count: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));\n"
			code += "            pointer += 4;\n"
			code += "            const items = try allocator.alloc(" + zigType(inner, classes) + ", count);\n"
			code += "            errdefer allocator.free(items);\n"
			code += "            for (items) |*item| {\n"
			code += zigDecodeArrayItem(inner, classes)
			code += "            }\n"
			code += "            self." + name + " = items;\n"
			code += "        }\n"
			return code
		}
		if isCustomType(attr.Type, classes) {
			return "        {\n" +
				"            if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
				"            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));\n" +
				"            pointer += 4;\n" +
				"            if (pointer + len > data.len) return error.UnexpectedEnd;\n" +
				"            try self." + name + ".decode(data[pointer .. pointer + len], allocator);\n" +
				"            pointer += len;\n" +
				"        }\n"
		}
	}
	return ""
}

func zigDecodeArrayItem(innerType string, classes []*Class) string {
	switch innerType {
	case "int8":
		return "                if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = @bitCast(data[pointer]);\n" +
			"                pointer += 1;\n"
	case "uint8", "byte":
		return "                if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = data[pointer];\n" +
			"                pointer += 1;\n"
	case "int16":
		return "                if (pointer + 2 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(i16, data[pointer..][0..2], .little);\n" +
			"                pointer += 2;\n"
	case "uint16":
		return "                if (pointer + 2 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(u16, data[pointer..][0..2], .little);\n" +
			"                pointer += 2;\n"
	case "int32":
		return "                if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(i32, data[pointer..][0..4], .little);\n" +
			"                pointer += 4;\n"
	case "uint32":
		return "                if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(u32, data[pointer..][0..4], .little);\n" +
			"                pointer += 4;\n"
	case "int64":
		return "                if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(i64, data[pointer..][0..8], .little);\n" +
			"                pointer += 8;\n"
	case "uint64":
		return "                if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = std.mem.readInt(u64, data[pointer..][0..8], .little);\n" +
			"                pointer += 8;\n"
	case "float32":
		return "                if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = @bitCast(std.mem.readInt(u32, data[pointer..][0..4], .little));\n" +
			"                pointer += 4;\n"
	case "float64":
		return "                if (pointer + 8 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = @bitCast(std.mem.readInt(u64, data[pointer..][0..8], .little));\n" +
			"                pointer += 8;\n"
	case "bool":
		return "                if (pointer + 1 > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = data[pointer] != 0;\n" +
			"                pointer += 1;\n"
	case "string":
		return "                if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
			"                const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));\n" +
			"                pointer += 4;\n" +
			"                if (pointer + len > data.len) return error.UnexpectedEnd;\n" +
			"                item.* = try allocator.dupe(u8, data[pointer .. pointer + len]);\n" +
			"                pointer += len;\n"
	default:
		if isCustomType(innerType, classes) {
			return "                if (pointer + 4 > data.len) return error.UnexpectedEnd;\n" +
				"                const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));\n" +
				"                pointer += 4;\n" +
				"                if (pointer + len > data.len) return error.UnexpectedEnd;\n" +
				"                item.* = .{};\n" +
				"                try item.decode(data[pointer .. pointer + len], allocator);\n" +
				"                pointer += len;\n"
		}
	}
	return ""
}

func zigEncodeBinary(attr *Attribute, classes []*Class) string {
	name := getValueNameFromLable(attr.Label)
	switch attr.Type {
	case "int8":
		return "        try list.append(allocator, @bitCast(self." + name + "));\n"
	case "uint8", "byte":
		return "        try list.append(allocator, self." + name + ");\n"
	case "int16":
		return "        {\n            var buf: [2]u8 = undefined;\n            std.mem.writeInt(i16, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "uint16":
		return "        {\n            var buf: [2]u8 = undefined;\n            std.mem.writeInt(u16, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "int32":
		return "        {\n            var buf: [4]u8 = undefined;\n            std.mem.writeInt(i32, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "uint32":
		return "        {\n            var buf: [4]u8 = undefined;\n            std.mem.writeInt(u32, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "int64":
		return "        {\n            var buf: [8]u8 = undefined;\n            std.mem.writeInt(i64, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "uint64":
		return "        {\n            var buf: [8]u8 = undefined;\n            std.mem.writeInt(u64, &buf, self." + name + ", .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "float32":
		return "        {\n            var buf: [4]u8 = undefined;\n            std.mem.writeInt(u32, &buf, @bitCast(self." + name + "), .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "float64":
		return "        {\n            var buf: [8]u8 = undefined;\n            std.mem.writeInt(u64, &buf, @bitCast(self." + name + "), .little);\n            try list.appendSlice(allocator, &buf);\n        }\n"
	case "bool":
		return "        try list.append(allocator, if (self." + name + ") 1 else 0);\n"
	case "string":
		return "        {\n            var len_buf: [4]u8 = undefined;\n            std.mem.writeInt(i32, &len_buf, @intCast(self." + name + ".len), .little);\n            try list.appendSlice(allocator, &len_buf);\n            try list.appendSlice(allocator, self." + name + ");\n        }\n"
	default:
		if strings.HasPrefix(attr.Type, "[]") {
			inner := attr.Type[2:]
			var code string
			code += "        {\n"
			code += "            var count_buf: [4]u8 = undefined;\n"
			code += "            std.mem.writeInt(i32, &count_buf, @intCast(self." + name + ".len), .little);\n"
			code += "            try list.appendSlice(allocator, &count_buf);\n"
			code += "            for (self." + name + ") |item| {\n"
			code += zigEncodeArrayItem(inner, classes)
			code += "            }\n"
			code += "        }\n"
			return code
		}
		if isCustomType(attr.Type, classes) {
			return "        {\n" +
				"            const nested = try self." + name + ".encodeBinary(allocator);\n" +
				"            defer allocator.free(nested);\n" +
				"            var len_buf: [4]u8 = undefined;\n" +
				"            std.mem.writeInt(i32, &len_buf, @intCast(nested.len), .little);\n" +
				"            try list.appendSlice(allocator, &len_buf);\n" +
				"            try list.appendSlice(allocator, nested);\n" +
				"        }\n"
		}
	}
	return ""
}

func zigEncodeArrayItem(innerType string, classes []*Class) string {
	switch innerType {
	case "int8":
		return "                try list.append(allocator, @bitCast(item));\n"
	case "uint8", "byte":
		return "                try list.append(allocator, item);\n"
	case "int16":
		return "                {\n                    var buf: [2]u8 = undefined;\n                    std.mem.writeInt(i16, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "uint16":
		return "                {\n                    var buf: [2]u8 = undefined;\n                    std.mem.writeInt(u16, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "int32":
		return "                {\n                    var buf: [4]u8 = undefined;\n                    std.mem.writeInt(i32, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "uint32":
		return "                {\n                    var buf: [4]u8 = undefined;\n                    std.mem.writeInt(u32, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "int64":
		return "                {\n                    var buf: [8]u8 = undefined;\n                    std.mem.writeInt(i64, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "uint64":
		return "                {\n                    var buf: [8]u8 = undefined;\n                    std.mem.writeInt(u64, &buf, item, .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "float32":
		return "                {\n                    var buf: [4]u8 = undefined;\n                    std.mem.writeInt(u32, &buf, @bitCast(item), .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "float64":
		return "                {\n                    var buf: [8]u8 = undefined;\n                    std.mem.writeInt(u64, &buf, @bitCast(item), .little);\n                    try list.appendSlice(allocator, &buf);\n                }\n"
	case "bool":
		return "                try list.append(allocator, if (item) 1 else 0);\n"
	case "string":
		return "                {\n                    var len_buf: [4]u8 = undefined;\n                    std.mem.writeInt(i32, &len_buf, @intCast(item.len), .little);\n                    try list.appendSlice(allocator, &len_buf);\n                    try list.appendSlice(allocator, item);\n                }\n"
	default:
		if isCustomType(innerType, classes) {
			return "                {\n" +
				"                    const nested = try item.encodeBinary(allocator);\n" +
				"                    defer allocator.free(nested);\n" +
				"                    var len_buf: [4]u8 = undefined;\n" +
				"                    std.mem.writeInt(i32, &len_buf, @intCast(nested.len), .little);\n" +
				"                    try list.appendSlice(allocator, &len_buf);\n" +
				"                    try list.appendSlice(allocator, nested);\n" +
				"                }\n"
		}
	}
	return ""
}
