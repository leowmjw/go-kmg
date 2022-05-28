package incident

import "time"

//go:generate stringer -type=IncidentType
type IncidentType int

const (
	UNKNOWN_IT IncidentType = iota
	TEST_IT
	ACTIVE_IT
)

//go:generate stringer -type=IncidentStatus
type IncidentStatus int

const (
	UNKNOWN_IS IncidentStatus = iota
	ACTIVE_IS
	RESOLVED_IS
	ARCHIVED_IS
)

type Responder string

type Incident struct {
	ID            string // xid ..
	Type          IncidentType
	Status        IncidentStatus
	IncidentStart time.Time
	IncidentEnd   time.Time
	Commander     Responder
	PICs          []Responder
	Others        []Responder
}
