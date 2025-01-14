// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: game.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_OK                     ErrorCode = 0
	ErrorCode_ROOM_NOT_FOUND         ErrorCode = 1
	ErrorCode_ROOM_FULL              ErrorCode = 2
	ErrorCode_PLAYER_NOT_FOUND       ErrorCode = 3
	ErrorCode_PLAYER_ALREADY_IN_ROOM ErrorCode = 4
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "OK",
		1: "ROOM_NOT_FOUND",
		2: "ROOM_FULL",
		3: "PLAYER_NOT_FOUND",
		4: "PLAYER_ALREADY_IN_ROOM",
	}
	ErrorCode_value = map[string]int32{
		"OK":                     0,
		"ROOM_NOT_FOUND":         1,
		"ROOM_FULL":              2,
		"PLAYER_NOT_FOUND":       3,
		"PLAYER_ALREADY_IN_ROOM": 4,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type MessageId int32

const (
	MessageId_LOGIN_REQUEST           MessageId = 0
	MessageId_LOGIN_RESPONSE          MessageId = 1
	MessageId_GET_ROOM_LIST_REQUEST   MessageId = 2
	MessageId_GET_ROOM_LIST_RESPONSE  MessageId = 3
	MessageId_CREATE_ROOM_REQUEST     MessageId = 4
	MessageId_CREATE_ROOM_RESPONSE    MessageId = 5
	MessageId_JOIN_ROOM_REQUEST       MessageId = 6
	MessageId_JOIN_ROOM_RESPONSE      MessageId = 7
	MessageId_MOVE_REQUEST            MessageId = 8
	MessageId_MOVE_RESPONSE           MessageId = 9
	MessageId_LEAVE_ROOM_REQUEST      MessageId = 10
	MessageId_LEAVE_ROOM_RESPONSE     MessageId = 11
	MessageId_ROOM_STATE_NOTIFICATION MessageId = 12
)

// Enum value maps for MessageId.
var (
	MessageId_name = map[int32]string{
		0:  "LOGIN_REQUEST",
		1:  "LOGIN_RESPONSE",
		2:  "GET_ROOM_LIST_REQUEST",
		3:  "GET_ROOM_LIST_RESPONSE",
		4:  "CREATE_ROOM_REQUEST",
		5:  "CREATE_ROOM_RESPONSE",
		6:  "JOIN_ROOM_REQUEST",
		7:  "JOIN_ROOM_RESPONSE",
		8:  "MOVE_REQUEST",
		9:  "MOVE_RESPONSE",
		10: "LEAVE_ROOM_REQUEST",
		11: "LEAVE_ROOM_RESPONSE",
		12: "ROOM_STATE_NOTIFICATION",
	}
	MessageId_value = map[string]int32{
		"LOGIN_REQUEST":           0,
		"LOGIN_RESPONSE":          1,
		"GET_ROOM_LIST_REQUEST":   2,
		"GET_ROOM_LIST_RESPONSE":  3,
		"CREATE_ROOM_REQUEST":     4,
		"CREATE_ROOM_RESPONSE":    5,
		"JOIN_ROOM_REQUEST":       6,
		"JOIN_ROOM_RESPONSE":      7,
		"MOVE_REQUEST":            8,
		"MOVE_RESPONSE":           9,
		"LEAVE_ROOM_REQUEST":      10,
		"LEAVE_ROOM_RESPONSE":     11,
		"ROOM_STATE_NOTIFICATION": 12,
	}
)

func (x MessageId) Enum() *MessageId {
	p := new(MessageId)
	*p = x
	return p
}

func (x MessageId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageId) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[1].Descriptor()
}

func (MessageId) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[1]
}

func (x MessageId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageId.Descriptor instead.
func (MessageId) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	mi := &file_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position *Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players []*Player `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *Room) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_game_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetRoomListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoomListRequest) Reset() {
	*x = GetRoomListRequest{}
	mi := &file_game_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomListRequest) ProtoMessage() {}

func (x *GetRoomListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomListRequest.ProtoReflect.Descriptor instead.
func (*GetRoomListRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

type GetRoomListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret   ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=game.ErrorCode" json:"ret,omitempty"`
	Rooms []*Room   `protobuf:"bytes,2,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetRoomListResponse) Reset() {
	*x = GetRoomListResponse{}
	mi := &file_game_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomListResponse) ProtoMessage() {}

func (x *GetRoomListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomListResponse.ProtoReflect.Descriptor instead.
func (*GetRoomListResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *GetRoomListResponse) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_OK
}

func (x *GetRoomListResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_game_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=game.ErrorCode" json:"ret,omitempty"`
	Room *Room     `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_game_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{8}
}

func (x *CreateRoomResponse) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_OK
}

func (x *CreateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	RoomId uint64  `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	mi := &file_game_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{9}
}

func (x *JoinRoomRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *JoinRoomRequest) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=game.ErrorCode" json:"ret,omitempty"`
	Room *Room     `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	mi := &file_game_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{10}
}

func (x *JoinRoomResponse) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_OK
}

func (x *JoinRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string    `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Position *Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	mi := &file_game_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{11}
}

