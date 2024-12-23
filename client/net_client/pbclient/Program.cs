﻿using System;
using System.Net.Sockets;
using System.Text;
using Google.Protobuf;
using Game; // 这里的命名空间应与生成的C#代码中的命名空间一致

class Program
{
    static void Main(string[] args)
    {
        string server = "127.0.0.1";
        int port = 12345;

        try
        {
            using (TcpClient client = new TcpClient(server, port))
            using (NetworkStream stream = client.GetStream())
            {
                Console.WriteLine("Connected to server");

                // 创建一个CreateRoomRequest消息
                var createRoomRequest = new CreateRoomRequest
                {
                    Name = "Test Room"
                };

                // 序列化CreateRoomRequest消息
                byte[] requestData = createRoomRequest.ToByteArray();

                // 创建一个Message消息
                var message = new Message
                {
                    Uuid = Guid.NewGuid().ToString(),
                    Id = MessageId.CreateRoomRequest,
                    Data = ByteString.CopyFrom(requestData)
                };

                // 序列化Message消息
                byte[] data = message.ToByteArray();

                // 发送消息
                stream.Write(data, 0, data.Length);
                Console.WriteLine("Message sent");

                // 确认数据已发送
                Console.WriteLine($"Sent {data.Length} bytes to server");

                // 接收响应
                byte[] buffer = new byte[1024];
                int bytesRead = stream.Read(buffer, 0, buffer.Length);
                if (bytesRead > 0)
                {
                    var response = Message.Parser.ParseFrom(buffer, 0, bytesRead);
                    Console.WriteLine($"Received response: {response.Id}, {response.Uuid}");
                }
                else
                {
                    Console.WriteLine("No data received from server.");
                }
            }
        }
        catch (Exception e)
        {
            Console.WriteLine($"Exception: {e.Message}");
        }
    }
}