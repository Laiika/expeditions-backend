package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocationService_GetAllLocations(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockLocationRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Locations
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockLocationRepo, args args) {
				m.EXPECT().GetAllLocations(args.ctx, args.client).
					Return(entity.Locations{
						&entity.Location{
							Id:          1,
							Name:        "aaa",
							Country:     "bbb",
							NearestTown: "ccc",
						},
					}, nil)
			},
			want: entity.Locations{
				&entity.Location{
					Id:          1,
					Name:        "aaa",
					Country:     "bbb",
					NearestTown: "ccc",
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
			locationRepo := mocks.NewMockLocationRepo(ctrl)
			tc.mockBehavior(locationRepo, tc.args)

			// init service
			s := NewLocationService(locationRepo)

			// run test
			got, err := s.GetAllLocations(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLocationService_CreateLocation(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLocationInput
	}

	type MockBehavior func(m *mocks.MockLocationRepo, args args)

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
				input: &entity.CreateLocationInput{
					Name:        "aaa",
					Country:     "bbb",
					NearestTown: "ccc",
				},
			},
			mockBehavior: func(m *mocks.MockLocationRepo, args args) {
				m.EXPECT().CreateLocation(args.ctx, args.client, &entity.Location{
					Name:        args.input.Name,
					Country:     args.input.Country,
					NearestTown: args.input.NearestTown,
				}).
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
			locationRepo := mocks.NewMockLocationRepo(ctrl)
			tc.mockBehavior(locationRepo, tc.args)

			// init service
			s := NewLocationService(locationRepo)

			// run test
			got, err := s.CreateLocation(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLocationService_DeleteLocation(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockLocationRepo, args args)

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
			mockBehavior: func(m *mocks.MockLocationRepo, args args) {
				m.EXPECT().DeleteLocation(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "location not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockLocationRepo, args args) {
				m.EXPECT().DeleteLocation(args.ctx, args.client, args.id).
					Return(ErrLocationNotFound)
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
			locationRepo := mocks.NewMockLocationRepo(ctrl)
			tc.mockBehavior(locationRepo, tc.args)

			// init service
			s := NewLocationService(locationRepo)

			// run test
			err := s.DeleteLocation(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
