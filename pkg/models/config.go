package models

type InitAnwsers struct {
	Username       string
	ConcourseUrl   string
	ConcourseTeams []string
}

type Config struct {
	ConcourseConfig ConcourseConfig `yaml:"concourse"`
}

type ConcourseConfig struct {
	Username string   `yaml:"username"`
	Url      string   `yaml:"url"`
	Teams    []string `yaml:"teams"`
}

func InitAnwsersToConfig(a InitAnwsers) *Config {
	return &Config{
		ConcourseConfig: ConcourseConfig{
			Username: a.Username,
			Url:      a.ConcourseUrl,
			Teams:    a.ConcourseTeams,
		},
	}
}
