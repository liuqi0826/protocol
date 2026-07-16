// Auto-generated C3 protocol code (C3 0.8+)
// Binary: little-endian, self-contained.
// JSON: encode via json::marshal; decode via json::parse + Object getters.
module protocol;

import std::io;
import std::encoding::json;
import std::collections::object;

struct ProtocolLogin
{
    ichar a;
    char b;
    short c;
    ushort d;
    int e;
    uint f;
    long g;
    ulong h;
    float i;
    double j;
    bool k;
    char l;
    String m;
    ichar[] n;
    String[] o;
    Account q;
    Account[] r;
}

fn void ProtocolLogin.init(&self)
{
    self.a = 0;
    self.b = 0;
    self.c = 0;
    self.d = 0;
    self.e = 0;
    self.f = 0;
    self.g = 0;
    self.h = 0;
    self.i = 0.0;
    self.j = 0.0;
    self.k = false;
    self.l = 0;
    self.m = "";
    self.n = {};
    self.o = {};
    self.q.init();
    self.r = {};
}

fn void ProtocolLogin.free(&self)
{
    free(self.n);
    self.n = {};
    free(self.o);
    self.o = {};
    self.q.free();
    foreach (&item : self.r)
    {
        item.free();
    }
    free(self.r);
    self.r = {};
}

fn void ProtocolLogin.decode_from_object(&self, Object* obj)
{
    if (try v = obj.get_ichar("a")) self.a = v;
    if (try v = obj.get_char("b")) self.b = v;
    if (try v = obj.get_short("c")) self.c = v;
    if (try v = obj.get_ushort("d")) self.d = v;
    if (try v = obj.get_int("e")) self.e = v;
    if (try v = obj.get_uint("f")) self.f = v;
    if (try v = obj.get_long("g")) self.g = v;
    if (try v = obj.get_ulong("h")) self.h = v;
    if (try v = obj.get_float("i")) self.i = (float)v;
    if (try v = obj.get_float("j")) self.j = v;
    if (try v = obj.get_bool("k")) self.k = v;
    if (try v = obj.get_char("l")) self.l = v;
    if (try v = obj.get_string("m")) self.m = v;
    if (try arr = obj.get("n"))
    {
        sz count = arr.get_len();
        self.n = mem::new_array(ichar, count);
        for (sz i = 0; i < count; i++)
        {
            if (try v = arr.get_ichar_at(i)) self.n[i] = v;
        }
    }
    if (try arr = obj.get("o"))
    {
        sz count = arr.get_len();
        self.o = mem::new_array(String, count);
        for (sz i = 0; i < count; i++)
        {
            if (try v = arr.get_string_at(i)) self.o[i] = v;
        }
    }
    if (try nested = obj.get("q"))
    {
        self.q.init();
        self.q.decode_from_object(nested);
    }
    if (try arr = obj.get("r"))
    {
        sz count = arr.get_len();
        self.r = mem::new_array(Account, count);
        for (sz i = 0; i < count; i++)
        {
            self.r[i].init();
            self.r[i].decode_from_object(arr.get_at(i));
        }
    }
}

