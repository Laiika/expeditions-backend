package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/repoerrs"
	"db_cp_6/pkg/postgres"
	"fmt"
)

type EquipmentRepo struct {
}

func NewEquipmentRepo() *EquipmentRepo {
	return &EquipmentRepo{}
}

func (r *EquipmentRepo) GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, expedition_id, name, amount
		FROM equipments
		WHERE expedition_id = $1
	`
	rows, err := pgClient.Query(ctx, q, expeditionId)
	if err != nil {
		return nil, fmt.Errorf("EquipmentRepo GetExpeditionEquipments: %v", err)
	}

	equipments := make(entity.Equipments, 0)
	for rows.Next() {
		var eq entity.Equipment

		err = rows.Scan(&eq.Id, &eq.ExpeditionId, &eq.Name, &eq.Amount)
		if err != nil {
			return nil, fmt.Errorf("EquipmentRepo GetExpeditionEquipments: %v", err)
		}

		equipments = append(equipments, &eq)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("EquipmentRepo GetExpeditionEquipments: %v", err)
	}

	return equipments, nil
}

func (r *EquipmentRepo) GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, expedition_id, name, amount
		FROM equipments
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("EquipmentRepo GetAllEquipments: %v", err)
	}

	equipments := make(entity.Equipments, 0)
	for rows.Next() {
		var eq entity.Equipment

		err = rows.Scan(&eq.Id, &eq.ExpeditionId, &eq.Name, &eq.Amount)
		if err != nil {
			return nil, fmt.Errorf("EquipmentRepo GetAllEquipments: %v", err)
		}

		equipments = append(equipments, &eq)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("EquipmentRepo GetAllEquipments: %v", err)
	}

	return equipments, nil
}

func (r *EquipmentRepo) CreateEquipment(ctx context.Context, client any, equipment *entity.Equipment) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO equipments
		    (expedition_id, name, amount) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, equipment.ExpeditionId, equipment.Name, equipment.Amount).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("EquipmentRepo CreateEquipment: %v", err)
	}

	return id, nil
}

func (r *EquipmentRepo) DeleteEquipment(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM equipments
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("EquipmentRepo DeleteEquipment: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
