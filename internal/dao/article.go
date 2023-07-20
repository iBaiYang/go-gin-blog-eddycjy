package dao

import (
	"github.com/iBaiYang/go-gin-blog-eddycjy/internal/model"
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/app"
)

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticleListByIDs(ids []uint32, state uint8) ([]*model.Article, error) {
	article := model.Article{State: state}
	return article.ListByIDs(d.engine, ids)
}

func (d *Dao) GetArticle(id uint32) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title string, desc string, coverImageUrl string, content string, state uint8, createdBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: coverImageUrl,
		Content:       content,
		State:         state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, desc string, coverImageUrl string, content string, state uint8, modifiedBy string) error {
	//article := model.Article{
	//	Title:  title,
	//	Desc:  desc,
	//	CoverImageUrl:  coverImageUrl,
	//	Content:  content,
	//	State: state,
	//	Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	//}
	//
	//return article.Update(d.engine)

	article := model.Article{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = title
	}
	if coverImageUrl != "" {
		values["cover_image_url"] = title
	}
	if content != "" {
		values["content"] = title
	}

	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
