// Auto-generated Zig protocol code
const std = @import("std");

pub const ProtocolLogin = struct {
    a: i8 = 0,
    b: u8 = 0,
    c: i16 = 0,
    d: u16 = 0,
    e: i32 = 0,
    f: u32 = 0,
    g: i64 = 0,
    h: u64 = 0,
    i: f32 = 0,
    j: f64 = 0,
    k: bool = false,
    l: u8 = 0,
    m: []const u8 = "",
    n: []i8 = &.{},
    o: [][]const u8 = &.{},
    q: Account = .{},
    r: []Account = &.{},

    pub fn deinit(self: *ProtocolLogin, allocator: std.mem.Allocator) void {
        if (self.m.len > 0) allocator.free(self.m);
        {
            if (self.n.len > 0) allocator.free(self.n);
        }
        {
            for (self.o) |item| {
                if (item.len > 0) allocator.free(item);
            }
            if (self.o.len > 0) allocator.free(self.o);
        }
        self.q.deinit(allocator);
        {
            for (self.r) |*item| {
                item.deinit(allocator);
            }
            if (self.r.len > 0) allocator.free(self.r);
        }
        self.* = .{};
    }

    pub fn decode(self: *ProtocolLogin, data: []const u8, allocator: std.mem.Allocator) !void {
        if (data.len == 0) return;
        if (data[0] == '{') {
            const parsed = try std.json.parseFromSlice(ProtocolLogin, allocator, data, .{ .allocate = .alloc_always });
            defer parsed.deinit();
            self.deinit(allocator);
            self.* = try cloneProtocolLogin(parsed.value, allocator);
            return;
        }
        var pointer: usize = 0;
        if (pointer + 1 > data.len) return error.UnexpectedEnd;
        self.a = @bitCast(data[pointer]);
        pointer += 1;
        if (pointer + 1 > data.len) return error.UnexpectedEnd;
        self.b = data[pointer];
        pointer += 1;
        if (pointer + 2 > data.len) return error.UnexpectedEnd;
        self.c = std.mem.readInt(i16, data[pointer..][0..2], .little);
        pointer += 2;
        if (pointer + 2 > data.len) return error.UnexpectedEnd;
        self.d = std.mem.readInt(u16, data[pointer..][0..2], .little);
        pointer += 2;
        if (pointer + 4 > data.len) return error.UnexpectedEnd;
        self.e = std.mem.readInt(i32, data[pointer..][0..4], .little);
        pointer += 4;
        if (pointer + 4 > data.len) return error.UnexpectedEnd;
        self.f = std.mem.readInt(u32, data[pointer..][0..4], .little);
        pointer += 4;
        if (pointer + 8 > data.len) return error.UnexpectedEnd;
        self.g = std.mem.readInt(i64, data[pointer..][0..8], .little);
        pointer += 8;
        if (pointer + 8 > data.len) return error.UnexpectedEnd;
        self.h = std.mem.readInt(u64, data[pointer..][0..8], .little);
        pointer += 8;
        if (pointer + 4 > data.len) return error.UnexpectedEnd;
        self.i = @bitCast(std.mem.readInt(u32, data[pointer..][0..4], .little));
        pointer += 4;
        if (pointer + 8 > data.len) return error.UnexpectedEnd;
        self.j = @bitCast(std.mem.readInt(u64, data[pointer..][0..8], .little));
        pointer += 8;
        if (pointer + 1 > data.len) return error.UnexpectedEnd;
        self.k = data[pointer] != 0;
        pointer += 1;
        if (pointer + 1 > data.len) return error.UnexpectedEnd;
        self.l = data[pointer];
        pointer += 1;
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.m = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const count: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            const items = try allocator.alloc(i8, count);
            errdefer allocator.free(items);
            for (items) |*item| {
                if (pointer + 1 > data.len) return error.UnexpectedEnd;
                item.* = @bitCast(data[pointer]);
                pointer += 1;
            }
            self.n = items;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const count: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            const items = try allocator.alloc([]const u8, count);
            errdefer allocator.free(items);
            for (items) |*item| {
                if (pointer + 4 > data.len) return error.UnexpectedEnd;
                const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
                pointer += 4;
                if (pointer + len > data.len) return error.UnexpectedEnd;
                item.* = try allocator.dupe(u8, data[pointer .. pointer + len]);
                pointer += len;
            }
            self.o = items;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            try self.q.decode(data[pointer .. pointer + len], allocator);
            pointer += len;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const count: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            const items = try allocator.alloc(Account, count);
            errdefer allocator.free(items);
            for (items) |*item| {
                if (pointer + 4 > data.len) return error.UnexpectedEnd;
                const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
                pointer += 4;
                if (pointer + len > data.len) return error.UnexpectedEnd;
                item.* = .{};
                try item.decode(data[pointer .. pointer + len], allocator);
                pointer += len;
            }
            self.r = items;
        }
    }

    pub fn encodeBinary(self: *const ProtocolLogin, allocator: std.mem.Allocator) ![]u8 {
        var list: std.ArrayList(u8) = .empty;
        errdefer list.deinit(allocator);
        try list.append(allocator, @bitCast(self.a));
        try list.append(allocator, self.b);
        {
            var buf: [2]u8 = undefined;
            std.mem.writeInt(i16, &buf, self.c, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [2]u8 = undefined;
            std.mem.writeInt(u16, &buf, self.d, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &buf, self.e, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [4]u8 = undefined;
            std.mem.writeInt(u32, &buf, self.f, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [8]u8 = undefined;
            std.mem.writeInt(i64, &buf, self.g, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [8]u8 = undefined;
            std.mem.writeInt(u64, &buf, self.h, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [4]u8 = undefined;
            std.mem.writeInt(u32, &buf, @bitCast(self.i), .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var buf: [8]u8 = undefined;
            std.mem.writeInt(u64, &buf, @bitCast(self.j), .little);
            try list.appendSlice(allocator, &buf);
        }
        try list.append(allocator, if (self.k) 1 else 0);
        try list.append(allocator, self.l);
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.m.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.m);
        }
        {
            var count_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &count_buf, @intCast(self.n.len), .little);
            try list.appendSlice(allocator, &count_buf);
            for (self.n) |item| {
                try list.append(allocator, @bitCast(item));
            }
        }
        {
            var count_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &count_buf, @intCast(self.o.len), .little);
            try list.appendSlice(allocator, &count_buf);
            for (self.o) |item| {
                {
                    var len_buf: [4]u8 = undefined;
                    std.mem.writeInt(i32, &len_buf, @intCast(item.len), .little);
                    try list.appendSlice(allocator, &len_buf);
                    try list.appendSlice(allocator, item);
                }
            }
        }
        {
            const nested = try self.q.encodeBinary(allocator);
            defer allocator.free(nested);
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(nested.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, nested);
        }
        {
            var count_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &count_buf, @intCast(self.r.len), .little);
            try list.appendSlice(allocator, &count_buf);
            for (self.r) |item| {
                {
                    const nested = try item.encodeBinary(allocator);
                    defer allocator.free(nested);
                    var len_buf: [4]u8 = undefined;
                    std.mem.writeInt(i32, &len_buf, @intCast(nested.len), .little);
                    try list.appendSlice(allocator, &len_buf);
                    try list.appendSlice(allocator, nested);
                }
            }
        }
        return try list.toOwnedSlice(allocator);
    }

    pub fn encodeJson(self: *const ProtocolLogin, allocator: std.mem.Allocator) ![]u8 {
        return try std.json.stringifyAlloc(allocator, self.*, .{});
    }
};

