package gpxsettings

import (
	"github.com/msbranco/goconfig"
	"strconv"
)

type Settings struct {
	RestHR int
	HRZone1 int
	HRZone2 int
	HRZone3 int
	HRZone4 int
	HRZone5 int
}

func LoadSettings(setupFile string) (*Settings, error) {
	c, error   := goconfig.ReadConfigFile("config.cfg")
	if error != nil {
		return nil, error
	}
	
	hrzone1 := 50
	p, _   := c.GetString("Config","HRZone1")
	if hrzone1, error = strconv.Atoi(p); error == nil {
		hrzone1  = 50
	}
	
	hrzone2  := 60
	p, _   = c.GetString("Config","HRZone2")
	if hrzone2, error = strconv.Atoi(p); error == nil {
		hrzone2  = 60
	}

	hrzone3  := 70
	p, _   = c.GetString("Config","HRZone3")
	if hrzone3, error = strconv.Atoi(p); error == nil {
		hrzone3  = 70
	}

	hrzone4  := 80
	p, _   = c.GetString("Config","HRZone4")
	if hrzone4, error = strconv.Atoi(p); error == nil {
		hrzone4  = 80
	}

	hrzone5  := 90
	p, _   = c.GetString("Config","HRZone5")
	if hrzone5, error = strconv.Atoi(p); error == nil {
		hrzone5  = 90
	}

	resthr  := 50
	p, _   = c.GetString("Config","RestHR")
	if resthr, error = strconv.Atoi(p); error == nil {
		resthr  = 50
	}

	return &Settings{resthr, hrzone1,hrzone2,hrzone3,hrzone4,hrzone5}, nil	
} 

