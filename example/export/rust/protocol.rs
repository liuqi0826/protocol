// Auto-generated Rust protocol code
use serde::{Deserialize, Serialize};
use std::collections::BTreeMap;
use std::io::{Read, Write};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProtocolLogin {
    pub a: i8,
    pub b: u8,
    pub c: i16,
    pub d: u16,
    pub e: i32,
    pub f: u32,
    pub g: i64,
    pub h: u64,
    pub i: f32,
    pub j: f64,
    pub k: bool,
    pub l: u8,
    pub m: String,
    pub n: Vec<i8>,
    pub o: Vec<String>,
    pub q: Account,
    pub r: Vec<Account>,
}

impl ProtocolLogin {
    pub fn new() -> Self {
        ProtocolLogin {
            a: 0,
            b: 0,
            c: 0,
            d: 0,
            e: 0,
            f: 0,
            g: 0,
            h: 0,
            i: 0.0,
            j: 0.0,
            k: false,
            l: 0,
            m: String::new(),
            n: Vec::new(),
            o: Vec::new(),
            q: Account::new(),
            r: Vec::new(),
        }
    }

    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {
        if data.is_empty() {
            return Err("Empty data".to_string());
        }
        if data[0] == b'{' {
            match serde_json::from_slice::<ProtocolLogin>(data) {
                Ok(obj) => {
                    *self = obj;
                    Ok(())
                },
                Err(e) => Err(format!("JSON parse error: {}", e)),
            }
        } else {
            let mut cursor = std::io::Cursor::new(data);
            self.a = {
                let mut buf = [0u8; 1];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read int8 failed".to_string()); }
                buf[0] as i8
            };
            self.b = {
                let mut buf = [0u8; 1];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read uint8 failed".to_string()); }
                buf[0]
            };
            self.c = {
                let mut buf = [0u8; 2];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read i16 failed".to_string()); }
                i16::from_le_bytes(buf)
            };
            self.d = {
                let mut buf = [0u8; 2];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read u16 failed".to_string()); }
                u16::from_le_bytes(buf)
            };
            self.e = {
                let mut buf = [0u8; 4];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read i32 failed".to_string()); }
                i32::from_le_bytes(buf)
            };
            self.f = {
                let mut buf = [0u8; 4];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read u32 failed".to_string()); }
                u32::from_le_bytes(buf)
            };
            self.g = {
                let mut buf = [0u8; 8];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read i64 failed".to_string()); }
                i64::from_le_bytes(buf)
            };
            self.h = {
                let mut buf = [0u8; 8];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read u64 failed".to_string()); }
                u64::from_le_bytes(buf)
            };
            self.i = {
                let mut buf = [0u8; 4];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read f32 failed".to_string()); }
                f32::from_le_bytes(buf)
            };
            self.j = {
                let mut buf = [0u8; 8];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read f64 failed".to_string()); }
                f64::from_le_bytes(buf)
            };
            self.k = {
                let mut buf = [0u8; 1];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read bool failed".to_string()); }
                buf[0] != 0
            };
            self.l = {
                let mut buf = [0u8; 1];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read uint8 failed".to_string()); }
                buf[0]
            };
            self.m = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            self.n = {
                let mut count_buf = [0u8; 4];
                if cursor.read_exact(&mut count_buf).is_err() { return Err("Read array count failed".to_string()); }
                let count = i32::from_le_bytes(count_buf) as usize;
                let mut vec = Vec::with_capacity(count);
                for _ in 0..count {
                    let mut buf = [0u8; 1];
                    if cursor.read_exact(&mut buf).is_err() { return Err("Read int8 failed".to_string()); }
                    vec.push(buf[0] as i8);
                }
                vec
            };
            self.o = {
                let mut count_buf = [0u8; 4];
                if cursor.read_exact(&mut count_buf).is_err() { return Err("Read array count failed".to_string()); }
                let count = i32::from_le_bytes(count_buf) as usize;
                let mut vec = Vec::with_capacity(count);
                for _ in 0..count {
                    let mut len_buf = [0u8; 4];
                    if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                    let len = i32::from_le_bytes(len_buf) as usize;
                    let mut str_buf = vec![0u8; len];
                    if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                    vec.push(String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?);
                }
                vec
            };
            self.q = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read object length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut obj_buf = vec![0u8; len];
                if cursor.read_exact(&mut obj_buf).is_err() { return Err("Read object data failed".to_string()); }
                let mut obj = Account::new();
                obj.decode(&obj_buf)?;
                obj
            };
            self.r = {
                let mut count_buf = [0u8; 4];
                if cursor.read_exact(&mut count_buf).is_err() { return Err("Read array count failed".to_string()); }
                let count = i32::from_le_bytes(count_buf) as usize;
                let mut vec = Vec::with_capacity(count);
                for _ in 0..count {
                    let mut len_buf = [0u8; 4];
                    if cursor.read_exact(&mut len_buf).is_err() { return Err("Read object length failed".to_string()); }
                    let len = i32::from_le_bytes(len_buf) as usize;
                    let mut obj_buf = vec![0u8; len];
                    if cursor.read_exact(&mut obj_buf).is_err() { return Err("Read object data failed".to_string()); }
                    let mut obj = Account::new();
                    obj.decode(&obj_buf)?;
                    vec.push(obj);
                }
                vec
            };
            Ok(())
        }
    }

    pub fn encode_binary(&self) -> Vec<u8> {
        let mut buffer = Vec::new();
        buffer.push(*&self.a as u8);
        buffer.push(*&self.b);
        buffer.extend_from_slice(&(*&self.c).to_le_bytes());
        buffer.extend_from_slice(&(*&self.d).to_le_bytes());
        buffer.extend_from_slice(&(*&self.e).to_le_bytes());
        buffer.extend_from_slice(&(*&self.f).to_le_bytes());
        buffer.extend_from_slice(&(*&self.g).to_le_bytes());
        buffer.extend_from_slice(&(*&self.h).to_le_bytes());
        buffer.extend_from_slice(&(*&self.i).to_le_bytes());
        buffer.extend_from_slice(&(*&self.j).to_le_bytes());
        buffer.push(if *&self.k { 1 } else { 0 });
        buffer.push(*&self.l);
        {
            let len = &self.m.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.m.as_bytes());
        }
        {
            let count = &self.n.len() as i32;
            buffer.extend_from_slice(&count.to_le_bytes());
            for item in &self.n {
                buffer.push(*item as u8);
            }
        }
        {
            let count = &self.o.len() as i32;
            buffer.extend_from_slice(&count.to_le_bytes());
            for item in &self.o {
                let len = item.len() as i32;
                buffer.extend_from_slice(&len.to_le_bytes());
                buffer.extend_from_slice(item.as_bytes());
            }
        }
        {
            let data = &self.q.encode_binary();
            let len = data.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&data);
        }
        {
            let count = &self.r.len() as i32;
            buffer.extend_from_slice(&count.to_le_bytes());
            for item in &self.r {
                let data = item.encode_binary();
                let len = data.len() as i32;
                buffer.extend_from_slice(&len.to_le_bytes());
                buffer.extend_from_slice(&data);
            }
        }
        buffer
    }

    pub fn encode_json(&self) -> Result<String, String> {
        match serde_json::to_string(self) {
            Ok(s) => Ok(s),
            Err(e) => Err(format!("JSON encode error: {}", e)),
        }
    }
}

