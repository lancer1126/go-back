package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  MySql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
}
