package business

import (
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/ozonmp/omp-bot/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// ErrUnableToConvertEquipmentRequestStatus is a unable to convert model error
	ErrUnableToConvertEquipmentRequestStatus = errors.New("unable to convert equipment request status")
)

//ConvertEquipmentRequestStatusToPb - convert EquipmentRequestStatus to protobuf EquipmentRequestStatus enum
func ConvertEquipmentRequestStatusToPb(equipmentRequestStatus EquipmentRequestStatus) (*pb.EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_value[string(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestPbStatus := pb.EquipmentRequestStatus(status)

	return &equipmentRequestPbStatus, nil
}

//ConvertEquipmentRequestToCreatePbRequest - convert EquipmentRequest to protobuf CreateEquipmentRequestV1Request message
func ConvertEquipmentRequestToCreatePbRequest(equipmentRequest *EquipmentRequest) (*pb.CreateEquipmentRequestV1Request, error) {
	equipmentRequestStatus, err := ConvertEquipmentRequestStatusToPb(equipmentRequest.EquipmentRequestStatus)

	if err != nil {
		return nil, err
	}

	return &pb.CreateEquipmentRequestV1Request{
		EmployeeId:             equipmentRequest.EmployeeID,
		EquipmentId:            equipmentRequest.EquipmentID,
		CreatedAt:              timestamppb.New(equipmentRequest.CreatedAt),
		UpdatedAt:              model.ConvertNullableTimeToPbTime(equipmentRequest.UpdatedAt),
		DoneAt:                 model.ConvertNullableTimeToPbTime(equipmentRequest.DoneAt),
		DeletedAt:              model.ConvertNullableTimeToPbTime(equipmentRequest.DeletedAt),
		EquipmentRequestStatus: *equipmentRequestStatus,
	}, nil
}
