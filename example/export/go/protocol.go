package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
)

const (
	protocolMaxByteLen = 16 << 20 // 单段字符串/嵌套二进制最大 16MiB
	protocolMaxCount   = 1 << 20  // 数组/map 最大元素数
)

type ProtocolLogin struct {
	A int8      `json:"a"`
	B uint8     `json:"b"`
	C int16     `json:"c"`
	D uint16    `json:"d"`
	E int32     `json:"e"`
	F uint32    `json:"f"`
	G int64     `json:"g"`
	H uint64    `json:"h"`
	I float32   `json:"i"`
	J float64   `json:"j"`
	K bool      `json:"k"`
	L byte      `json:"l"`
	M string    `json:"m"`
	N []int8    `json:"n"`
	O []string  `json:"o"`
	Q Account   `json:"q"`
	R []Account `json:"r"`
}

func (this *ProtocolLogin) Decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		if err := json.Unmarshal(data, this); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return nil
	}
	var pointer int
	if pointer+1 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.A = int8(data[pointer])
	pointer += 1
	if pointer+1 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.B = uint8(data[pointer])
	pointer += 1
	if pointer+2 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.C = int16(binary.LittleEndian.Uint16(data[pointer : pointer+2]))
	pointer += 2
	if pointer+2 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.D = binary.LittleEndian.Uint16(data[pointer : pointer+2])
	pointer += 2
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.E = int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
	pointer += 4
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.F = binary.LittleEndian.Uint32(data[pointer : pointer+4])
	pointer += 4
	if pointer+8 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.G = int64(binary.LittleEndian.Uint64(data[pointer : pointer+8]))
	pointer += 8
	if pointer+8 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.H = binary.LittleEndian.Uint64(data[pointer : pointer+8])
	pointer += 8
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.I = math.Float32frombits(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
	pointer += 4
	if pointer+8 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.J = math.Float64frombits(binary.LittleEndian.Uint64(data[pointer : pointer+8]))
	pointer += 8
	if pointer+1 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.K = data[pointer] != 0
	pointer += 1
	if pointer+1 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.L = byte(data[pointer])
	pointer += 1
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.M = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		count := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if count < 0 {
			return fmt.Errorf("protocol: negative count")
		}
		if int(count) > protocolMaxCount {
			return fmt.Errorf("protocol: count %d exceeds limit %d", count, protocolMaxCount)
		}
		this.N = this.N[:0]
		for i := 0; i < int(count); i++ {
			if pointer+1 > len(data) {
				return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
			}
			this.N = append(this.N, int8(data[pointer]))
			pointer += 1
		}
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		count := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if count < 0 {
			return fmt.Errorf("protocol: negative count")
		}
		if int(count) > protocolMaxCount {
			return fmt.Errorf("protocol: count %d exceeds limit %d", count, protocolMaxCount)
		}
		this.O = this.O[:0]
		for i := 0; i < int(count); i++ {
			if pointer+4 > len(data) {
				return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
			}
			{
				length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
				pointer += 4
				if length < 0 {
					return fmt.Errorf("protocol: negative length")
				}
				if int(length) > protocolMaxByteLen {
					return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
				}
				if pointer+int(length) > len(data) {
					return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
				}
				this.O = append(this.O, string(data[pointer:pointer+int(length)]))
				pointer += int(length)
			}
		}
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Q = Account{}
		if err := this.Q.Decode(data[pointer : pointer+int(length)]); err != nil {
			return err
		}
		pointer += int(length)
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		count := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if count < 0 {
			return fmt.Errorf("protocol: negative count")
		}
		if int(count) > protocolMaxCount {
			return fmt.Errorf("protocol: count %d exceeds limit %d", count, protocolMaxCount)
		}
		this.R = this.R[:0]
		for i := 0; i < int(count); i++ {
			if pointer+4 > len(data) {
				return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
			}
			{
				length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
				pointer += 4
				if length < 0 {
					return fmt.Errorf("protocol: negative length")
				}
				if int(length) > protocolMaxByteLen {
					return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
				}
				if pointer+int(length) > len(data) {
					return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
				}
				var value = Account{}
				if err := value.Decode(data[pointer : pointer+int(length)]); err != nil {
					return err
				}
				this.R = append(this.R, value)
				pointer += int(length)
			}
		}
	}
	return nil
}

