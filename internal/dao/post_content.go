package dao

import "github.com/rocboss/paopao-ce/internal/model"

func (d *dataServant) CreatePostContent(content *model.PostContent) (*model.PostContent, error) {
	return content.Create(d.engine)
}

func (d *dataServant) GetPostContentsByIDs(ids []int64) ([]*model.PostContent, error) {
	return (&model.PostContent{}).List(d.engine, &model.ConditionsT{
		"post_id IN ?": ids,
		"ORDER":        "sort ASC",
	}, 0, 0)
}

func (d *dataServant) GetPostContentByID(id int64) (*model.PostContent, error) {
	return (&model.PostContent{
		Model: &model.Model{
			ID: id,
		},
	}).Get(d.engine)
}
