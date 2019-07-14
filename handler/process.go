package handler

import (
	"M-GateDBConfig/configuration"
	"M-GateDBConfig/dbhandler"
	"M-GateDBConfig/model"
)

// SetParam to input value to DB
func (a *AppHandler) SetParam(DB dbhandler.SetDBValueInterface) error {

	simple := configuration.GetParamValue()

	data := &model.Parameters{
		Key:   simple.Key,
		Value: simple.Value,
	}

	err := DB.SetParameter(data)

	return err
}