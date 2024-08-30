package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArtifactService_GetLocationArtifacts(t *testing.T) {
	type args struct {
		ctx        context.Context
		client     any
		locationId int
	}

	type MockBehavior func(m *mocks.MockArtifactRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Artifacts
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:        context.Background(),
				client:     nil,
				locationId: 1,
			},
			mockBehavior: func(m *mocks.MockArtifactRepo, args args) {
				m.EXPECT().GetLocationArtifacts(args.ctx, args.client, args.locationId).
					Return(entity.Artifacts{
						&entity.Artifact{
							Id:         1,
							LocationId: 1,
							Name:       "aaa",
							Age:        10000,
						},
					}, nil)
			},
			want: entity.Artifacts{
				&entity.Artifact{
					Id:         1,
					LocationId: 1,
					Name:       "aaa",
					Age:        10000,
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
			artifactRepo := mocks.NewMockArtifactRepo(ctrl)
			tc.mockBehavior(artifactRepo, tc.args)

			// init service
			s := NewArtifactService(artifactRepo)

			// run test
			got, err := s.GetLocationArtifacts(tc.args.ctx, tc.args.client, tc.args.locationId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestArtifactService_GetAllArtifacts(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockArtifactRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Artifacts
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockArtifactRepo, args args) {
				m.EXPECT().GetAllArtifacts(args.ctx, args.client).
					Return(entity.Artifacts{
						&entity.Artifact{
							Id:         1,
							LocationId: 1,
							Name:       "aaa",
							Age:        10000,
						},
					}, nil)
			},
			want: entity.Artifacts{
				&entity.Artifact{
					Id:         1,
					LocationId: 1,
					Name:       "aaa",
					Age:        10000,
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
			artifactRepo := mocks.NewMockArtifactRepo(ctrl)
			tc.mockBehavior(artifactRepo, tc.args)

			// init service
			s := NewArtifactService(artifactRepo)

			// run test
			got, err := s.GetAllArtifacts(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestArtifactService_CreateArtifact(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateArtifactInput
	}

	type MockBehavior func(m *mocks.MockArtifactRepo, args args)

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
				input: &entity.CreateArtifactInput{
					LocationId: 1,
					Name:       "aaa",
					Age:        10000,
				},
			},
			mockBehavior: func(m *mocks.MockArtifactRepo, args args) {
				m.EXPECT().CreateArtifact(args.ctx, args.client, &entity.Artifact{
					LocationId: args.input.LocationId,
					Name:       args.input.Name,
					Age:        args.input.Age,
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
			artifactRepo := mocks.NewMockArtifactRepo(ctrl)
			tc.mockBehavior(artifactRepo, tc.args)

			// init service
			s := NewArtifactService(artifactRepo)

			// run test
			got, err := s.CreateArtifact(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
