package main

import pb "server/src/proto"

func RoomsToProto(rooms []*Room) []*pb.Room {
	var protoRooms []*pb.Room
	for _, room := range rooms {
		protoRooms = append(protoRooms, &pb.Room{
			Id:   room.ID,
			Name: room.Name,
		})
	}
	return protoRooms
}
