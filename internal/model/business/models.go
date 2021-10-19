package business

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/app/helpers"
)

const (
	dataBaseDateFormat = "2006-01-02T15:04:05"
	outputDateFormat   = "2 Jan 2006 15:04:05"
)

type EquipmentRequest struct {
	Id            uint64 `json:"-"`
	EmployeeId    uint64 `json:"employee_id"`
	EquipmentType string `json:"equipment_type"`
	EquipmentId   uint64 `json:"equipment_id"`
	CreatedAt     string `json:"created_at"`
	DoneAt        string `json:"done_at"`
	Status        bool   `json:"status"`
}

type EquipmentRequestList struct {
	LastId uint64
	List   []EquipmentRequest
}

func (equipmentRequest *EquipmentRequest) String() string {
	createdAtDate := helpers.DateFormatter(equipmentRequest.CreatedAt, dataBaseDateFormat, outputDateFormat)
	doneAtDate := helpers.DateFormatter(equipmentRequest.DoneAt, dataBaseDateFormat, outputDateFormat)

	return fmt.Sprintf("EquipmentRequest ID: %d,\nEmployee ID: %d,\nEquipment Type: %s,\nEquipment ID: %d,\nCreated at: %s,\nDone at: %s,\nStatus: %t \n",
		equipmentRequest.Id,
		equipmentRequest.EmployeeId,
		equipmentRequest.EquipmentType,
		equipmentRequest.EquipmentId,
		createdAtDate,
		doneAtDate,
		equipmentRequest.Status,
	)
}

var AllEquipmentRequests = EquipmentRequestList{
	6,
	[]EquipmentRequest{
		{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		{Id: 2, EmployeeId: 1, EquipmentType: "Keyboard", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		{Id: 3, EmployeeId: 1, EquipmentType: "Keyboard", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		{Id: 4, EmployeeId: 2, EquipmentType: "Laptop", EquipmentId: 2, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		{Id: 5, EmployeeId: 3, EquipmentType: "Keyboard", EquipmentId: 2, CreatedAt: "2021-01-19T10:00:00", DoneAt: "2021-01-19T10:00:00", Status: true},
		{Id: 6, EmployeeId: 4, EquipmentType: "Mouse", EquipmentId: 1, CreatedAt: "2021-01-19T10:00:00", DoneAt: "2021-01-19T10:00:00", Status: true},
	},
}
