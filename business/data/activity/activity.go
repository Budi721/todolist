package activity

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetActivities(ctx context.Context) ([]Activity, error)
	GetActivity(ctx context.Context, id int) (*Activity, error)
	CreateActivity(ctx context.Context, activity *NewActivity) (*Activity, error)
	UpdateActivity(ctx context.Context, id int, activity *UpdateActivity) (*Activity, error)
	DeleteActivity(ctx context.Context, id int) (int, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetActivities(ctx context.Context) ([]Activity, error) {
	var activities []Activity

	if err := r.db.WithContext(ctx).Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

func (r *repository) GetActivity(ctx context.Context, id int) (*Activity, error) {
	var activity Activity

	if err := r.db.WithContext(ctx).First(&activity, id).Error; err != nil {
		return nil, err
	}

	return &activity, nil
}

func (r *repository) CreateActivity(ctx context.Context, newActivity *NewActivity) (*Activity, error) {
	activity := Activity{
		Email: newActivity.Email,
		Title: newActivity.Title,
	}

	if err := r.db.WithContext(ctx).Create(&activity).Error; err != nil {
		return nil, err
	}

	return &activity, nil
}

func (r *repository) UpdateActivity(ctx context.Context, id int, updateActivity *UpdateActivity) (*Activity, error) {
	activity := Activity{}
	if err := r.db.WithContext(ctx).First(&activity, id).Error; err != nil {
		return nil, err
	}

	activity.UpdatedAt = time.Now()
	activity.Title = updateActivity.Title

	if err := r.db.WithContext(ctx).Save(&activity).Error; err != nil {
		return nil, err
	}

	return &activity, nil
}

func (r *repository) DeleteActivity(ctx context.Context, id int) (int, error) {
	activity := Activity{}

	if err := r.db.WithContext(ctx).First(&activity, id).Error; err != nil {
		return -1, err
	}

	if err := r.db.WithContext(ctx).Delete(&activity, id).Error; err != nil {
		return -1, err
	}

	return activity.Id, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
