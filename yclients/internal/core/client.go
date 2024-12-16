package core

import (
	"time"

	"superadmin.ru/internal/id"
)

type ClientId string
type ClientName string
type ExtClientId string
type VisitId string
type VisitDate time.Time
type ExtVisitId string

type Client struct {
	id            ClientId
	applicationId ApplicationId
	extClientId   ExtClientId
	name          ClientName
}

func (a *YClientsApp) RegisterClient(extClientId ClientId, name ClientName) Client {
	// repos.CreateUniqClient(client)

	return Client{
		id:            ClientId(id.Random()),
		applicationId: a.Id,
		extClientId:   ExtClientId(extClientId),
		name:          name,
	}
}
