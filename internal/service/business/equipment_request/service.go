package equipment_request

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/app/helpers"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type EquipmentRequestService interface {
	Get(equipmentRequestId uint64) (*business.EquipmentRequest, error)
	List(page uint64, perPage uint64) ([]business.EquipmentRequest, error)
	Create(equipmentRequest business.EquipmentRequest) (uint64, error)
	Update(equipmentRequestId uint64, equipmentRequest business.EquipmentRequest) error
	Remove(equipmentRequestId uint64) (bool, error)
	Count() uint64
}

type DummyEquipmentRequestService struct{}

func NewDummyEquipmentRequestService() *DummyEquipmentRequestService {
	return &DummyEquipmentRequestService{}
}

func (s *DummyEquipmentRequestService) List(page uint64, perPage uint64) ([]business.EquipmentRequest, error) {
	start := page * perPage
	end := start + perPage
	count := s.Count()
	end = helpers.Min(end, s.Count())

	if start > end || start > count {
		return nil, errors.New(fmt.Sprintf("out of bounds with page %d and per page %d", page, perPage))
	}
	return business.AllEquipmentRequests.List[start:end], nil
}

func (s *DummyEquipmentRequestService) Get(equipmentRequestId uint64) (*business.EquipmentRequest, error) {
	return s.get(equipmentRequestId)
}

func (s *DummyEquipmentRequestService) Remove(equipmentRequestId uint64) (bool, error) {
	for i := range business.AllEquipmentRequests.List {
		if business.AllEquipmentRequests.List[i].Id == equipmentRequestId {
			business.AllEquipmentRequests.List = append(business.AllEquipmentRequests.List[:i], business.AllEquipmentRequests.List[i+1:]...)

			return true, nil
		}
	}
	return false, errors.New(fmt.Sprintf("index out of range %d", equipmentRequestId))
}
func (s *DummyEquipmentRequestService) Create(equipmentRequest business.EquipmentRequest) (uint64, error) {
	itemId := business.AllEquipmentRequests.LastId + 1
	equipmentRequest.Id = itemId
	business.AllEquipmentRequests.List = append(business.AllEquipmentRequests.List, equipmentRequest)
	business.AllEquipmentRequests.LastId = itemId
	return itemId, nil
}

func (s *DummyEquipmentRequestService) Update(equipmentRequestId uint64, equipmentRequest business.EquipmentRequest) error {
	item, err := s.get(equipmentRequestId)
	if err != nil {
		return err
	}

	item.DoneAt = equipmentRequest.DoneAt
	item.CreatedAt = equipmentRequest.CreatedAt
	item.EquipmentId = equipmentRequest.EquipmentId
	item.EquipmentType = equipmentRequest.EquipmentType
	item.Status = equipmentRequest.Status
	item.EmployeeId = equipmentRequest.EmployeeId

	return nil
}

func (s *DummyEquipmentRequestService) Count() uint64 {
	return uint64(len(business.AllEquipmentRequests.List))
}

func (s *DummyEquipmentRequestService) get(equipmentRequestId uint64) (*business.EquipmentRequest, error) {
	for i := range business.AllEquipmentRequests.List {
		if business.AllEquipmentRequests.List[i].Id == equipmentRequestId {
			return &business.AllEquipmentRequests.List[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("index out of range %d", equipmentRequestId))
}