pub const Account = struct {
    nickname: []const u8 = "",
    password: []const u8 = "",

    pub fn deinit(self: *Account, allocator: std.mem.Allocator) void {
        if (self.nickname.len > 0) allocator.free(self.nickname);
        if (self.password.len > 0) allocator.free(self.password);
        self.* = .{};
    }

    pub fn decode(self: *Account, data: []const u8, allocator: std.mem.Allocator) !void {
        if (data.len == 0) return;
        if (data[0] == '{') {
            const parsed = try std.json.parseFromSlice(Account, allocator, data, .{ .allocate = .alloc_always });
            defer parsed.deinit();
            self.deinit(allocator);
            self.* = try cloneAccount(parsed.value, allocator);
            return;
        }
        var pointer: usize = 0;
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.nickname = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.password = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
    }

    pub fn encodeBinary(self: *const Account, allocator: std.mem.Allocator) ![]u8 {
        var list: std.ArrayList(u8) = .empty;
        errdefer list.deinit(allocator);
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.nickname.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.nickname);
        }
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.password.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.password);
        }
        return try list.toOwnedSlice(allocator);
    }

    pub fn encodeJson(self: *const Account, allocator: std.mem.Allocator) ![]u8 {
        return try std.json.stringifyAlloc(allocator, self.*, .{});
    }
};

