package types

type KustomizegenConfiguration struct {
	Version    string                 `yaml:"version"`
	NameSuffix KustomizegenNameSuffix `yaml:"nameSuffix"`
	NamePrefix KustomizegenNamePrefix `yaml:"namePrefix"`
	Overlays   []KustomizegenOverlay  `yaml:"overlays"`
}

type KustomizegenNamePrefix struct {
	Value string                      `yaml:"value"`
	Rules []ResourceRuleConfiguration `yaml:"rules"`
}

type KustomizegenNameSuffix struct {
	Value string                      `yaml:"value"`
	Rules []ResourceRuleConfiguration `yaml:"rules"`
}

type KustomizegenOverlay struct {
	Name       string                 `yaml:"name"`
	NamePrefix KustomizegenNamePrefix `yaml:"namePrefix"`
	NameSuffix KustomizegenNameSuffix `yaml:"nameSuffix"`
}

type KustomizegenContext struct {
	Namespace string
}
