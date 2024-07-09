package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestLeaderService_GetLeaderById(t *testing.T) {
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
		want         *entity.Leader
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
				m.EXPECT().GetLeaderById(args.ctx, args.client, args.id).
					Return(&entity.Leader{
						Id:          1,
						Name:        "aaa",
						PhoneNumber: "+79021061232",
						Login:       "dhhjds",
						Password:    "jdskjdsjk",
					}, nil)
			},
			want: &entity.Leader{
				Id:          1,
				Name:        "aaa",
				PhoneNumber: "+79021061232",
				Login:       "dhhjds",
				Password:    "jdskjdsjk",
			},
			wantErr: false,
		},
		{
			name: "leader not found error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				id:     1,
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().GetLeaderById(args.ctx, args.client, args.id).
					Return(nil, ErrLeaderNotFound)
			},
			want:    nil,
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
			got, err := s.GetLeaderById(tc.args.ctx, tc.args.client, tc.args.id)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLeaderService_GetExpeditionLeaders(t *testing.T) {
	type args struct {
		ctx          context.Context
		client       any
		expeditionId int
	}

	type MockBehavior func(m *mocks.MockLeaderRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Leaders
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:          context.Background(),
				client:       nil,
				expeditionId: 1,
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().GetExpeditionLeaders(args.ctx, args.client, args.expeditionId).
					Return(entity.Leaders{
						&entity.Leader{
							Id:          1,
							Name:        "aaa",
							PhoneNumber: "+79021061232",
							Login:       "dhhjds",
							Password:    "jdskjdsjk",
						},
					}, nil)
			},
			want: entity.Leaders{
				&entity.Leader{
					Id:          1,
					Name:        "aaa",
					PhoneNumber: "+79021061232",
					Login:       "dhhjds",
					Password:    "jdskjdsjk",
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
			leaderRepo := mocks.NewMockLeaderRepo(ctrl)
			tc.mockBehavior(leaderRepo, tc.args)

			// init service
			s := NewLeaderService(leaderRepo)

			// run test
			got, err := s.GetExpeditionLeaders(tc.args.ctx, tc.args.client, tc.args.expeditionId)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLeaderService_GetAllLeaders(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
	}

	type MockBehavior func(m *mocks.MockLeaderRepo, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         entity.Leaders
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:    context.Background(),
				client: nil,
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().GetAllLeaders(args.ctx, args.client).
					Return(entity.Leaders{
						&entity.Leader{
							Id:          1,
							Name:        "aaa",
							PhoneNumber: "+79021061232",
							Login:       "dhhjds",
							Password:    "jdskjdsjk",
						},
					}, nil)
			},
			want: entity.Leaders{
				&entity.Leader{
					Id:          1,
					Name:        "aaa",
					PhoneNumber: "+79021061232",
					Login:       "dhhjds",
					Password:    "jdskjdsjk",
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
			leaderRepo := mocks.NewMockLeaderRepo(ctrl)
			tc.mockBehavior(leaderRepo, tc.args)

			// init service
			s := NewLeaderService(leaderRepo)

			// run test
			got, err := s.GetAllLeaders(tc.args.ctx, tc.args.client)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestLeaderService_CreateLeader(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLeaderInput
	}

	type MockBehavior func(m *mocks.MockLeaderRepo, args args)

	bytes, _ := bcrypt.GenerateFromPassword([]byte("ddd"), 14)

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
				input: &entity.CreateLeaderInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().CreateLeader(args.ctx, args.client, &entity.Leader{
					Name:        args.input.Name,
					PhoneNumber: args.input.PhoneNumber,
					Login:       args.input.Login,
					Password:    string(bytes),
				}).
					Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "leader already exists error",
			args: args{
				ctx:    context.Background(),
				client: nil,
				input: &entity.CreateLeaderInput{
					Name:        "aaa",
					PhoneNumber: "bbb",
					Login:       "ccc",
					Password:    "ddd",
				},
			},
			mockBehavior: func(m *mocks.MockLeaderRepo, args args) {
				m.EXPECT().CreateLeader(args.ctx, args.client, &entity.Leader{
					Name:        args.input.Name,
					PhoneNumber: args.input.PhoneNumber,
					Login:       args.input.Login,
					Password:    string(bytes),
				}).
					Return(nil, ErrLeaderAlreadyExists)
			},
			want:    0,
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
			got, err := s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

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
		want         error
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
			want:    nil,
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
			want:    ErrLeaderNotFound,
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
			assert.Equal(t, tc.want, err)
		})
	}
}
