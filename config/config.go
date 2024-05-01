package config

type secret struct {
	JwtSecret string `env:"JWT_SECRET" envDefault:"secret"`
	DBUrl     string `env:"DB_URL" envDefault:"mysql://root:@tcp(localhost:3306)/mf_development?parseTime=true"`
}

type setting struct {
	TestConfig string `yaml:"test_config" default:"abc"`
	Env        string `yaml:"env"`
}