impl Default for ProtocolLogin {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Account {
    pub nickname: String,
    pub password: String,
}

impl Account {
    pub fn new() -> Self {
        Account {
            nickname: String::new(),
            password: String::new(),
        }
    }

    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {
        if data.is_empty() {
            return Err("Empty data".to_string());
        }
        if data[0] == b'{' {
            match serde_json::from_slice::<Account>(data) {
                Ok(obj) => {
                    *self = obj;
                    Ok(())
                },
                Err(e) => Err(format!("JSON parse error: {}", e)),
            }
        } else {
            let mut cursor = std::io::Cursor::new(data);
            self.nickname = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            self.password = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            Ok(())
        }
    }

    pub fn encode_binary(&self) -> Vec<u8> {
        let mut buffer = Vec::new();
        {
            let len = &self.nickname.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.nickname.as_bytes());
        }
        {
            let len = &self.password.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.password.as_bytes());
        }
        buffer
    }

    pub fn encode_json(&self) -> Result<String, String> {
        match serde_json::to_string(self) {
            Ok(s) => Ok(s),
            Err(e) => Err(format!("JSON encode error: {}", e)),
        }
    }
}

impl Default for Account {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProtocolServerLogin {
    pub id: String,
    pub token: String,
}

impl ProtocolServerLogin {
    pub fn new() -> Self {
        ProtocolServerLogin {
            id: String::new(),
            token: String::new(),
        }
    }

    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {
        if data.is_empty() {
            return Err("Empty data".to_string());
        }
        if data[0] == b'{' {
            match serde_json::from_slice::<ProtocolServerLogin>(data) {
                Ok(obj) => {
                    *self = obj;
                    Ok(())
                },
                Err(e) => Err(format!("JSON parse error: {}", e)),
            }
        } else {
            let mut cursor = std::io::Cursor::new(data);
            self.id = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            self.token = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            Ok(())
        }
    }