fn void ProtocolLogin.decode(&self, char[] data)
{
    if (data.len == 0) return;
    if (data[0] == '{')
    {
        String s = (String)data;
        Object*? parsed = json::tparse(s);
        if (catch parsed) return;
        Object* obj = parsed;
        defer obj.free();
        self.decode_from_object(obj);
        return;
    }
    usz pointer = 0;
    if (pointer + 1 > data.len) return;
    self.a = (ichar)data[pointer];
    pointer += 1;
    if (pointer + 1 > data.len) return;
    self.b = data[pointer];
    pointer += 1;
    if (pointer + 2 > data.len) return;
    self.c = protocol_read_short(data, pointer);
    pointer += 2;
    if (pointer + 2 > data.len) return;
    self.d = protocol_read_ushort(data, pointer);
    pointer += 2;
    if (pointer + 4 > data.len) return;
    self.e = protocol_read_int(data, pointer);
    pointer += 4;
    if (pointer + 4 > data.len) return;
    self.f = protocol_read_uint(data, pointer);
    pointer += 4;
    if (pointer + 8 > data.len) return;
    self.g = protocol_read_long(data, pointer);
    pointer += 8;
    if (pointer + 8 > data.len) return;
    self.h = protocol_read_ulong(data, pointer);
    pointer += 8;
    if (pointer + 4 > data.len) return;
    self.i = protocol_read_float(data, pointer);
    pointer += 4;
    if (pointer + 8 > data.len) return;
    self.j = protocol_read_double(data, pointer);
    pointer += 8;
    if (pointer + 1 > data.len) return;
    self.k = data[pointer] != 0;
    pointer += 1;
    if (pointer + 1 > data.len) return;
    self.l = data[pointer];
    pointer += 1;
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.m = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
    {
        if (pointer + 4 > data.len) return;
        int count = protocol_read_int(data, pointer);
        pointer += 4;
        if (count < 0) return;
        self.n = mem::new_array(ichar, (usz)count);
        for (usz i = 0; i < (usz)count; i++)
        {
            if (pointer + 1 > data.len) return;
            self.n[i] = (ichar)data[pointer];
            pointer += 1;
        }
    }
    {
        if (pointer + 4 > data.len) return;
        int count = protocol_read_int(data, pointer);
        pointer += 4;
        if (count < 0) return;
        self.o = mem::new_array(String, (usz)count);
        for (usz i = 0; i < (usz)count; i++)
        {
            if (pointer + 4 > data.len) return;
            int len = protocol_read_int(data, pointer);
            pointer += 4;
            if (len < 0 || pointer + (usz)len > data.len) return;
            self.o[i] = (String)data[pointer:(usz)len];
            pointer += (usz)len;
        }
    }
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.q.init();
        self.q.decode(data[pointer:(usz)len]);
        pointer += (usz)len;
    }
    {
        if (pointer + 4 > data.len) return;
        int count = protocol_read_int(data, pointer);
        pointer += 4;
        if (count < 0) return;
        self.r = mem::new_array(Account, (usz)count);
        for (usz i = 0; i < (usz)count; i++)
        {
            if (pointer + 4 > data.len) return;
            int len = protocol_read_int(data, pointer);
            pointer += 4;
            if (len < 0 || pointer + (usz)len > data.len) return;
            self.r[i].init();
            self.r[i].decode(data[pointer:(usz)len]);
            pointer += (usz)len;
        }
    }
}

fn String ProtocolLogin.encode_binary(&self, Allocator allocator)
{
    DString buffer = dstring::temp();
    protocol_write_char(&buffer, (char)self.a);
    protocol_write_char(&buffer, self.b);
    protocol_write_short(&buffer, self.c);
    protocol_write_ushort(&buffer, self.d);
    protocol_write_int(&buffer, self.e);
    protocol_write_uint(&buffer, self.f);
    protocol_write_long(&buffer, self.g);
    protocol_write_ulong(&buffer, self.h);
    protocol_write_float(&buffer, self.i);
    protocol_write_double(&buffer, self.j);
    protocol_write_char(&buffer, self.k ? 1 : 0);
    protocol_write_char(&buffer, self.l);
    protocol_write_int(&buffer, (int)self.m.len);
    buffer.append(self.m);
    protocol_write_int(&buffer, (int)self.n.len);
    foreach (item : self.n)
    {
        protocol_write_char(&buffer, (char)item);
    }
    protocol_write_int(&buffer, (int)self.o.len);
    foreach (item : self.o)
    {
        protocol_write_int(&buffer, (int)item.len);
        buffer.append(item);
    }
    {
        String nested = self.q.encode_binary(allocator);
        protocol_write_int(&buffer, (int)nested.len);
        buffer.append(nested);
        free(nested);
    }
    protocol_write_int(&buffer, (int)self.r.len);
    foreach (item : self.r)
    {
        {
            String nested = item.encode_binary(allocator);
            protocol_write_int(&buffer, (int)nested.len);
            buffer.append(nested);
            free(nested);
        }
    }
    return buffer.copy_str(allocator);
}

fn String ProtocolLogin.encode_json(&self, Allocator allocator)
{
    return json::marshal(allocator, *self);
}

struct Account
{
    String nickname;
    String password;
}

fn void Account.init(&self)
{
    self.nickname = "";
    self.password = "";
}

fn void Account.free(&self)
{
}

