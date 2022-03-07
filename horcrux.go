package horcrux

import (
	"os"
	"html/template"
)

type Horcrux struct {
	Prompt string `json:"prompt"`
	Encrypted template.JS `json:"encrypted"`
	Treasure template.JS `json:"treasure"`
}

func GetEnvironmentHorcrux() (horcrux Horcrux){
	return Horcrux{
		Prompt: os.Getenv("HORCRUX_PROMPT"),
		Encrypted: template.JS(os.Getenv("HORCRUX_ENCRYPTED")),
		Treasure: template.JS(os.Getenv("HORCRUX_TREASURE")),
	}
}