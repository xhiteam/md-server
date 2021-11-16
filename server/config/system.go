package config

type System struct {
	Env        string `mapstructure:"env" json:"env" yaml:"env"`                       // 环境值
	DomainName string `mapstructure:"domain-name" json:"domainName" yaml:"DomainName"` //域名
	Port       string `mapstructure:"port" json:"port" yaml:"port"`                    // 端口值
	DbType     string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`            // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