fn void Account.decode_from_object(&self, Object* obj)
{
    if (try v = obj.get_string("nickname")) self.nickname = v;
    if (try v = obj.get_string("password")) self.password = v;
}

fn void Account.decode(&self, char[] data)
{
    if (data.len == 0) return;
    if (data[0] == '{')
    {
        String s = (String)data;
        Object*? parsed = json::tparse(s);
        if (catch parsed) return;
        Object* obj = parsed;
        defer obj.free();
        self.decode_from_object(obj);
        return;
    }
    usz pointer = 0;
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.nickname = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.password = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
}

fn String Account.encode_binary(&self, Allocator allocator)
{
    DString buffer = dstring::temp();
    protocol_write_int(&buffer, (int)self.nickname.len);
    buffer.append(self.nickname);
    protocol_write_int(&buffer, (int)self.password.len);
    buffer.append(self.password);
    return buffer.copy_str(allocator);
}

fn String Account.encode_json(&self, Allocator allocator)
{
    return json::marshal(allocator, *self);
}

struct ProtocolServerLogin
{
    String id;
    String token;
}

fn void ProtocolServerLogin.init(&self)
{
    self.id = "";
    self.token = "";
}

fn void ProtocolServerLogin.free(&self)
{
}

fn void ProtocolServerLogin.decode_from_object(&self, Object* obj)
{
    if (try v = obj.get_string("id")) self.id = v;
    if (try v = obj.get_string("token")) self.token = v;
}

fn void ProtocolServerLogin.decode(&self, char[] data)
{
    if (data.len == 0) return;
    if (data[0] == '{')
    {
        String s = (String)data;
        Object*? parsed = json::tparse(s);
        if (catch parsed) return;
        Object* obj = parsed;
        defer obj.free();
        self.decode_from_object(obj);
        return;
    }
    usz pointer = 0;
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.id = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.token = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
}

fn String ProtocolServerLogin.encode_binary(&self, Allocator allocator)
{
    DString buffer = dstring::temp();
    protocol_write_int(&buffer, (int)self.id.len);
    buffer.append(self.id);
    protocol_write_int(&buffer, (int)self.token.len);
    buffer.append(self.token);
    return buffer.copy_str(allocator);
}

fn String ProtocolServerLogin.encode_json(&self, Allocator allocator)
{
    return json::marshal(allocator, *self);
}

struct ProtocolServerState
{
    ushort state;
    String value;
}

fn void ProtocolServerState.init(&self)
{
    self.state = 0;
    self.value = "";
}

fn void ProtocolServerState.free(&self)
{
}

fn void ProtocolServerState.decode_from_object(&self, Object* obj)
{
    if (try v = obj.get_ushort("state")) self.state = v;
    if (try v = obj.get_string("value")) self.value = v;
}

fn void ProtocolServerState.decode(&self, char[] data)
{
    if (data.len == 0) return;
    if (data[0] == '{')
    {
        String s = (String)data;
        Object*? parsed = json::tparse(s);
        if (catch parsed) return;
        Object* obj = parsed;
        defer obj.free();
        self.decode_from_object(obj);
        return;
    }
    usz pointer = 0;
    if (pointer + 2 > data.len) return;
    self.state = protocol_read_ushort(data, pointer);
    pointer += 2;
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.value = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
}

fn String ProtocolServerState.encode_binary(&self, Allocator allocator)
{
    DString buffer = dstring::temp();
    protocol_write_ushort(&buffer, self.state);
    protocol_write_int(&buffer, (int)self.value.len);
    buffer.append(self.value);
    return buffer.copy_str(allocator);
}

fn String ProtocolServerState.encode_json(&self, Allocator allocator)
{
    return json::marshal(allocator, *self);
}

struct ProtocolServerCommand
{
    ushort command;
    String value;
}

fn void ProtocolServerCommand.init(&self)
{
    self.command = 0;
    self.value = "";
}

fn void ProtocolServerCommand.free(&self)
{
}

fn void ProtocolServerCommand.decode_from_object(&self, Object* obj)
{
    if (try v = obj.get_ushort("command")) self.command = v;
    if (try v = obj.get_string("value")) self.value = v;
}

