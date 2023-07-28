# KustomizeGen

KustomizeGen is a command-line tool that simplifies the generation and management of Kustomize overlays for Kubernetes manifests using the [Kustomize](https://github.com/kubernetes-sigs/kustomize) tool.

## Installation

### Method 1: Go command
To install KustomizeGen, you can use the following command:

```bash
go install github.com/josephrodriguez/kustomizegen@latest
```

### Method 2: Download Binary Release

#### Linux:

1. Download the binary for your Linux architecture using `curl` and extract the downloaded archive:

   ```bash
   curl -LO https://github.com/josephrodriguez/kustomizegen/releases/latest/download/kustomizegen_amd64.tar.gz
   tar -xzvf kustomizegen_amd64.tar.gz
   ```

2. Move the binary to a location included in your system's PATH:

   ```bash
   sudo mv kustomizegen /usr/local/bin/
   ```

3. Now, you can use `kustomizegen` from the terminal.

Please ensure that the destination folder (`/usr/local/bin/`) is included in your system's PATH environment variable. This will allow you to execute `kustomizegen` from any location on your Linux system.

#### Windows:

1. Download the binary for Windows using `curl` and extract the downloaded archive:

   ```powershell
   curl -LO https://github.com/josephrodriguez/kustomizegen/releases/latest/download/kustomizegen_amd64.zip
   Expand-Archive -Path kustomizegen_amd64.zip -DestinationPath .
   ```

2. Move the binary to a location included in your system's PATH:

   ```powershell
   Move-Item -Path .\kustomizegen.exe -Destination C:\Windows\System32
   ```

3. Now, you can use `kustomizegen` from the command prompt.

## Usage

KustomizeGen offers the following commands:

- `generate-overlays`: Creates the Kustomize overlays for your Kubernetes resources.
- `print-build-command`: Generates a shell script with build commands for the configured overlays.
- `destroy-overlays`: Destroys the generated Kustomize overlays.
- `version`: Displays the compiled version of KustomizeGen.

### 1. Generate Overlays Command

This command allows you to generate Kustomization overlays based on the configuration provided. It creates overlay files that you can further customize for different environments. The command takes the base folder path for the Kustomization as a parameter.

```bash
kustomizegen generate-overlays --root /path/to/kustomization/base
```

### 2. Generate Build Command

The "print-build-command" command generates a shell script with build commands for the configured overlays. It provides an option to enable Helm while generating the build script. The command takes the base folder path for the Kustomization as a parameter.

```bash
kustomizegen print-build-command --root /path/to/kustomization/base [--enable-helm]
```

### 3. Destroy Overlays Command

The "destroy-overlays" command destroys the generated Kustomization overlays. It removes the overlay files and cleans up the environment. The command takes the base folder path for the Kustomization as a parameter.

```bash
kustomizegen destroy-overlays --root /path/to/kustomization/base
```

### 4. Version Command

The "version" command displays the compiled version of KustomizeGen.

```bash
kustomizegen version
```

## Contributing

We welcome contributions from the community. If you encounter any issues, have suggestions, or want to contribute to KustomizeGen, feel free to create an issue or submit a pull request on GitHub.

## License

KustomizeGen is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.