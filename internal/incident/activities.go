package incident

type IncidentLifecycleActivity struct {
	IsIncidentActive        bool // Behavior diff based on during incident or post-incident
	ChangeChannelFunc       func() error
	MarkToRecordFunc        func() error
	AddActionOrReminderFunc func() error
}

func (ila IncidentLifecycleActivity) Setup() {

}

func (ila IncidentLifecycleActivity) ChangeChannel() {

}

func (ila IncidentLifecycleActivity) MarkToRecord() {

}

func (ila IncidentLifecycleActivity) AddAction() {

}

func (ila IncidentLifecycleActivity) AddReminder() {

}

type IncidentCommsActivity struct {
	PublicPublishFunc  func([]string) error
	PrivatePublishFunc func([]string) error
}

func (ica IncidentCommsActivity) Setup() {

}

func (ica IncidentCommsActivity) PublishPublic() {

}

func (ica IncidentCommsActivity) PublishPrivate() {

}

type IncidentWarRoomActivity struct {
	InitiateWarRoomFunc func() error
	TearDownWarRoomFunc func() error
}

func (iwa IncidentWarRoomActivity) Setup() {

}

func (iwa IncidentWarRoomActivity) InitiateWarRoom() {

}

func (iwa IncidentWarRoomActivity) TearDownWarRoom() {

}
