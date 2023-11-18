package config

type MySql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *MySql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
