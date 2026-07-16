class_name ProtocolServerLogin
extends RefCounted
var ID: String
var Token: String

func _init():
	ID = ""
	Token = ""

func decode(data):
	if data == null or data.size() == 0:
		return
	var pointer = 0
	if data[0] == 123:  # '{' 的 ASCII 码是 123
		var json = JSON.new()
		if json.parse(data.get_string_from_utf8()) == 0:
			var obj = json.data
			ID = obj["ID"]
			Token = obj["Token"]
			return
		else:
			push_error("Failed to parse JSON data")
			return
	var ID_len = data.decode_u32(pointer)
	pointer += 4
	ID = data.slice(pointer, pointer + ID_len).get_string_from_utf8()
	pointer += ID_len
	var Token_len = data.decode_u32(pointer)
	pointer += 4
	Token = data.slice(pointer, pointer + Token_len).get_string_from_utf8()
	pointer += Token_len

func encode_json() -> String:
	var obj = {}
	obj["ID"] = ID
	obj["Token"] = Token
	return JSON.stringify(obj)

func encode_binary() -> PackedByteArray:
	var buffer = PackedByteArray()
	var ID_bytes = ID.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(ID_bytes))
	buffer.append_array(ID_bytes)
	var Token_bytes = Token.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(Token_bytes))
	buffer.append_array(Token_bytes)
	return buffer

