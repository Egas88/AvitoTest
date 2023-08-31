package database

import (
	"github.com/Egas88/AvitoTest/src/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.Preload("Segments").Find(&user)
	return user
}

func (db *UserRepository) SelectSegmentsByUserId(id string) (segments []string, err error) {
	user := domain.User{}
	if err := db.Preload("Segments").First(&user, id).Error; err != nil {
		return nil, err
	}
	var names []string
	for _, segment := range user.Segments {
		names = append(names, segment.Name)
	}
	return names, nil
}

func (db *UserRepository) UpdateUserSegments(SegmentsToAdd []string, SegmentsToDelete []string, idUser string) error {
	user := domain.User{}
	if err := db.Preload("Segments").Where("Id = ?", idUser).First(&user, idUser).Error; err != nil {
		return err
	}
	segmentsToAdd := []domain.Segment{}
	if err := db.Where("name IN ?", SegmentsToAdd).First(&segmentsToAdd).Error; err != nil {
		return err
	}
	segmentsToDelete := []domain.Segment{}
	if err := db.Where("name IN ?", SegmentsToDelete).First(&segmentsToDelete).Error; err != nil {
		return err
	}
	db.Model(&user).Association("Segments").Delete(&segmentsToDelete)
	db.Model(&user).Association("Segments").Unscoped().Clear()
	db.Model(&user).Association("Segments").Append(&segmentsToAdd)
	return nil
}

func (db *UserRepository) Delete(id string) error {
	user := []domain.User{}
	if err := db.Where("Id = ?", id).First(&user, id).Error; err != nil {
		return err
	}
	db.DeleteById(&user, id)
	return nil
}
