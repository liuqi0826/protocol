const std = @import("std");
const protocol = @import("protocol.zig");

// Run from example/export/zig:
//   zig build-exe test_zig.zig && ./test_zig
pub fn main() !void {
    const allocator = std.heap.page_allocator;

    var login: protocol.ProtocolLogin = .{};
    defer login.deinit(allocator);

    login.a = -10;
    login.b = 20;
    login.m = try allocator.dupe(u8, "Hello World");
    login.n = try allocator.dupe(i8, &[_]i8{ 1, -2, 3, -4 });
    login.q.nickname = try allocator.dupe(u8, "user123");
    login.q.password = try allocator.dupe(u8, "pass456");

    const binary = try login.encodeBinary(allocator);
    defer allocator.free(binary);

    var decoded: protocol.ProtocolLogin = .{};
    defer decoded.deinit(allocator);
    try decoded.decode(binary, allocator);

    if (decoded.a != login.a or !std.mem.eql(u8, decoded.m, login.m)) {
        std.debug.print("Binary round-trip FAILED\n", .{});
        return error.TestFailed;
    }
    std.debug.print("Zig binary round-trip PASSED ({d} bytes)\n", .{binary.len});
}
