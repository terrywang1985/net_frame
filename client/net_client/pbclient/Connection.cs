using System;
using System.Collections.Concurrent;
using System.Net.Sockets;
using System.Threading.Tasks;
using Google.Protobuf;
using Game;
using System.Collections.Generic;

public class Connection : IDisposable
{
    private readonly string _server;
    private readonly int _port;

    private int _msgSerialNo = 0;

    private TcpClient _client;
    private NetworkStream _stream;
    private readonly Dictionary<MessageId, Action<Message>> _messageHandlers;

    private readonly ConcurrentQueue<byte[]> _sendQueue = new ConcurrentQueue<byte[]>();
    private readonly ConcurrentQueue<byte[]> _receiveQueue = new ConcurrentQueue<byte[]>();
    private readonly ConcurrentDictionary<int, TaskCompletionSource<Message>> _responseTasks = new ConcurrentDictionary<int, TaskCompletionSource<Message>>();

    public Connection(string server, int port)
    {
        _server = server;
        _port = port;
        _messageHandlers = new Dictionary<MessageId, Action<Message>>();

        _client = new TcpClient(); // ��ʼ�� _client
        Connect();
    }

    private void Connect()
    {
        _client.Connect(_server, _port);
        _stream = _client.GetStream();
        Task.Run(() => ProcessSendQueue());
        Task.Run(() => ProcessReceiveQueue());
    }

    public void RegisterMessageHandler(MessageId messageId, Action<Message> handler)
    {
        _messageHandlers[messageId] = handler;
    }

    public async Task<TResponse> SendRequestAsync<TRequest, TResponse>(MessageId messageId, TRequest request)
        where TRequest : IMessage<TRequest>, new()
        where TResponse : IMessage<TResponse>, new()
    {
        byte[] requestData = request.ToByteArray();

        var message = new Message
        {
            MsgSerialNo = _msgSerialNo++,
            ClientId = Guid.NewGuid().ToString(),
            Id = messageId,
            Data = ByteString.CopyFrom(requestData)
        };

        var gameMessage = new GameMessage(message.ToByteArray());
        byte[] data = gameMessage.Serialize();

        // ����һ�� TaskCompletionSource ���ȴ���Ӧ
        var tcs = new TaskCompletionSource<Message>();
        _responseTasks[message.MsgSerialNo] = tcs;

        // ����Ϣ���뷢�Ͷ���
        _sendQueue.Enqueue(data);

        // �ȴ���Ӧ��Ϣ
        var responseMessage = await tcs.Task;

        // ������Ӧ��Ϣ
        var response = new TResponse();
        response.MergeFrom(responseMessage.Data);
        return response;
    }

    private async Task<byte[]> ReceiveMessageAsync()
    {
        while (true)
        {
            if (_receiveQueue.TryDequeue(out var data))
            {
                return data;
            }
            await Task.Delay(10); // �ȴ����ݵ���
        }
    }

    private async Task ProcessSendQueue()
    {
        while (true)
        {
            if (_sendQueue.TryDequeue(out var data))
            {
                await _stream.WriteAsync(data, 0, data.Length);
                Console.WriteLine($"Sent {data.Length} bytes to server");
            }
            await Task.Delay(10); // ���Ʒ���Ƶ��
        }
    }

    private async Task ProcessReceiveQueue()
    {
        byte[] buffer = new byte[1024];
        int bufferOffset = 0;

        while (true)
        {
            int bytesRead = await _stream.ReadAsync(buffer, bufferOffset, buffer.Length - bufferOffset);
            if (bytesRead > 0)
            {
                bufferOffset += bytesRead;

                while (bufferOffset >= 4)
                {
                    int messageLength = BitConverter.ToInt32(buffer, 0);
                    if (bufferOffset >= messageLength + 4)
                    {
                        byte[] messageBuffer = new byte[messageLength];
                        Array.Copy(buffer, 4, messageBuffer, 0, messageLength);

                        // ����Ϣ������ն���
                        _receiveQueue.Enqueue(messageBuffer);

                        bufferOffset -= (messageLength + 4);
                        Array.Copy(buffer, messageLength + 4, buffer, 0, bufferOffset);
                    }
                    else
                    {
                        break;
                    }
                }
            }
            else
            {
                Console.WriteLine("Connection closed by server.");
                await ReconnectAsync();
                break;
            }
        }
    }

    public async Task ListenForNotificationsAsync()
    {
        while (true)
        {
            byte[] messageBuffer = await ReceiveMessageAsync();
            if (messageBuffer != null)
            {
                var message = Message.Parser.ParseFrom(messageBuffer);

                // ����Ƿ��еȴ�����Ϣ������
                if (_responseTasks.TryRemove(message.MsgSerialNo, out var tcs))
                {
                    tcs.SetResult(message);
                }
                else if (_messageHandlers.TryGetValue(message.Id, out var handler))
                {
                    handler(message);
                }
                else
                {
                    Console.WriteLine($"No handler registered for message ID: {message.Id}");
                }
            }
        }
    }

    private async Task ReconnectAsync()
    {
        while (true)
        {
            try
            {
                Console.WriteLine("Attempting to reconnect...");
                Connect();
                Console.WriteLine("Reconnected successfully.");
                break;
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Reconnect failed: {ex.Message}");
                await Task.Delay(5000); // �ȴ�5�������
            }
        }
    }

    public void Dispose()
    {
        _stream?.Dispose();
        _client?.Dispose();
    }
}
