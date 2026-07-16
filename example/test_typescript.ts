import { ProtocolLogin, Account } from './export/typescript/protocol';

import * as fs from 'fs';

function verifyLogin(original: ProtocolLogin, decoded: ProtocolLogin): boolean {
    if (original.a !== decoded.a ||
        original.b !== decoded.b ||
        original.c !== decoded.c ||
        original.d !== decoded.d ||
        original.e !== decoded.e ||
        original.f !== decoded.f ||
        original.g !== decoded.g ||
        original.h !== decoded.h ||
        Math.abs(original.i - decoded.i) > 0.001 ||
        Math.abs(original.j - decoded.j) > 0.001 ||
        original.k !== decoded.k ||
        original.l !== decoded.l ||
        original.m !== decoded.m) {
        return false;
    }

    if (!original.n || !decoded.n || original.n.length !== decoded.n.length) {
        return false;
    }
    for (let i = 0; i < original.n.length; i++) {
        if (original.n[i] !== decoded.n[i]) {
            return false;
        }
    }

    if (!original.o || !decoded.o || original.o.length !== decoded.o.length) {
        return false;
    }
    for (let i = 0; i < original.o.length; i++) {
        if (original.o[i] !== decoded.o[i]) {
            return false;
        }
    }

    if (!original.q || !decoded.q ||
        original.q.nickname !== decoded.q.nickname ||
        original.q.password !== decoded.q.password) {
        return false;
    }

    if (!original.r || !decoded.r || original.r.length !== decoded.r.length) {
        return false;
    }
    for (let i = 0; i < original.r.length; i++) {
        if (original.r[i].nickname !== decoded.r[i].nickname ||
            original.r[i].password !== decoded.r[i].password) {
            return false;
        }
    }

    return true;
}

function main() {
    console.log("=== TypeScript Protocol Test ===\n");

    // 创建测试数据
    const login = new ProtocolLogin();
    login.a = -10;
    login.b = 20;
    login.c = -300;
    login.d = 400;
    login.e = -5000;
    login.f = 6000;
    login.g = -70000;
    login.h = 80000;
    login.i = 3.14;
    login.j = 2.718;
    login.k = true;
    login.l = 255;
    login.m = "Hello World";
    login.n = [1, -2, 3, -4];
    login.o = ["test1", "test2", "test3"];
    
    login.q = new Account();
    login.q.nickname = "user123";
    login.q.password = "pass456";
    
    login.r = [
        (() => { const a = new Account(); a.nickname = "user1"; a.password = "pass1"; return a; })(),
        (() => { const a = new Account(); a.nickname = "user2"; a.password = "pass2"; return a; })()
    ];

    // 测试 JSON 序列化
    console.log("1. Testing JSON Serialization...");
    const jsonData = new TextEncoder().encode(login.encodeJson());
    console.log(`   JSON encoded: ${jsonData.length} bytes`);
    fs.writeFileSync("test_output_typescript.json", jsonData);

    // 测试 JSON 反序列化
    console.log("2. Testing JSON Deserialization...");
    const loginFromJson = new ProtocolLogin();
    loginFromJson.decode(jsonData);
    if (verifyLogin(login, loginFromJson)) {
        console.log("   ✓ JSON serialization/deserialization PASSED");
    } else {
        console.log("   ✗ JSON serialization/deserialization FAILED");
    }

    // 测试二进制序列化
    console.log("3. Testing Binary Serialization...");
    const binaryData = login.encodeBinary();
    console.log(`   Binary encoded: ${binaryData.length} bytes`);
    fs.writeFileSync("test_output_typescript.bin", binaryData);

    // 测试二进制反序列化
    console.log("4. Testing Binary Deserialization...");
    const loginFromBinary = new ProtocolLogin();
    loginFromBinary.decode(binaryData);
    if (verifyLogin(login, loginFromBinary)) {
        console.log("   ✓ Binary serialization/deserialization PASSED");
    } else {
        console.log("   ✗ Binary serialization/deserialization FAILED");
    }

    // 测试自动格式识别
    console.log("5. Testing Auto Format Detection...");
    const loginAuto1 = new ProtocolLogin();
    loginAuto1.decode(jsonData);
    if (verifyLogin(login, loginAuto1)) {
        console.log("   ✓ Auto detection (JSON) PASSED");
    } else {
        console.log("   ✗ Auto detection (JSON) FAILED");
    }

    const loginAuto2 = new ProtocolLogin();
    loginAuto2.decode(binaryData);
    if (verifyLogin(login, loginAuto2)) {
        console.log("   ✓ Auto detection (Binary) PASSED");
    } else {
        console.log("   ✗ Auto detection (Binary) FAILED");
    }

    console.log("\n=== Test Complete ===");
}

main();

