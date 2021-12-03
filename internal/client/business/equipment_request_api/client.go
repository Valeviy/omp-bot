package equipment_request_api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

//BssEquipmentRequestAPIServiceClient is a client for equipment request api
type BssEquipmentRequestAPIServiceClient interface {
	CreateEquipmentRequest(ctx context.Context, equipmentRequest business.EquipmentRequest) (uint64, error)
	UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status business.EquipmentRequestStatus) (bool, error)
	UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error)
}

type bssEquipmentRequestAPIServiceClient struct {
	grpcClient pb.BssEquipmentRequestApiServiceClient
}

//NewBssEquipmentRequestAPIServiceClient returns a new equipment request api client
func NewBssEquipmentRequestAPIServiceClient(grpcClient pb.BssEquipmentRequestApiServiceClient) BssEquipmentRequestAPIServiceClient {
	return &bssEquipmentRequestAPIServiceClient{
		grpcClient: grpcClient,
	}
}

func (c *bssEquipmentRequestAPIServiceClient) CreateEquipmentRequest(ctx context.Context, equipmentRequest business.EquipmentRequest) (uint64, error) {
	newRequest, err := business.ConvertEquipmentRequestToCreatePbRequest(&equipmentRequest)

	if err != nil {
		return 0, err
	}

	response, err := c.grpcClient.CreateEquipmentRequestV1(ctx, newRequest)

	if err != nil {
		return 0, err
	}

	return response.EquipmentRequestId, nil
}

func (c *bssEquipmentRequestAPIServiceClient) UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status business.EquipmentRequestStatus) (bool, error) {
	pbStatus, err := business.ConvertEquipmentRequestStatusToPb(status)

	if err != nil {
		return false, err
	}

	newRequest := pb.UpdateStatusEquipmentRequestV1Request{
		EquipmentRequestId:     equipmentRequestID,
		EquipmentRequestStatus: *pbStatus,
	}

	response, err := c.grpcClient.UpdateStatusEquipmentRequestV1(ctx, &newRequest)

	if err != nil {
		return false, err
	}

	return response.Updated, nil
}

func (c *bssEquipmentRequestAPIServiceClient) UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error) {

	newRequest := pb.UpdateEquipmentIDEquipmentRequestV1Request{
		EquipmentRequestId: equipmentRequestID,
		EquipmentId:        equipmentID,
	}

	response, err := c.grpcClient.UpdateEquipmentIDEquipmentRequestV1(ctx, &newRequest)

	if err != nil {
		return false, err
	}

	return response.Updated, nil
}

func (c *bssEquipmentRequestAPIServiceClient) RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error) {

	newRequest := pb.RemoveEquipmentRequestV1Request{
		EquipmentRequestId: equipmentRequestID,
	}

	response, err := c.grpcClient.RemoveEquipmentRequestV1(ctx, &newRequest)

	if err != nil {
		return false, err
	}

	return response.Removed, nil
}
