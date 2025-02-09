func TestLeaderService_DeleteLeader(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		id     int
	}

	type MockBehavior func(m *mocks.MockLeaderRepo, args args)

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
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().DeleteLeader(args.ctx, args.client, args.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "leader not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     100,
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().DeleteLeader(args.ctx, args.client, args.id).
					Return(ErrLeaderNotFound)
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
			leaderRepo := mocks.NewMockLeaderRepo(ctrl)
			tc.mockBehavior(leaderRepo, tc.args)

			// init service
			s := NewLeaderService(leaderRepo)

			// run test
			err := s.DeleteLeader(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
