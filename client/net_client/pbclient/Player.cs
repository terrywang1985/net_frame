using System;
using System.Threading.Tasks;
using Game;

public class Player
{
    public string Id { get; private set; }
    public string Name { get; private set; }
    public Position Position { get; private set; }
    public Room Room { get; set; }
    private Connection Connection { get; set; }

    public Player(string name, Connection connection)
    {
        Id = Guid.NewGuid().ToString();
        Name = name;
        Position = new Position { X = 0, Y = 0, Z = 0 };
        Connection = connection;
    }

    public async Task<JoinRoomResponse> JoinRoom(Room room)
    {
        var joinRoomRequest = new JoinRoomRequest
        {
            Player = new Game.Player
            {
                Id = this.Id,
                Name = this.Name,
                Position = this.Position
            },
            RoomId = room.Id.ToString()
        };

        var response = await Connection.SendRequestAsync<JoinRoomRequest, JoinRoomResponse>(MessageId.JoinRoomRequest, joinRoomRequest);
        if (response.Ret == (int)ErrorCode.Ok)
        {
            // 加入房间成功
            Console.WriteLine("Join room successfully.");
        }
        return response;
    }


    public async Task<CreateRoomResponse> CreateRoom(string roomName)
    {
        var createRoomRequest = new CreateRoomRequest
        {
            Name = roomName
        };

        return await Connection.SendRequestAsync<CreateRoomRequest, CreateRoomResponse>(MessageId.CreateRoomRequest, createRoomRequest);
    }

    public async Task<GetRoomListResponse> GetRoomList()
    {
        return await Connection.SendRequestAsync<GetRoomListRequest, GetRoomListResponse>(MessageId.GetRoomListRequest, new GetRoomListRequest());
    }

}