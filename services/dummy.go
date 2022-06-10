package services

import "telecommunication/entities"

type DummyService interface {
	Dummy() entities.Dummy
}

type dummyService struct {
	dummy entities.Dummy
}

func New() DummyService {
	return &dummyService{}
}

func (s *dummyService) Dummy() entities.Dummy {
	s.dummy.Dummy = "dummy-true"
	return s.dummy
}
