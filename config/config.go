package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  MySql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
}
