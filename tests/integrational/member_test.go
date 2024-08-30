package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgMemberService_GetExpeditionMembers(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
		want    entity.Members
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:          context.Background(),
				client:       pgClient,
				expeditionId: 100,
			},
			s:       service.NewMemberService(pgRepo.MemberRepo),
			want:    entity.Members{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetExpeditionMembers(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgMemberService_GetAllMembers(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
		want    entity.Members
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewMemberService(pgRepo.MemberRepo),
			want:    entity.Members{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllMembers(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgMemberService_CreateMember(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateMemberInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateMemberInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Role:        "aaa",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			s:       service.NewMemberService(pgRepo.MemberRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				id, err := tc.s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
				assert.NoError(t, err)

				_, err = tc.s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
				assert.Error(t, err)

				err = tc.s.DeleteMember(tc.args.ctx, tc.args.client, id)
				assert.NoError(t, err)
				return
			}

			id, err := tc.s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteMember(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

func TestPgCuratorService_CreateMemberExpedition(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
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
			s:       service.NewMemberService(pgRepo.MemberRepo),
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

			id, err := tc.s.CreateMember(tc.args.ctx, tc.args.client, &entity.CreateMemberInput{
				Name:        "aaa",
				PhoneNumber: "bbb",
				Role:        "aaa",
				Login:       "ccc",
				Password:    "ddd",
			})
			assert.NoError(t, err)

			_, err = tc.s.CreateMemberExpedition(tc.args.ctx, tc.args.client, id, expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			err = tc.s.DeleteMember(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)

			err = tc.ls.DeleteLocation(tc.args.ctx, tc.args.client, locationId)
			assert.NoError(t, err)
		})
	}
}

func TestPgMemberService_DeleteMember(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateMemberInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateMemberInput{
					Name:        "aaa",
					PhoneNumber: "aaa",
					Role:        "aaa",
					Login:       "aaa",
					Password:    "aaa",
				},
			},
			s:       service.NewMemberService(pgRepo.MemberRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteMember(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteMember(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}
