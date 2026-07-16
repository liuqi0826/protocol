using System;
using System.IO;
using System.Text;
using protocol;

class TestCSharp
{
    static void Main()
    {
        Console.WriteLine("=== C# Protocol Test ===\n");

        // 创建测试数据
        var login = new ProtocolLogin
        {
            a = -10,
            b = 20,
            c = -300,
            d = 400,
            e = -5000,
            f = 6000,
            g = -70000,
            h = 80000,
            i = 3.14f,
            j = 2.718,
            k = true,
            l = 255,
            m = "Hello World",
            n = new sbyte[] { 1, -2, 3, -4 },
            o = new string[] { "test1", "test2", "test3" },
            q = new Account
            {
                nickname = "user123",
                password = "pass456"
            },
            r = new Account[]
            {
                new Account { nickname = "user1", password = "pass1" },
                new Account { nickname = "user2", password = "pass2" }
            }
        };

        // 测试 JSON 序列化
        Console.WriteLine("1. Testing JSON Serialization...");
        byte[] jsonData = login.EncodeJson();
        Console.WriteLine($"   JSON encoded: {jsonData.Length} bytes");
        File.WriteAllBytes("test_output_csharp.json", jsonData);

        // 测试 JSON 反序列化
        Console.WriteLine("2. Testing JSON Deserialization...");
        var loginFromJson = new ProtocolLogin();
        loginFromJson.Decoder(jsonData);
        if (VerifyLogin(login, loginFromJson))
        {
            Console.WriteLine("   ✓ JSON serialization/deserialization PASSED");
        }
        else
        {
            Console.WriteLine("   ✗ JSON serialization/deserialization FAILED");
        }

        // 测试二进制序列化
        Console.WriteLine("3. Testing Binary Serialization...");
        byte[] binaryData = login.EncodeBinary();
        Console.WriteLine($"   Binary encoded: {binaryData.Length} bytes");
        File.WriteAllBytes("test_output_csharp.bin", binaryData);

        // 测试二进制反序列化
        Console.WriteLine("4. Testing Binary Deserialization...");
        var loginFromBinary = new ProtocolLogin();
        loginFromBinary.Decoder(binaryData);
        if (VerifyLogin(login, loginFromBinary))
        {
            Console.WriteLine("   ✓ Binary serialization/deserialization PASSED");
        }
        else
        {
            Console.WriteLine("   ✗ Binary serialization/deserialization FAILED");
        }

        // 测试自动格式识别
        Console.WriteLine("5. Testing Auto Format Detection...");
        var loginAuto1 = new ProtocolLogin();
        loginAuto1.Decoder(jsonData);
        if (VerifyLogin(login, loginAuto1))
        {
            Console.WriteLine("   ✓ Auto detection (JSON) PASSED");
        }
        else
        {
            Console.WriteLine("   ✗ Auto detection (JSON) FAILED");
        }

        var loginAuto2 = new ProtocolLogin();
        loginAuto2.Decoder(binaryData);
        if (VerifyLogin(login, loginAuto2))
        {
            Console.WriteLine("   ✓ Auto detection (Binary) PASSED");
        }
        else
        {
            Console.WriteLine("   ✗ Auto detection (Binary) FAILED");
        }

        Console.WriteLine("\n=== Test Complete ===");
    }

    static bool VerifyLogin(ProtocolLogin original, ProtocolLogin decoded)
    {
        if (original.a != decoded.a ||
            original.b != decoded.b ||
            original.c != decoded.c ||
            original.d != decoded.d ||
            original.e != decoded.e ||
            original.f != decoded.f ||
            original.g != decoded.g ||
            original.h != decoded.h ||
            Math.Abs(original.i - decoded.i) > 0.001f ||
            Math.Abs(original.j - decoded.j) > 0.001 ||
            original.k != decoded.k ||
            original.l != decoded.l ||
            original.m != decoded.m)
        {
            return false;
        }

        if (original.n == null || decoded.n == null || original.n.Length != decoded.n.Length)
        {
            return false;
        }
        for (int i = 0; i < original.n.Length; i++)
        {
            if (original.n[i] != decoded.n[i])
            {
                return false;
            }
        }

        if (original.o == null || decoded.o == null || original.o.Length != decoded.o.Length)
        {
            return false;
        }
        for (int i = 0; i < original.o.Length; i++)
        {
            if (original.o[i] != decoded.o[i])
            {
                return false;
            }
        }

        if (original.q == null || decoded.q == null ||
            original.q.nickname != decoded.q.nickname ||
            original.q.password != decoded.q.password)
        {
            return false;
        }

        if (original.r == null || decoded.r == null || original.r.Length != decoded.r.Length)
        {
            return false;
        }
        for (int i = 0; i < original.r.Length; i++)
        {
            if (original.r[i].nickname != decoded.r[i].nickname ||
                original.r[i].password != decoded.r[i].password)
            {
                return false;
            }
        }

        return true;
    }
}

