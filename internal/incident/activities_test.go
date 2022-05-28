package incident

import "testing"

func TestIncidentCommsActivity_PublishPrivate(t *testing.T) {
	type fields struct {
		PublicPublishFunc  func([]string) error
		PrivatePublishFunc func([]string) error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ica := IncidentCommsActivity{
				PublicPublishFunc:  tt.fields.PublicPublishFunc,
				PrivatePublishFunc: tt.fields.PrivatePublishFunc,
			}
			ica.PublishPrivate()
		})
	}
}

func TestIncidentCommsActivity_PublishPublic(t *testing.T) {
	type fields struct {
		PublicPublishFunc  func([]string) error
		PrivatePublishFunc func([]string) error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ica := IncidentCommsActivity{
				PublicPublishFunc:  tt.fields.PublicPublishFunc,
				PrivatePublishFunc: tt.fields.PrivatePublishFunc,
			}
			ica.PublishPublic()
		})
	}
}

func TestIncidentCommsActivity_Setup(t *testing.T) {
	type fields struct {
		PublicPublishFunc  func([]string) error
		PrivatePublishFunc func([]string) error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ica := IncidentCommsActivity{
				PublicPublishFunc:  tt.fields.PublicPublishFunc,
				PrivatePublishFunc: tt.fields.PrivatePublishFunc,
			}
			ica.Setup()
		})
	}
}

func TestIncidentLifecycleActivity_AddAction(t *testing.T) {
	type fields struct {
		IsIncidentActive        bool
		ChangeChannelFunc       func() error
		MarkToRecordFunc        func() error
		AddActionOrReminderFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ila := IncidentLifecycleActivity{
				IsIncidentActive:        tt.fields.IsIncidentActive,
				ChangeChannelFunc:       tt.fields.ChangeChannelFunc,
				MarkToRecordFunc:        tt.fields.MarkToRecordFunc,
				AddActionOrReminderFunc: tt.fields.AddActionOrReminderFunc,
			}
			ila.AddAction()
		})
	}
}

func TestIncidentLifecycleActivity_AddReminder(t *testing.T) {
	type fields struct {
		IsIncidentActive        bool
		ChangeChannelFunc       func() error
		MarkToRecordFunc        func() error
		AddActionOrReminderFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ila := IncidentLifecycleActivity{
				IsIncidentActive:        tt.fields.IsIncidentActive,
				ChangeChannelFunc:       tt.fields.ChangeChannelFunc,
				MarkToRecordFunc:        tt.fields.MarkToRecordFunc,
				AddActionOrReminderFunc: tt.fields.AddActionOrReminderFunc,
			}
			ila.AddReminder()
		})
	}
}

func TestIncidentLifecycleActivity_ChangeChannel(t *testing.T) {
	type fields struct {
		IsIncidentActive        bool
		ChangeChannelFunc       func() error
		MarkToRecordFunc        func() error
		AddActionOrReminderFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ila := IncidentLifecycleActivity{
				IsIncidentActive:        tt.fields.IsIncidentActive,
				ChangeChannelFunc:       tt.fields.ChangeChannelFunc,
				MarkToRecordFunc:        tt.fields.MarkToRecordFunc,
				AddActionOrReminderFunc: tt.fields.AddActionOrReminderFunc,
			}
			ila.ChangeChannel()
		})
	}
}

func TestIncidentLifecycleActivity_MarkToRecord(t *testing.T) {
	type fields struct {
		IsIncidentActive        bool
		ChangeChannelFunc       func() error
		MarkToRecordFunc        func() error
		AddActionOrReminderFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ila := IncidentLifecycleActivity{
				IsIncidentActive:        tt.fields.IsIncidentActive,
				ChangeChannelFunc:       tt.fields.ChangeChannelFunc,
				MarkToRecordFunc:        tt.fields.MarkToRecordFunc,
				AddActionOrReminderFunc: tt.fields.AddActionOrReminderFunc,
			}
			ila.MarkToRecord()
		})
	}
}

func TestIncidentLifecycleActivity_Setup(t *testing.T) {
	type fields struct {
		IsIncidentActive        bool
		ChangeChannelFunc       func() error
		MarkToRecordFunc        func() error
		AddActionOrReminderFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ila := IncidentLifecycleActivity{
				IsIncidentActive:        tt.fields.IsIncidentActive,
				ChangeChannelFunc:       tt.fields.ChangeChannelFunc,
				MarkToRecordFunc:        tt.fields.MarkToRecordFunc,
				AddActionOrReminderFunc: tt.fields.AddActionOrReminderFunc,
			}
			ila.Setup()
		})
	}
}

func TestIncidentWarRoomActivity_InitiateWarRoom(t *testing.T) {
	type fields struct {
		InitiateWarRoomFunc func() error
		TearDownWarRoomFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iwa := IncidentWarRoomActivity{
				InitiateWarRoomFunc: tt.fields.InitiateWarRoomFunc,
				TearDownWarRoomFunc: tt.fields.TearDownWarRoomFunc,
			}
			iwa.InitiateWarRoom()
		})
	}
}

func TestIncidentWarRoomActivity_Setup(t *testing.T) {
	type fields struct {
		InitiateWarRoomFunc func() error
		TearDownWarRoomFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iwa := IncidentWarRoomActivity{
				InitiateWarRoomFunc: tt.fields.InitiateWarRoomFunc,
				TearDownWarRoomFunc: tt.fields.TearDownWarRoomFunc,
			}
			iwa.Setup()
		})
	}
}

func TestIncidentWarRoomActivity_TearDownWarRoom(t *testing.T) {
	type fields struct {
		InitiateWarRoomFunc func() error
		TearDownWarRoomFunc func() error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iwa := IncidentWarRoomActivity{
				InitiateWarRoomFunc: tt.fields.InitiateWarRoomFunc,
				TearDownWarRoomFunc: tt.fields.TearDownWarRoomFunc,
			}
			iwa.TearDownWarRoom()
		})
	}
}