fn void ProtocolServerCommand.decode(&self, char[] data)
{
    if (data.len == 0) return;
    if (data[0] == '{')
    {
        String s = (String)data;
        Object*? parsed = json::tparse(s);
        if (catch parsed) return;
        Object* obj = parsed;
        defer obj.free();
        self.decode_from_object(obj);
        return;
    }
    usz pointer = 0;
    if (pointer + 2 > data.len) return;
    self.command = protocol_read_ushort(data, pointer);
    pointer += 2;
    {
        if (pointer + 4 > data.len) return;
        int len = protocol_read_int(data, pointer);
        pointer += 4;
        if (len < 0 || pointer + (usz)len > data.len) return;
        self.value = (String)data[pointer:(usz)len];
        pointer += (usz)len;
    }
}

fn String ProtocolServerCommand.encode_binary(&self, Allocator allocator)
{
    DString buffer = dstring::temp();
    protocol_write_ushort(&buffer, self.command);
    protocol_write_int(&buffer, (int)self.value.len);
    buffer.append(self.value);
    return buffer.copy_str(allocator);
}

fn String ProtocolServerCommand.encode_json(&self, Allocator allocator)
{
    return json::marshal(allocator, *self);
}


fn void protocol_write_char(DString* buffer, char value)
{
    buffer.append_char(value);
}

fn void protocol_write_short(DString* buffer, short value)
{
    char[2] bytes = { (char)(value & 0xff), (char)((value >> 8) & 0xff) };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_ushort(DString* buffer, ushort value)
{
    protocol_write_short(buffer, (short)value);
}

fn void protocol_write_int(DString* buffer, int value)
{
    char[4] bytes = {
        (char)(value & 0xff),
        (char)((value >> 8) & 0xff),
        (char)((value >> 16) & 0xff),
        (char)((value >> 24) & 0xff)
    };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_uint(DString* buffer, uint value)
{
    protocol_write_int(buffer, (int)value);
}

fn void protocol_write_long(DString* buffer, long value)
{
    char[8] bytes = {
        (char)(value & 0xff),
        (char)((value >> 8) & 0xff),
        (char)((value >> 16) & 0xff),
        (char)((value >> 24) & 0xff),
        (char)((value >> 32) & 0xff),
        (char)((value >> 40) & 0xff),
        (char)((value >> 48) & 0xff),
        (char)((value >> 56) & 0xff)
    };
    buffer.append_bytes(&bytes);
}

fn void protocol_write_ulong(DString* buffer, ulong value)
{
    protocol_write_long(buffer, (long)value);
}

fn void protocol_write_float(DString* buffer, float value)
{
    uint bits = bitcast(value, uint);
    protocol_write_uint(buffer, bits);
}

fn void protocol_write_double(DString* buffer, double value)
{
    ulong bits = bitcast(value, ulong);
    protocol_write_ulong(buffer, bits);
}

fn short protocol_read_short(char[] data, usz pointer)
{
    return (short)((ushort)data[pointer] | ((ushort)data[pointer + 1] << 8));
}

fn ushort protocol_read_ushort(char[] data, usz pointer)
{
    return (ushort)protocol_read_short(data, pointer);
}

fn int protocol_read_int(char[] data, usz pointer)
{
    return (int)((uint)data[pointer]
        | ((uint)data[pointer + 1] << 8)
        | ((uint)data[pointer + 2] << 16)
        | ((uint)data[pointer + 3] << 24));
}

fn uint protocol_read_uint(char[] data, usz pointer)
{
    return (uint)protocol_read_int(data, pointer);
}

fn long protocol_read_long(char[] data, usz pointer)
{
    return (long)((ulong)data[pointer]
        | ((ulong)data[pointer + 1] << 8)
        | ((ulong)data[pointer + 2] << 16)
        | ((ulong)data[pointer + 3] << 24)
        | ((ulong)data[pointer + 4] << 32)
        | ((ulong)data[pointer + 5] << 40)
        | ((ulong)data[pointer + 6] << 48)
        | ((ulong)data[pointer + 7] << 56));
}

fn ulong protocol_read_ulong(char[] data, usz pointer)
{
    return (ulong)protocol_read_long(data, pointer);
}

fn float protocol_read_float(char[] data, usz pointer)
{
    return bitcast(protocol_read_uint(data, pointer), float);
}

fn double protocol_read_double(char[] data, usz pointer)
{
    return bitcast(protocol_read_ulong(data, pointer), double);
}
