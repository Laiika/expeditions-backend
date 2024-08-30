package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgEquipmentService_GetExpeditionEquipments(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.EquipmentService
		want    entity.Equipments
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:          context.Background(),
				client:       pgClient,
				expeditionId: 100,
			},
			s:       service.NewEquipmentService(pgRepo.EquipmentRepo),
			want:    entity.Equipments{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetExpeditionEquipments(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgEquipmentService_GetAllEquipments(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.EquipmentService
		want    entity.Equipments
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewEquipmentService(pgRepo.EquipmentRepo),
			want:    entity.Equipments{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllEquipments(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgEquipmentService_CreateEquipment(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateEquipmentInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.EquipmentService
		ls      *service.LocationService
		es      *service.ExpeditionService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateEquipmentInput{
					Name:   "aaa",
					Amount: 10000,
				},
			},
			s:       service.NewEquipmentService(pgRepo.EquipmentRepo),
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
			tc.args.input.ExpeditionId = expeditionId

			_, err = tc.s.CreateEquipment(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgMemberService_DeleteEquipment(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateEquipmentInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.EquipmentService
		ls      *service.LocationService
		es      *service.ExpeditionService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateEquipmentInput{
					Name:   "aaa",
					Amount: 10000,
				},
			},
			s:       service.NewEquipmentService(pgRepo.EquipmentRepo),
			ls:      service.NewLocationService(pgRepo.LocationRepo),
			es:      service.NewExpeditionService(pgRepo.ExpeditionRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteEquipment(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

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
			tc.args.input.ExpeditionId = expeditionId

			id, err := tc.s.CreateEquipment(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteEquipment(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}
