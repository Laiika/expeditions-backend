package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgAdminService_GetAllAdmins(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.AdminService
		want    entity.Admins
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
			},
			s:       service.NewAdminService(pgRepo.AdminRepo),
			want:    entity.Admins{},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.GetAllAdmins(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPgAdminService_CreateAdmin(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateAdminInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.AdminService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateAdminInput{
					Name:     "aaa",
					Login:    "ccc",
					Password: "ddd",
				},
			},
			s:       service.NewAdminService(pgRepo.AdminRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				id, err := tc.s.CreateAdmin(tc.args.ctx, tc.args.client, tc.args.input)
				assert.NoError(t, err)

				_, err = tc.s.CreateAdmin(tc.args.ctx, tc.args.client, tc.args.input)
				assert.Error(t, err)

				err = tc.s.DeleteAdmin(tc.args.ctx, tc.args.client, id)
				assert.NoError(t, err)
				return
			}

			id, err := tc.s.CreateAdmin(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteAdmin(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

func TestPgAdminService_DeleteAdmin(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateAdminInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.AdminService
		wantErr bool
	}{
		{
			name: "Simple positive test",
			args: args{
				ctx:    context.Background(),
				client: pgClient,
				input: &entity.CreateAdminInput{
					Name:     "aaa",
					Login:    "aaa",
					Password: "aaa",
				},
			},
			s:       service.NewAdminService(pgRepo.AdminRepo),
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				err := tc.s.DeleteAdmin(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateAdmin(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)

			err = tc.s.DeleteAdmin(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}
