package usecase

import "github.com/Egas88/AvitoTest/src/domain"

type SegmentInteractor struct {
	SegmentRepository SegmentRepository
}

func (interactor *SegmentInteractor) Add(u domain.Segment) {
	interactor.SegmentRepository.Store(u)
}

func (interactor *SegmentInteractor) GetInfo() []domain.Segment {
	return interactor.SegmentRepository.Select()
}

func (interactor *SegmentInteractor) Delete(id string) error {
	return interactor.SegmentRepository.Delete(id)
}

func (interactor *SegmentInteractor) DeleteByName(name string) error {
	return interactor.SegmentRepository.DeleteByName(name)
}