    pub fn encode_binary(&self) -> Vec<u8> {
        let mut buffer = Vec::new();
        {
            let len = &self.id.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.id.as_bytes());
        }
        {
            let len = &self.token.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.token.as_bytes());
        }
        buffer
    }

    pub fn encode_json(&self) -> Result<String, String> {
        match serde_json::to_string(self) {
            Ok(s) => Ok(s),
            Err(e) => Err(format!("JSON encode error: {}", e)),
        }
    }
}

impl Default for ProtocolServerLogin {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProtocolServerState {
    pub state: u16,
    pub value: String,
}

impl ProtocolServerState {
    pub fn new() -> Self {
        ProtocolServerState {
            state: 0,
            value: String::new(),
        }
    }

    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {
        if data.is_empty() {
            return Err("Empty data".to_string());
        }
        if data[0] == b'{' {
            match serde_json::from_slice::<ProtocolServerState>(data) {
                Ok(obj) => {
                    *self = obj;
                    Ok(())
                },
                Err(e) => Err(format!("JSON parse error: {}", e)),
            }
        } else {
            let mut cursor = std::io::Cursor::new(data);
            self.state = {
                let mut buf = [0u8; 2];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read u16 failed".to_string()); }
                u16::from_le_bytes(buf)
            };
            self.value = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            Ok(())
        }
    }

    pub fn encode_binary(&self) -> Vec<u8> {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(&(*&self.state).to_le_bytes());
        {
            let len = &self.value.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.value.as_bytes());
        }
        buffer
    }

    pub fn encode_json(&self) -> Result<String, String> {
        match serde_json::to_string(self) {
            Ok(s) => Ok(s),
            Err(e) => Err(format!("JSON encode error: {}", e)),
        }
    }
}

impl Default for ProtocolServerState {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProtocolServerCommand {
    pub command: u16,
    pub value: String,
}

impl ProtocolServerCommand {
    pub fn new() -> Self {
        ProtocolServerCommand {
            command: 0,
            value: String::new(),
        }
    }

    pub fn decode(&mut self, data: &[u8]) -> Result<(), String> {
        if data.is_empty() {
            return Err("Empty data".to_string());
        }
        if data[0] == b'{' {
            match serde_json::from_slice::<ProtocolServerCommand>(data) {
                Ok(obj) => {
                    *self = obj;
                    Ok(())
                },
                Err(e) => Err(format!("JSON parse error: {}", e)),
            }
        } else {
            let mut cursor = std::io::Cursor::new(data);
            self.command = {
                let mut buf = [0u8; 2];
                if cursor.read_exact(&mut buf).is_err() { return Err("Read u16 failed".to_string()); }
                u16::from_le_bytes(buf)
            };
            self.value = {
                let mut len_buf = [0u8; 4];
                if cursor.read_exact(&mut len_buf).is_err() { return Err("Read string length failed".to_string()); }
                let len = i32::from_le_bytes(len_buf) as usize;
                let mut str_buf = vec![0u8; len];
                if cursor.read_exact(&mut str_buf).is_err() { return Err("Read string data failed".to_string()); }
                String::from_utf8(str_buf).map_err(|e| format!("Invalid UTF-8: {}", e))?
            };
            Ok(())
        }
    }

    pub fn encode_binary(&self) -> Vec<u8> {
        let mut buffer = Vec::new();
        buffer.extend_from_slice(&(*&self.command).to_le_bytes());
        {
            let len = &self.value.len() as i32;
            buffer.extend_from_slice(&len.to_le_bytes());
            buffer.extend_from_slice(&self.value.as_bytes());
        }
        buffer
    }

    pub fn encode_json(&self) -> Result<String, String> {
        match serde_json::to_string(self) {
            Ok(s) => Ok(s),
            Err(e) => Err(format!("JSON encode error: {}", e)),
        }
    }
}

impl Default for ProtocolServerCommand {
    fn default() -> Self {
        Self::new()
    }
}

