# 协议测试说明

本目录包含所有语言的协议序列化/反序列化测试。

## 测试步骤

### 1. 生成协议代码

首先运行协议生成器生成所有语言的代码：

```bash
cd ..
go run main.go
# 或使用编译后的程序
./protocol.exe -i source/source.go -o export/ -t all
```

### 2. 运行各语言测试

#### Go 测试

```bash
# 方法1：直接使用生成的代码
cd export/go
go run ../../tests/test_go_standalone.go

# 方法2：复制协议代码到测试目录
cp export/go/protocol.go tests/
cd tests
go run test_go.go
```

#### C# 测试

```bash
# 编译协议库
csc /target:library /out:Protocol.dll ../export/csharp/Protocol.cs

# 编译并运行测试
csc /reference:Protocol.dll test_csharp.cs
./test_csharp.exe
```

#### TypeScript 测试

```bash
# 安装依赖
npm install

# 运行测试
npx ts-node test_typescript.ts
# 或
tsc test_typescript.ts && node test_typescript.js
```

#### JavaScript 测试

```bash
node test_javascript.js
```

#### Rust 测试

```bash
# 创建 Rust 项目
cd ../export/rust
cargo init --name protocol

# 将 protocol.rs 内容复制到 src/lib.rs
# 添加依赖到 Cargo.toml

# 编译库
cargo build

# 运行测试
cd ../../tests
rustc --extern protocol=../export/rust/target/debug/libprotocol.rlib test_rust.rs
./test_rust
```

## 测试验证

每个测试都会验证：

1. ✓ JSON 序列化/反序列化
2. ✓ 二进制序列化/反序列化  
3. ✓ 自动格式识别（JSON）
4. ✓ 自动格式识别（二进制）

所有测试通过后会输出 PASSED，失败会输出 FAILED 并显示详细信息。

## 测试数据

所有测试使用相同的测试数据以确保跨语言兼容性：

- 整数：-10, 20, -300, 400, -5000, 6000, -70000, 80000
- 浮点：3.14, 2.718
- 布尔：true
- 字符串："Hello World"
- 数组：[1, -2, 3, -4], ["test1", "test2", "test3"]
- 嵌套对象：Account 和 Account 数组

