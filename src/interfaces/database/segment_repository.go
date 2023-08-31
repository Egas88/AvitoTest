package database

import "github.com/Egas88/AvitoTest/src/domain"

type SegmentRepository struct {
	SqlHandler
}

func (db *SegmentRepository) Store(s domain.Segment) {
	db.Create(&s)
}

func (db *SegmentRepository) Select() []domain.Segment {
	segment := []domain.Segment{}
	db.FindAll(&segment)
	return segment
}

func (db *SegmentRepository) Delete(id string) error {
	segment := []domain.Segment{}
	if err := db.Where("Id = ?", id).First(&segment).Error; err != nil {
		return err
	}
	db.DeleteById(&segment, id)
	return nil
}

func (db *SegmentRepository) DeleteByName(name string) error {
	segment := []domain.Segment{}
	if err := db.Where("Name = ?", name).First(&segment).Error; err != nil {
		return err
	}
	db.Where("name = ?", name).Delete(segment)
	return nil
}
