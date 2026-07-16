class_name ProtocolServerCommand
extends RefCounted
var Command: int
var Value: String

func _init():
	Command = 0
	Value = ""

func decode(data):
	if data == null or data.size() == 0:
		return
	var pointer = 0
	if data[0] == 123:  # '{' 的 ASCII 码是 123
		var json = JSON.new()
		if json.parse(data.get_string_from_utf8()) == 0:
			var obj = json.data
			Command = obj["Command"]
			Value = obj["Value"]
			return
		else:
			push_error("Failed to parse JSON data")
			return
	Command = data.decode_u32(pointer)
	pointer += 4
	var Value_len = data.decode_u32(pointer)
	pointer += 4
	Value = data.slice(pointer, pointer + Value_len).get_string_from_utf8()
	pointer += Value_len

func encode_json() -> String:
	var obj = {}
	obj["Command"] = Command
	obj["Value"] = Value
	return JSON.stringify(obj)

func encode_binary() -> PackedByteArray:
	var buffer = PackedByteArray()
	buffer.resize(buffer.size() + 2)
	buffer.encode_u16(buffer.size() - 2, Command)
	var Value_bytes = Value.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(Value_bytes))
	buffer.append_array(Value_bytes)
	return buffer

