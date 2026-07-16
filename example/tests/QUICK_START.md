# 快速开始测试指南

## 一键测试流程

### 步骤 1: 生成协议代码

```bash
cd ..
go run main.go
```

这会自动生成所有语言的协议代码到 `export/` 目录。

### 步骤 2: 运行各语言测试

#### Go 测试

```bash
# 将协议代码设置为可导入的包
cd export/go
go mod init protocol 2>/dev/null || true
cd ../../tests
go run test_go_standalone.go
```

#### C# 测试

```bash
cd ../export/csharp
csc /target:library /out:../../tests/Protocol.dll Protocol.cs
cd ../../tests
csc /reference:Protocol.dll test_csharp.cs
./test_csharp.exe
```

#### TypeScript 测试

```bash
cd tests
npm install -g typescript ts-node
npx ts-node test_typescript.ts
```

#### JavaScript 测试

```bash
cd tests
node test_javascript.js
```

#### Rust 测试

```bash
# 首次运行需要设置项目
cd ../export/rust
cargo init --name protocol 2>/dev/null || true
# 手动将 protocol.rs 内容复制到 src/lib.rs
# 添加依赖到 Cargo.toml:
# [dependencies]
# serde = { version = "1.0", features = ["derive"] }
# serde_json = "1.0"

cargo build --lib
cd ../../tests
rustc --edition 2021 --extern protocol=../export/rust/target/debug/libprotocol.rlib test_rust.rs
./test_rust
```

## 预期输出

所有测试成功后会看到：

```
=== [Language] Protocol Test ===

1. Testing JSON Serialization...
   JSON encoded: XXX bytes
2. Testing JSON Deserialization...
   ✓ JSON serialization/deserialization PASSED
3. Testing Binary Serialization...
   Binary encoded: XXX bytes
4. Testing Binary Deserialization...
   ✓ Binary serialization/deserialization PASSED
5. Testing Auto Format Detection...
   ✓ Auto detection (JSON) PASSED
   ✓ Auto detection (Binary) PASSED

=== Test Complete ===
```

## 验证跨语言兼容性

测试会生成输出文件，可以用于验证跨语言兼容性：

- `test_output_<language>.json` - 可以用于验证 JSON 格式一致性
- `test_output_<language>.bin` - 可以用于验证二进制格式一致性

## 故障排除

### 常见问题

1. **协议代码未生成**
   - 确保已运行 `go run main.go`
   - 检查 `export/` 目录下是否有对应语言的代码

2. **导入错误**
   - Go: 检查包路径和模块设置
   - TypeScript/JavaScript: 检查文件路径
   - Rust: 检查 Cargo.toml 和模块结构

3. **编译错误**
   - 确保安装了对应语言的工具链
   - 检查依赖项是否正确安装

## 下一步

- 查看 [ALL_TESTS.md](ALL_TESTS.md) 了解详细测试说明
- 查看 [SUMMARY.md](SUMMARY.md) 了解测试文件总结
- 查看 [README.md](README.md) 了解测试结构

