package conf

import "fmt"

var SupportedType = map[string]bool{
	"top":      true,
	"guonei":   true,
	"guoji":    true,
	"yule":     true,
	"tiyu":     true,
	"junshi":   true,
	"keji":     true,
	"caijing":  true,
	"youxi":    true,
	"qiche":    true,
	"jiankang": true,
}

type NewsConfig struct {
	ApiUri   string
	ApiKey   string
	Type     string
	Page     int
	Size     int
	IsFilter int
}

func (conf *NewsConfig) Check() error {
	if conf.ApiUri == "" {
		return fmt.Errorf("api url cannot be null")
	}

	if conf.ApiKey == "" {
		return fmt.Errorf("api key cannot be null")
	}

	if _, ok := SupportedType[conf.Type]; !ok {
		return fmt.Errorf("type is not supported")
	}

	if conf.Page > 50 {
		return fmt.Errorf("page cannot be larger than 50")
	}

	if conf.Size > 30 {
		return fmt.Errorf("page size cannot be larger than 30")
	}

	return nil
}
