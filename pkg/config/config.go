package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	InfoLog        *log.Logger
	Session        *scs.SessionManager
	CSRFToken      string
	AdminTemplates *template.Template
	UITemplates    *template.Template
}
