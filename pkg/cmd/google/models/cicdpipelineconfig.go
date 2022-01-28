package models

type CiCdPipelineConfig struct {
	csrConfig CSRConfig
	cbtConfig CBTConfig
}

func NewCiCdPipelineConfig(csrConfig CSRConfig, cbtConfig CBTConfig) *CiCdPipelineConfig {
	cfg := new(CiCdPipelineConfig)

	cfg.SetCSRConfig(csrConfig)
	cfg.SetCBTConfig(cbtConfig)

	return cfg
}

func (cicdConfig *CiCdPipelineConfig) SetCSRConfig(csrConfig CSRConfig) {
	cicdConfig.csrConfig = csrConfig
}

func (cicdConfig CiCdPipelineConfig) GetCSRConfig() CSRConfig {
	return cicdConfig.csrConfig
}

func (cicdConfig *CiCdPipelineConfig) SetCBTConfig(cbtConfig CBTConfig) {
	cicdConfig.cbtConfig = cbtConfig
}

func (cicdConfig CiCdPipelineConfig) GetCBTConfig() CBTConfig {
	return cicdConfig.cbtConfig
}
