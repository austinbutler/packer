package command

import (
	"bytes"
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer/builder/file"
	"github.com/hashicorp/packer/builder/null"
	hcppackerimagedatasource "github.com/hashicorp/packer/datasource/hcp-packer-image"
	hcppackeriterationdatasource "github.com/hashicorp/packer/datasource/hcp-packer-iteration"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/post-processor/manifest"
	shell_local_pp "github.com/hashicorp/packer/post-processor/shell-local"
	filep "github.com/hashicorp/packer/provisioner/file"
	"github.com/hashicorp/packer/provisioner/shell"
	shell_local "github.com/hashicorp/packer/provisioner/shell-local"
)

// Utils to use in other tests

// TestMetaFile creates a Meta object that includes a file builder
func TestMetaFile(t *testing.T) Meta {
	var out, err bytes.Buffer
	return Meta{
		CoreConfig: testCoreConfigBuilder(t),
		Ui: &packersdk.BasicUi{
			Writer:      &out,
			ErrorWriter: &err,
		},
	}
}

// testCoreConfigBuilder creates a packer CoreConfig that has a file builder
// available. This allows us to test a builder that writes files to disk.
func testCoreConfigBuilder(t *testing.T) *packer.CoreConfig {
	components := packer.ComponentFinder{
		PluginConfig: &packer.PluginConfig{
			Builders: packer.MapOfBuilder{
				"file": func() (packersdk.Builder, error) { return &file.Builder{}, nil },
				"null": func() (packersdk.Builder, error) { return &null.Builder{}, nil },
			},
			Provisioners: packer.MapOfProvisioner{
				"shell-local": func() (packersdk.Provisioner, error) { return &shell_local.Provisioner{}, nil },
				"shell":       func() (packersdk.Provisioner, error) { return &shell.Provisioner{}, nil },
				"file":        func() (packersdk.Provisioner, error) { return &filep.Provisioner{}, nil },
			},
			PostProcessors: packer.MapOfPostProcessor{
				"shell-local": func() (packersdk.PostProcessor, error) { return &shell_local_pp.PostProcessor{}, nil },
				"manifest":    func() (packersdk.PostProcessor, error) { return &manifest.PostProcessor{}, nil },
			},
			DataSources: packer.MapOfDatasource{
				"mock":                 func() (packersdk.Datasource, error) { return &packersdk.MockDatasource{}, nil },
				"hcp-packer-image":     func() (packersdk.Datasource, error) { return &hcppackerimagedatasource.Datasource{}, nil },
				"hcp-packer-iteration": func() (packersdk.Datasource, error) { return &hcppackeriterationdatasource.Datasource{}, nil },
			},
		},
	}
	return &packer.CoreConfig{
		Components: components,
	}
}
