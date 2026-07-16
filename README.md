# Protocol Generator

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

一个强大的跨语言协议代码生成器，可以从 Go 结构体定义自动生成多种编程语言的协议序列化/反序列化代码。

> **注意**：本项目支持从 Go 结构体自动生成 10 种编程语言的协议代码，支持 JSON 和二进制两种序列化格式，并具备自动格式识别功能。

## 📑 目录

- [项目简介](#-项目简介)
- [核心特性](#核心特性)
- [支持的语言](#-支持的语言)
- [安装](#-安装)
- [快速开始](#-快速开始)
- [命令行参数](#-命令行参数)
- [支持的数据类型](#-支持的数据类型)
- [序列化格式](#-序列化格式)
- [自动格式识别](#-自动格式识别)
- [项目结构](#-项目结构)
- [测试](#-测试)
- [使用场景](#-使用场景)
- [工作流程](#-工作流程)
- [配置选项](#️-配置选项)
- [故障排除](#-故障排除)
- [开发指南](#-开发指南)
- [贡献](#-贡献)
- [许可证](#-许可证)

## 📖 项目简介

Protocol Generator 是一个用 Go 编写的代码生成工具，它能够解析 Go 结构体定义，并自动生成支持 JSON 和二进制两种序列化格式的协议代码。生成的代码支持多种编程语言，确保跨语言通信的兼容性和一致性。

### 核心特性

- 🔄 **双模式序列化**：自动支持 JSON 和二进制两种序列化格式
- 🎯 **智能格式识别**：根据数据首字节自动识别序列化格式（`{` 表示 JSON）
- 🌍 **多语言支持**：支持 8 种主流编程语言
- 🔒 **类型安全**：生成类型安全的代码，减少运行时错误
- 📦 **零依赖**：生成的代码尽可能使用语言标准库
- 🚀 **易于集成**：生成的代码可直接使用，无需额外配置

## 🎯 支持的语言

| 语言 | 状态 | 文件扩展名 | 依赖库 |
|------|------|-----------|--------|
| **Go** | ✅ 完整支持 | `.go` | 标准库 |
| **C#** | ✅ 完整支持 | `.cs` | System.Text.Json |
| **TypeScript** | ✅ 完整支持 | `.ts` | 无（浏览器/Node.js） |
| **JavaScript** | ✅ 完整支持 | `.js` | 无（浏览器/Node.js） |
| **Rust** | ✅ 完整支持 | `.rs` | serde, serde_json |
| **Zig** | ✅ 完整支持 | `.zig` | 标准库（std.json） |
| **C3** | ✅ 完整支持 | `.c3` | 二进制无依赖；JSON 用 `std::encoding::json`（含 decode） |
| **C** | ✅ 完整支持 | `.h` | cJSON |
| **C++** | ✅ 完整支持 | `.h` | nlohmann/json |
| **GDScript** | ✅ 完整支持 | `.gd` | Godot 引擎 |

## 📦 安装

### 方式一：从源码编译

```bash
# 克隆仓库
git clone <repository-url>
cd protocol

# 编译
go build -o protocol.exe example/main.go

# 或直接运行
go run example/main.go
```

### 方式二：使用预编译二进制

下载对应平台的二进制文件并添加到 PATH。

## 🚀 快速开始

### 1. 定义协议结构

创建 Go 源文件定义你的协议结构：

```go
// source.go
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
```

### 2. 生成协议代码

```bash
# 生成 Go 代码
.\protocol.exe -i source.go -o protocol.go -t go

# 生成 C# 代码
.\protocol.exe -i source.go -o Protocol.cs -t csharp

# 生成 TypeScript 代码
.\protocol.exe -i source.go -o protocol.ts -t typescript

# 生成 JavaScript 代码
.\protocol.exe -i source.go -o protocol.js -t javascript

# 生成 Rust 代码
.\protocol.exe -i source.go -o protocol.rs -t rust

# 生成 Zig 代码
.\protocol.exe -i source.go -o protocol.zig -t zig

# 生成 C3 代码
.\protocol.exe -i source.go -o protocol.c3 -t c3

# 生成 C 代码
.\protocol.exe -i source.go -o protocol.h -t c

# 生成 C++ 代码
.\protocol.exe -i source.go -o protocol.h -t cpp

# 生成 GDScript 代码（多文件）
.\protocol.exe -i source.go -o gdscript/ -t gdscript
```

### 3. 使用生成的代码

#### Go 示例

```go
import "your-module/protocol"

// 创建对象
login := &protocol.ProtocolLogin{
    A: -10,
    B: 20,
    M: "Hello World",
    // ... 其他字段
}

// JSON 序列化
jsonData := login.EncodeJson()

// 二进制序列化
binaryData := login.EncodeBinary()

// 自动识别格式并反序列化，返回 error
var decoded protocol.ProtocolLogin
if err := decoded.Decode(jsonData); err != nil {  // 自动识别为 JSON
    // 处理错误
}
if err := decoded.Decode(binaryData); err != nil { // 自动识别为二进制，截断/损坏会返回错误
    // 处理错误
}
```

#### TypeScript 示例

```typescript
import { ProtocolLogin } from './protocol';

// 创建对象
const login = new ProtocolLogin();
login.a = -10;
login.b = 20;
login.m = "Hello World";

// JSON 序列化
const jsonStr = login.encodeJson();

// 二进制序列化
const binaryData = login.encodeBinary();

// 自动识别格式并反序列化
const decoded = new ProtocolLogin();
decoded.decode(jsonData);  // 自动识别格式
```

#### Rust 示例

```rust
use protocol::ProtocolLogin;

// 创建对象
let mut login = ProtocolLogin::new();
login.a = -10;
login.b = 20;
login.m = "Hello World".to_string();

// JSON 序列化
let json_str = login.encode_json()?;

// 二进制序列化
let binary_data = login.encode_binary();

// 自动识别格式并反序列化
let mut decoded = ProtocolLogin::new();
decoded.decode(&json_data)?;  // 自动识别格式
```

## 📚 命令行参数

```
Usage: protocol.exe [options]

Options:
  -i <file>     输入文件路径（Go 源文件）
  -o <path>     输出文件或目录路径
  -t <language> 目标语言
  -h            显示帮助信息

支持的语言：
  go, csharp, typescript (ts), javascript (js), 
  rust (rs), zig, c3, c, cpp (c++), gdscript, all
```

### 示例

```bash
# 基本用法
.\protocol.exe -i source/source.go -o export/protocol.go -t go

# 使用简写
.\protocol.exe -i source.go -o protocol.ts -t ts
.\protocol.exe -i source.go -o protocol.js -t js
.\protocol.exe -i source.go -o protocol.rs -t rs
.\protocol.exe -i source.go -o protocol.zig -t zig
.\protocol.exe -i source.go -o protocol.c3 -t c3

# 生成到目录（GDScript 会生成多个文件）
.\protocol.exe -i source.go -o gdscript/ -t gdscript
```

## 🔧 支持的数据类型

### 基本类型

- `int8`, `uint8` - 8 位整数
- `int16`, `uint16` - 16 位整数
- `int32`, `uint32` - 32 位整数
- `int64`, `uint64` - 64 位整数
- `float32`, `float64` - 浮点数
- `bool` - 布尔值
- `byte` - 字节
- `string` - 字符串

### 复合类型

- `[]Type` - 数组/切片
- 自定义结构体 - 嵌套对象
- `[]CustomType` - 自定义类型数组
- `*T` - 可选字段（二进制：1 字节是否存在 + 值；Go/TS/JS/Rust 完整支持）
- `map[string]T` - 字符串键映射（二进制：int32 数量 + 排序后的 key/value 对；Go/TS/JS/Rust 完整支持）

### 二进制扩展约定

| 类型 | 布局 |
|------|------|
| `*T` | `uint8 present`（0/1），若为 1 再写 `T` |
| `map[string]T` | `int32 count` + 重复 `string` key + `T` value；编码时按 key 字典序排序 |

## 🎨 序列化格式

### JSON 格式

使用标准的 JSON 格式，字段名来自 `json` 标签：

```json
{
  "a": -10,
  "b": 20,
  "m": "Hello World",
  "q": {
    "nickname": "user123",
    "password": "pass456"
  }
}
```

### 二进制格式

使用小端序（Little-Endian）字节序：

- 整数类型：直接写入对应字节数
- 浮点数：IEEE 754 标准
- 字符串：4 字节长度 + UTF-8 字节
- 数组：4 字节元素数量 + 元素数据
- 自定义类型：4 字节长度 + 序列化后的数据

## 🔍 自动格式识别

所有生成的代码都支持自动格式识别。解码函数会检查数据首字节：

- 如果首字节是 `{` (ASCII 123)，则按 JSON 解析
- 否则按二进制格式解析

这确保了：
- ✅ 不会将二进制数据误解析为 JSON
- ✅ 无需手动指定格式
- ✅ 向后兼容

## 📁 项目结构

```
protocol/
├── protocol.go          # 核心导出逻辑（go/ast 解析）
├── ir.go               # 字段中间表示（FieldIR）
├── golang.go           # Go 代码生成（基于 IR）
├── typescript.go       # TypeScript 代码生成（基于 IR）
├── javascript.go       # JavaScript 代码生成（基于 IR）
├── jscode.go           # TS/JS 共用二进制编解码
├── rust.go             # Rust 代码生成（基于 IR）
├── csharp.go           # C# 代码生成
├── zig.go              # Zig 代码生成
├── c3.go               # C3 代码生成（基于 IR，含 JSON decode）
├── c.go                # C 代码生成
├── cpp.go              # C++ 代码生成
├── gdscript.go         # GDScript 代码生成
├── protocol_test.go    # 单元测试
├── ir_test.go          # IR 单元测试
├── scripts/            # 本地 CI 脚本
├── .github/workflows/  # GitHub Actions
├── example/            # 示例和测试
│   ├── main.go         # 主程序入口
│   ├── source/         # 协议定义源文件
│   ├── export/         # 生成的协议代码
│   └── tests/          # 各语言测试文件
└── README.md           # 本文档
```

## 🧪 测试

### 单元测试与本地 CI

```bash
# 运行生成器单元测试
go test ./...

# 一键本地 CI（生成 + Go/JS 往返）
# Windows:
.\scripts\ci_local.bat
# Linux/macOS:
bash scripts/ci_local.sh
```

项目包含完整的测试套件，验证所有语言的序列化/反序列化功能。

### 运行示例测试

```bash
cd example

# 生成所有语言的协议代码
go run main.go

# 运行测试（查看 tests/README.md 了解详情）
cd tests
# 根据语言运行对应测试
```

详细测试说明请查看 [example/tests/README.md](example/tests/README.md)

## 💡 使用场景

- **游戏开发**：客户端-服务器通信协议
- **微服务架构**：服务间数据交换
- **跨平台应用**：多语言客户端统一协议
- **API 开发**：自动生成客户端 SDK
- **数据持久化**：统一的序列化格式

## 🔄 工作流程

```
Go 结构体定义
    ↓
Protocol Generator
    ↓
多语言协议代码
    ↓
JSON/二进制序列化
    ↓
跨语言通信
```

## ⚙️ 配置选项

### 自动文件扩展名

如果输出路径没有扩展名，工具会自动添加：

- Go → `.go`
- C# → `.cs`
- TypeScript → `.ts`
- JavaScript → `.js`
- Rust → `.rs`
- C/C++ → `.h`
- GDScript → `.gd`

### 自动目录创建

工具会自动创建不存在的目录。

### Rust 项目支持

生成 Rust 代码时，如果输出到目录，会自动生成 `Cargo.toml` 配置文件。

## 🐛 故障排除

### 常见问题

1. **协议代码生成失败**
   - 检查源文件语法是否正确
   - 确保结构体字段有 `json` 标签
   - 检查输出路径权限

2. **生成的代码编译错误**
   - 检查是否安装了必要的依赖
   - 查看生成代码的注释说明
   - 验证目标语言的版本要求

3. **序列化/反序列化不匹配**
   - 确保所有语言使用相同的协议定义
   - 检查字节序设置（使用小端序）
   - 验证数据类型映射是否正确

## 📝 开发指南

### 添加新语言支持

1. 创建新的代码生成文件（如 `python.go`）
2. 实现 `codingToPython()` 方法
3. 在 `protocol.go` 中添加语言常量
4. 在 `Export()` 函数中添加 case
5. 在 `main.go` 中添加命令行支持
6. 创建测试文件验证功能

### 代码生成模式

所有语言生成器遵循相同的模式：

1. **结构体/类定义** - 定义数据结构
2. **构造函数** - 初始化默认值
3. **Decode 方法** - 自动识别格式并解码
4. **EncodeJson 方法** - JSON 序列化
5. **EncodeBinary 方法** - 二进制序列化

## 🤝 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证。详情请查看 LICENSE 文件。

## 🙏 致谢

感谢所有贡献者和使用者的支持！

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 Issue
- 发送 Pull Request
- 查看示例代码：`example/` 目录

## 🔗 相关资源

- [示例代码](example/)
- [测试文档](example/tests/README.md)
- [快速开始指南](example/tests/QUICK_START.md)
- [完整测试文档](example/tests/ALL_TESTS.md)

## 📋 版本历史

### v1.0.0

- ✅ 支持 10 种编程语言（含 Zig、C3）
- ✅ JSON 和二进制双模式序列化
- ✅ 自动格式识别
- ✅ 完整的测试套件
- ✅ 跨平台支持

## ⚠️ 注意事项

1. **依赖管理**：生成的代码需要根据目标语言的要求进行适当的依赖管理
   - Rust: 需要 serde 和 serde_json
   - Zig: 使用标准库 std.json（decode/encodeJson 需要 Allocator）
   - C3: 二进制无额外依赖；JSON 使用标准库 `std::encoding::json`（含 decode）
   - C: 需要 cJSON 库
   - C++: 需要 nlohmann/json 库
   - 其他语言使用标准库

2. **字节序**：所有二进制序列化使用小端序（Little-Endian）

3. **字符编码**：字符串使用 UTF-8 编码

4. **长度上限（防 OOM）**：解码时限制单段字符串/嵌套块最大 16MiB，数组/map 最大约 100 万元素（Go / TS / JS）
5. **浮点精度**：浮点数比较时建议使用容差（如 0.001）
6. **Go Decode**：返回 `error`；截断或超限数据不会 panic，而是返回错误
7. **TS/JS Decode**：二进制越界或超限会 `throw Error`

## 🎓 学习资源

- [Go 结构体定义指南](example/source/source.go)
- [各语言使用示例](example/tests/)
- [序列化格式规范](#序列化格式)

---

**License**: MIT | **Author**: Protocol Generator Team

**Made with ❤️ for cross-language communication**

