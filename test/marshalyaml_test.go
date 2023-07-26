package kustomizegen

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
	"josephrodriguez.github.com/kustomizegen/src/serialization"
	"josephrodriguez.github.com/kustomizegen/src/types"
)

func TestMarshalToYAML(t *testing.T) {
	// Define the test data (Kustomization instance)
	data := NewKustomization()

	// Define the filename for the test YAML file
	filename := "test_kustomization.yaml"

	// Cleanup the test file after the test finishes
	defer func() {
		_ = os.Remove(filename)
	}()

	// Call the function being tested
	err := serialization.MarshalToYAML(data, filename)
	if err != nil {
		t.Errorf("Error while marshaling to YAML: %v", err)
		return
	}

	// Read the content of the written file
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("Error while opening the test file: %v", err)
		return
	}
	defer file.Close()

	// Decode the content of the written file to check if it's correct
	var decodedData types.Kustomization
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&decodedData); err != nil {
		t.Errorf("Error while decoding the written file: %v", err)
		return
	}

	// Compare the data
	if !compareKustomizations(data, &decodedData) { // Convert decodedData to a pointer
		t.Errorf("Decoded data does not match the original data")
		return
	}
}

// NewKustomization creates a new test Kustomization instance
func NewKustomization() *types.Kustomization {
	return &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Resources:  []string{"deployment.yaml", "service.yaml"},
		CommonLabels: map[string]string{
			"app":  "myapp",
			"env":  "prod",
			"tier": "frontend",
		},
		// Add other fields as needed for your tests
	}
}

// Helper function to compare two Kustomization instances
func compareKustomizations(k1, k2 *types.Kustomization) bool {
	// Compare the relevant fields of the Kustomization instances
	// Here, we're comparing only some fields for simplicity, you can add more fields as needed.
	return k1.APIVersion == k2.APIVersion &&
		k1.Kind == k2.Kind &&
		compareStringSlice(k1.Resources, k2.Resources) &&
		compareStringMap(k1.CommonLabels, k2.CommonLabels)
}

// Helper function to compare two string slices
func compareStringSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

// Helper function to compare two string maps
func compareStringMap(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if val, ok := map2[k]; !ok || val != v {
			return false
		}
	}
	return true
}
