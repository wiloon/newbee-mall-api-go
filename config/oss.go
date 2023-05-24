package config

type Local struct {
	Path           string `mapstructure:"path" json:"path" yaml:"path"`                // 本地文件路径
	UrlPrefix      string `mapstructure:"urlprefix" json:"urlprefix" yaml:"urlprefix"` // 本地文件路径
	Imageurlprefix string `mapstructure:"imageurlprefix" json:"imageurlprefix" yaml:"imageurlprefix"`
	MallUrl        string `mapstructure:"mallurl" json:"mallurl" yaml:"mallurl"`
}
