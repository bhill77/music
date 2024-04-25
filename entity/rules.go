package entity

import "github.com/thedevsaddam/govalidator"

type HasRules interface {
	Rules() govalidator.MapData
}
