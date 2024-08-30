package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgExpeditionService_GetAllExpeditions(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ExpeditionService
		want    entity.Expeditions
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewExpeditionService(pgRepo.ExpeditionRepo),
			want:    entity.Expeditions{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllExpeditions(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgExpeditionService_GetLocationExpeditions(t *testing.T) {
	type args struct {
		ctx        context.Context
		client     any
		locationId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ExpeditionService
		want    entity.Expeditions
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:        context.Background(),
				client:     pgClient,
				locationId: 1,
			},
			s:       service.NewExpeditionService(pgRepo.ExpeditionRepo),
			want:    entity.Expeditions{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetLocationExpeditions(tc.args.ctx, tc.args.client, tc.args.locationId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgExpeditionService_CreateExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateExpeditionInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ExpeditionService
		ls      *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateExpeditionInput{
					StartDate: "2024-07-01",
					EndDate:   "2024-08-01",
				},
			},
			s:       service.NewExpeditionService(pgRepo.ExpeditionRepo),
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

			_, err = tc.s.CreateExpedition(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgMemberService_UpdateExpeditionDates(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateExpeditionInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ExpeditionService
		ls      *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateExpeditionInput{
					StartDate: "2024-07-01",
					EndDate:   "2024-08-01",
				},
			},
			s:       service.NewExpeditionService(pgRepo.ExpeditionRepo),
			ls:      service.NewLocationService(pgRepo.LocationRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.UpdateExpeditionDates(tc.args.ctx, tc.args.client, 1, "2024-08-01", "2024-09-01")
				assert.Error(t, err)
				return
			}

			locationId, err := tc.ls.CreateLocation(tc.args.ctx, tc.args.client, &entity.CreateLocationInput{
				Name:        "aaa",
				Country:     "aaa",
				NearestTown: "aaa",
			})
			assert.NoError(t, err)
			tc.args.input.LocationId = locationId

			id, err := tc.s.CreateExpedition(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.UpdateExpeditionDates(tc.args.ctx, tc.args.client, id, "2024-08-01", "2024-09-01")
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgMemberService_DeleteExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateExpeditionInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.ExpeditionService
		ls      *service.LocationService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateExpeditionInput{
					StartDate: "2024-07-01",
					EndDate:   "2024-08-01",
				},
			},
			s:       service.NewExpeditionService(pgRepo.ExpeditionRepo),
			ls:      service.NewLocationService(pgRepo.LocationRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteExpedition(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			locationId, err := tc.ls.CreateLocation(tc.args.ctx, tc.args.client, &entity.CreateLocationInput{
				Name:        "aaa",
				Country:     "aaa",
				NearestTown: "aaa",
			})
			assert.NoError(t, err)
			tc.args.input.LocationId = locationId

			id, err := tc.s.CreateExpedition(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteExpedition(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}
