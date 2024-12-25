using System;
using System.Threading.Tasks;
using Game;

class Program
{
    static Player? player;
    static async Task Main(string[] args)
    {
        string server = "127.0.0.1";
        int port = 12345;

        try
        {
            using (var connection = new Connection(server, port))
            {
                Console.WriteLine("Connected to server");

                connection.RegisterMessageHandler(MessageId.RoomStateNotification, OnRoomStateNotification);
                var listenTask = connection.ListenForNotificationsAsync();

                // 创建玩家
                player = new Player("Player1", connection);

                // 获取房间列表
                var getRoomListResponse = await player.GetRoomList();

                if (getRoomListResponse.Rooms.Count > 0)
                {
                    Console.WriteLine("Rooms:");
                    foreach (var roomInfo in getRoomListResponse.Rooms)
                    {
                        Console.WriteLine($"- {roomInfo.Name}");
                    }

                    var roomId = getRoomListResponse.Rooms[0].Id;
                    var roomName = getRoomListResponse.Rooms[0].Name;

                    // 创建房间对象
                    var room = new Room { Id = roomId, Name = roomName };

                    // 玩家加入房间
                    var joinRoomResponse = await player.JoinRoom(room);

                    if (joinRoomResponse.Ret == (int)ErrorCode.Ok)
                    {
                        Console.WriteLine("Joined room successfully.");
                    }
                    else
                    {
                        Console.WriteLine("Failed to join room.");
                    }
                }
                else
                {
                    Console.WriteLine("No rooms found. Creating a new room...");

                    // 创建房间
                    var createRoomResponse = await player.CreateRoom("Test Room");

                    if (createRoomResponse.Ret == (int)ErrorCode.Ok)
                    {
                        Console.WriteLine("Room created successfully.");
                    }
                    else
                    {
                        Console.WriteLine("Failed to create room.");
                    }
                }

                // 等待监听任务完成
                await listenTask;
            }
        }
        catch (Exception e)
        {
            Console.WriteLine($"Exception: {e.Message}");
        }
    }

    static void OnRoomStateNotification(Message message)
    {
        var notification = RoomStateNotification.Parser.ParseFrom(message.Data);
        if (player != null)
        {
            if (player.Room == null)
            {
                player.Room = new Room();
            }
            player.Room = notification.Room;
        }

        Console.WriteLine($"Received room state notification:");

        Console.WriteLine($"Received room state notification:");
    }
}
