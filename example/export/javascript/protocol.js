// Auto-generated JavaScript protocol code

class ProtocolLogin {
    constructor() {
        this.a = 0;
        this.b = 0;
        this.c = 0;
        this.d = 0;
        this.e = 0;
        this.f = 0;
        this.g = 0;
        this.h = 0;
        this.i = 0.0;
        this.j = 0.0;
        this.k = false;
        this.l = 0;
        this.m = "";
        this.n = [];
        this.o = [];
        this.q = new Account();
        this.r = [];
    }

    decode(data) {
        if (!data || data.length === 0) return;
        if (data.length > 0 && data[0] === 123) { // '{'
            try {
                const jsonStr = new TextDecoder().decode(data);
                const obj = JSON.parse(jsonStr);
                if (Object.prototype.hasOwnProperty.call(obj, 'a')) {
                	this.a = Number(obj.a);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'b')) {
                	this.b = Number(obj.b);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'c')) {
                	this.c = Number(obj.c);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'd')) {
                	this.d = Number(obj.d);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'e')) {
                	this.e = Number(obj.e);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'f')) {
                	this.f = Number(obj.f);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'g')) {
                	this.g = Number(obj.g);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'h')) {
                	this.h = Number(obj.h);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'i')) {
                	this.i = Number(obj.i);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'j')) {
                	this.j = Number(obj.j);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'k')) {
                	this.k = Boolean(obj.k);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'l')) {
                	this.l = Number(obj.l);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'm')) {
                	this.m = String(obj.m);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'n')) {
                	this.n = Array.isArray(obj.n) ? obj.n.map(n => Number(n)) : [];
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'o')) {
                	this.o = Array.isArray(obj.o) ? obj.o.map(s => String(s)) : [];
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'q')) {
                	this.q.decode(new TextEncoder().encode(JSON.stringify(obj.q)));
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'r')) {
                	this.r = [];
                	if (Array.isArray(obj.r)) {
                		for (const item of obj.r) {
                			const itemObj = new Account();
                			itemObj.decode(new TextEncoder().encode(JSON.stringify(item)));
                			this.r.push(itemObj);
                		}
                	}
                }
            } catch (e) {
                console.error('JSON parse error:', e);
            }
            return;
        }
        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
        let pointer = 0;
        if (pointer + 1 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.a = view.getInt8(pointer);
        pointer += 1;
        if (pointer + 1 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.b = view.getUint8(pointer);
        pointer += 1;
        if (pointer + 2 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.c = view.getInt16(pointer, true);
        pointer += 2;
        if (pointer + 2 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.d = view.getUint16(pointer, true);
        pointer += 2;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.e = view.getInt32(pointer, true);
        pointer += 4;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.f = view.getUint32(pointer, true);
        pointer += 4;
        if (pointer + 8 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.g = Number(view.getBigInt64(pointer, true));
        pointer += 8;
        if (pointer + 8 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.h = Number(view.getBigUint64(pointer, true));
        pointer += 8;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.i = view.getFloat32(pointer, true);
        pointer += 4;
        if (pointer + 8 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.j = view.getFloat64(pointer, true);
        pointer += 8;
        if (pointer + 1 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.k = data[pointer] !== 0;
        pointer += 1;
        if (pointer + 1 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.l = view.getUint8(pointer);
        pointer += 1;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.m = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const count = view.getInt32(pointer, true);
        	pointer += 4;
        	if (count < 0 || count > 1048576) throw new Error('protocol: count out of range: ' + count);
        	this.n = [];
        	for (let i = 0; i < count; i++) {
        		if (pointer + 1 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        		this.n.push(view.getInt8(pointer));
        		pointer += 1;
        	}
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const count = view.getInt32(pointer, true);
        	pointer += 4;
        	if (count < 0 || count > 1048576) throw new Error('protocol: count out of range: ' + count);
        	this.o = [];
        	for (let i = 0; i < count; i++) {
        		if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        		{
        			const len = view.getInt32(pointer, true);
        			pointer += 4;
        			if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        			if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        			this.o.push(new TextDecoder().decode(data.slice(pointer, pointer + len)));
        			pointer += len;
        		}
        	}
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.q = new Account();
        	this.q.decode(new Uint8Array(data.buffer, data.byteOffset + pointer, len));
        	pointer += len;
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const count = view.getInt32(pointer, true);
        	pointer += 4;
        	if (count < 0 || count > 1048576) throw new Error('protocol: count out of range: ' + count);
        	this.r = [];
        	for (let i = 0; i < count; i++) {
        		if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        		{
        			const len = view.getInt32(pointer, true);
        			pointer += 4;
        			if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        			if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        			const value = new Account();
        			value.decode(new Uint8Array(data.buffer, data.byteOffset + pointer, len));
        			this.r.push(value);
        			pointer += len;
        		}
        	}
        }
    }

    encodeBinary() {
        const buffer = [];
        {
        	const buf = new ArrayBuffer(1);
        	new DataView(buf).setInt8(0, this.a);
        	buffer.push(...new Uint8Array(buf));
        }
        buffer.push(this.b & 0xFF);
        {
        	const buf = new ArrayBuffer(2);
        	new DataView(buf).setInt16(0, this.c, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(2);
        	new DataView(buf).setUint16(0, this.d, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(4);
        	new DataView(buf).setInt32(0, this.e, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(4);
        	new DataView(buf).setUint32(0, this.f, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(8);
        	new DataView(buf).setBigInt64(0, BigInt(this.g), true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(8);
        	new DataView(buf).setBigUint64(0, BigInt(this.h), true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(4);
        	new DataView(buf).setFloat32(0, this.i, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const buf = new ArrayBuffer(8);
        	new DataView(buf).setFloat64(0, this.j, true);
        	buffer.push(...new Uint8Array(buf));
        }
        buffer.push(this.k ? 1 : 0);
        buffer.push(this.l & 0xFF);
        {
        	const bytes = new TextEncoder().encode(this.m);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        {
        	const countBuf = new ArrayBuffer(4);
        	new DataView(countBuf).setInt32(0, this.n.length, true);
        	buffer.push(...new Uint8Array(countBuf));
        	for (const item of this.n) {
        		{
        			const buf = new ArrayBuffer(1);
        			new DataView(buf).setInt8(0, item);
        			buffer.push(...new Uint8Array(buf));
        		}
        	}
        }
        {
        	const countBuf = new ArrayBuffer(4);
        	new DataView(countBuf).setInt32(0, this.o.length, true);
        	buffer.push(...new Uint8Array(countBuf));
        	for (const item of this.o) {
        		{
        			const bytes = new TextEncoder().encode(item);
        			const lenBuf = new ArrayBuffer(4);
        			new DataView(lenBuf).setInt32(0, bytes.length, true);
        			buffer.push(...new Uint8Array(lenBuf));
        			buffer.push(...bytes);
        		}
        	}
        }
        {
        	const bytes = this.q.encodeBinary();
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        {
        	const countBuf = new ArrayBuffer(4);
        	new DataView(countBuf).setInt32(0, this.r.length, true);
        	buffer.push(...new Uint8Array(countBuf));
        	for (const item of this.r) {
        		{
        			const bytes = item.encodeBinary();
        			const lenBuf = new ArrayBuffer(4);
        			new DataView(lenBuf).setInt32(0, bytes.length, true);
        			buffer.push(...new Uint8Array(lenBuf));
        			buffer.push(...bytes);
        		}
        	}
        }
        return new Uint8Array(buffer);
    }

    encodeJson() {
        return JSON.stringify(this);
    }
}

class Account {
    constructor() {
        this.nickname = "";
        this.password = "";
    }

    decode(data) {
        if (!data || data.length === 0) return;
        if (data.length > 0 && data[0] === 123) { // '{'
            try {
                const jsonStr = new TextDecoder().decode(data);
                const obj = JSON.parse(jsonStr);
                if (Object.prototype.hasOwnProperty.call(obj, 'nickname')) {
                	this.nickname = String(obj.nickname);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'password')) {
                	this.password = String(obj.password);
                }
            } catch (e) {
                console.error('JSON parse error:', e);
            }
            return;
        }
        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
        let pointer = 0;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.nickname = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.password = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
    }

    encodeBinary() {
        const buffer = [];
        {
        	const bytes = new TextEncoder().encode(this.nickname);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        {
        	const bytes = new TextEncoder().encode(this.password);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        return new Uint8Array(buffer);
    }

    encodeJson() {
        return JSON.stringify(this);
    }
}

class ProtocolServerLogin {
    constructor() {
        this.id = "";
        this.token = "";
    }

    decode(data) {
        if (!data || data.length === 0) return;
        if (data.length > 0 && data[0] === 123) { // '{'
            try {
                const jsonStr = new TextDecoder().decode(data);
                const obj = JSON.parse(jsonStr);
                if (Object.prototype.hasOwnProperty.call(obj, 'id')) {
                	this.id = String(obj.id);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'token')) {
                	this.token = String(obj.token);
                }
            } catch (e) {
                console.error('JSON parse error:', e);
            }
            return;
        }
        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
        let pointer = 0;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.id = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.token = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
    }

    encodeBinary() {
        const buffer = [];
        {
        	const bytes = new TextEncoder().encode(this.id);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        {
        	const bytes = new TextEncoder().encode(this.token);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        return new Uint8Array(buffer);
    }

    encodeJson() {
        return JSON.stringify(this);
    }
}

class ProtocolServerState {
    constructor() {
        this.state = 0;
        this.value = "";
    }

    decode(data) {
        if (!data || data.length === 0) return;
        if (data.length > 0 && data[0] === 123) { // '{'
            try {
                const jsonStr = new TextDecoder().decode(data);
                const obj = JSON.parse(jsonStr);
                if (Object.prototype.hasOwnProperty.call(obj, 'state')) {
                	this.state = Number(obj.state);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'value')) {
                	this.value = String(obj.value);
                }
            } catch (e) {
                console.error('JSON parse error:', e);
            }
            return;
        }
        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
        let pointer = 0;
        if (pointer + 2 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.state = view.getUint16(pointer, true);
        pointer += 2;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.value = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
    }

    encodeBinary() {
        const buffer = [];
        {
        	const buf = new ArrayBuffer(2);
        	new DataView(buf).setUint16(0, this.state, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const bytes = new TextEncoder().encode(this.value);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        return new Uint8Array(buffer);
    }

    encodeJson() {
        return JSON.stringify(this);
    }
}

class ProtocolServerCommand {
    constructor() {
        this.command = 0;
        this.value = "";
    }

    decode(data) {
        if (!data || data.length === 0) return;
        if (data.length > 0 && data[0] === 123) { // '{'
            try {
                const jsonStr = new TextDecoder().decode(data);
                const obj = JSON.parse(jsonStr);
                if (Object.prototype.hasOwnProperty.call(obj, 'command')) {
                	this.command = Number(obj.command);
                }
                if (Object.prototype.hasOwnProperty.call(obj, 'value')) {
                	this.value = String(obj.value);
                }
            } catch (e) {
                console.error('JSON parse error:', e);
            }
            return;
        }
        const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
        let pointer = 0;
        if (pointer + 2 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        this.command = view.getUint16(pointer, true);
        pointer += 2;
        if (pointer + 4 > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        {
        	const len = view.getInt32(pointer, true);
        	pointer += 4;
        	if (len < 0 || len > 16777216) throw new Error('protocol: length out of range: ' + len);
        	if (pointer + len > data.length) throw new Error('protocol: unexpected end of data at offset ' + pointer);
        	this.value = new TextDecoder().decode(data.slice(pointer, pointer + len));
        	pointer += len;
        }
    }

    encodeBinary() {
        const buffer = [];
        {
        	const buf = new ArrayBuffer(2);
        	new DataView(buf).setUint16(0, this.command, true);
        	buffer.push(...new Uint8Array(buf));
        }
        {
        	const bytes = new TextEncoder().encode(this.value);
        	const lenBuf = new ArrayBuffer(4);
        	new DataView(lenBuf).setInt32(0, bytes.length, true);
        	buffer.push(...new Uint8Array(lenBuf));
        	buffer.push(...bytes);
        }
        return new Uint8Array(buffer);
    }

    encodeJson() {
        return JSON.stringify(this);
    }
}

// Export for Node.js/ES6 modules
if (typeof module !== 'undefined' && module.exports) {
    module.exports = {
        ProtocolLogin: ProtocolLogin,
        Account: Account,
        ProtocolServerLogin: ProtocolServerLogin,
        ProtocolServerState: ProtocolServerState,
        ProtocolServerCommand: ProtocolServerCommand
    };
}

// Export for browser/global scope
if (typeof window !== 'undefined') {
    window.ProtocolLogin = ProtocolLogin;
    window.Account = Account;
    window.ProtocolServerLogin = ProtocolServerLogin;
    window.ProtocolServerState = ProtocolServerState;
    window.ProtocolServerCommand = ProtocolServerCommand;
}
