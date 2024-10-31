package entity

type Config struct {
	ApiParam struct {
		Fixmedins_code string `yaml:"fixmedins_code"`
		Infosyscode    string `yaml:"infosyscode"`
		Infosyssign    string `yaml:"infosyssign"`
		Url            string `yaml:"url"`
	} `yaml:"apiParam"`

	Port string `yaml:"port"`

	EntityMstrId int64 `yaml:"entitymstr_id"`
}
