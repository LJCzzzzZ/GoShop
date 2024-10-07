package common

// 前台参数
var (
	FontedHost = "172.28.71.242"
	FontedPort = "8080"
)

// 后台参数
var (
	BackendHost = "172.28.71.242"
	BackendPort = "8081"
)

// validate
var (
	ValidateHost1 = "172.28.71.242"
	ValidatePort  = "8083"

	// 集群地址
	ClusterHostArray = []string{"172.28.71.242", "172.28.71.242"}

	// 每个用户抢购间隔时间
	Interval = 5
)

// GetProduct 接口参数
var (
	// 数量控制接口服务器内网 IP
	GetProductIP   = "172.28.71.242"
	GetProductPort = "12345"
)

// MySQL 连接参数
var (
	MySQLUserName = "root"          // 账号
	MySQLPassWord = "123456"        // 密码
	MySQLHost     = "47.121.209.33" // 数据库地址，可以是Ip或者域名
	MySQLPort     = 3306            // 数据库端口
	MySQLDbName   = "productshop"   // 数据库名
)

// url格式：amqp://账号:密码 @RabbitMQ 服务器地址:端口号/vHost
// RabbitMQ 连接参数
var (
	RMQUser  = "guest"
	RMQPawsd = "guest"
	RMQHost  = "47.121.209.33"
	RMQPort  = "5672"
	RMQVHost = "productshop"
)

// redis 配置
var (
	// redis 地址
	RedisHost = "47.121.209.33"
	RedisPort = "6379"
	// redis 密码
	RedisPassword = "123456"
)
