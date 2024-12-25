using System;
using System.IO;

public class GameMessage
{
    public int Length { get; private set; }
    public byte[] Body { get; private set; }

    public GameMessage(byte[] body)
    {
        Body = body;
        Length = body.Length;
    }

    public byte[] Serialize()
    {
        using (var memoryStream = new MemoryStream())
        {
            // 写入包头（消息长度）
            byte[] lengthBytes = BitConverter.GetBytes(Length);
            memoryStream.Write(lengthBytes, 0, lengthBytes.Length);

            // 写入包体
            memoryStream.Write(Body, 0, Body.Length);

            return memoryStream.ToArray();
        }
    }

    public static GameMessage Deserialize(byte[] data)
    {
        using (var memoryStream = new MemoryStream(data))
        {
            // 读取包头（消息长度）
            byte[] lengthBytes = new byte[4];
            memoryStream.Read(lengthBytes, 0, lengthBytes.Length);
            int length = BitConverter.ToInt32(lengthBytes, 0);

            // 读取包体
            byte[] body = new byte[length];
            memoryStream.Read(body, 0, body.Length);

            return new GameMessage(body);
        }
    }
}