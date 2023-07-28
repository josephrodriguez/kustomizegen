// package kustomizegen

// import (
// 	"path/filepath"
// 	"testing"

// 	"github.com/josephrodriguez/kustomizegen/filesystem"
// 	"github.com/josephrodriguez/kustomizegen/types"
// 	"github.com/stretchr/testify/assert"
// )

// // Mock implementation of getConfig
// func mockGetConfig(configFilePath string) (*types.Configuration, error) {
// 	// Return a dummy configuration for testing
// 	return &types.Configuration{
// 		Version: "1.0.0",
// 		Namespaces: []types.Namespace{
// 			{
// 				Name:      "develop",
// 				UnsetOnly: true,
// 			},
// 			{
// 				Name:      "staging",
// 				UnsetOnly: false,
// 			},
// 			{
// 				Name:      "production",
// 				UnsetOnly: false,
// 			},
// 		},
// 	}, nil
// }

// // Mock implementation of getOverlayPath
// func mockGetOverlayPath(rootPath string, namespace *types.Namespace) (string, error) {
// 	// Return a dummy overlay path for testing
// 	return "/path/to/overlay", nil
// }

// // Mock implementation of generateKustomization
// func mockGenerateKustomization(namespaceName, resourcesPath string) (*types.Kustomization, error) {
// 	// Return a dummy Kustomization for testing
// 	return &types.Kustomization{
// 		APIVersion: "kustomize.config.k8s.io/v1beta1",
// 		Kind:       "Kustomization",
// 		Resources: []string{
// 			"resources.yaml",
// 		},
// 	}, nil
// }

// func TestConfigureCommand(t *testing.T) {
// 	namespace := &types.KustomizationNamespace{Name: "test"}

// 	// Create a temporary test directory
// 	tempDir, err := filesystem.CreateTempDir()
// 	assert.NoError(t, err)
// 	defer filesystem.RemoveTempDir(tempDir)

// 	// Create kustomization.yaml file with test content
// 	kustomizationContent := `apiVersion: kustomize.config.k8s.io/v1beta1
// kind: Kustomization
// helmCharts:
// - name: argo-cd
//   releaseName: argocd
//   repo: https://argoproj.github.io/argo-helm
//   version: 5.40.0`

// 	kustomizationFile := filepath.Join(tempDir, "kustomization.yaml")
// 	err = filesystem.WriteFile(kustomizationFile, []byte(kustomizationContent))
// 	assert.NoError(t, err)

// 	// Run the configure command
// 	err = configureCommand(tempDir, namespace)
// 	assert.NoError(t, err)

// 	// Validate the created overlay path
// 	parent := filepath.Dir(filepath.Join(tempDir, "overlays", namespace.Name))
// 	overlayPath := filepath.Join(parent, "overlays", namespace.Name)
// 	exists, err := filesystem.ExistFolder(overlayPath)
// 	assert.NoError(t, err)
// 	assert.True(t, exists)
// }