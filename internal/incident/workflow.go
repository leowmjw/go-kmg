package incident

import (
	"github.com/davecgh/go-spew/spew"
	"go.temporal.io/sdk/workflow"
)

type IncidentLifecycleReq struct {
	ID          string // IncidentID: xid
	ChannelName string
	SlackID     string
}

// IncidentLifecycleWorkflow tracks lifecycle of incident; pattern is: YYYY-MM-DD-<description>
func IncidentLifecycleWorkflow(ctx workflow.Context, req IncidentLifecycleReq) error {
	// Start; with WorkflowID being the channel-incident as per standard pattern; original

	thisIncident := Incident{}
	thisIncident.Status = ACTIVE_IS
	spew.Dump(thisIncident)
	// Re-open of previously closed incident; accidentally; within same day, more than 1 day not allowed!

	// Action: Appoint Roles
	// Action: Start war-room
	// Action: End war-room

	// Action: Mark for record

	// Action: Mark for next action/reminder during incident

	// Action: Channel name change; but SlackID remain

	// Action: Close Incident
	thisIncident.Status = RESOLVED_IS
	// Action: Mark for next action/reminder post-incident

	// Action: All outstanding actions done
	thisIncident.Status = ARCHIVED_IS
	return nil
}

type IncidentCommsReq struct {
	IncidentID string
}

// IncidentCommsWorkflow is to manage comms both publicly and privately
func IncidentCommsWorkflow(ctx workflow.Context, req IncidentCommsReq) error {
	var publicContent, privateContent []string
	spew.Dump(publicContent)
	spew.Dump(privateContent)
	// Above can be queried; and adjusted .. start with append only
	// 	next time can move things around ..

	// Action: Public Publish; the recommended changes public; may need approval of Incident Commander
	//	this kicks off the static content rebuild + republish ..
	// Action: Private
	return nil
}

type IncidentWarRoomReq struct {
	IncidentID string
}

// IncidentWarRoomWorkflow is to manage the video + chat during an incident; including recording next time
func IncidentWarRoomWorkflow(ctx workflow.Context, req IncidentWarRoomReq) error {
	// Setup the initial warroom ..
	// Action: Re-open if needed ..
	// Action: Close down
	//	<TODO>: Recording ...

	return nil
}
