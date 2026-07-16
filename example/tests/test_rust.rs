use std::fs;
use protocol::ProtocolLogin;
use protocol::Account;

fn verify_login(original: &ProtocolLogin, decoded: &ProtocolLogin) -> bool {
    if original.a != decoded.a ||
        original.b != decoded.b ||
        original.c != decoded.c ||
        original.d != decoded.d ||
        original.e != decoded.e ||
        original.f != decoded.f ||
        original.g != decoded.g ||
        original.h != decoded.h ||
        (original.i - decoded.i).abs() > 0.001 ||
        (original.j - decoded.j).abs() > 0.001 ||
        original.k != decoded.k ||
        original.l != decoded.l ||
        original.m != decoded.m {
        return false;
    }

    if original.n.len() != decoded.n.len() {
        return false;
    }
    for i in 0..original.n.len() {
        if original.n[i] != decoded.n[i] {
            return false;
        }
    }

    if original.o.len() != decoded.o.len() {
        return false;
    }
    for i in 0..original.o.len() {
        if original.o[i] != decoded.o[i] {
            return false;
        }
    }

    if original.q.nickname != decoded.q.nickname ||
        original.q.password != decoded.q.password {
        return false;
    }

    if original.r.len() != decoded.r.len() {
        return false;
    }
    for i in 0..original.r.len() {
        if original.r[i].nickname != decoded.r[i].nickname ||
            original.r[i].password != decoded.r[i].password {
            return false;
        }
    }

    true
}

fn main() {
    println!("=== Rust Protocol Test ===\n");

    // 创建测试数据
    let mut login = ProtocolLogin::new();
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
    login.m = "Hello World".to_string();
    login.n = vec![1, -2, 3, -4];
    login.o = vec!["test1".to_string(), "test2".to_string(), "test3".to_string()];
    
    login.q.nickname = "user123".to_string();
    login.q.password = "pass456".to_string();
    
    login.r = vec![
        {
            let mut a = Account::new();
            a.nickname = "user1".to_string();
            a.password = "pass1".to_string();
            a
        },
        {
            let mut a = Account::new();
            a.nickname = "user2".to_string();
            a.password = "pass2".to_string();
            a
        }
    ];

    // 测试 JSON 序列化
    println!("1. Testing JSON Serialization...");
    match login.encode_json() {
        Ok(json_str) => {
            let json_data = json_str.as_bytes();
            println!("   JSON encoded: {} bytes", json_data.len());
            fs::write("test_output_rust.json", json_data).unwrap();
        }
        Err(e) => {
            println!("   ✗ JSON encoding failed: {}", e);
            return;
        }
    }

    // 测试 JSON 反序列化
    println!("2. Testing JSON Deserialization...");
    let json_data = fs::read("test_output_rust.json").unwrap();
    let mut login_from_json = ProtocolLogin::new();
    match login_from_json.decode(&json_data) {
        Ok(_) => {
            if verify_login(&login, &login_from_json) {
                println!("   ✓ JSON serialization/deserialization PASSED");
            } else {
                println!("   ✗ JSON serialization/deserialization FAILED");
            }
        }
        Err(e) => {
            println!("   ✗ JSON decoding failed: {}", e);
        }
    }

    // 测试二进制序列化
    println!("3. Testing Binary Serialization...");
    let binary_data = login.encode_binary();
    println!("   Binary encoded: {} bytes", binary_data.len());
    fs::write("test_output_rust.bin", &binary_data).unwrap();

    // 测试二进制反序列化
    println!("4. Testing Binary Deserialization...");
    let mut login_from_binary = ProtocolLogin::new();
    match login_from_binary.decode(&binary_data) {
        Ok(_) => {
            if verify_login(&login, &login_from_binary) {
                println!("   ✓ Binary serialization/deserialization PASSED");
            } else {
                println!("   ✗ Binary serialization/deserialization FAILED");
            }
        }
        Err(e) => {
            println!("   ✗ Binary decoding failed: {}", e);
        }
    }

    // 测试自动格式识别
    println!("5. Testing Auto Format Detection...");
    let mut login_auto1 = ProtocolLogin::new();
    match login_auto1.decode(&json_data) {
        Ok(_) => {
            if verify_login(&login, &login_auto1) {
                println!("   ✓ Auto detection (JSON) PASSED");
            } else {
                println!("   ✗ Auto detection (JSON) FAILED");
            }
        }
        Err(e) => {
            println!("   ✗ Auto detection (JSON) failed: {}", e);
        }
    }

    let mut login_auto2 = ProtocolLogin::new();
    match login_auto2.decode(&binary_data) {
        Ok(_) => {
            if verify_login(&login, &login_auto2) {
                println!("   ✓ Auto detection (Binary) PASSED");
            } else {
                println!("   ✗ Auto detection (Binary) FAILED");
            }
        }
        Err(e) => {
            println!("   ✗ Auto detection (Binary) failed: {}", e);
        }
    }

    println!("\n=== Test Complete ===");
}

