package business

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	outputDateFormat = "2 Jan 2006 15:04:05"
)

// EquipmentRequest is a request for equipment
type EquipmentRequest struct {
	ID                     uint64                 `json:"-"`
	EmployeeID             uint64                 `json:"employee_id"`
	EquipmentID            uint64                 `json:"equipment_id"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              sql.NullTime           `json:"updated_at,omitempty"`
	DeletedAt              sql.NullTime           `json:"deleted_at,omitempty"`
	DoneAt                 sql.NullTime           `json:"done_at,omitempty"`
	EquipmentRequestStatus EquipmentRequestStatus `json:"equipment_request_status"`
}

// EquipmentRequestStatus is a status of request for equipment
type EquipmentRequestStatus string

//EquipmentRequestStatuses
const (
	// Do is an equipment request to do
	Do EquipmentRequestStatus = "EQUIPMENT_REQUEST_STATUS_DO"
	// InProgress is a equipment request in progress
	InProgress = "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"
	// Done is a done equipment request
	Done = "EQUIPMENT_REQUEST_STATUS_DONE"
	// Cancelled is a cancelled equipment request
	Cancelled = "EQUIPMENT_REQUEST_STATUS_CANCELLED"
)

//EquipmentRequestStatuses is a slice of existing equipment request statuses
var EquipmentRequestStatuses = []EquipmentRequestStatus{Do, InProgress, Done, Cancelled}

func (es EquipmentRequestStatus) String() string {
	return string(es)
}

func (e *EquipmentRequest) String() string {
	return fmt.Sprintf("EquipmentRequest ID: %d,\nEmployee ID: %d,\nEquipment ID: %d,\nCreated at: %s,\nUpdated at: %s,\nDeleted at: %s,\nDone at: %s,\nStatus: %s \n",
		e.ID,
		e.EmployeeID,
		e.EquipmentID,
		e.CreatedAt.Format(outputDateFormat),
		e.UpdatedAt.Time.Format(outputDateFormat),
		e.DeletedAt.Time.Format(outputDateFormat),
		e.DoneAt.Time.Format(outputDateFormat),
		e.EquipmentRequestStatus,
	)
}

//EquipmentRequestStatusesContains checks if there is a status in the system
func EquipmentRequestStatusesContains(status EquipmentRequestStatus) bool {
	for _, v := range EquipmentRequestStatuses {
		if v == status {
			return true
		}
	}
	return false
}
