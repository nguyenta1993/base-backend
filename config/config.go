package config

type AppConfig struct {
	ServiceName string            `mapstructure:"serviceName"`
	Development bool              `mapstructure:"development"`
	Logger      *LoggerConfig     `mapstructure:"logger"`
	Http        *HttpConfig       `mapstructure:"http"`
	GRPC        *GrpcConfig       `mapstructure:"grpc"`
	Database    *DatabaseConfig   `mapstructure:"database"`
	Tracing     *TracingConfig    `mapstructure:"tracing"`
	Kafka       *KafkaConfig      `mapstructure:"kafka"`
	Redis       *RedisConfig      `mapstructure:"redis"`
	Heathcheck  *HeathcheckConfig `mapstructure:"heathcheck"`
	Metrics     *MetricsConfig    `mapstructure:"metrics"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

type HttpConfig struct {
	Port            string `mapstructure:"port"`
	Development     bool   `mapstructure:"development"`
	ShutdownTimeout int    `mapstructure:"shutdownTimeout"`

	Resources    []string            `mapstructure:"resources"`
	RateLimiting *RateLimitingConfig `mapstructure:"rateLimiting"`
}

type RateLimitingConfig struct {
	RateFormat string `mapstructure:"rateFormat"`
}

type GrpcConfig struct {
	Port              string `mapstructure:"port"`
	Development       bool   `mapstructure:"development"`
	MaxConnectionIdle int    `mapstructure:"maxConnectionIdle"`
	Timeout           int    `mapstructure:"timeout"`
	MaxConnectionAge  int    `mapstructure:"maxConnectionAge"`
	Time              int    `mapstructure:"time"`
}
type DatabaseConfig struct {
	ReadDbCfg  *ReadDbConfig  `mapstructure:"readDb"`
	WriteDbCfg *WriteDbConfig `mapstructure:"writeDb"`
}

type ReadDbConfig struct {
	DbType            string `mapstructure:"dbType"`
	ConnectionString  string `mapstructure:"connectionString"`
	MigrationFilePath string `mapstructure:"migrationFilePath"`
}

type WriteDbConfig struct {
	DbType            string `mapstructure:"dbType"`
	ConnectionString  string `mapstructure:"connectionString"`
	MigrationFilePath string `mapstructure:"migrationFilePath"`
}

type TracingConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	HostPort    string `mapstructure:"hostPort"`
	Enable      bool   `mapstructure:"enable"`
	LogSpans    bool   `mapstructure:"logSpans"`
}

type KafkaConfig struct {
	Config *KafkaConfigDetail `mapstructure:"config"`
	Topics *KafkaTopics       `mapstructure:"topics"`
	Dialer *KafkaDialer       `mapstructure:"dialer"`
}

type KafkaDialer struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type KafkaConfigDetail struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupID    string   `mapstructure:"groupID"`
	InitTopics bool     `mapstructure:"initTopics"`
	NumWorker  int      `mapstructure:"numWorker"`
}

type KafkaTopics struct {
	UserUpdate  KafkaTopicConfig `mapstructure:"userUpdate"`
	UserUpdated KafkaTopicConfig `mapstructure:"userUpdated"`
}

type KafkaTopicConfig struct {
	TopicName         string `mapstructure:"topicName"`
	NumPartitions     int    `mapstructure:"numPartitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}

type RedisConfig struct {
	Addrs    []string `mapstructure:"addrs"`
	Password string   `mapstructure:"password"`
	DB       int      `mapstructure:"db"`
	PoolSize int      `mapstructure:"poolSize"`
}

type HeathcheckConfig struct {
	Interval           int    `mapstructure:"interval"`
	Port               string `mapstructure:"port"`
	GoroutineThreshold int    `mapstructure:"goroutineThreshold"`
}

type MetricsConfig struct {
	PrometheusPath string `mapstructure:"prometheusPath"`
	PrometheusPort string `mapstructure:"prometheusPort"`
}
