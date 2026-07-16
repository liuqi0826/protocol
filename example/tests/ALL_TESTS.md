# 完整测试套件

本目录包含所有支持语言的协议序列化/反序列化测试。

## 支持的语言

- ✅ Go
- ✅ C#
- ✅ TypeScript
- ✅ JavaScript
- ✅ Rust
- ⚠️ C (需要手动编写，依赖 cJSON)
- ⚠️ C++ (需要手动编写，依赖 nlohmann/json)
- ⚠️ GDScript (需要在 Godot 引擎中运行)

## 快速开始

### 1. 生成所有语言的协议代码

```bash
cd ..
go run main.go
```

这会生成所有语言的协议代码到 `export/` 目录。

### 2. 运行测试

#### Go 测试

```bash
# 需要将 export/go/protocol.go 放在 protocol 包中
cd export/go
go mod init protocol  # 如果还没有
go run ../../tests/test_go_standalone.go
```

#### C# 测试

```bash
# 编译协议库
csc /target:library /out:../tests/Protocol.dll Protocol.cs

# 编译测试
cd ../../tests
csc /reference:Protocol.dll test_csharp.cs
test_csharp.exe
```

#### TypeScript 测试

```bash
cd tests
npm install typescript @types/node
npx ts-node test_typescript.ts
```

#### JavaScript 测试

```bash
cd tests
node test_javascript.js
```

#### Rust 测试

```bash
# 设置 Rust 项目
cd ../export/rust
cargo init --name protocol
# 将 protocol.rs 复制到 src/lib.rs
# 添加依赖到 Cargo.toml

# 编译
cargo build --lib

# 运行测试
cd ../../tests
rustc --edition 2021 --extern protocol=../export/rust/target/debug/libprotocol.rlib test_rust.rs
./test_rust
```

## 测试验证项

每个测试都会验证以下功能：

1. **JSON 序列化** - 对象 → JSON 字节数组
2. **JSON 反序列化** - JSON 字节数组 → 对象（数据完整性）
3. **二进制序列化** - 对象 → 二进制字节数组
4. **二进制反序列化** - 二进制字节数组 → 对象（数据完整性）
5. **自动格式识别** - 根据首字节 '{' 自动识别 JSON 或二进制

## 测试数据规范

所有语言使用相同的测试数据以确保兼容性：

```json
{
  "a": -10,      // int8
  "b": 20,       // uint8
  "c": -300,     // int16
  "d": 400,      // uint16
  "e": -5000,    // int32
  "f": 6000,     // uint32
  "g": -70000,   // int64
  "h": 80000,    // uint64
  "i": 3.14,     // float32
  "j": 2.718,    // float64
  "k": true,     // bool
  "l": 255,      // byte
  "m": "Hello World",  // string
  "n": [1, -2, 3, -4],  // []int8
  "o": ["test1", "test2", "test3"],  // []string
  "q": {
    "nickname": "user123",
    "password": "pass456"
  },
  "r": [
    {"nickname": "user1", "password": "pass1"},
    {"nickname": "user2", "password": "pass2"}
  ]
}
```

## 输出文件

测试会生成以下文件用于验证：

- `test_output_<language>.json` - JSON 序列化结果
- `test_output_<language>.bin` - 二进制序列化结果

这些文件可以用于：
- 跨语言兼容性验证
- 调试序列化问题
- 性能对比

## 故障排除

### Go 测试失败

- 确保 `export/go/protocol.go` 在正确的包中
- 检查导入路径是否正确

### C# 测试失败

- 确保 .NET SDK 已安装
- 检查 Protocol.cs 是否编译成功

### TypeScript/JavaScript 测试失败

- 确保 Node.js 已安装
- 检查协议文件路径是否正确
- TypeScript 需要安装类型定义

### Rust 测试失败

- 确保 Rust 工具链已安装
- 检查 Cargo.toml 依赖是否正确
- 确保 protocol.rs 在正确的模块结构中

## 跨语言兼容性

所有语言生成的协议代码应该能够：

1. **互相读取 JSON 数据** - 任何语言生成的 JSON 都能被其他语言解析
2. **互相读取二进制数据** - 任何语言生成的二进制数据都能被其他语言解析
3. **自动识别格式** - 根据数据首字节自动选择解析方式

## 性能测试

可以扩展这些测试来测量：

- 序列化速度
- 反序列化速度
- 数据大小（JSON vs Binary）
- 内存使用

## 贡献

添加新语言的测试时，请：

1. 遵循相同的测试结构
2. 使用相同的测试数据
3. 实现完整的验证函数
4. 更新本文档

