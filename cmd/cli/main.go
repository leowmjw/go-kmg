package main

import (
	"app/api/rootly"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

type FakerRootlyServer struct{}

func (m FakerRootlyServer) GetScimUsers(w http.ResponseWriter, r *http.Request, params rootly.GetScimUsersParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) PostScimUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteScimUsersId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetScimUsersId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) PatchScimUsersId(w http.ResponseWriter, r *http.Request, id int) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentActionItem(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentActionItems(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentActionItem(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListAlerts(w http.ResponseWriter, r *http.Request, params rootly.ListAlertsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateAlert(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetAlert(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListCauses(w http.ResponseWriter, r *http.Request, params rootly.ListCausesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateCause(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteCause(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetCause(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateCause(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteCustomFieldOption(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetCustomFieldOption(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateCustomFieldOption(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListCustomFields(w http.ResponseWriter, r *http.Request, params rootly.ListCustomFieldsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateCustomField(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListCustomFieldOptions(w http.ResponseWriter, r *http.Request, customFieldId string, params rootly.ListCustomFieldOptionsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateCustomFieldOption(w http.ResponseWriter, r *http.Request, customFieldId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteCustomField(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetCustomField(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateCustomField(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListEnvironments(w http.ResponseWriter, r *http.Request, params rootly.ListEnvironmentsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me

	spew.Dump(rootly.EnvironmentList{
		Data: nil,
		Links: struct {
			rootly.Links `yaml:",inline"`
		}{},
	})

	panic("implement me")
}

func (m FakerRootlyServer) DeleteEnvironment(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetEnvironment(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateEnvironment(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentEvent(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentEvents(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentEvent(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentFeedback(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentFeedbacks(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentFeedback(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListFunctionalities(w http.ResponseWriter, r *http.Request, params rootly.ListFunctionalitiesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateFunctionality(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteFunctionality(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetFunctionality(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateFunctionality(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentRoleTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentRoleTasks(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentRoleTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentRoles(w http.ResponseWriter, r *http.Request, params rootly.ListIncidentRolesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentRole(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentRole(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentRole(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentRole(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentRoleTasks(w http.ResponseWriter, r *http.Request, incidentRoleId string, params rootly.ListIncidentRoleTasksParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentRoleTask(w http.ResponseWriter, r *http.Request, incidentRoleId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentTypes(w http.ResponseWriter, r *http.Request, params rootly.ListIncidentTypesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentType(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentType(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentType(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentType(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidents(w http.ResponseWriter, r *http.Request, params rootly.ListIncidentsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncident(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CancelIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentCustomFieldSelections(w http.ResponseWriter, r *http.Request, id string, params rootly.ListIncidentCustomFieldSelectionsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) MitigateIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ResolveIncident(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentActionItems(w http.ResponseWriter, r *http.Request, incidentId string, params rootly.ListIncidentActionItemsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentActionItem(w http.ResponseWriter, r *http.Request, incidentId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListAlert(w http.ResponseWriter, r *http.Request, incidentId string, params rootly.ListAlertParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) AttachAlert(w http.ResponseWriter, r *http.Request, incidentId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentEvents(w http.ResponseWriter, r *http.Request, incidentId string, params rootly.ListIncidentEventsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentEvent(w http.ResponseWriter, r *http.Request, incidentId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentFeedbacks(w http.ResponseWriter, r *http.Request, incidentId string, params rootly.ListIncidentFeedbacksParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentFeedback(w http.ResponseWriter, r *http.Request, incidentId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentStatusPages(w http.ResponseWriter, r *http.Request, incidentId string, params rootly.ListIncidentStatusPagesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateIncidentStatusPage(w http.ResponseWriter, r *http.Request, incidentId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeletePlaybookTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetPlaybookTasks(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdatePlaybookTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListPlaybooks(w http.ResponseWriter, r *http.Request, params rootly.ListPlaybooksParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreatePlaybook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeletePlaybook(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetPlaybook(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdatePlaybook(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListPlaybookTasks(w http.ResponseWriter, r *http.Request, playbookId string, params rootly.ListPlaybookTasksParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreatePlaybookTask(w http.ResponseWriter, r *http.Request, playbookId string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListPostmortemTemplates(w http.ResponseWriter, r *http.Request, params rootly.ListPostmortemTemplatesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreatePostmortemTemplate(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeletePostmortemTemplate(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetPostmortemTemplate(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdatePostmortemTemplate(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentPostMortems(w http.ResponseWriter, r *http.Request, params rootly.ListIncidentPostMortemsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListIncidentPostmortem(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentPostmortem(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListPulses(w http.ResponseWriter, r *http.Request, params rootly.ListPulsesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreatePulse(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetPulse(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdatePulse(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListServices(w http.ResponseWriter, r *http.Request, params rootly.ListServicesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateService(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteService(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetService(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateService(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListSeverities(w http.ResponseWriter, r *http.Request, params rootly.ListSeveritiesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateSeverity(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteSeverity(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetSeverity(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateSeverity(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteIncidentStatusPage(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetIncidentStatusPages(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateIncidentStatusPage(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListStatusPages(w http.ResponseWriter, r *http.Request, params rootly.ListStatusPagesParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateStatusPage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteStatusPage(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetStatusPage(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateStatusPage(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListTeams(w http.ResponseWriter, r *http.Request, params rootly.ListTeamsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateTeam(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteTeam(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetTeam(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateTeam(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteWorkflowCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetWorkflowCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateWorkflowCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteWorkflowTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetWorkflowTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateWorkflowTask(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListWorkflows(w http.ResponseWriter, r *http.Request, params rootly.ListWorkflowsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) DeleteWorkflow(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) GetWorkflow(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) UpdateWorkflow(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListWorkflowCustomFieldSelections(w http.ResponseWriter, r *http.Request, id string, params rootly.ListWorkflowCustomFieldSelectionsParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateWorkflowCustomFieldSelection(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) ListWorkflowTasks(w http.ResponseWriter, r *http.Request, workflowId string, params rootly.ListWorkflowTasksParams) {
	//TODO implement me
	panic("implement me")
}

func (m FakerRootlyServer) CreateWorkflowTask(w http.ResponseWriter, r *http.Request, workflowId string) {
	//TODO implement me
	panic("implement me")
}

func main() {
	fmt.Println("Welcome to Kejadian Menjadi Ganguan (KMG)!!")
	// Faker Service is in-mem implementation that pass all Behavior TDDs
	// Used by other teams; can setup pre-req scenarios like happy path, sad path, mutiples
	// even have basic state tracking?
	frs := FakerRootlyServer{}
	http.ListenAndServe(":8080", rootly.Handler(frs))

}
