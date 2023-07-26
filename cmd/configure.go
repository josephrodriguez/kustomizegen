package cmd

import (
	"fmt"
	"os"

	"github.com/josephrodriguez/kustomizegen/serialization"
	"github.com/josephrodriguez/kustomizegen/types"
	"github.com/spf13/cobra"
)

func Configure(cmd *cobra.Command, args []string) {
	namespace, _ := cmd.Flags().GetString("namespace")
	unsetOnly, _ := cmd.Flags().GetBool("unsetOnly")

	nsTransformer := types.NewNamespaceTransformer(namespace, unsetOnly)

	// Serialize the nsTransformer to a YAML string
	serializedTransformer, err := serialization.MarshalToYAML(nsTransformer)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	kustomization := types.Kustomization{
		APIVersion: "kustomize.config.k8s.io/v1beta1",
		Kind:       "Kustomization",
		Metadata:   types.Metadata{},
	}

	// Add the serialized nsTransformer to the Transformers array of Kustomization
	kustomization.Transformers = append(kustomization.Transformers, serializedTransformer)

	kustomizationYAML, err := serialization.MarshalToYAML(kustomization)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(kustomizationYAML)
}
