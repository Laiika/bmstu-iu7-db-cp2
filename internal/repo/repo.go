package repo

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/pgdb"
	"time"
)

type LeaderRepo interface {
	GetLeaderById(ctx context.Context, client any, id int) (*entity.Leader, error)
	GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error)
	GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error)
	CreateLeader(ctx context.Context, client any, leader *entity.Leader) (int, error)
	DeleteLeader(ctx context.Context, client any, id int) error
}

type MemberRepo interface {
	GetMemberById(ctx context.Context, client any, id int) (*entity.Member, error)
	GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error)
	GetAllMembers(ctx context.Context, client any) (entity.Members, error)
	CreateMember(ctx context.Context, client any, member *entity.Member) (int, error)
	DeleteMember(ctx context.Context, client any, id int) error
}

type CuratorRepo interface {
	GetCuratorById(ctx context.Context, client any, id int) (*entity.Curator, error)
	GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error)
	GetAllCurators(ctx context.Context, client any) (entity.Curators, error)
	CreateCurator(ctx context.Context, client any, curator *entity.Curator) (int, error)
	DeleteCurator(ctx context.Context, client any, id int) error
}

type LocationRepo interface {
	GetLocationById(ctx context.Context, client any, id int) (*entity.Location, error)
	GetAllLocations(ctx context.Context, client any) (entity.Locations, error)
	CreateLocation(ctx context.Context, client any, location *entity.Location) (int, error)
	DeleteLocation(ctx context.Context, client any, id int) error
}

type ExpeditionRepo interface {
	GetExpeditionById(ctx context.Context, client any, id int) (*entity.Expedition, error)
	GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error)
	CreateExpedition(ctx context.Context, client any, expedition *entity.Expedition) (int, error)
	UpdateExpeditionDates(ctx context.Context, client any, id int, start time.Time, end time.Time) error
	DeleteExpedition(ctx context.Context, client any, id int) error
}

type ArtifactRepo interface {
	GetArtifactById(ctx context.Context, client any, id int) (*entity.Artifact, error)
	GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error)
	GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error)
	CreateArtifact(ctx context.Context, client any, location *entity.Artifact) (int, error)
}

type EquipmentRepo interface {
	GetEquipmentById(ctx context.Context, client any, id int) (*entity.Equipment, error)
	GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error)
	GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error)
	CreateEquipment(ctx context.Context, client any, location *entity.Equipment) (int, error)
	DeleteEquipment(ctx context.Context, client any, id int) error
}

type Repositories struct {
	LeaderRepo
	MemberRepo
	CuratorRepo
	LocationRepo
	ExpeditionRepo
	ArtifactRepo
	EquipmentRepo
}

func NewRepositories() *Repositories {
	return &Repositories{
		LeaderRepo:     pgdb.NewLeaderRepo(),
		MemberRepo:     pgdb.NewMemberRepo(),
		CuratorRepo:    pgdb.NewCuratorRepo(),
		LocationRepo:   pgdb.NewLocationRepo(),
		ExpeditionRepo: pgdb.NewExpeditionRepo(),
		ArtifactRepo:   pgdb.NewArtifactRepo(),
		EquipmentRepo:  pgdb.NewEquipmentRepo(),
	}
}
