package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/netascode/terraform-provider-iosxr/internal/provider"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// Donwload YANG models.
//go:generate go run gen/load_models.go

// Run the resource and datasource generation tool.
//go:generate go run gen/generator.go

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

// Update documentation categories.
//go:generate go run gen/doc_category.go

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	opts := tfsdk.ServeOpts{
		Name: "registry.terraform.io/netascode/iosxr",
	}

	err := tfsdk.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
