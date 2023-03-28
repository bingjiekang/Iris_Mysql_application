package conf

const (
	// ConfigEnv 配置环境
	ConfigEnv = "CONFIG"
	// ConfigFile 配置文件
	ConfigFile = "conf/config.yaml"
)

// 数据库结构体
type MySQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Database     string `mapstructure:"database" json:"database" yaml:"database"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	Debug        bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	LogLevel     string `mapstructure:"logLevel" json:"logLevel" yaml:"logLevel"`
	MaxOpenConns string `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
	MaxIdleConns string `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
	Timeout      string `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	ParseTime    bool   `mapstructure:"parseTime" json:"parseTime" yaml:"parseTime"`
}

// Redis结构体
type Redis struct {
	Network   string `mapstructure:"Network" json:"Network" yaml:"Network"`
	Addr      string `mapstructure:"Addr" json:"Addr" yaml:"Addr"`
	Timeout   int    `mapstructure:"Timeout" json:"Timeout" yaml:"Timeout"`
	MaxActive int    `mapstructure:"MaxActive" json:"MaxActive" yaml:"MaxActive"`
	Password  string `mapstructure:"Password" json:"Password" yaml:"Password"`
	Database  string `mapstructure:"Database" json:"Database" yaml:"Database"`
	Prefix    string `mapstructure:"Prefix" json:"Prefix" yaml:"Prefix"`
	Delim     string `mapstructure:"Delim" json:"Delim" yaml:"Delim"`
}

// 系统配置
type EGAdmin struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Addr    string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	Debug   bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Image   string `mapstructure:"image" json:"image" yaml:"image"`
	// Uploads string `mapstructure:"uploads" json:"uploads" yaml:"uploads"`
}

// 全局配置结构体
type Config struct {
	Mysql MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"Redis" json:"Redis" yaml:"Redis"`
	// Attachment config.Attachment `mapstructure:"attachment" json:"attachment" yaml:"attachment"`
	EGAdmin EGAdmin `mapstructure:"easygoadmin" json:"easygoadmin" yaml:"easygoadmin"`
}
