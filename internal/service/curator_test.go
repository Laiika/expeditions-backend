package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCuratorService_GetExpeditionCurators(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockCuratorRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Curators
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:          context.Background(),
				client:       nil,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().GetExpeditionCurators(args.ctx, args.client, args.expeditionId).
					Return(entity.Curators{
						&entity.Curator{
							Id:   1,
							Name: "aaa",
						},
					}, nil)
			},
			want: entity.Curators{
				&entity.Curator{
					Id:   1,
					Name: "aaa",
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
			curatorRepo := mocks.NewMockCuratorRepo(ctrl)
			tc.mockBehavior(curatorRepo, tc.args)

			// init service
			s := NewCuratorService(curatorRepo)

			// run test
			got, err := s.GetExpeditionCurators(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCuratorService_GetAllCurators(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockCuratorRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Curators
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().GetAllCurators(args.ctx, args.client).
					Return(entity.Curators{
						&entity.Curator{
							Id:   1,
							Name: "aaa",
						},
					}, nil)
			},
			want: entity.Curators{
				&entity.Curator{
					Id:   1,
					Name: "aaa",
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
			curatorRepo := mocks.NewMockCuratorRepo(ctrl)
			tc.mockBehavior(curatorRepo, tc.args)

			// init service
			s := NewCuratorService(curatorRepo)

			// run test
			got, err := s.GetAllCurators(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCuratorService_CreateCurator(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateCuratorInput
	}

	type MockBehavior func(m *mocks.MockCuratorRepo, args args)

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
				input: &entity.CreateCuratorInput{
					Name: "aaa",
				},
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().CreateCurator(args.ctx, args.client, &entity.Curator{
					Name: args.input.Name,
				}).
					Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "curator already exists error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				input: &entity.CreateCuratorInput{
					Name: "aaa",
				},
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().CreateCurator(args.ctx, args.client, &entity.Curator{
					Name: args.input.Name,
				}).
					Return(0, ErrCuratorAlreadyExists)
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
			curatorRepo := mocks.NewMockCuratorRepo(ctrl)
			tc.mockBehavior(curatorRepo, tc.args)

			// init service
			s := NewCuratorService(curatorRepo)

			// run test
			got, err := s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCuratorService_CreateCuratorExpedition(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		curatorId    int
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockCuratorRepo, args args)

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
				curatorId:    1,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().CreateCuratorExpedition(args.ctx, args.client, args.curatorId, args.expeditionId).
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
			curatorRepo := mocks.NewMockCuratorRepo(ctrl)
			tc.mockBehavior(curatorRepo, tc.args)

			// init service
			s := NewCuratorService(curatorRepo)

			// run test
			got, err := s.CreateCuratorExpedition(tc.args.ctx, tc.args.client, tc.args.curatorId, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCuratorService_DeleteCurator(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockCuratorRepo, args args)

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
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().DeleteCurator(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "curator not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockCuratorRepo, args args) {
				m.EXPECT().DeleteCurator(args.ctx, args.client, args.id).
					Return(ErrCuratorNotFound)
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
			curatorRepo := mocks.NewMockCuratorRepo(ctrl)
			tc.mockBehavior(curatorRepo, tc.args)

			// init service
			s := NewCuratorService(curatorRepo)

			// run test
			err := s.DeleteCurator(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
