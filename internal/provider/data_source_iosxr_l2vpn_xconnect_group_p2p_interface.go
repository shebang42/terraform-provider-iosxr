// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceL2VPNXconnectGroupP2PInterfaceType struct{}

func (t dataSourceL2VPNXconnectGroupP2PInterfaceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the L2VPN Xconnect Group P2P Interface configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"group_name": {
				MarkdownDescription: "Specify the group the cross connects belong to",
				Type:                types.StringType,
				Computed:            true,
			},
			"p2p_xconnect_name": {
				MarkdownDescription: "Configure point to point cross connect commands",
				Type:                types.StringType,
				Computed:            true,
			},
			"interface_name": {
				MarkdownDescription: "Specify (sub-)interface name to cross connect",
				Type:                types.StringType,
				Required:            true,
			},
		},
	}, nil
}

func (t dataSourceL2VPNXconnectGroupP2PInterfaceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceL2VPNXconnectGroupP2PInterface{
		provider: provider,
	}, diags
}

type dataSourceL2VPNXconnectGroupP2PInterface struct {
	provider provider
}

func (d dataSourceL2VPNXconnectGroupP2PInterface) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state L2VPNXconnectGroupP2PInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.provider.client.Get(ctx, config.Device.Value, config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	state.fromPlan(config)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
