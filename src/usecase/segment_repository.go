package usecase

import "github.com/Egas88/AvitoTest/src/domain"

type SegmentRepository interface {
	Store(domain.Segment)
	Select() []domain.Segment
	Delete(id string) error
	DeleteByName(name string) error
}
