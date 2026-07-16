package protocol

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

type Account struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type ProtocolServerLogin struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
type ProtocolServerState struct {
	State uint16 `json:"state"`
	Value string `json:"value"`
}
type ProtocolServerCommand struct {
	Command uint16 `json:"command"`
	Value   string `json:"value"`
}
