package core

import (
  "time"
)

type VisitId string
type VisitDate time.Time
type ExtVisitId string

type Visit struct {
	id         VisitId
	clientId   ClientId
	extVisitId ExtVisitId
	date       VisitDate
}
