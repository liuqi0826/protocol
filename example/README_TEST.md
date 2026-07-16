# 协议序列化/反序列化测试说明

本目录包含所有语言的协议序列化与反序列化测试示例。

## 测试内容

每个测试文件都会验证以下功能：

1. **JSON 序列化** - 将对象编码为 JSON 格式
2. **JSON 反序列化** - 从 JSON 数据解码对象
3. **二进制序列化** - 将对象编码为二进制格式
4. **二进制反序列化** - 从二进制数据解码对象
5. **自动格式识别** - 根据数据首字节自动识别 JSON 或二进制格式

## 测试数据

测试使用以下数据：

- 整数类型：int8(-10), uint8(20), int16(-300), uint16(400), int32(-5000), uint32(6000), int64(-70000), uint64(80000)
- 浮点类型：float32(3.14), float64(2.718)
- 布尔类型：true
- 字节类型：255
- 字符串："Hello World"
- 数组：int8数组 [1, -2, 3, -4]，字符串数组 ["test1", "test2", "test3"]
- 自定义类型：Account {nickname: "user123", password: "pass456"}
- 自定义类型数组：两个 Account 对象

## 各语言测试说明

### Go

```bash
# 1. 先导出协议代码
go run main.go  # 或使用编译后的 protocol.exe

# 2. 运行测试
cd export/go
go run ../../test_go.go
```

### C#

```bash
# 1. 编译协议代码
csc /out:Protocol.dll export/csharp/Protocol.cs

# 2. 编译并运行测试
csc /reference:Protocol.dll test_csharp.cs
test_csharp.exe
```

### TypeScript

```bash
# 1. 安装依赖
npm install

# 2. 编译 TypeScript
tsc test_typescript.ts

# 3. 运行测试
node test_typescript.js
```

### JavaScript

```bash
# 直接运行（Node.js 环境）
node test_javascript.js
```

### Rust

```bash
# 1. 创建 Rust 项目（如果还没有）
cd export/rust
cargo init --name protocol

# 2. 将生成的 protocol.rs 复制到 src/lib.rs
# 3. 添加依赖到 Cargo.toml
# [dependencies]
# serde = { version = "1.0", features = ["derive"] }
# serde_json = "1.0"

# 4. 运行测试
cd ../../..
rustc --edition 2021 --extern protocol=../export/rust/target/debug/libprotocol.rlib test_rust.rs
./test_rust
```

### C/C++

C 和 C++ 的测试需要手动编写，因为需要链接 cJSON 或 nlohmann/json 库。

## 验证结果

每个测试都会：

1. 创建测试数据对象
2. 序列化为 JSON 和二进制格式
3. 反序列化并验证数据完整性
4. 测试自动格式识别功能

如果所有测试通过，会输出：
- ✓ JSON serialization/deserialization PASSED
- ✓ Binary serialization/deserialization PASSED
- ✓ Auto detection (JSON) PASSED
- ✓ Auto detection (Binary) PASSED

## 输出文件

测试会生成以下文件：

- `test_output_<language>.json` - JSON 序列化输出
- `test_output_<language>.bin` - 二进制序列化输出

这些文件可以用于跨语言兼容性测试。

## 注意事项

1. 确保已先运行协议生成器生成各语言的协议代码
2. 某些语言需要安装额外的依赖（如 Rust 的 serde，TypeScript 的类型定义）
3. 浮点数比较使用容差（0.001）来避免精度问题
4. 所有测试使用小端序（little-endian）进行二进制编码

