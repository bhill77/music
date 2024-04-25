package helper

import (
	"net/url"

	"github.com/bhill77/music/entity"
	"github.com/thedevsaddam/govalidator"
)

func Validate(rules govalidator.MapData, data *entity.Song) url.Values {
	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()

	return e
}
