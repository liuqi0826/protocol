@echo off
setlocal enabledelayedexpansion
cd /d "%~dp0.."

echo === Unit tests ===
go test ./...
if errorlevel 1 exit /b 1

echo === Generate all languages ===
cd example
go run main.go
if errorlevel 1 exit /b 1
cd ..

echo === Vet/build generated Go ===
go vet ./example/export/go/...
if errorlevel 1 exit /b 1
go build -o NUL ./example/export/go/...
if errorlevel 1 exit /b 1

echo === Go round-trip ===
cd example\tests
go run test_go.go
if errorlevel 1 exit /b 1
cd ..\..

echo === JavaScript round-trip ===
cd example
node test_javascript.js
if errorlevel 1 exit /b 1
cd ..

echo === All CI checks passed ===
endlocal
