package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpeditionService_GetAllExpeditions(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockExpeditionRepo, args args)

	layout := "2000-01-01"
	start, _ := time.Parse(layout, "2024-07-01")
	end, _ := time.Parse(layout, "2024-08-01")

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Expeditions
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().GetAllExpeditions(args.ctx, args.client).
					Return(entity.Expeditions{
						&entity.Expedition{
							Id:         1,
							LocationId: 1,
							StartDate:  start,
							EndDate:    end,
						},
					}, nil)
			},
			want: entity.Expeditions{
				&entity.Expedition{
					Id:         1,
					LocationId: 1,
					StartDate:  start,
					EndDate:    end,
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
			expeditionRepo := mocks.NewMockExpeditionRepo(ctrl)
			tc.mockBehavior(expeditionRepo, tc.args)

			// init service
			s := NewExpeditionService(expeditionRepo)

			// run test
			got, err := s.GetAllExpeditions(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestExpeditionService_GetLocationExpeditions(t *testing.T) {
	type args struct {
		ctx        context.Context
		client     any
		locationId int
	}

	type MockBehavior func(m *mocks.MockExpeditionRepo, args args)

	layout := "2000-01-01"
	start, _ := time.Parse(layout, "2024-07-01")
	end, _ := time.Parse(layout, "2024-08-01")

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Expeditions
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:        context.Background(),
				client:     nil,
				locationId: 1,
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().GetLocationExpeditions(args.ctx, args.client, args.locationId).
					Return(entity.Expeditions{
						&entity.Expedition{
							Id:         1,
							LocationId: 1,
							StartDate:  start,
							EndDate:    end,
						},
					}, nil)
			},
			want: entity.Expeditions{
				&entity.Expedition{
					Id:         1,
					LocationId: 1,
					StartDate:  start,
					EndDate:    end,
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
			expeditionRepo := mocks.NewMockExpeditionRepo(ctrl)
			tc.mockBehavior(expeditionRepo, tc.args)

			// init service
			s := NewExpeditionService(expeditionRepo)

			// run test
			got, err := s.GetLocationExpeditions(tc.args.ctx, tc.args.client, tc.args.locationId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestExpeditionService_CreateExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateExpeditionInput
	}

	type MockBehavior func(m *mocks.MockExpeditionRepo, args args)

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
				input: &entity.CreateExpeditionInput{
					LocationId: 1,
					StartDate:  "2024-07-01",
					EndDate:    "2024-08-01",
				},
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().CreateExpedition(args.ctx, args.client, gomock.Any()).
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
			expeditionRepo := mocks.NewMockExpeditionRepo(ctrl)
			tc.mockBehavior(expeditionRepo, tc.args)

			// init service
			s := NewExpeditionService(expeditionRepo)

			// run test
			got, err := s.CreateExpedition(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestExpeditionService_UpdateExpedition(t *testing.T) {
	type args struct {
		ctx       context.Context
		client    any
		id        int
		startDate string
		endDate   string
	}

	type MockBehavior func(m *mocks.MockExpeditionRepo, args args)

	start := time.Date(2024, 7, 1, 00, 00, 00, 00, time.UTC)
	end := time.Date(2024, 8, 1, 00, 00, 00, 00, time.UTC)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:       context.Background(),
				client:    nil,
				id:        1,
				startDate: "2024-07-01",
				endDate:   "2024-08-01",
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().UpdateExpeditionDates(args.ctx, args.client, args.id, start, end).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "expedition not found error",
			args: args{
				ctx:       context.Background(),
				client:    nil,
				id:        100,
				startDate: "2024-07-01",
				endDate:   "2024-08-01",
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().UpdateExpeditionDates(args.ctx, args.client, args.id, start, end).
					Return(ErrExpeditionNotFound)
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
			expeditionRepo := mocks.NewMockExpeditionRepo(ctrl)
			tc.mockBehavior(expeditionRepo, tc.args)

			// init service
			s := NewExpeditionService(expeditionRepo)

			// run test
			err := s.UpdateExpeditionDates(tc.args.ctx, tc.args.client, tc.args.id, tc.args.startDate, tc.args.endDate)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestExpeditionService_DeleteExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockExpeditionRepo, args args)

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
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().DeleteExpedition(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "expedition not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockExpeditionRepo, args args) {
				m.EXPECT().DeleteExpedition(args.ctx, args.client, args.id).
					Return(ErrExpeditionNotFound)
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
			expeditionRepo := mocks.NewMockExpeditionRepo(ctrl)
			tc.mockBehavior(expeditionRepo, tc.args)

			// init service
			s := NewExpeditionService(expeditionRepo)

			// run test
			err := s.DeleteExpedition(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
