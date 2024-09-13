package comment_repository

import (
	"avitoTest/data/entities"
	"context"

	"gorm.io/gorm"
)

type commentRepositoryGorm struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepositoryGorm{db: db}
}

func (r *commentRepositoryGorm) Create(ctx context.Context, comment *entities.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *commentRepositoryGorm) FindByUsername(ctx context.Context, username string) ([]*entities.Comment, error) {
	var comments []*entities.Comment
	if err := r.db.WithContext(ctx).Where("user_id = (SELECT id FROM users WHERE username = ?)", username).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepositoryGorm) FindByOrganizationID(ctx context.Context, organizationID int) ([]*entities.Comment, error) {
	var comments []*entities.Comment
	if err := r.db.WithContext(ctx).Where("organization_id = ?", organizationID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepositoryGorm) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entities.Comment{}, id).Error
}
