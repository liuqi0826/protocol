package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	protocol "protocol/example/export/go"
)

// 可运行的 Go 协议往返测试。
// 用法（在 example/tests 目录）：
//
//	go run test_go.go
//
// 请先生成协议代码：
//
//	cd .. && go run main.go
func main() {
	fmt.Println("=== Go Protocol Test ===")
	fmt.Println()

	protocolPath := filepath.Join("..", "export", "go", "protocol.go")
	if _, err := os.Stat(protocolPath); os.IsNotExist(err) {
		fmt.Printf("Error: Protocol file not found at %s\n", protocolPath)
		fmt.Println("Please run the protocol generator first:")
		fmt.Println("  cd .. && go run main.go")
		os.Exit(1)
	}

	login := &protocol.ProtocolLogin{
		A: -10,
		B: 20,
		C: -300,
		D: 400,
		E: -5000,
		F: 6000,
		G: -70000,
		H: 80000,
		I: 3.14,
		J: 2.718,
		K: true,
		L: 255,
		M: "Hello World",
		N: []int8{1, -2, 3, -4},
		O: []string{"test1", "test2", "test3"},
		Q: protocol.Account{Nickname: "user123", Password: "pass456"},
		R: []protocol.Account{
			{Nickname: "user1", Password: "pass1"},
			{Nickname: "user2", Password: "pass2"},
		},
	}

	fmt.Println("1. Testing JSON...")
	jsonData := login.EncodeJson()
	var fromJSON protocol.ProtocolLogin
	if err := fromJSON.Decode(jsonData); err != nil {
		fmt.Printf("   ✗ JSON decode error: %v\n", err)
		os.Exit(1)
	}
	if !verify(login, &fromJSON) {
		fmt.Println("   ✗ JSON FAILED")
		os.Exit(1)
	}
	fmt.Println("   ✓ JSON PASSED")

	fmt.Println("2. Testing Binary...")
	binaryData := login.EncodeBinary()
	var fromBinary protocol.ProtocolLogin
	if err := fromBinary.Decode(binaryData); err != nil {
		fmt.Printf("   ✗ Binary decode error: %v\n", err)
		os.Exit(1)
	}
	if !verify(login, &fromBinary) {
		fmt.Println("   ✗ Binary FAILED")
		os.Exit(1)
	}
	fmt.Println("   ✓ Binary PASSED")

	fmt.Println("3. Testing empty data...")
	var empty protocol.ProtocolLogin
	if err := empty.Decode([]byte{}); err != nil {
		fmt.Printf("   ✗ Empty data error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("   ✓ Empty data: no panic")

	fmt.Println("4. Testing truncated binary...")
	var truncated protocol.ProtocolLogin
	if err := truncated.Decode(binaryData[:8]); err == nil {
		fmt.Println("   ✗ Truncated binary should return error")
		os.Exit(1)
	}
	fmt.Println("   ✓ Truncated binary returns error")

	fmt.Println()
	fmt.Println("=== Test Complete ===")
}

func verify(want, got *protocol.ProtocolLogin) bool {
	if got.A != want.A || got.B != want.B || got.C != want.C || got.D != want.D ||
		got.E != want.E || got.F != want.F || got.G != want.G || got.H != want.H ||
		got.M != want.M || got.K != want.K || got.L != want.L {
		return false
	}
	if math.Abs(float64(got.I-want.I)) > 0.001 || math.Abs(got.J-want.J) > 0.001 {
		return false
	}
	if len(got.N) != len(want.N) || len(got.O) != len(want.O) || len(got.R) != len(want.R) {
		return false
	}
	for i := range want.N {
		if got.N[i] != want.N[i] {
			return false
		}
	}
	for i := range want.O {
		if got.O[i] != want.O[i] {
			return false
		}
	}
	if got.Q.Nickname != want.Q.Nickname || got.Q.Password != want.Q.Password {
		return false
	}
	for i := range want.R {
		if got.R[i].Nickname != want.R[i].Nickname || got.R[i].Password != want.R[i].Password {
			return false
		}
	}
	return true
}
