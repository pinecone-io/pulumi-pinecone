// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pinecone

import (
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/pinecone-io/pulumi-pinecone/provider/v2/pkg/version"
	"github.com/pinecone-io/terraform-provider-pinecone/shim"
	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "pinecone"
	// modules:
	mainMod = "index" // the pinecone module
)

//go:embed cmd/pulumi-resource-pinecone/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	// Create a Pulumi provider mapping
	p := pftfbridge.ShimProvider(shim.NewProvider())
	prov := tfbridge.ProviderInfo{
		P:       p,
		Name:    "pinecone",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Pinecone",
		Description: "A Pulumi package for creating and managing Pinecone resources.",
		Keywords: []string{
			"pulumi",
			"pinecone",
			"category/utility",
			"kind/native",
		},
		Homepage:          "https://www.pinecone.io",
		License:           "Apache-2.0",
		Repository:        "https://github.com/pinecone-io/pulumi-pinecone",
		PluginDownloadURL: "github://api.github.com/pinecone-io/pulumi-pinecone",
		Publisher:         "pinecone-io",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "pinecone-io",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PINECONE_API_KEY"},
				},
			},
			"client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PINECONE_CLIENT_ID"},
				},
			},
			"client_secret": {
				Secret: tfbridge.BoolRef(true),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PINECONE_CLIENT_SECRET"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			"pinecone_index": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Index")},
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pinecone-database/pulumi",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pinecone_pulumi",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pinecone-io/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "PineconeDatabase",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("pinecone_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}