pub const ProtocolServerLogin = struct {
    id: []const u8 = "",
    token: []const u8 = "",

    pub fn deinit(self: *ProtocolServerLogin, allocator: std.mem.Allocator) void {
        if (self.id.len > 0) allocator.free(self.id);
        if (self.token.len > 0) allocator.free(self.token);
        self.* = .{};
    }

    pub fn decode(self: *ProtocolServerLogin, data: []const u8, allocator: std.mem.Allocator) !void {
        if (data.len == 0) return;
        if (data[0] == '{') {
            const parsed = try std.json.parseFromSlice(ProtocolServerLogin, allocator, data, .{ .allocate = .alloc_always });
            defer parsed.deinit();
            self.deinit(allocator);
            self.* = try cloneProtocolServerLogin(parsed.value, allocator);
            return;
        }
        var pointer: usize = 0;
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.id = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.token = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
    }

    pub fn encodeBinary(self: *const ProtocolServerLogin, allocator: std.mem.Allocator) ![]u8 {
        var list: std.ArrayList(u8) = .empty;
        errdefer list.deinit(allocator);
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.id.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.id);
        }
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.token.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.token);
        }
        return try list.toOwnedSlice(allocator);
    }

    pub fn encodeJson(self: *const ProtocolServerLogin, allocator: std.mem.Allocator) ![]u8 {
        return try std.json.stringifyAlloc(allocator, self.*, .{});
    }
};

pub const ProtocolServerState = struct {
    state: u16 = 0,
    value: []const u8 = "",

    pub fn deinit(self: *ProtocolServerState, allocator: std.mem.Allocator) void {
        if (self.value.len > 0) allocator.free(self.value);
        self.* = .{};
    }

    pub fn decode(self: *ProtocolServerState, data: []const u8, allocator: std.mem.Allocator) !void {
        if (data.len == 0) return;
        if (data[0] == '{') {
            const parsed = try std.json.parseFromSlice(ProtocolServerState, allocator, data, .{ .allocate = .alloc_always });
            defer parsed.deinit();
            self.deinit(allocator);
            self.* = try cloneProtocolServerState(parsed.value, allocator);
            return;
        }
        var pointer: usize = 0;
        if (pointer + 2 > data.len) return error.UnexpectedEnd;
        self.state = std.mem.readInt(u16, data[pointer..][0..2], .little);
        pointer += 2;
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.value = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
    }

    pub fn encodeBinary(self: *const ProtocolServerState, allocator: std.mem.Allocator) ![]u8 {
        var list: std.ArrayList(u8) = .empty;
        errdefer list.deinit(allocator);
        {
            var buf: [2]u8 = undefined;
            std.mem.writeInt(u16, &buf, self.state, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.value.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.value);
        }
        return try list.toOwnedSlice(allocator);
    }

    pub fn encodeJson(self: *const ProtocolServerState, allocator: std.mem.Allocator) ![]u8 {
        return try std.json.stringifyAlloc(allocator, self.*, .{});
    }
};

