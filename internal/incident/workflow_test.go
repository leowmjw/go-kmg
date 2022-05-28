package incident

import (
	"go.temporal.io/sdk/workflow"
	"testing"
)

func TestIncidentCommsWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		req IncidentCommsReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncidentCommsWorkflow(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("IncidentCommsWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIncidentLifecycleWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		req IncidentLifecycleReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncidentLifecycleWorkflow(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("IncidentLifecycleWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIncidentWarRoomWorkflow(t *testing.T) {
	type args struct {
		ctx workflow.Context
		req IncidentWarRoomReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncidentWarRoomWorkflow(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("IncidentWarRoomWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
