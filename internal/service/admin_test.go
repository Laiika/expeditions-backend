package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminService_GetAllAdmins(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockAdminRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Admins
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockAdminRepo, args args) {
				m.EXPECT().GetAllAdmins(args.ctx, args.client).
					Return(entity.Admins{
						&entity.Admin{
							Id:       1,
							Name:     "aaa",
							Login:    "dhhjds",
							Password: "jdskjdsjk",
						},
					}, nil)
			},
			want: entity.Admins{
				&entity.Admin{
					Id:       1,
					Name:     "aaa",
					Login:    "dhhjds",
					Password: "jdskjdsjk",
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// init mocks
			adminRepo := mocks.NewMockAdminRepo(ctrl)
			tc.mockBehavior(adminRepo, tc.args)

			// init service
			s := NewAdminService(adminRepo)

			// run test
			got, err := s.GetAllAdmins(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAdminService_CreateAdmin(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateAdminInput
	}

	type MockBehavior func(m *mocks.MockAdminRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
				input: &entity.CreateAdminInput{
					Name:     "aaa",
					Login:    "ccc",
					Password: "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockAdminRepo, args args) {
				m.EXPECT().CreateAdmin(args.ctx, args.client, gomock.Any()).
					Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "admin already exists error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				input: &entity.CreateAdminInput{
					Name:     "aaa",
					Login:    "ccc",
					Password: "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockAdminRepo, args args) {
				m.EXPECT().CreateAdmin(args.ctx, args.client, gomock.Any()).
					Return(0, ErrAdminAlreadyExists)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// init mocks
			adminRepo := mocks.NewMockAdminRepo(ctrl)
			tc.mockBehavior(adminRepo, tc.args)

			// init service
			s := NewAdminService(adminRepo)

			// run test
			got, err := s.CreateAdmin(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAdminService_DeleteAdmin(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockAdminRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     1,
			},
			mockBehavior: func(m *mocks.MockAdminRepo, args args) {
				m.EXPECT().DeleteAdmin(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "admin not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockAdminRepo, args args) {
				m.EXPECT().DeleteAdmin(args.ctx, args.client, args.id).
					Return(ErrAdminNotFound)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// init mocks
			adminRepo := mocks.NewMockAdminRepo(ctrl)
			tc.mockBehavior(adminRepo, tc.args)

			// init service
			s := NewAdminService(adminRepo)

			// run test
			err := s.DeleteAdmin(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
