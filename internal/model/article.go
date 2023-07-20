package model

import (
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var Articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&Articles).Error; err != nil {
		return nil, err
	}

	return Articles, nil
}

func (a Article) ListByIDs(db *gorm.DB, ids []uint32) ([]*Article, error) {
	var Articles []*Article
	db = db.Where("state = ? AND is_del = ?", a.State, 0)
	err := db.Where("id IN (?)", ids).Find(&Articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return Articles, nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var Article Article
	err := db.Where("id = ? AND is_del = ?", t.ID, 0).First(&Article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Article, err
	}

	return Article, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
}
