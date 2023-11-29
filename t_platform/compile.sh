# protoc --proto_path=pb --go_opt=appinfo.proto=github.com/aheadIV/textcharge/t_platform/pb 


protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative proto/app/info.proto
