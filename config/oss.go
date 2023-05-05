package config

type Local struct {
	Path           string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
	Imageurlprefix string `mapstructure:"imageurlprefix" json:"imageurlprefix" yaml:"imageurlprefix"`
}
