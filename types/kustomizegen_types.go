package types

type KustomizegenConfiguration struct {
	Version    string                `yaml:"version"`
	NameSuffix string                `yaml:"nameSuffix"`
	NamePrefix string                `yaml:"namePrefix"`
	Overlays   []KustomizegenOverlay `yaml:"overlays"`
}

type KustomizegenOverlay struct {
	Name       string `yaml:"name"`
	NamePrefix string `yaml:"namePrefix"`
	NameSuffix string `yaml:"nameSuffix"`
}

type KustomizegenContext struct {
	Namespace string
}

type KustomizegenBuildCommandContext struct {
	Path       string
	EnableHelm bool
}
