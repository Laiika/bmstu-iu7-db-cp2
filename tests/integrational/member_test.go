package integrational

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgMemberService_GetMemberById(t *testing.T) {
	type args struct {
		ctx    context.Context
		client any
		input  *entity.CreateMemberInput
	}

	testCases := []struct {
		name    string
		args    args
		s       *service.MemberService
		want    *entity.Member
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
					Login:       "aaa",
					Password:    "aaa",
				},
			},
			s: service.NewMemberService(pgRepo.MemberRepo),
			want: &entity.Member{
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
				_, err := tc.s.GetMemberById(tc.args.ctx, tc.args.client, 1)
				assert.Error(t, err)
				return
			}

			id, err := tc.s.CreateMember(tc.args.ctx, tc.args.client, tc.args.input)
			assert.NoError(t, err)
			tc.want.Id = id

			got, err := tc.s.GetMemberById(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			err = tc.s.DeleteMember(tc.args.ctx, tc.args.client, id)
			assert.NoError(t, err)
		})
	}
}

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
