package types

type Kustomization struct {
	APIVersion            string                 `yaml:"apiVersion"`
	Kind                  string                 `yaml:"kind"`
	Metadata              Metadata               `yaml:"metadata,omitempty"`
	Resources             []string               `yaml:"resources,omitempty"`
	CommonLabels          map[string]string      `yaml:"commonLabels,omitempty"`
	CommonAnnotations     map[string]string      `yaml:"commonAnnotations,omitempty"`
	NamePrefix            string                 `yaml:"namePrefix,omitempty"`
	NameSuffix            string                 `yaml:"nameSuffix,omitempty"`
	Namespace             string                 `yaml:"namespace,omitempty"`
	Bases                 []string               `yaml:"bases,omitempty"`
	Crds                  []string               `yaml:"crds,omitempty"`
	PatchesStrategicMerge []string               `yaml:"patchesStrategicMerge,omitempty"`
	PatchesJson6902       []PatchJson6902        `yaml:"patchesJson6902,omitempty"`
	Patches               []Patch                `yaml:"patches,omitempty"`
	VarReference          []VarReference         `yaml:"varReference,omitempty"`
	ConfigMapGenerator    []ConfigMapArgs        `yaml:"configMapGenerator,omitempty"`
	SecretGenerator       []SecretArgs           `yaml:"secretGenerator,omitempty"`
	GeneratorOptions      GeneratorOptions       `yaml:"generatorOptions,omitempty"`
	Transformers          []string               `yaml:"transformers,omitempty"`
	Vars                  []Var                  `yaml:"vars,omitempty"`
	Images                []Image                `yaml:"images,omitempty"`
	Replicas              []Replica              `yaml:"replicas,omitempty"`
	Inventory             Inventory              `yaml:"inventory,omitempty"`
	Configurations        []string               `yaml:"configurations,omitempty"`
	HelmCharts            []HelmChart            `yaml:"helmCharts,omitempty"`
	configStruct          map[string]interface{} `yaml:"-"`
}

type Metadata struct {
	Name      string            `yaml:"name,omitempty"`
	Namespace string            `yaml:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

type PatchJson6902 struct {
	Target  Target `yaml:"target,omitempty"`
	Path    string `yaml:"path,omitempty"`
	Content string `yaml:"content,omitempty"`
}

type Target struct {
	Group     string `yaml:"group,omitempty"`
	Version   string `yaml:"version,omitempty"`
	Kind      string `yaml:"kind,omitempty"`
	Name      string `yaml:"name,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
}

type Patch struct {
	Target Target `yaml:"target,omitempty"`
	Patch  string `yaml:"patch,omitempty"`
}

type VarReference struct {
	Path     string        `yaml:"path,omitempty"`
	FieldRef FieldSelector `yaml:"fieldRef,omitempty"`
}

type FieldSelector struct {
	FieldPath string `yaml:"fieldPath,omitempty"`
}

type ConfigMapArgs struct {
	Name           string          `yaml:"name,omitempty"`
	Files          []string        `yaml:"files,omitempty"`
	LiteralSources []LiteralSource `yaml:"literalSources,omitempty"`
}

type SecretArgs struct {
	Name           string          `yaml:"name,omitempty"`
	Files          []string        `yaml:"files,omitempty"`
	LiteralSources []LiteralSource `yaml:"literalSources,omitempty"`
}

type LiteralSource struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type GeneratorOptions struct {
	DisableNameSuffixHash bool `yaml:"disableNameSuffixHash,omitempty"`
}

type Var struct {
	Name     string        `yaml:"name,omitempty"`
	Objref   Objref        `yaml:"objref,omitempty"`
	Fieldref FieldSelector `yaml:"fieldref,omitempty"`
	Literal  string        `yaml:"literal,omitempty"`
}

type Objref struct {
	APIVersion string `yaml:"apiVersion,omitempty"`
	Kind       string `yaml:"kind,omitempty"`
	Name       string `yaml:"name,omitempty"`
}

type Image struct {
	Name    string `yaml:"name,omitempty"`
	NewName string `yaml:"newName,omitempty"`
	Digest  string `yaml:"digest,omitempty"`
	NewTag  string `yaml:"newTag,omitempty"`
}

type Replica struct {
	Name  string `yaml:"name,omitempty"`
	Count int    `yaml:"count,omitempty"`
}

type Inventory struct {
	Type          string        `yaml:"type,omitempty"`
	ConfigMap     string        `yaml:"configMap,omitempty"`
	LabelSelector LabelSelector `yaml:"labelSelector,omitempty"`
}

type LabelSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels,omitempty"`
}

type HelmChart struct {
	Name        string            `yaml:"name,omitempty"`
	ReleaseName string            `yaml:"releaseName,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Chart       string            `yaml:"chart,omitempty"`
	Repository  string            `yaml:"repository,omitempty"`
	Namespace   string            `yaml:"namespace,omitempty"`
	Values      map[string]string `yaml:"values,omitempty"`
	ValueFiles  []string          `yaml:"valueFiles,omitempty"`
	FileSources []FileSource      `yaml:"fileSources,omitempty"`
}

type FileSource struct {
	Target string `yaml:"target,omitempty"`
	Source string `yaml:"source,omitempty"`
}

type NamespaceTransformer struct {
	APIVersion             string      `yaml:"apiVersion"`
	Kind                   string      `yaml:"kind"`
	Metadata               Metadata    `yaml:"metadata"`
	SetRoleBindingSubjects string      `yaml:"setRoleBindingSubjects,omitempty"`
	UnsetOnly              bool        `yaml:"unsetOnly,omitempty"`
	FieldSpecs             []FieldSpec `yaml:"fieldSpecs,omitempty"`
}

type FieldSpec struct {
	Path   string `yaml:"path,omitempty"`
	Kind   string `yaml:"kind,omitempty"`
	Create bool   `yaml:"create,omitempty"`
}

func PrototypeKustomization() *Kustomization {
	kustomization := &Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
	}

	return kustomization
}

func NewNamespaceTransformer(namespace string, unsetOnly ...bool) NamespaceTransformer {
	ns := NamespaceTransformer{
		APIVersion: "builtin",
		Kind:       "NamespaceTransformer",
		Metadata: Metadata{
			Name:      "namespace-transformer",
			Namespace: namespace,
		},
		FieldSpecs: []FieldSpec{
			{
				Path:   "metadata/name",
				Kind:   "Namespace",
				Create: true,
			},
		},
	}

	if len(unsetOnly) > 0 {
		ns.UnsetOnly = unsetOnly[0]
	} else {
		ns.UnsetOnly = false
	}

	return ns
}
