package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/service/auth"
)

type Auth interface {
	GetSession(token string) bool
	GetClient(token string) (any, error)
}

type Leader interface {
	GetLeaderById(ctx context.Context, client any, id int) (*entity.Leader, error)
	GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error)
	GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error)
	CreateLeader(ctx context.Context, client any, input *entity.CreateLeaderInput) (int, error)
	DeleteLeader(ctx context.Context, client any, id int) error
}

type Member interface {
	GetMemberById(ctx context.Context, client any, id int) (*entity.Member, error)
	GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error)
	GetAllMembers(ctx context.Context, client any) (entity.Members, error)
	CreateMember(ctx context.Context, client any, input *entity.CreateMemberInput) (int, error)
	DeleteMember(ctx context.Context, client any, id int) error
}

type Curator interface {
	GetCuratorById(ctx context.Context, client any, id int) (*entity.Curator, error)
	GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error)
	GetAllCurators(ctx context.Context, client any) (entity.Curators, error)
	CreateCurator(ctx context.Context, client any, input *entity.CreateCuratorInput) (int, error)
	DeleteCurator(ctx context.Context, client any, id int) error
}

type Location interface {
	GetLocationById(ctx context.Context, client any, id int) (*entity.Location, error)
	GetAllLocations(ctx context.Context, client any) (entity.Locations, error)
	CreateLocation(ctx context.Context, client any, input *entity.CreateLocationInput) (int, error)
	DeleteLocation(ctx context.Context, client any, id int) error
}

type Expedition interface {
	GetExpeditionById(ctx context.Context, client any, id int) (*entity.Expedition, error)
	GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error)
	CreateExpedition(ctx context.Context, client any, input *entity.CreateExpeditionInput) (int, error)
	UpdateExpeditionDates(ctx context.Context, client any, id int, startDate string, endDate string) error
	DeleteExpedition(ctx context.Context, client any, id int) error
}

type Artifact interface {
	GetArtifactById(ctx context.Context, client any, id int) (*entity.Artifact, error)
	GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error)
	GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error)
	CreateArtifact(ctx context.Context, client any, input *entity.CreateArtifactInput) (int, error)
}

type Equipment interface {
	GetEquipmentById(ctx context.Context, client any, id int) (*entity.Equipment, error)
	GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error)
	GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error)
	CreateEquipment(ctx context.Context, client any, input *entity.CreateEquipmentInput) (int, error)
	DeleteEquipment(ctx context.Context, client any, id int) error
}

type Services struct {
	Auth       Auth
	Leader     Leader
	Member     Member
	Curator    Curator
	Location   Location
	Expedition Expedition
	Artifact   Artifact
	Equipment  Equipment
}

func NewServices(repos *repo.Repositories, admin any, leader any, member any) *Services {
	return &Services{
		Auth:       auth.NewAuthService(member, leader, admin),
		Leader:     NewLeaderService(repos.LeaderRepo),
		Member:     NewMemberService(repos.MemberRepo),
		Curator:    NewCuratorService(repos.CuratorRepo),
		Location:   NewLocationService(repos.LocationRepo),
		Expedition: NewExpeditionService(repos.ExpeditionRepo),
		Artifact:   NewArtifactService(repos.ArtifactRepo),
		Equipment:  NewEquipmentService(repos.EquipmentRepo),
	}
}
