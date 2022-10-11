package protojson

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Marshal(m proto.Message) ([]byte, error) {
	marshaler := protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true}
	return marshaler.Marshal(m)
}
