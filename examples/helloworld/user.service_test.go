package example

import (
	context "context"
	reflect "reflect"
	testing "testing"

	pb "github.com/erda-project/erda-infra/examples/protocol/pb"
)

func Test_userService_GetUser(t *testing.T) {
	type fields struct {
		p *provider
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.GetUserResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				p: tt.fields.p,
			}
			gotResp, err := s.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("userService.GetUser() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	type fields struct {
		p *provider
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *pb.UpdateUserResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				p: tt.fields.p,
			}
			gotResp, err := s.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("userService.UpdateUser() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
