package categoryService

import (
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/tagService"
)

func GetCategoryWithTag() (rs []model.ConCategoryTags, err error) {
	category := model.CategoryList{}
	err = category.GetList()
	if err != nil {
		return rs, err
	}

	l := len(category)
	if l == 0 {
		return []model.ConCategoryTags{}, nil
	}

	tagsWithCid, err := tagService.GetTagsMapWithCid()
	if err != nil {
		return rs, err
	}

	result := make([]model.ConCategoryTags, l)
	for i, v := range category {
		cate := model.ConCategoryTags{
			Id:          v.Id,
			Title:       v.Title,
			ShortName:   v.ShortName,
			Description: v.Description,
			Sort:        v.Sort,
			Tags:        []model.ConTag{},
		}
		v, exist := tagsWithCid[v.Id]
		if exist {
			cate.Tags = v
		}

		result[i] = cate
	}

	return result, nil
}
