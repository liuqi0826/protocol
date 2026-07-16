#!/bin/bash

echo "=== Running All Protocol Tests ==="
echo ""

# 检查是否已生成协议代码
if [ ! -f "export/go/protocol.go" ]; then
    echo "Error: Protocol code not generated. Please run main.go first."
    exit 1
fi

# Go 测试
echo "--- Testing Go ---"
if command -v go &> /dev/null; then
    cd tests
    go run test_go.go
    cd ..
else
    echo "Go not found, skipping Go test"
fi
echo ""

# TypeScript 测试
echo "--- Testing TypeScript ---"
if command -v node &> /dev/null && command -v npm &> /dev/null; then
    if [ -f "export/typescript/protocol.ts" ]; then
        cd tests
        npx tsx test_typescript.ts 2>/dev/null || npx ts-node test_typescript.ts 2>/dev/null || echo "TypeScript test requires tsx or ts-node"
        cd ..
    else
        echo "TypeScript protocol code not found"
    fi
else
    echo "Node.js/npm not found, skipping TypeScript test"
fi
echo ""

# JavaScript 测试
echo "--- Testing JavaScript ---"
if command -v node &> /dev/null; then
    if [ -f "export/javascript/protocol.js" ]; then
        cd tests
        node test_javascript.js
        cd ..
    else
        echo "JavaScript protocol code not found"
    fi
else
    echo "Node.js not found, skipping JavaScript test"
fi
echo ""

echo "=== All Tests Complete ==="

