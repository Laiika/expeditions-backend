package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgArtifactService_GetLocationArtifacts(t *testing.T) {
	type args struct {
		ctx        context.Context
		client     any
		locationId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ArtifactService
		want    entity.Artifacts
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:        context.Background(),
				client:     pgClient,
				locationId: 100,
			},
			s:       service.NewArtifactService(pgRepo.ArtifactRepo),
			want:    entity.Artifacts{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetLocationArtifacts(tc.args.ctx, tc.args.client, tc.args.locationId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgArtifactService_GetAllArtifacts(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ArtifactService
		want    entity.Artifacts
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewArtifactService(pgRepo.ArtifactRepo),
			want:    entity.Artifacts{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllArtifacts(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgArtifactService_CreateArtifact(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateArtifactInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ArtifactService
		ls      *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateArtifactInput{
					Name: "aaa",
					Age:  10000,
				},
			},
			s:       service.NewArtifactService(pgRepo.ArtifactRepo),
			ls:      service.NewLocationService(pgRepo.LocationRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			locationId, err := tc.ls.CreateLocation(tc.args.ctx, tc.args.client, &entity.CreateLocationInput{
				Name:        "aaa",
				Country:     "aaa",
				NearestTown: "aaa",
			})
			assert.NoError(t, err)
			tc.args.input.LocationId = locationId

			_, err = tc.s.CreateArtifact(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}
