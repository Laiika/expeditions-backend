package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgLeaderService_GetExpeditionLeaders(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
		want    entity.Leaders
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:          context.Background(),
				client:       pgClient,
				expeditionId: 100,
			},
			s:       service.NewLeaderService(pgRepo.LeaderRepo),
			want:    entity.Leaders{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetExpeditionLeaders(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgLeaderService_GetAllLeaders(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
		want    entity.Leaders
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewLeaderService(pgRepo.LeaderRepo),
			want:    entity.Leaders{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllLeaders(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgLeaderService_CreateLeader(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLeaderInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateLeaderInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			s:       service.NewLeaderService(pgRepo.LeaderRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				id, err := tc.s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
				assert.NoError(t, err)

				_, err = tc.s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
				assert.Error(t, err)

				err = tc.s.DeleteLeader(tc.args.ctx, tc.args.client, id)
				assert.NoError(t, err)
				return
			}

			id, err := tc.s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteLeader(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

func TestPgCuratorService_CreateLeaderExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
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
			s:       service.NewLeaderService(pgRepo.LeaderRepo),
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

			id, err := tc.s.CreateLeader(tc.args.ctx, tc.args.client, &entity.CreateLeaderInput{
				Name:        "aaa",
				PhoneNumber: "bbb",
				Login:       "ccc",
				Password:    "ddd",
			})
			assert.NoError(t, err)

			_, err = tc.s.CreateLeaderExpedition(tc.args.ctx, tc.args.client, id, expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			err = tc.s.DeleteLeader(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgLeaderService_DeleteLeader(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLeaderInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateLeaderInput{
					Name:        "aaa",
					PhoneNumber: "aaa",
					Login:       "aaa",
					Password:    "aaa",
				},
			},
			s:       service.NewLeaderService(pgRepo.LeaderRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteLeader(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteLeader(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}
