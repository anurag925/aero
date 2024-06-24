package config

type secret struct {
	JwtSecret string `env:"JWT_SECRET" envDefault:"secret"`
	DBUrl     string `env:"DB_URL" envDefault:"mysql://root:@tcp(localhost:3306)/mf_development?parseTime=true"`

	// mongo_db
	MongoDBUrl  string `env:"MONGO_DB_URL"`
	MongoDBName string `env:"MONGO_DB_NAME"`

	// worker redis
	WorkerRedisPassword string `env:"WORKER_REDIS_PASSWORD"`
	WorkerRedisPort     int    `env:"WORKER_REDIS_PORT"`
	WorkerRedisHost     string `env:"WORKER_REDIS_HOST"`
	WorkerRedisDBName   string `env:"WORKER_REDIS_DB_NAME"`
	WorkerRedisPoolSize int    `env:"WORKER_REDIS_POOL_SIZE"`
	WorkerRedisMaxRetry int    `env:"WORKER_REDIS_MAX_RETRY"`
	WorkerRedisDB       int    `env:"WORKER_REDIS_DB"`
}

type setting struct {
	TestConfig string `yaml:"test_config" default:"abc"`
	Env        string `yaml:"env"`
}
