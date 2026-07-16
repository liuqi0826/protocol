@echo off
echo === Running All Protocol Tests ===
echo.

REM 检查是否已生成协议代码
if not exist "export\go\protocol.go" (
    echo Error: Protocol code not generated. Please run main.go first.
    exit /b 1
)

REM Go 测试
echo --- Testing Go ---
if exist "export\go\protocol.go" (
    cd tests
    go run test_go.go
    cd ..
) else (
    echo Go protocol code not found
)
echo.

REM JavaScript 测试
echo --- Testing JavaScript ---
if exist "export\javascript\protocol.js" (
    node test_javascript.js
) else (
    echo JavaScript protocol code not found
)
echo.

echo === All Tests Complete ===
pause

