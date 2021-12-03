package business

import (
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"

	"github.com/ozonmp/omp-bot/internal/model"
)

//ConvertRepeatedFacadePbToEquipmentRequests - convert slice of protobuf EquipmentRequest messages to EquipmentRequests
func ConvertRepeatedFacadePbToEquipmentRequests(equipmentRequestsPb []*facadepb.EquipmentRequest) ([]*EquipmentRequest, error) {
	var equipmentRequests []*EquipmentRequest

	for i := range equipmentRequestsPb {
		equipmentRequest, err := ConvertFacadePbToEquipmentRequest(equipmentRequestsPb[i])
		if err != nil {
			return nil, err
		}
		equipmentRequests = append(equipmentRequests, equipmentRequest)
	}

	return equipmentRequests, nil
}

//ConvertFacadePbToEquipmentRequest - convert protobuf EquipmentRequest message to EquipmentRequest
func ConvertFacadePbToEquipmentRequest(equipmentRequest *facadepb.EquipmentRequest) (*EquipmentRequest, error) {
	equipmentRequestStatus, err := ConvertFacadePbEquipmentRequestStatus(equipmentRequest.EquipmentRequestStatus)

	if err != nil {
		return nil, err
	}

	return &EquipmentRequest{
		ID:                     equipmentRequest.Id,
		EmployeeID:             equipmentRequest.EmployeeId,
		EquipmentID:            equipmentRequest.EquipmentId,
		CreatedAt:              equipmentRequest.CreatedAt.AsTime(),
		UpdatedAt:              model.ConvertPbTimeToNullableTime(equipmentRequest.UpdatedAt),
		DoneAt:                 model.ConvertPbTimeToNullableTime(equipmentRequest.DoneAt),
		DeletedAt:              model.ConvertPbTimeToNullableTime(equipmentRequest.DeletedAt),
		EquipmentRequestStatus: *equipmentRequestStatus,
	}, nil
}

//ConvertFacadePbEquipmentRequestStatus - convert protobuf EquipmentRequestStatus enum to EquipmentRequestStatus
func ConvertFacadePbEquipmentRequestStatus(equipmentRequestStatus facadepb.EquipmentRequestStatus) (*EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_name[int32(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestModelStatus := EquipmentRequestStatus(status)

	return &equipmentRequestModelStatus, nil
}
