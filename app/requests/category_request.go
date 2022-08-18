package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `json:"name,omitempty" valid:"name"`
	Description string `json:"description,omitempty" valid:"description"`
	State       string `json:"state,omitempty" valid:"state"`
	ImageUrl    string `json:"image_url,omitempty" valid:"image_url"`
	Sort        string `json:"sort,omitempty" valid:"sort"`
	ParentID    string `json:"parent_id,omitempty" valid:"parent_id"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"max_cn:255"},
		"state":       []string{"in:0,1"},
		"image_url":   []string{"max_cn:255"},
		"sort":        []string{"numeric_between:1,99999"},
		"parent_id":   []string{"exists:categories,id"},
	}
	message := govalidator.MapData{
		"name": []string{
			"required:标题为必填项，参数名称 name",
			"min_cn:标题长度需大于2个字",
			"max_cn:标题长度不能超过8个字",
			"not_exists:分类标题已存在，请重新填写",
		},
		"description": []string{
			"max_cn:描述的最大长度不能超过255个字",
		},
		"state": []string{
			"in:分类状态只支持 0,1",
		},
		"image_url": []string{
			"max_cn:图片路径长度需小于 255个字符",
		},
		"sort": []string{
			"numeric_between:排序数字介于1~99999之间",
		},
		"parent_id": []string{
			"exists:父级标识不存在",
		},
	}
	return validate(data, rules, message)
}

func CategoryUpdate(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":        []string{"min_cn:2", "max_cn:8", "not_exists:categories,name," + c.Param("id")},
		"description": []string{"max_cn:255"},
		"state":       []string{"in:0,1"},
		"image_url":   []string{"max_cn:255"},
		"sort":        []string{"numeric_between:1,99999"},
		"parent_id":   []string{"exists:categories,id"},
	}
	message := govalidator.MapData{
		"name": []string{
			"required:标题为必填项，参数名称 name",
			"min_cn:标题长度需大于2个字",
			"max_cn:标题长度不能超过8个字",
			"not_exists:分类标题已存在，请重新填写",
		},
		"description": []string{
			"max_cn:描述的最大长度不能超过255个字",
		},
		"state": []string{
			"in:分类状态只支持 0,1",
		},
		"image_url": []string{
			"max_cn:图片路径长度需小于 255个字符",
		},
		"sort": []string{
			"numeric_between:排序数字介于1~99999之间",
		},
		"parent_id": []string{
			"exists:父级标识不存在",
		},
	}
	return validate(data, rules, message)
}
