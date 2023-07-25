package kustomizegen

import (
	"reflect"
	"testing"

	"josephrodriguez.github.com/kustomizegen/src/types"
)

func TestNewKustomizationWithEmptyResources(t *testing.T) {
	resources := []string{}
	expected := &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Resources:  resources,
		// Initialize other fields here as needed
	}

	result := types.NewKustomization(resources)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewKustomizationWithEmptyResources failed. Expected: %+v, Got: %+v", expected, result)
	}
}

func TestNewKustomizationWithNonEmptyResources(t *testing.T) {
	resources := []string{"deployment.yaml", "service.yaml", "configmap.yaml"}
	expected := &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Resources:  resources,
		// Initialize other fields here as needed
	}

	result := types.NewKustomization(resources)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewKustomizationWithNonEmptyResources failed. Expected: %+v, Got: %+v", expected, result)
	}
}

func TestNewKustomizationWithNilResources(t *testing.T) {
	expected := &types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		// Initialize other fields here as needed
	}

	result := types.NewKustomization(nil)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestNewKustomizationWithNilResources failed. Expected: %+v, Got: %+v", expected, result)
	}
}
