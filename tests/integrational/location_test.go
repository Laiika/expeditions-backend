package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgLocationService_GetAllLocations(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LocationService
		want    entity.Locations
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewLocationService(pgRepo.LocationRepo),
			want:    entity.Locations{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllLocations(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgLocationService_CreateLocation(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLocationInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateLocationInput{
					Name:        "aaa",
					Country:     "aaa",
					NearestTown: "aaa",
				},
			},
			s:       service.NewLocationService(pgRepo.LocationRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.s.CreateLocation(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteLocation(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

func TestPgLocationService_DeleteLocation(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLocationInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateLocationInput{
					Name:        "aaa",
					Country:     "aaa",
					NearestTown: "aaa",
				},
			},
			s:       service.NewLocationService(pgRepo.LocationRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteLocation(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateLocation(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteLocation(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}
