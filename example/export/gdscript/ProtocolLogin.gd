class_name ProtocolLogin
extends RefCounted
const Account = preload("Account.gd")

var A: int
var B: int
var C: int
var D: int
var E: int
var F: int
var G: int
var H: int
var I: float
var J: float
var K: bool
var L: int
var M: String
var N: Array
var O: Array
var Q: Account
var R: Array[Account]

func _init():
	A = 0
	B = 0
	C = 0
	D = 0
	E = 0
	F = 0
	G = 0
	H = 0
	I = 0.0
	J = 0.0
	K = false
	L = 0
	M = ""
	N = []
	O = []
	Q = Account.new()
	R = []

func decode(data):
	if data == null or data.size() == 0:
		return
	var pointer = 0
	if data[0] == 123:  # '{' 的 ASCII 码是 123
		var json = JSON.new()
		if json.parse(data.get_string_from_utf8()) == 0:
			var obj = json.data
			A = obj["A"]
			B = obj["B"]
			C = obj["C"]
			D = obj["D"]
			E = obj["E"]
			F = obj["F"]
			G = obj["G"]
			H = obj["H"]
			I = obj["I"]
			J = obj["J"]
			K = obj["K"]
			L = obj["L"]
			M = obj["M"]
			N = obj["N"]
			O = obj["O"]
			Q = obj["Q"]
			R = obj["R"]
			return
		else:
			push_error("Failed to parse JSON data")
			return
	A = data.decode_u32(pointer)
	pointer += 4
	B = data.decode_u32(pointer)
	pointer += 4
	C = data.decode_u32(pointer)
	pointer += 4
	D = data.decode_u32(pointer)
	pointer += 4
	E = data.decode_u32(pointer)
	pointer += 4
	F = data.decode_u32(pointer)
	pointer += 4
	G = data.decode_u32(pointer)
	pointer += 4
	H = data.decode_u32(pointer)
	pointer += 4
	I = data.decode_float(pointer)
	pointer += 4
	J = data.decode_float(pointer)
	pointer += 4
	K = data[pointer] != 0
	pointer += 1
	L = data[pointer]
	pointer += 1
	var M_len = data.decode_u32(pointer)
	pointer += 4
	M = data.slice(pointer, pointer + M_len).get_string_from_utf8()
	pointer += M_len
	var N_count = data.decode_u32(pointer)
	pointer += 4
	N = []
	for i in range(N_count):
		N.append(data.decode_u32(pointer))
		pointer += 4
	var O_count = data.decode_u32(pointer)
	pointer += 4
	O = []
	for i in range(O_count):
		var len = data.decode_u32(pointer)
		pointer += 4
		O.append(data.slice(pointer, pointer + len).get_string_from_utf8())
		pointer += len
	var Q_len = data.decode_u32(pointer)
	pointer += 4
	Q = Account.new()
	Q.decode(data.slice(pointer, pointer + Q_len))
	pointer += Q_len
	var R_count = data.decode_u32(pointer)
	pointer += 4
	R = []
	for i in range(R_count):
		var len = data.decode_u32(pointer)
		pointer += 4
		var value = Account.new()
		value.decode(data.slice(pointer, pointer + len))
		R.append(value)
		pointer += len

func encode_json() -> String:
	var obj = {}
	obj["A"] = A
	obj["B"] = B
	obj["C"] = C
	obj["D"] = D
	obj["E"] = E
	obj["F"] = F
	obj["G"] = G
	obj["H"] = H
	obj["I"] = I
	obj["J"] = J
	obj["K"] = K
	obj["L"] = L
	obj["M"] = M
	obj["N"] = N
	obj["O"] = O
	obj["Q"] = Q
	obj["R"] = R
	return JSON.stringify(obj)

func encode_binary() -> PackedByteArray:
	var buffer = PackedByteArray()
	buffer.resize(buffer.size() + 1)
	buffer.encode_s8(buffer.size() - 1, A)
	buffer.resize(buffer.size() + 1)
	buffer.encode_u8(buffer.size() - 1, B)
	buffer.resize(buffer.size() + 2)
	buffer.encode_s16(buffer.size() - 2, C)
	buffer.resize(buffer.size() + 2)
	buffer.encode_u16(buffer.size() - 2, D)
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, E)
	buffer.resize(buffer.size() + 4)
	buffer.encode_u32(buffer.size() - 4, F)
	buffer.resize(buffer.size() + 8)
	buffer.encode_s64(buffer.size() - 8, G)
	buffer.resize(buffer.size() + 8)
	buffer.encode_u64(buffer.size() - 8, H)
	buffer.resize(buffer.size() + 4)
	buffer.encode_float(buffer.size() - 4, I)
	buffer.resize(buffer.size() + 8)
	buffer.encode_double(buffer.size() - 8, J)
	buffer.append_bool(K)
	buffer.append(L)
	var M_bytes = M.to_utf8_buffer()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(M_bytes))
	buffer.append_array(M_bytes)
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(N))
	for v in N:
		buffer.append(v)
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(O))
	for v in O:
		var v_bytes = v.to_utf8_buffer()
		buffer.resize(buffer.size() + 4)
		buffer.encode_s32(buffer.size() - 4, len(v_bytes))
		buffer.append_array(v_bytes)
	var Q_bytes = Q.encode_binary()
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(Q_bytes))
	buffer.append_array(Q_bytes)
	buffer.resize(buffer.size() + 4)
	buffer.encode_s32(buffer.size() - 4, len(R))
	for v in R:
		var v_bytes = v.encode_binary()
		buffer.resize(buffer.size() + 4)
		buffer.encode_s32(buffer.size() - 4, len(v_bytes))
		buffer.append_array(v_bytes)
	return buffer

