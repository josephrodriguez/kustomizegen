package kustomizegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/josephrodriguez/kustomizegen/cmd"
	"github.com/josephrodriguez/kustomizegen/filesystem"
	"github.com/josephrodriguez/kustomizegen/serialization"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func mockKustomizegen() *types.KustomizegenConfiguration {

	return &types.KustomizegenConfiguration{
		Version: "1.0.0",
		Overlays: []types.KustomizegenOverlay{
			{
				Name: "develop",
			},
			{
				Name: "staging",
			},
			{
				Name: "production",
			},
		},
	}
}

func mockKustomization() *types.Kustomization {
	return &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Resources: []string{
			"resources.yaml",
		},
		HelmCharts: []types.HelmChart{
			{
				Name:        "argo-cd",
				ReleaseName: "argocd",
				Repository:  "https:argoproj.github.io/argo-helm",
				Version:     "5.40.0",
			},
		},
	}
}

func TestConfigureCommand(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "kustomizegen-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	baseDirPath := filepath.Join(tempDir, "base")

	kustomizationPath := filepath.Join(baseDirPath, "kustomization.yaml")
	err = serialization.MarshalToYAMLFile(mockKustomization(), kustomizationPath)
	assert.NoError(t, err)

	kustomizegenPath := filepath.Join(baseDirPath, "kustomizegen.yaml")
	kustomizegen := mockKustomizegen()
	err = serialization.MarshalToYAMLFile(kustomizegen, kustomizegenPath)
	assert.NoError(t, err)

	command := &cobra.Command{}
	command.Flags().String("root", "", "Path to the Kustomization base folder")
	command.Flags().Set("root", baseDirPath)

	cmd.GenerateOverlaysCommand(command, nil)
	assert.NoError(t, err)

	for _, overlay := range kustomizegen.Overlays {
		overlayPath := filepath.Join(tempDir, "overlays", overlay.Name)
		exists, err := filesystem.ExistFolder(overlayPath)
		assert.NoError(t, err)
		assert.True(t, exists)
	}
}
