package repo

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/pgdb"
	"time"
)

type AdminRepo interface {
	GetAdminByLogin(ctx context.Context, client any, login string) (*entity.Admin, error)
	GetAllAdmins(ctx context.Context, client any) (entity.Admins, error)
	CreateAdmin(ctx context.Context, client any, leader *entity.Admin) (int, error)
	DeleteAdmin(ctx context.Context, client any, id int) error
}

type LeaderRepo interface {
	GetLeaderByLogin(ctx context.Context, client any, login string) (*entity.Leader, error)
	GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error)
	GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error)
	CreateLeader(ctx context.Context, client any, leader *entity.Leader) (int, error)
	CreateLeaderExpedition(ctx context.Context, client any, leaderId int, expeditionId int) (int, error)
	DeleteLeader(ctx context.Context, client any, id int) error
}

type MemberRepo interface {
	GetMemberByLogin(ctx context.Context, client any, login string) (*entity.Member, error)
	GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error)
	GetAllMembers(ctx context.Context, client any) (entity.Members, error)
	CreateMember(ctx context.Context, client any, member *entity.Member) (int, error)
	CreateMemberExpedition(ctx context.Context, client any, memberId int, expeditionId int) (int, error)
	DeleteMember(ctx context.Context, client any, id int) error
}

type CuratorRepo interface {
	GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error)
	GetAllCurators(ctx context.Context, client any) (entity.Curators, error)
	CreateCurator(ctx context.Context, client any, curator *entity.Curator) (int, error)
	CreateCuratorExpedition(ctx context.Context, client any, curatorId int, expeditionId int) (int, error)
	DeleteCurator(ctx context.Context, client any, id int) error
}

type LocationRepo interface {
	GetAllLocations(ctx context.Context, client any) (entity.Locations, error)
	CreateLocation(ctx context.Context, client any, location *entity.Location) (int, error)
	DeleteLocation(ctx context.Context, client any, id int) error
}

type ExpeditionRepo interface {
	GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error)
	GetLocationExpeditions(ctx context.Context, client any, locationId int) (entity.Expeditions, error)
	CreateExpedition(ctx context.Context, client any, expedition *entity.Expedition) (int, error)
	UpdateExpeditionDates(ctx context.Context, client any, id int, start time.Time, end time.Time) error
	DeleteExpedition(ctx context.Context, client any, id int) error
}

type ArtifactRepo interface {
	GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error)
	GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error)
	CreateArtifact(ctx context.Context, client any, location *entity.Artifact) (int, error)
}

type EquipmentRepo interface {
	GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error)
	GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error)
	CreateEquipment(ctx context.Context, client any, location *entity.Equipment) (int, error)
	DeleteEquipment(ctx context.Context, client any, id int) error
}

type Repositories struct {
	AdminRepo
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
		AdminRepo:      pgdb.NewAdminRepo(),
		LeaderRepo:     pgdb.NewLeaderRepo(),
		MemberRepo:     pgdb.NewMemberRepo(),
		CuratorRepo:    pgdb.NewCuratorRepo(),
		LocationRepo:   pgdb.NewLocationRepo(),
		ExpeditionRepo: pgdb.NewExpeditionRepo(),
		ArtifactRepo:   pgdb.NewArtifactRepo(),
		EquipmentRepo:  pgdb.NewEquipmentRepo(),
	}
}
