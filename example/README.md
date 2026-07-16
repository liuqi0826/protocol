# 协议生成器示例和测试

本目录包含协议生成器的使用示例和所有语言的完整测试套件。

## 目录结构

```
example/
├── source/              # 协议定义源文件
│   └── source.go
├── export/              # 生成的各语言协议代码
│   ├── go/
│   ├── csharp/
│   ├── typescript/
│   ├── javascript/
│   ├── rust/
│   ├── c/
│   ├── cpp/
│   └── gdscript/
├── tests/               # 所有语言的测试文件
│   ├── test_go_standalone.go
│   ├── test_csharp.cs
│   ├── test_typescript.ts
│   ├── test_javascript.js
│   ├── test_rust.rs
│   ├── README.md
│   ├── ALL_TESTS.md
│   └── SUMMARY.md
├── main.go              # 协议生成器主程序
└── README.md            # 本文件
```

## 快速开始

### 1. 生成所有语言的协议代码

```bash
go run main.go
```

或者使用编译后的程序：

```bash
.\protocol.exe -i source/source.go -o export/go/protocol.go -t go
.\protocol.exe -i source/source.go -o export/csharp/Protocol.cs -t csharp
.\protocol.exe -i source/source.go -o export/typescript/protocol.ts -t typescript
.\protocol.exe -i source/source.go -o export/javascript/protocol.js -t javascript
.\protocol.exe -i source/source.go -o export/rust/protocol.rs -t rust
.\protocol.exe -i source/source.go -o export/c/protocol.h -t c
.\protocol.exe -i source/source.go -o export/cpp/protocol.h -t cpp
.\protocol.exe -i source/source.go -o export/gdscript/ -t gdscript
```

### 2. 运行测试

详细测试说明请查看 [tests/README.md](tests/README.md)

## 协议定义

源文件 `source/source.go` 定义了以下协议结构：

- `ProtocolLogin` - 登录协议（包含所有基本类型和嵌套类型）
- `Account` - 账户信息
- `ProtocolServerLogin` - 服务器登录协议
- `ProtocolServerState` - 服务器状态协议
- `ProtocolServerCommand` - 服务器命令协议

## 测试覆盖

所有语言的测试都验证：

1. ✅ JSON 序列化
2. ✅ JSON 反序列化
3. ✅ 二进制序列化
4. ✅ 二进制反序列化
5. ✅ 自动格式识别（根据首字节 '{' 判断）

## 支持的语言

| 语言 | 状态 | 测试文件 | 说明 |
|------|------|----------|------|
| Go | ✅ | test_go_standalone.go | 完整支持 |
| C# | ✅ | test_csharp.cs | 完整支持 |
| TypeScript | ✅ | test_typescript.ts | 完整支持 |
| JavaScript | ✅ | test_javascript.js | 完整支持 |
| Rust | ✅ | test_rust.rs | 完整支持 |
| C | ⚠️ | - | 需要手动编写，依赖 cJSON |
| C++ | ⚠️ | - | 需要手动编写，依赖 nlohmann/json |
| GDScript | ⚠️ | - | 需要在 Godot 引擎中测试 |

## 跨语言兼容性

所有语言生成的协议代码都支持：

- **互操作性** - 任何语言生成的 JSON 可以被其他语言解析
- **二进制兼容** - 任何语言生成的二进制数据可以被其他语言解析
- **自动识别** - 根据数据首字节自动选择解析方式

## 更多信息

- [测试详细说明](tests/README.md)
- [完整测试文档](tests/ALL_TESTS.md)
- [测试总结](tests/SUMMARY.md)

