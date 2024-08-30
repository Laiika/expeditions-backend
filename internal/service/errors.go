package service

import "errors"

var (
	ErrSessionNotExists = errors.New("session not exists")

	ErrAdminAlreadyExists = errors.New("admin already exists")
	ErrAdminNotFound      = errors.New("admin not found")

	ErrLeaderAlreadyExists = errors.New("leader already exists")
	ErrLeaderNotFound      = errors.New("leader not found")

	ErrMemberAlreadyExists = errors.New("member already exists")
	ErrMemberNotFound      = errors.New("member not found")

	ErrCuratorAlreadyExists = errors.New("curator already exists")
	ErrCuratorNotFound      = errors.New("curator not found")

	ErrLocationNotFound = errors.New("location not found")

	ErrExpeditionNotFound = errors.New("expedition not found")

	ErrArtifactNotFound = errors.New("artifact not found")

	ErrEquipmentNotFound = errors.New("equipment not found")
)
