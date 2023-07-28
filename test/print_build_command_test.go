package kustomizegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/josephrodriguez/kustomizegen/cmd"
	"github.com/josephrodriguez/kustomizegen/serialization"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func mockBuildCommandKustomization() *types.Kustomization {
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

func mockBuildCommandKustomizegen() *types.KustomizegenConfiguration {

	return &types.KustomizegenConfiguration{
		Version: "1.0.0",
		Overlays: []types.KustomizegenOverlay{
			{
				Name: "overlay1",
			},
			{
				Name: "overlay2",
			},
			{
				Name: "overlayN",
			},
		},
	}
}

func TestGenerateBuildCommand(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "kustomizegen-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	baseDirPath := filepath.Join(tmpDir, "base")

	kustomizationPath := filepath.Join(baseDirPath, "kustomization.yaml")
	err = serialization.MarshalToYAMLFile(mockBuildCommandKustomization(), kustomizationPath)
	assert.NoError(t, err)

	kustomizegenPath := filepath.Join(baseDirPath, "kustomizegen.yaml")
	kustomizegen := mockBuildCommandKustomizegen()
	err = serialization.MarshalToYAMLFile(kustomizegen, kustomizegenPath)
	assert.NoError(t, err)

	command := &cobra.Command{}
	command.Flags().String("root", "", "Path to the Kustomization base folder")
	command.Flags().Bool("enable-helm", true, "Enable Helm")
	command.Flags().Set("root", baseDirPath)

	tempFile, err := ioutil.TempFile("", "kustomizegen_test")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Save the original value of os.Stdout
	stdout := os.Stdout

	// Redirect os.Stdout to the buffer
	os.Stdout = tempFile

	// Call the GenerateBuildCommand method
	cmd.GenerateBuildCommand(command, nil)
	os.Stdout = stdout

	out, _ := ioutil.ReadFile(tempFile.Name())
	output := string(out)

	// Assert the expected output
	expectedOutput := "kustomize build --enable-helm " + filepath.Join(tmpDir, "overlays", "overlay1") + "\n" +
		"kustomize build --enable-helm " + filepath.Join(tmpDir, "overlays", "overlay2") + "\n" +
		"kustomize build --enable-helm " + filepath.Join(tmpDir, "overlays", "overlayN") + "\n"

	assert.Equal(t, expectedOutput, output)
}
