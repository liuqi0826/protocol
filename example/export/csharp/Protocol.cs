using System;
using System.IO;
using System.Text;
using System.Text.Json;

namespace protocol
{
public class ProtocolLogin
{
public sbyte a;
public byte b;
public Int16 c;
public UInt16 d;
public Int32 e;
public UInt32 f;
public Int64 g;
public UInt64 h;
public float i;
public double j;
public bool k;
public byte l;
public string m;
public sbyte[] n;
public string[] o;
public Account q;
public Account[] r;
    public ProtocolLogin()
    {
a = 0;
b = 0;
c = 0;
d = 0;
e = 0;
f = 0;
g = 0;
h = 0;
i = 0;
j = 0;
k = false;
l = 0;
m = "";
n = new sbyte[]{};
o = new string[]{};
q = new Account();
r = new Account[]{};
    }
    public void Decoder(byte[] data)
    {
        if (data == null || data.Length == 0) return;
        if (Encoding.Default.GetString(new byte[] { data[0] }) == "{")
        {
            var protocol = JsonSerializer.Deserialize<ProtocolLogin>(Encoding.Default.GetString(data));
a = protocol.a;
b = protocol.b;
c = protocol.c;
d = protocol.d;
e = protocol.e;
f = protocol.f;
g = protocol.g;
h = protocol.h;
i = protocol.i;
j = protocol.j;
k = protocol.k;
l = protocol.l;
m = protocol.m;
n = protocol.n;
o = protocol.o;
q = protocol.q;
r = protocol.r;
        }
    	else
    	{
            MemoryStream steam = new MemoryStream(data);
            BinaryReader reader = new BinaryReader(steam);
a = reader.ReadSByte();
b = reader.ReadByte();
c = reader.ReadInt16();
d = reader.ReadUInt16();
e = reader.ReadInt32();
f = reader.ReadUInt32();
g = reader.ReadInt64();
h = reader.ReadUInt64();
i = reader.ReadSingle();
j = reader.ReadDouble();
k = reader.ReadBoolean();
l = reader.ReadByte();
            var mLenght = reader.ReadInt32();
            var mReader = new BinaryReader(new MemoryStream(reader.ReadBytes(mLenght)));
            this.m = mReader.ReadString();
            var nCount = reader.ReadInt32();
            this.n = new sbyte[nCount];
            for (var i = 0; i < nCount; i++)
            {
this.n[i] = reader.ReadSByte();
            }
            var oCount = reader.ReadInt32();
            this.o = new string[oCount];
            for (var i = 0; i < oCount; i++)
            {
                var l = reader.ReadInt32();
                var r = new BinaryReader(new MemoryStream(reader.ReadBytes(l)));
                this.o[i] = r.ReadString();
            }
            var qLenght = reader.ReadInt32();
            var qByte = reader.ReadBytes(qLenght);
            this.q = new Account();
            this.q.Decoder(qByte);
            var rCount = reader.ReadInt32();
            this.r = new Account[rCount];
            for (var i = 0; i < rCount; i++)
            {
                var l = reader.ReadInt32();
                var d = reader.ReadBytes(l);
                this.r[i] = new Account();
                this.r[i].Decoder(d);
            }
    	}
    }
    public byte[] EncodeBinary()
    {
        MemoryStream steam = new MemoryStream();
        BinaryWriter writer = new BinaryWriter(steam);
writer.Write(this.a);
writer.Write(this.b);
writer.Write(this.c);
writer.Write(this.d);
writer.Write(this.e);
writer.Write(this.f);
writer.Write(this.g);
writer.Write(this.h);
writer.Write(this.i);
writer.Write(this.j);
writer.Write(this.k);
writer.Write(this.l);
        var mWriter = new BinaryWriter(new MemoryStream());
        mWriter.Write(this.m);
        writer.Write((Int32)mWriter.BaseStream.Length);
        writer.Write(this.m);
writer.Write((Int32)this.n.Length);
for (var i = 0; i < this.n.Length; i++)
{
            writer.Write(this.n[i]);
}
writer.Write((Int32)this.o.Length);
for (var i = 0; i < this.o.Length; i++)
{
            var w = new BinaryWriter(new MemoryStream());
            w.Write(this.o[i]);
            writer.Write((Int32)w.BaseStream.Length);
            writer.Write(this.o[i]);
}
        var qByteArray = this.q.EncodeBinary();
        writer.Write((Int32)qByteArray.Length);
        writer.Write(qByteArray);
writer.Write((Int32)this.r.Length);
for (var i = 0; i < this.r.Length; i++)
{
            var array = this.r[i].EncodeBinary();
            writer.Write((Int32)array.Length);
            writer.Write(array);
}
    return steam.ToArray();
    }
    public byte[] EncodeJson()
    {
        return Encoding.Default.GetBytes(JsonSerializer.Serialize(this));
    }
}
public class Account
{
public string nickname;
public string password;
    public Account()
    {
nickname = "";
password = "";
    }
    public void Decoder(byte[] data)
    {
        if (data == null || data.Length == 0) return;
        if (Encoding.Default.GetString(new byte[] { data[0] }) == "{")
        {
            var protocol = JsonSerializer.Deserialize<Account>(Encoding.Default.GetString(data));
nickname = protocol.nickname;
password = protocol.password;
        }
    	else
    	{
            MemoryStream steam = new MemoryStream(data);
            BinaryReader reader = new BinaryReader(steam);
            var nicknameLenght = reader.ReadInt32();
            var nicknameReader = new BinaryReader(new MemoryStream(reader.ReadBytes(nicknameLenght)));
            this.nickname = nicknameReader.ReadString();
            var passwordLenght = reader.ReadInt32();
            var passwordReader = new BinaryReader(new MemoryStream(reader.ReadBytes(passwordLenght)));
            this.password = passwordReader.ReadString();
    	}
    }
    public byte[] EncodeBinary()
    {
        MemoryStream steam = new MemoryStream();
        BinaryWriter writer = new BinaryWriter(steam);
        var nicknameWriter = new BinaryWriter(new MemoryStream());
        nicknameWriter.Write(this.nickname);
        writer.Write((Int32)nicknameWriter.BaseStream.Length);
        writer.Write(this.nickname);
        var passwordWriter = new BinaryWriter(new MemoryStream());
        passwordWriter.Write(this.password);
        writer.Write((Int32)passwordWriter.BaseStream.Length);
        writer.Write(this.password);
    return steam.ToArray();
    }
    public byte[] EncodeJson()
    {
        return Encoding.Default.GetBytes(JsonSerializer.Serialize(this));
    }
}
public class ProtocolServerLogin
{
public string id;
public string token;
    public ProtocolServerLogin()
    {
id = "";
token = "";
    }
    public void Decoder(byte[] data)
    {
        if (data == null || data.Length == 0) return;
        if (Encoding.Default.GetString(new byte[] { data[0] }) == "{")
        {
            var protocol = JsonSerializer.Deserialize<ProtocolServerLogin>(Encoding.Default.GetString(data));
id = protocol.id;
token = protocol.token;
        }
    	else
    	{
            MemoryStream steam = new MemoryStream(data);
            BinaryReader reader = new BinaryReader(steam);
            var idLenght = reader.ReadInt32();
            var idReader = new BinaryReader(new MemoryStream(reader.ReadBytes(idLenght)));
            this.id = idReader.ReadString();
            var tokenLenght = reader.ReadInt32();
            var tokenReader = new BinaryReader(new MemoryStream(reader.ReadBytes(tokenLenght)));
            this.token = tokenReader.ReadString();
    	}
    }
    public byte[] EncodeBinary()
    {
        MemoryStream steam = new MemoryStream();
        BinaryWriter writer = new BinaryWriter(steam);
        var idWriter = new BinaryWriter(new MemoryStream());
        idWriter.Write(this.id);
        writer.Write((Int32)idWriter.BaseStream.Length);
        writer.Write(this.id);
        var tokenWriter = new BinaryWriter(new MemoryStream());
        tokenWriter.Write(this.token);
        writer.Write((Int32)tokenWriter.BaseStream.Length);
        writer.Write(this.token);
    return steam.ToArray();
    }
    public byte[] EncodeJson()
    {
        return Encoding.Default.GetBytes(JsonSerializer.Serialize(this));
    }
}
public class ProtocolServerState
{
public UInt16 state;
public string value;
    public ProtocolServerState()
    {
state = 0;
value = "";
    }
    public void Decoder(byte[] data)
    {
        if (data == null || data.Length == 0) return;
        if (Encoding.Default.GetString(new byte[] { data[0] }) == "{")
        {
            var protocol = JsonSerializer.Deserialize<ProtocolServerState>(Encoding.Default.GetString(data));
state = protocol.state;
value = protocol.value;
        }
    	else
    	{
            MemoryStream steam = new MemoryStream(data);
            BinaryReader reader = new BinaryReader(steam);
state = reader.ReadUInt16();
            var valueLenght = reader.ReadInt32();
            var valueReader = new BinaryReader(new MemoryStream(reader.ReadBytes(valueLenght)));
            this.value = valueReader.ReadString();
    	}
    }
    public byte[] EncodeBinary()
    {
        MemoryStream steam = new MemoryStream();
        BinaryWriter writer = new BinaryWriter(steam);
writer.Write(this.state);
        var valueWriter = new BinaryWriter(new MemoryStream());
        valueWriter.Write(this.value);
        writer.Write((Int32)valueWriter.BaseStream.Length);
        writer.Write(this.value);
    return steam.ToArray();
    }
    public byte[] EncodeJson()
    {
        return Encoding.Default.GetBytes(JsonSerializer.Serialize(this));
    }
}
public class ProtocolServerCommand
{
public UInt16 command;
public string value;
    public ProtocolServerCommand()
    {
command = 0;
value = "";
    }
    public void Decoder(byte[] data)
    {
        if (data == null || data.Length == 0) return;
        if (Encoding.Default.GetString(new byte[] { data[0] }) == "{")
        {
            var protocol = JsonSerializer.Deserialize<ProtocolServerCommand>(Encoding.Default.GetString(data));
command = protocol.command;
value = protocol.value;
        }
    	else
    	{
            MemoryStream steam = new MemoryStream(data);
            BinaryReader reader = new BinaryReader(steam);
command = reader.ReadUInt16();
            var valueLenght = reader.ReadInt32();
            var valueReader = new BinaryReader(new MemoryStream(reader.ReadBytes(valueLenght)));
            this.value = valueReader.ReadString();
    	}
    }
    public byte[] EncodeBinary()
    {
        MemoryStream steam = new MemoryStream();
        BinaryWriter writer = new BinaryWriter(steam);
writer.Write(this.command);
        var valueWriter = new BinaryWriter(new MemoryStream());
        valueWriter.Write(this.value);
        writer.Write((Int32)valueWriter.BaseStream.Length);
        writer.Write(this.value);
    return steam.ToArray();
    }
    public byte[] EncodeJson()
    {
        return Encoding.Default.GetBytes(JsonSerializer.Serialize(this));
    }
}
}
