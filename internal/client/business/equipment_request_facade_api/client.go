package equipment_request_facade_api

import (
	"context"
	"errors"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrNoListEquipmentRequest is an "empty equipment request list returned by this limit and offset" error
var ErrNoListEquipmentRequest = errors.New("empty equipment request list returned by this limit and offset")

// ErrNoExistsEquipmentRequest is a "equipment request not founded" error
var ErrNoExistsEquipmentRequest = errors.New("equipment request with this id does not exist")

//BssEquipmentRequestFacadeAPIServiceClient is a client for equipment request facade
type BssEquipmentRequestFacadeAPIServiceClient interface {
	ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]*business.EquipmentRequest, uint64, error)
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*business.EquipmentRequest, error)
}

type bssEquipmentRequestFacadeAPIServiceClient struct {
	grpcClient facadepb.BssEquipmentRequestFacadeApiServiceClient
}

//NewBssEquipmentRequestFacadeAPIServiceClient returns a new equipment request facade client
func NewBssEquipmentRequestFacadeAPIServiceClient(grpcClient facadepb.BssEquipmentRequestFacadeApiServiceClient) BssEquipmentRequestFacadeAPIServiceClient {
	return &bssEquipmentRequestFacadeAPIServiceClient{
		grpcClient: grpcClient,
	}
}

func (c *bssEquipmentRequestFacadeAPIServiceClient) ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]*business.EquipmentRequest, uint64, error) {
	newRequest := facadepb.ListEquipmentRequestFacadeV1Request{
		Limit:  limit,
		Offset: offset,
	}

	response, err := c.grpcClient.ListEquipmentRequestFacadeV1(ctx, &newRequest)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, 0, ErrNoListEquipmentRequest
		}

		return nil, 0, err
	}

	equipmentRequestPb, err := business.ConvertRepeatedFacadePbToEquipmentRequests(response.Items)

	if err != nil {
		return nil, 0, err
	}

	return equipmentRequestPb, response.Total, nil
}

func (c *bssEquipmentRequestFacadeAPIServiceClient) DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*business.EquipmentRequest, error) {
	newRequest := facadepb.DescribeEquipmentRequestFacadeV1Request{
		EquipmentRequestId: equipmentRequestID,
	}

	response, err := c.grpcClient.DescribeEquipmentRequestFacadeV1(ctx, &newRequest)

	if err != nil {
		return nil, err
	}

	request, err := business.ConvertFacadePbToEquipmentRequest(response.EquipmentRequest)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, ErrNoExistsEquipmentRequest
		}

		return nil, err
	}

	return request, nil
}
