package entity

type Config struct {
	ApiParam struct {
		Fixmedins_code string `json:"fixmedins_code" yaml:"fixmedins_code" gorm:"column:fixmedins_code"`
		Infosyscode    string `json:"infosyscode" yaml:"infosyscode" gorm:"column:infosyscode"`
		Infosyssign    string `json:"infosyssign" yaml:"infosyssign" gorm:"column:infosyssign"`
		Url            string `json:"url" yaml:"url" gorm:"column:url"`
	} `json:"api_param" yaml:"api_param" gorm:"embedded"`

	Port         string `json:"port" yaml:"port" gorm:"column:port"`
	EntityMstrId int64  `json:"entitymstr_id" yaml:"entitymstr_id" gorm:"column:entitymstr_id"`
}
