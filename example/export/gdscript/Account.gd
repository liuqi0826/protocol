class_name Account
extends RefCounted
var Nickname: String
var Password: String

func _init():
	Nickname = ""
	Password = ""

func decode(data):
	if data == null or data.size() == 0:
		return
	var pointer = 0
	if data[0] == 123:  # '{' 的 ASCII 码是 123
		var json = JSON.new()
		if json.parse(data.get_string_from_utf8()) == 0:
			var obj = json.data
			Nickname = obj["Nickname"]
			Password = obj["Password"]
			return
		else:
			push_error("Failed to parse JSON data")
			return
	var Nickname_len = data.decode_u32(pointer)
	pointer += 4
	Nickname = data.slice(pointer, pointer + Nickname_len).get_string_from_utf8()
	pointer += Nickname_len
	var Password_len = data.decode_u32(pointer)
	pointer += 4
	Password = data.slice(pointer, pointer + Password_len).get_string_from_utf8()
	pointer += Password_len

func encode_json() -> String:
	var obj = {}
	obj["Nickname"] = Nickname
	obj["Password"] = Password
	return JSON.stringify(obj)

func encode_binary() -> PackedByteArray:
	var buffer = PackedByteArray()
	var Nickname_bytes = Nickname.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(Nickname_bytes))
	buffer.append_array(Nickname_bytes)
	var Password_bytes = Password.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(Password_bytes))
	buffer.append_array(Password_bytes)
	return buffer