func (x *MoveRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *MoveRequest) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=game.ErrorCode" json:"ret,omitempty"`
	Room *Room     `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	mi := &file_game_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{12}
}

func (x *MoveResponse) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_OK
}

func (x *MoveResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type LeaveRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *LeaveRoomRequest) Reset() {
	*x = LeaveRoomRequest{}
	mi := &file_game_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomRequest) ProtoMessage() {}

func (x *LeaveRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomRequest.ProtoReflect.Descriptor instead.
func (*LeaveRoomRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{13}
}

func (x *LeaveRoomRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type LeaveRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=game.ErrorCode" json:"ret,omitempty"`
	Room *Room     `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *LeaveRoomResponse) Reset() {
	*x = LeaveRoomResponse{}
	mi := &file_game_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomResponse) ProtoMessage() {}

func (x *LeaveRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomResponse.ProtoReflect.Descriptor instead.
func (*LeaveRoomResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{14}
}

func (x *LeaveRoomResponse) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_OK
}

func (x *LeaveRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type RoomStateNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *RoomStateNotification) Reset() {
	*x = RoomStateNotification{}
	mi := &file_game_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomStateNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStateNotification) ProtoMessage() {}

func (x *RoomStateNotification) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStateNotification.ProtoReflect.Descriptor instead.
func (*RoomStateNotification) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{15}
}

func (x *RoomStateNotification) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    string    `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`          //客户端唯一标识
	MsgSerialNo int32     `protobuf:"varint,2,opt,name=msgSerialNo,proto3" json:"msgSerialNo,omitempty"`   //消息序列号, 每条消息加1
	Id          MessageId `protobuf:"varint,3,opt,name=id,proto3,enum=game.MessageId" json:"id,omitempty"` //消息ID
	Data        []byte    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                  //消息体
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_game_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{16}
}

func (x *Message) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Message) GetMsgSerialNo() int32 {
	if x != nil {
		return x.MsgSerialNo
	}
	return 0
}

func (x *Message) GetId() MessageId {
	if x != nil {
		return x.Id
	}
	return MessageId_LOGIN_REQUEST
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0x58, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x4f, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12,
	0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22,
	0x55, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x2e, 0x0a, 0x10, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x11, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x22, 0x37, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x7c, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x68, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d,
	0x10, 0x04, 0x2a, 0xbe, 0x02, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x54, 0x5f, 0x52,
	0x4f, 0x4f, 0x4d, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4c,
	0x49, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10,
	0x05, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x49, 0x4e,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x07,
	0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x52,
	0x4f, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x0a, 0x12, 0x17, 0x0a,
	0x13, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0x0b, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x0c, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_game_proto_goTypes = []any{
	(ErrorCode)(0),                // 0: game.ErrorCode
	(MessageId)(0),                // 1: game.MessageId
	(*Position)(nil),              // 2: game.Position
	(*Player)(nil),                // 3: game.Player
	(*Room)(nil),                  // 4: game.Room
	(*LoginRequest)(nil),          // 5: game.LoginRequest
	(*LoginResponse)(nil),         // 6: game.LoginResponse
	(*GetRoomListRequest)(nil),    // 7: game.GetRoomListRequest
	(*GetRoomListResponse)(nil),   // 8: game.GetRoomListResponse
	(*CreateRoomRequest)(nil),     // 9: game.CreateRoomRequest
	(*CreateRoomResponse)(nil),    // 10: game.CreateRoomResponse
	(*JoinRoomRequest)(nil),       // 11: game.JoinRoomRequest
	(*JoinRoomResponse)(nil),      // 12: game.JoinRoomResponse
	(*MoveRequest)(nil),           // 13: game.MoveRequest
	(*MoveResponse)(nil),          // 14: game.MoveResponse
	(*LeaveRoomRequest)(nil),      // 15: game.LeaveRoomRequest
	(*LeaveRoomResponse)(nil),     // 16: game.LeaveRoomResponse
	(*RoomStateNotification)(nil), // 17: game.RoomStateNotification
	(*Message)(nil),               // 18: game.Message
}
var file_game_proto_depIdxs = []int32{
	2,  // 0: game.Player.position:type_name -> game.Position
	3,  // 1: game.Room.players:type_name -> game.Player
	0,  // 2: game.GetRoomListResponse.ret:type_name -> game.ErrorCode
	4,  // 3: game.GetRoomListResponse.rooms:type_name -> game.Room
	0,  // 4: game.CreateRoomResponse.ret:type_name -> game.ErrorCode
	4,  // 5: game.CreateRoomResponse.room:type_name -> game.Room
	3,  // 6: game.JoinRoomRequest.player:type_name -> game.Player
	0,  // 7: game.JoinRoomResponse.ret:type_name -> game.ErrorCode
	4,  // 8: game.JoinRoomResponse.room:type_name -> game.Room
	2,  // 9: game.MoveRequest.position:type_name -> game.Position
	0,  // 10: game.MoveResponse.ret:type_name -> game.ErrorCode
	4,  // 11: game.MoveResponse.room:type_name -> game.Room
	0,  // 12: game.LeaveRoomResponse.ret:type_name -> game.ErrorCode
	4,  // 13: game.LeaveRoomResponse.room:type_name -> game.Room
	4,  // 14: game.RoomStateNotification.room:type_name -> game.Room
	1,  // 15: game.Message.id:type_name -> game.MessageId
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
