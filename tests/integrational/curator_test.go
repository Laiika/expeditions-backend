package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgCuratorService_GetExpeditionCurators(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		want    entity.Curators
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:          context.Background(),
				client:       pgClient,
				expeditionId: 100,
			},
			s:       service.NewCuratorService(pgRepo.CuratorRepo),
			want:    entity.Curators{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetExpeditionCurators(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgCuratorService_GetAllCurators(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		want    entity.Curators
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewCuratorService(pgRepo.CuratorRepo),
			want:    entity.Curators{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllCurators(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgCuratorService_CreateCurator(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateCuratorInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateCuratorInput{
					Name: "aaa",
				},
			},
			s:       service.NewCuratorService(pgRepo.CuratorRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				id, err := tc.s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
				assert.NoError(t, err)

				_, err = tc.s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
				assert.Error(t, err)

				err = tc.s.DeleteCurator(tc.args.ctx, tc.args.client, id)
				assert.NoError(t, err)
				return
			}

			id, err := tc.s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteCurator(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

func TestPgCuratorService_CreateCuratorExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		ls      *service.LocationService
		es      *service.ExpeditionService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewCuratorService(pgRepo.CuratorRepo),
			ls:      service.NewLocationService(pgRepo.LocationRepo),
			es:      service.NewExpeditionService(pgRepo.ExpeditionRepo),
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

			expeditionId, err := tc.es.CreateExpedition(tc.args.ctx, tc.args.client, &entity.CreateExpeditionInput{
				LocationId: locationId,
				StartDate:  "2024-07-01",
				EndDate:    "2024-08-01",
			})
			assert.NoError(t, err)

			id, err := tc.s.CreateCurator(tc.args.ctx, tc.args.client, &entity.CreateCuratorInput{
				Name: "aaa",
			})
			assert.NoError(t, err)

			_, err = tc.s.CreateCuratorExpedition(tc.args.ctx, tc.args.client, id, expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			err = tc.s.DeleteCurator(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgCuratorService_DeleteCurator(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateCuratorInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateCuratorInput{
					Name: "aaa",
				},
			},
			s:       service.NewCuratorService(pgRepo.CuratorRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteCurator(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteCurator(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}
