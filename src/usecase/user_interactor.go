package usecase

import "github.com/Egas88/AvitoTest/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) GetSegmentsByUserId(id string) (segments []string, err error) {
	return interactor.UserRepository.SelectSegmentsByUserId(id)
}

func (interactor *UserInteractor) Delete(id string) error {
	return interactor.UserRepository.Delete(id)
}

func (interactor *UserInteractor) UpdateUserSegments(SegmentsToAdd []string, SegmentsToDelete []string, idUser string) error {
	return interactor.UserRepository.UpdateUserSegments(SegmentsToAdd, SegmentsToDelete, idUser)
}
