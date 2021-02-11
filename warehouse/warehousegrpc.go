package warehouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	proto "github.com/rudderlabs/rudder-server/proto/warehouse"
	"time"
)

type warehousegrpc struct {
	proto.UnimplementedWarehouseServer
}

func (w *warehousegrpc) GetWHUploads(context context.Context, request *proto.GetWHUploadsRequest) (*proto.GetWHUploadsResponse, error) {
	from, err := ptypes.Timestamp(request.From)
	if err != nil {
		from = time.Time{}
	}
	to, err := ptypes.Timestamp(request.To)
	if err != nil {
		to = time.Time{}
	}
	uploadsReq := UploadsReqT{
		SourceID:           request.SourceId,
		DestinationID:      request.DestinationId,
		DestinationType:    request.DestinationType,
		Status:             request.Status,
		IncludeTablesInRes: request.IncludeTablesInRes,
		From:               from,
		To:                 to,
		API:                UploadAPI,
	}
	protoRes := &proto.GetWHUploadsResponse{}
	res, err := uploadsReq.GetWhUploads()
	if err != nil {
		return protoRes, err
	}
	bytes, err := json.Marshal(res)
	err = jsonpb.UnmarshalString(string(bytes), protoRes)
	fmt.Println(err)
	fmt.Println(protoRes)
	if err != nil {
		return protoRes, err
	}
	return protoRes, nil

}
