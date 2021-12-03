package equipment_request

import (
	"context"
	"github.com/opentracing/opentracing-go"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	facadepb "github.com/ozonmp/bss-equipment-request-facade/pkg/bss-equipment-request-facade"
	"github.com/ozonmp/omp-bot/internal/client/business/equipment_request_api"
	"github.com/ozonmp/omp-bot/internal/client/business/equipment_request_facade_api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"github.com/pkg/errors"
)

// ErrNoExistsEquipmentRequest is a "equipment request not founded" error
var ErrNoExistsEquipmentRequest = errors.Wrap(equipment_request_facade_api.ErrNoExistsEquipmentRequest, "EquipmentRequestService")

// ErrEmptyListEquipmentRequest is a "empty list returned" error
var ErrEmptyListEquipmentRequest = errors.Wrap(equipment_request_facade_api.ErrNoListEquipmentRequest, "EquipmentRequestService")

// EquipmentRequestService is a interface for equipment request bot
type EquipmentRequestService interface {
	Get(ctx context.Context, equipmentRequestID uint64) (*business.EquipmentRequest, error)
	Remove(ctx context.Context, equipmentRequestID uint64) (bool, error)
	Create(ctx context.Context, equipmentRequest business.EquipmentRequest) (uint64, error)
	UpdateStatus(ctx context.Context, equipmentRequestID uint64, status business.EquipmentRequestStatus) (bool, error)
	UpdateEquipmentID(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error)

	List(ctx context.Context, page uint64, perPage uint64) ([]*business.EquipmentRequest, uint64, error)
}

// BssEquipmentRequestAPIServiceClient is a interface for equipment request api client
type BssEquipmentRequestAPIServiceClient interface {
	CreateEquipmentRequest(ctx context.Context, equipmentRequest business.EquipmentRequest) (uint64, error)
	UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status business.EquipmentRequestStatus) (bool, error)
	UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error)
}

// BssEquipmentRequestFacadeAPIServiceClient is a interface for equipment request facade api client
type BssEquipmentRequestFacadeAPIServiceClient interface {
	ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]*business.EquipmentRequest, uint64, error)
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*business.EquipmentRequest, error)
}

type equipmentRequestService struct {
	apiClient    BssEquipmentRequestAPIServiceClient
	facadeClient BssEquipmentRequestFacadeAPIServiceClient
}

// NewEquipmentRequestService returns EquipmentRequestService
func NewEquipmentRequestService(apiClient pb.BssEquipmentRequestApiServiceClient, facadeClient facadepb.BssEquipmentRequestFacadeApiServiceClient) EquipmentRequestService {
	return &equipmentRequestService{
		apiClient:    equipment_request_api.NewBssEquipmentRequestAPIServiceClient(apiClient),
		facadeClient: equipment_request_facade_api.NewBssEquipmentRequestFacadeAPIServiceClient(facadeClient),
	}
}

func (s *equipmentRequestService) Get(ctx context.Context, equipmentRequestID uint64) (*business.EquipmentRequest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.Get")
	defer span.Finish()

	equipmentRequest, err := s.facadeClient.DescribeEquipmentRequest(ctx, equipmentRequestID)

	if err != nil {
		if errors.Is(err, equipment_request_facade_api.ErrNoExistsEquipmentRequest) {
			return nil, ErrNoExistsEquipmentRequest
		}
		return nil, err
	}

	return equipmentRequest, nil
}

func (s *equipmentRequestService) Remove(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.Remove")
	defer span.Finish()

	removed, err := s.apiClient.RemoveEquipmentRequest(ctx, equipmentRequestID)

	if err != nil {
		return false, err
	}

	return removed, nil
}

func (s *equipmentRequestService) Create(ctx context.Context, equipmentRequest business.EquipmentRequest) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.Create")
	defer span.Finish()

	requestID, err := s.apiClient.CreateEquipmentRequest(ctx, equipmentRequest)

	if err != nil {
		return 0, err
	}

	return requestID, nil
}

func (s *equipmentRequestService) UpdateStatus(ctx context.Context, equipmentRequestID uint64, status business.EquipmentRequestStatus) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.UpdateStatus")
	defer span.Finish()

	updated, err := s.apiClient.UpdateStatusEquipmentRequest(ctx, equipmentRequestID, status)

	if err != nil {
		return false, err
	}

	return updated, nil
}

func (s *equipmentRequestService) UpdateEquipmentID(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.UpdateEquipmentID")
	defer span.Finish()

	updated, err := s.apiClient.UpdateEquipmentIDEquipmentRequest(ctx, equipmentRequestID, equipmentID)

	if err != nil {
		return false, err
	}

	return updated, nil
}

func (s *equipmentRequestService) List(ctx context.Context, page uint64, perPage uint64) ([]*business.EquipmentRequest, uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "equipmentRequestService.List")
	defer span.Finish()

	offset := page * perPage

	requests, total, err := s.facadeClient.ListEquipmentRequest(ctx, perPage, offset)

	if err != nil {
		if errors.Is(err, equipment_request_facade_api.ErrNoListEquipmentRequest) {
			return nil, 0, ErrEmptyListEquipmentRequest
		}
		return nil, 0, err
	}

	return requests, total, nil
}
