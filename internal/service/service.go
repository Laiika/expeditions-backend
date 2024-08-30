package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
)

type Auth interface {
	Login(ctx context.Context, data *entity.LoginInput) (string, error)
	Logout(token string) error
	GetSession(token string) bool
	GetClient(token string) (any, error)
}

type Admin interface {
	GetAllAdmins(ctx context.Context, client any) (entity.Admins, error)
	CreateAdmin(ctx context.Context, client any, input *entity.CreateAdminInput) (int, error)
	DeleteAdmin(ctx context.Context, client any, id int) error
}

type Leader interface {
	GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error)
	GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error)
	CreateLeader(ctx context.Context, client any, input *entity.CreateLeaderInput) (int, error)
	CreateLeaderExpedition(ctx context.Context, client any, leaderId int, expeditionId int) (int, error)
	DeleteLeader(ctx context.Context, client any, id int) error
}

type Member interface {
	GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error)
	GetAllMembers(ctx context.Context, client any) (entity.Members, error)
	CreateMember(ctx context.Context, client any, input *entity.CreateMemberInput) (int, error)
	CreateMemberExpedition(ctx context.Context, client any, memberId int, expeditionId int) (int, error)
	DeleteMember(ctx context.Context, client any, id int) error
}

type Curator interface {
	GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error)
	GetAllCurators(ctx context.Context, client any) (entity.Curators, error)
	CreateCurator(ctx context.Context, client any, input *entity.CreateCuratorInput) (int, error)
	CreateCuratorExpedition(ctx context.Context, client any, curatorId int, expeditionId int) (int, error)
	DeleteCurator(ctx context.Context, client any, id int) error
}

type Location interface {
	GetAllLocations(ctx context.Context, client any) (entity.Locations, error)
	CreateLocation(ctx context.Context, client any, input *entity.CreateLocationInput) (int, error)
	DeleteLocation(ctx context.Context, client any, id int) error
}

type Expedition interface {
	GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error)
	GetLocationExpeditions(ctx context.Context, client any, locationId int) (entity.Expeditions, error)
	CreateExpedition(ctx context.Context, client any, input *entity.CreateExpeditionInput) (int, error)
	UpdateExpeditionDates(ctx context.Context, client any, id int, startDate string, endDate string) error
	DeleteExpedition(ctx context.Context, client any, id int) error
}

type Artifact interface {
	GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error)
	GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error)
	CreateArtifact(ctx context.Context, client any, input *entity.CreateArtifactInput) (int, error)
}

type Equipment interface {
	GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error)
	GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error)
	CreateEquipment(ctx context.Context, client any, input *entity.CreateEquipmentInput) (int, error)
	DeleteEquipment(ctx context.Context, client any, id int) error
}

type Services struct {
	Auth       Auth
	Admin      Admin
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
		Auth:       NewAuthService(member, leader, admin, repos.MemberRepo, repos.LeaderRepo, repos.AdminRepo),
		Admin:      NewAdminService(repos.AdminRepo),
		Leader:     NewLeaderService(repos.LeaderRepo),
		Member:     NewMemberService(repos.MemberRepo),
		Curator:    NewCuratorService(repos.CuratorRepo),
		Location:   NewLocationService(repos.LocationRepo),
		Expedition: NewExpeditionService(repos.ExpeditionRepo),
		Artifact:   NewArtifactService(repos.ArtifactRepo),
		Equipment:  NewEquipmentService(repos.EquipmentRepo),
	}
}
