package common

type ConfigurationVersionInfo struct {
	TPVersion *uint   `json:"mavpay_config_version,omitempty" yaml:"mavpay_config_version,omitempty"`
	Version   *string `json:"version,omitempty" yaml:"version,omitempty"`
}
