package repository

import (
	"log"
	"main/internal/entity"
	"main/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactRepository struct {
	Repository[entity.Contact]
	Log *logrus.Logger
}

func NewContactRepository(log *logrus.Logger) *ContactRepository {
	return &ContactRepository{
		Log: log,
	}
}
func (r *ContactRepository) FindByIdAndUserId(db *gorm.DB, contact *entity.Contact, id string, userId string) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(contact).Error
}

func (r *ContactRepository) Search(db *gorm.DB, request *model.SearchContactRequest) ([]entity.Contact, int64, error) {
	var contacts []entity.Contact
	if err := db.scopes(r.FilterContact(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}
	return contacts, total, nil
}

func (r *ContactRepository) FilterContact(reuqest *model.SearchContactRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", reuqest.UserId)

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
		}

		if phone := request.Phone; phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}

		if email := reuqest.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}
		return tx

	}
}
