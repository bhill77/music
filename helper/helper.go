package helper

import (
	"net/url"

	"github.com/bhill77/music/entity"
	"github.com/thedevsaddam/govalidator"
)

func Validate(data entity.HasRules) url.Values {
	opts := govalidator.Options{
		Data:  data,
		Rules: data.Rules(),
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()

	return e
}
