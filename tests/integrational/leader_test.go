package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgLeaderService_GetLeaderById(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateLeaderInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.LeaderService
		want    *entity.Leader
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
			s: service.NewLeaderService(pgRepo.LeaderRepo),
			want: &entity.Leader{
				Name:        "aaa",
				PhoneNumber: "aaa",
				Login:       "aaa",
				Password:    "aaa",
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				_, err := tc.s.GetLeaderById(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateLeader(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)
			tc.want.Id = id

			got, err := tc.s.GetLeaderById(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			err = tc.s.DeleteLeader(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

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
