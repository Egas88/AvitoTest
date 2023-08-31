package usecase

import "github.com/Egas88/AvitoTest/src/domain"

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	SelectSegmentsByUserId(id string) (segments []string, err error)
	Delete(id string) error
	UpdateUserSegments(SegmentsToAdd []string, SegmentsToDelete []string, idUser string) error
}
