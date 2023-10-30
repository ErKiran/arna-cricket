package controllers

import (
	"arna-cricket/api/controllers/cricket"
)

type controller struct {
	CricInfoController cricket.CrickInfoController
}

func NewController() *controller {
	return &controller{
		CricInfoController: cricket.NewCricInfoController(),
	}
}
