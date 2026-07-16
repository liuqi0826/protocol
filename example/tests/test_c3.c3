module test_c3;
import protocol;
import std::io;

fn void main()
{
    ProtocolLogin login;
    login.init();
    login.a = -10;
    login.b = 20;
    login.m = "Hello World";
    login.n = { 1, -2, 3, -4 };
    login.q.init();
    login.q.nickname = "user123";
    login.q.password = "pass456";

    String binary = login.encode_binary(mem);
    ProtocolLogin decoded;
    decoded.init();
    decoded.decode((char[])binary);

    if (decoded.a != login.a)
    {
        io::printfn("C3 binary FAILED");
        return;
    }
    io::printfn("C3 binary round-trip PASSED (%d bytes)", binary.len);
    free(binary);
}
