package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberService_GetExpeditionMembers(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockMemberRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Members
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:          context.Background(),
				client:       nil,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().GetExpeditionMembers(args.ctx, args.client, args.expeditionId).
					Return(entity.Members{
						&entity.Member{
							Id:          1,
							Name:        "aaa",
							PhoneNumber: "+79021061232",
							Role:        "aaa",
							Login:       "dhhjds",
							Password:    "jdskjdsjk",
						},
					}, nil)
			},
			want: entity.Members{
				&entity.Member{
					Id:          1,
					Name:        "aaa",
					PhoneNumber: "+79021061232",
					Role:        "aaa",
					Login:       "dhhjds",
					Password:    "jdskjdsjk",
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
			memberRepo := mocks.NewMockMemberRepo(ctrl)
			tc.mockBehavior(memberRepo, tc.args)

			// init service
			s := NewMemberService(memberRepo)

			// run test
			got, err := s.GetExpeditionMembers(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMemberService_GetAllMembers(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockMemberRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Members
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().GetAllMembers(args.ctx, args.client).
					Return(entity.Members{
						&entity.Member{
							Id:          1,
							Name:        "aaa",
							PhoneNumber: "+79021061232",
							Role:        "aaa",
							Login:       "dhhjds",
							Password:    "jdskjdsjk",
						},
					}, nil)
			},
			want: entity.Members{
				&entity.Member{
					Id:          1,
					Name:        "aaa",
					PhoneNumber: "+79021061232",
					Role:        "aaa",
					Login:       "dhhjds",
					Password:    "jdskjdsjk",
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
			memberRepo := mocks.NewMockMemberRepo(ctrl)
			tc.mockBehavior(memberRepo, tc.args)

			// init service
			s := NewMemberService(memberRepo)

			// run test
			got, err := s.GetAllMembers(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMemberService_CreateMember(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateMemberInput
	}

	type MockBehavior func(m *mocks.MockMemberRepo, args args)

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
				input: &entity.CreateMemberInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Role:        "aaa",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().CreateMember(args.ctx, args.client, gomock.Any()).
					Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "member already exists error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				input: &entity.CreateMemberInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Role:        "aaa",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().CreateMember(args.ctx, args.client, gomock.Any()).
					Return(0, ErrMemberAlreadyExists)
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
			memberRepo := mocks.NewMockMemberRepo(ctrl)
			tc.mockBehavior(memberRepo, tc.args)

			// init service
			s := NewMemberService(memberRepo)

			// run test
			got, err := s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCuratorService_CreateMemberExpedition(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		memberId     int
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockMemberRepo, args args)

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
				ctx:          context.Background(),
				client:       nil,
				memberId:     1,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().CreateMemberExpedition(args.ctx, args.client, args.memberId, args.expeditionId).
					Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// init mocks
			memberRepo := mocks.NewMockMemberRepo(ctrl)
			tc.mockBehavior(memberRepo, tc.args)

			// init service
			s := NewMemberService(memberRepo)

			// run test
			got, err := s.CreateMemberExpedition(tc.args.ctx, tc.args.client, tc.args.memberId, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMemberService_DeleteMember(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockMemberRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         error
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     1,
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().DeleteMember(args.ctx, args.client, args.id).
					Return(nil)
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "member not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockMemberRepo, args args) {
				m.EXPECT().DeleteMember(args.ctx, args.client, args.id).
					Return(ErrMemberNotFound)
			},
			want:    ErrMemberNotFound,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// init mocks
			memberRepo := mocks.NewMockMemberRepo(ctrl)
			tc.mockBehavior(memberRepo, tc.args)

			// init service
			s := NewMemberService(memberRepo)

			// run test
			err := s.DeleteMember(tc.args.ctx, tc.args.client, tc.args.id)
			assert.Equal(t, tc.want, err)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