pub const ProtocolServerCommand = struct {
    command: u16 = 0,
    value: []const u8 = "",

    pub fn deinit(self: *ProtocolServerCommand, allocator: std.mem.Allocator) void {
        if (self.value.len > 0) allocator.free(self.value);
        self.* = .{};
    }

    pub fn decode(self: *ProtocolServerCommand, data: []const u8, allocator: std.mem.Allocator) !void {
        if (data.len == 0) return;
        if (data[0] == '{') {
            const parsed = try std.json.parseFromSlice(ProtocolServerCommand, allocator, data, .{ .allocate = .alloc_always });
            defer parsed.deinit();
            self.deinit(allocator);
            self.* = try cloneProtocolServerCommand(parsed.value, allocator);
            return;
        }
        var pointer: usize = 0;
        if (pointer + 2 > data.len) return error.UnexpectedEnd;
        self.command = std.mem.readInt(u16, data[pointer..][0..2], .little);
        pointer += 2;
        {
            if (pointer + 4 > data.len) return error.UnexpectedEnd;
            const len: usize = @intCast(std.mem.readInt(i32, data[pointer..][0..4], .little));
            pointer += 4;
            if (pointer + len > data.len) return error.UnexpectedEnd;
            self.value = try allocator.dupe(u8, data[pointer .. pointer + len]);
            pointer += len;
        }
    }

    pub fn encodeBinary(self: *const ProtocolServerCommand, allocator: std.mem.Allocator) ![]u8 {
        var list: std.ArrayList(u8) = .empty;
        errdefer list.deinit(allocator);
        {
            var buf: [2]u8 = undefined;
            std.mem.writeInt(u16, &buf, self.command, .little);
            try list.appendSlice(allocator, &buf);
        }
        {
            var len_buf: [4]u8 = undefined;
            std.mem.writeInt(i32, &len_buf, @intCast(self.value.len), .little);
            try list.appendSlice(allocator, &len_buf);
            try list.appendSlice(allocator, self.value);
        }
        return try list.toOwnedSlice(allocator);
    }

    pub fn encodeJson(self: *const ProtocolServerCommand, allocator: std.mem.Allocator) ![]u8 {
        return try std.json.stringifyAlloc(allocator, self.*, .{});
    }
};

fn cloneProtocolLogin(src: ProtocolLogin, allocator: std.mem.Allocator) !ProtocolLogin {
    var dst: ProtocolLogin = .{};
    dst.a = src.a;
    dst.b = src.b;
    dst.c = src.c;
    dst.d = src.d;
    dst.e = src.e;
    dst.f = src.f;
    dst.g = src.g;
    dst.h = src.h;
    dst.i = src.i;
    dst.j = src.j;
    dst.k = src.k;
    dst.l = src.l;
    dst.m = try allocator.dupe(u8, src.m);
    {
        const items = try allocator.alloc(i8, src.n.len);
        for (src.n, 0..) |item, i| {
            items[i] = item;
        }
        dst.n = items;
    }
    {
        const items = try allocator.alloc([]const u8, src.o.len);
        for (src.o, 0..) |item, i| {
            items[i] = try allocator.dupe(u8, item);
        }
        dst.o = items;
    }
    dst.q = try cloneAccount(src.q, allocator);
    {
        const items = try allocator.alloc(Account, src.r.len);
        for (src.r, 0..) |item, i| {
            items[i] = try cloneAccount(item, allocator);
        }
        dst.r = items;
    }
    return dst;
}

fn cloneAccount(src: Account, allocator: std.mem.Allocator) !Account {
    var dst: Account = .{};
    dst.nickname = try allocator.dupe(u8, src.nickname);
    dst.password = try allocator.dupe(u8, src.password);
    return dst;
}

fn cloneProtocolServerLogin(src: ProtocolServerLogin, allocator: std.mem.Allocator) !ProtocolServerLogin {
    var dst: ProtocolServerLogin = .{};
    dst.id = try allocator.dupe(u8, src.id);
    dst.token = try allocator.dupe(u8, src.token);
    return dst;
}

fn cloneProtocolServerState(src: ProtocolServerState, allocator: std.mem.Allocator) !ProtocolServerState {
    var dst: ProtocolServerState = .{};
    dst.state = src.state;
    dst.value = try allocator.dupe(u8, src.value);
    return dst;
}

fn cloneProtocolServerCommand(src: ProtocolServerCommand, allocator: std.mem.Allocator) !ProtocolServerCommand {
    var dst: ProtocolServerCommand = .{};
    dst.command = src.command;
    dst.value = try allocator.dupe(u8, src.value);
    return dst;
}

