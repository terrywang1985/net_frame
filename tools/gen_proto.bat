protoc.exe --proto_path=../proto --go_out=../server/src/proto --go_opt=paths=source_relative ../proto/game.proto

protoc.exe -I=../proto --csharp_out=../client/net_client/pbclient  ../proto/game.proto