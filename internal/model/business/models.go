package business

import (
	"fmt"
	"time"
)

const (
	outputDateFormat = "2 Jan 2006 15:04:05"
)

type EquipmentRequest struct {
	Id            uint64                 `json:"-"`
	EmployeeId    uint64                 `json:"employee_id"`
	EquipmentType string                 `json:"equipment_type"`
	EquipmentId   uint64                 `json:"equipment_id"`
	CreatedAt     time.Time              `json:"created_at"`
	DoneAt        time.Time              `json:"done_at"`
	Status        EquipmentRequestStatus `json:"status"`
}

type EquipmentRequestList struct {
	LastId uint64
	List   []EquipmentRequest
}

type EquipmentRequestStatus int

const (
	Do EquipmentRequestStatus = iota
	InProgress
	Done
	Cancelled
)

func (es EquipmentRequestStatus) String() string {
	return [...]string{"Do", "In Progress", "Done", "Cancelled"}[es]
}

func (e *EquipmentRequest) String() string {
	return fmt.Sprintf("EquipmentRequest ID: %d,\nEmployee ID: %d,\nEquipment Type: %s,\nEquipment ID: %d,\nCreated at: %s,\nDone at: %s,\nStatus: %s \n",
		e.Id,
		e.EmployeeId,
		e.EquipmentType,
		e.EquipmentId,
		e.CreatedAt.Format(outputDateFormat),
		e.DoneAt.Format(outputDateFormat),
		e.Status,
	)
}

var AllEquipmentRequests = EquipmentRequestList{
	6,
	[]EquipmentRequest{
		{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: time.Now(), DoneAt: time.Now(), Status: Do},
		{Id: 2, EmployeeId: 1, EquipmentType: "Keyboard", EquipmentId: 1, CreatedAt: time.Now(), DoneAt: time.Now(), Status: InProgress},
		{Id: 3, EmployeeId: 1, EquipmentType: "Keyboard", EquipmentId: 1, CreatedAt: time.Now(), DoneAt: time.Now(), Status: Done},
		{Id: 4, EmployeeId: 2, EquipmentType: "Laptop", EquipmentId: 2, CreatedAt: time.Now(), DoneAt: time.Now(), Status: Cancelled},
		{Id: 5, EmployeeId: 3, EquipmentType: "Keyboard", EquipmentId: 2, CreatedAt: time.Now(), DoneAt: time.Now(), Status: Done},
		{Id: 6, EmployeeId: 4, EquipmentType: "Mouse", EquipmentId: 1, CreatedAt: time.Now(), DoneAt: time.Now(), Status: Do},
	},
}
