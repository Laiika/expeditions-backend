package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquipmentService_GetExpeditionEquipments(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockEquipmentRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Equipments
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:          context.Background(),
				client:       nil,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockEquipmentRepo, args args) {
				m.EXPECT().GetExpeditionEquipments(args.ctx, args.client, args.expeditionId).
					Return(entity.Equipments{
						&entity.Equipment{
							Id:           1,
							ExpeditionId: 1,
							Name:         "aaa",
							Amount:       10,
						},
					}, nil)
			},
			want: entity.Equipments{
				&entity.Equipment{
					Id:           1,
					ExpeditionId: 1,
					Name:         "aaa",
					Amount:       10,
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
			equipmentRepo := mocks.NewMockEquipmentRepo(ctrl)
			tc.mockBehavior(equipmentRepo, tc.args)

			// init service
			s := NewEquipmentService(equipmentRepo)

			// run test
			got, err := s.GetExpeditionEquipments(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEquipmentService_GetAllEquipments(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockEquipmentRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Equipments
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockEquipmentRepo, args args) {
				m.EXPECT().GetAllEquipments(args.ctx, args.client).
					Return(entity.Equipments{
						&entity.Equipment{
							Id:           1,
							ExpeditionId: 1,
							Name:         "aaa",
							Amount:       10,
						},
					}, nil)
			},
			want: entity.Equipments{
				&entity.Equipment{
					Id:           1,
					ExpeditionId: 1,
					Name:         "aaa",
					Amount:       10,
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
			equipmentRepo := mocks.NewMockEquipmentRepo(ctrl)
			tc.mockBehavior(equipmentRepo, tc.args)

			// init service
			s := NewEquipmentService(equipmentRepo)

			// run test
			got, err := s.GetAllEquipments(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEquipmentService_CreateEquipment(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateEquipmentInput
	}

	type MockBehavior func(m *mocks.MockEquipmentRepo, args args)

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
				input: &entity.CreateEquipmentInput{
					ExpeditionId: 1,
					Name:         "aaa",
					Amount:       10,
				},
			},
			mockBehavior: func(m *mocks.MockEquipmentRepo, args args) {
				m.EXPECT().CreateEquipment(args.ctx, args.client, &entity.Equipment{
					ExpeditionId: args.input.ExpeditionId,
					Name:         args.input.Name,
					Amount:       args.input.Amount,
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
			equipmentRepo := mocks.NewMockEquipmentRepo(ctrl)
			tc.mockBehavior(equipmentRepo, tc.args)

			// init service
			s := NewEquipmentService(equipmentRepo)

			// run test
			got, err := s.CreateEquipment(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEquipmentService_DeleteEquipment(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockEquipmentRepo, args args)

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
			mockBehavior: func(m *mocks.MockEquipmentRepo, args args) {
				m.EXPECT().DeleteEquipment(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "equipment not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockEquipmentRepo, args args) {
				m.EXPECT().DeleteEquipment(args.ctx, args.client, args.id).
					Return(ErrEquipmentNotFound)
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
			equipmentRepo := mocks.NewMockEquipmentRepo(ctrl)
			tc.mockBehavior(equipmentRepo, tc.args)

			// init service
			s := NewEquipmentService(equipmentRepo)

			// run test
			err := s.DeleteEquipment(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