func (this *ProtocolLogin) EncodeBinary() (data []byte) {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, this.A)
	binary.Write(buffer, binary.LittleEndian, this.B)
	binary.Write(buffer, binary.LittleEndian, this.C)
	binary.Write(buffer, binary.LittleEndian, this.D)
	binary.Write(buffer, binary.LittleEndian, this.E)
	binary.Write(buffer, binary.LittleEndian, this.F)
	binary.Write(buffer, binary.LittleEndian, this.G)
	binary.Write(buffer, binary.LittleEndian, this.H)
	binary.Write(buffer, binary.LittleEndian, this.I)
	binary.Write(buffer, binary.LittleEndian, this.J)
	binary.Write(buffer, binary.LittleEndian, this.K)
	binary.Write(buffer, binary.LittleEndian, this.L)
	{
		var b = []byte(this.M)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	{
		var count = int32(len(this.N))
		binary.Write(buffer, binary.LittleEndian, count)
		for _, item := range this.N {
			binary.Write(buffer, binary.LittleEndian, item)
		}
	}
	{
		var count = int32(len(this.O))
		binary.Write(buffer, binary.LittleEndian, count)
		for _, item := range this.O {
			var b = []byte(item)
			var l = int32(len(b))
			binary.Write(buffer, binary.LittleEndian, l)
			binary.Write(buffer, binary.LittleEndian, b)
		}
	}
	{
		var b = (this.Q).EncodeBinary()
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	{
		var count = int32(len(this.R))
		binary.Write(buffer, binary.LittleEndian, count)
		for _, item := range this.R {
			var b = item.EncodeBinary()
			var l = int32(len(b))
			binary.Write(buffer, binary.LittleEndian, l)
			binary.Write(buffer, binary.LittleEndian, b)
		}
	}
	data = buffer.Bytes()
	return
}

func (this *ProtocolLogin) EncodeJson() (data []byte) {
	data, _ = json.Marshal(this)
	return
}

type Account struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (this *Account) Decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		if err := json.Unmarshal(data, this); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return nil
	}
	var pointer int
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Nickname = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Password = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	return nil
}

func (this *Account) EncodeBinary() (data []byte) {
	var buffer = bytes.NewBuffer([]byte{})
	{
		var b = []byte(this.Nickname)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	{
		var b = []byte(this.Password)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	data = buffer.Bytes()
	return
}

func (this *Account) EncodeJson() (data []byte) {
	data, _ = json.Marshal(this)
	return
}

type ProtocolServerLogin struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (this *ProtocolServerLogin) Decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		if err := json.Unmarshal(data, this); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return nil
	}
	var pointer int
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.ID = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Token = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	return nil
}

func (this *ProtocolServerLogin) EncodeBinary() (data []byte) {
	var buffer = bytes.NewBuffer([]byte{})
	{
		var b = []byte(this.ID)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	{
		var b = []byte(this.Token)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	data = buffer.Bytes()
	return
}

func (this *ProtocolServerLogin) EncodeJson() (data []byte) {
	data, _ = json.Marshal(this)
	return
}

type ProtocolServerState struct {
	State uint16 `json:"state"`
	Value string `json:"value"`
}

func (this *ProtocolServerState) Decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		if err := json.Unmarshal(data, this); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return nil
	}
	var pointer int
	if pointer+2 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.State = binary.LittleEndian.Uint16(data[pointer : pointer+2])
	pointer += 2
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Value = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	return nil
}

func (this *ProtocolServerState) EncodeBinary() (data []byte) {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, this.State)
	{
		var b = []byte(this.Value)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	data = buffer.Bytes()
	return
}

func (this *ProtocolServerState) EncodeJson() (data []byte) {
	data, _ = json.Marshal(this)
	return
}

type ProtocolServerCommand struct {
	Command uint16 `json:"command"`
	Value   string `json:"value"`
}

func (this *ProtocolServerCommand) Decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '{' {
		if err := json.Unmarshal(data, this); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return nil
	}
	var pointer int
	if pointer+2 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	this.Command = binary.LittleEndian.Uint16(data[pointer : pointer+2])
	pointer += 2
	if pointer+4 > len(data) {
		return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
	}
	{
		length := int32(binary.LittleEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		if length < 0 {
			return fmt.Errorf("protocol: negative length")
		}
		if int(length) > protocolMaxByteLen {
			return fmt.Errorf("protocol: length %d exceeds limit %d", length, protocolMaxByteLen)
		}
		if pointer+int(length) > len(data) {
			return fmt.Errorf("protocol: unexpected end of data at offset %d", pointer)
		}
		this.Value = string(data[pointer : pointer+int(length)])
		pointer += int(length)
	}
	return nil
}

func (this *ProtocolServerCommand) EncodeBinary() (data []byte) {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, this.Command)
	{
		var b = []byte(this.Value)
		var l = int32(len(b))
		binary.Write(buffer, binary.LittleEndian, l)
		binary.Write(buffer, binary.LittleEndian, b)
	}
	data = buffer.Bytes()
	return
}

func (this *ProtocolServerCommand) EncodeJson() (data []byte) {
	data, _ = json.Marshal(this)
	return
}
