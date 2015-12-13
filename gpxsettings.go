package gpxsettings

import (
	"github.com/msbranco/goconfig"
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
	c, error   := goconfig.ReadConfigFile(setupFile)
	if error != nil {
		return nil, error
	}

	hrzone1, error := c.GetInt("Config","HRZone1")
	if error != nil {
		return nil, error
	}

	hrzone2, error := c.GetInt("Config","HRZone2")
	if error != nil {
		return nil, error
	}

	hrzone3, error := c.GetInt("Config","HRZone3")
	if error != nil {
		return nil, error
	}

	hrzone4, error := c.GetInt("Config","HRZone4")
	if error != nil {
		return nil, error
	}

	hrzone5, error := c.GetInt("Config","HRZone5")
	if error != nil {
		return nil, error
	}

	resthr, error := c.GetInt("Config","RestHR")
	if error != nil {
		return nil, error
	}

	return &Settings{resthr, hrzone1,hrzone2,hrzone3,hrzone4,hrzone5}, nil	
} 

