package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgCuratorService_GetCuratorById(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateCuratorInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.CuratorService
		want    *entity.Curator
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
			s: service.NewCuratorService(pgRepo.CuratorRepo),
			want: &entity.Curator{
				Name: "aaa",
			},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr {
				_, err := tc.s.GetCuratorById(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateCurator(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)
			tc.want.Id = id

			got, err := tc.s.GetCuratorById(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			err = tc.s.DeleteCurator(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

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
