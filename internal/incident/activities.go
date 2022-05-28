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

func NewRealIncidentWarRoomActivity() IncidentWarRoomActivity {
	// Connect to 100ms
	// Take credentials from ENV ..
	// STart with hard coded first ..
	return IncidentWarRoomActivity{}
}

func NewFakerIncidentWarRoomActivity() IncidentWarRoomActivity {
	// depending on the scenarios .. we can return different behavior
	return IncidentWarRoomActivity{
		InitiateWarRoomFunc: func() error {
			return nil
		},
		TearDownWarRoomFunc: func() error {
			return nil
		},
	}
}

func (iwa IncidentWarRoomActivity) Setup() {

}

func (iwa IncidentWarRoomActivity) InitiateWarRoom() {
	err := iwa.InitiateWarRoomFunc()
	if err != nil {
		panic(err)
	}
}

func (iwa IncidentWarRoomActivity) TearDownWarRoom() {

}
