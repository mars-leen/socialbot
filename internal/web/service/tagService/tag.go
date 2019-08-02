package tagService

import "socialbot/internal/web/model"

func GetTagsMapWithCid() (rs map[int][]model.ConTag, err error) {
	list := model.TagList{}
	err = list.GetList()
	if err != nil {
		return rs, err
	}

	l := len(list)
	if l == 0 {
		return map[int][]model.ConTag{}, nil
	}

	result := make(map[int][]model.ConTag)
	for _, t := range list {
		conTag := model.ConTag{
			Id:          t.Id,
			Title:       t.Title,
			ShortName:   t.ShortName,
			Description: t.Description,
			BoardName:   t.BoardName,
		}

		if _, exist := result[t.Cid]; !exist {
			result[t.Cid] = []model.ConTag{conTag}
		} else {
			result[t.Cid] = append(result[t.Cid], conTag)
		}
	}

	return result, nil
}
