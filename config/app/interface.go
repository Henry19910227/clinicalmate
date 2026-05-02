package app

import model "clinicalmate/internal/model/config"

type Config interface {
	Config() *model.Config
}
