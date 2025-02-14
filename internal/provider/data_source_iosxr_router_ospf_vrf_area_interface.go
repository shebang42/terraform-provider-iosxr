// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RouterOSPFVRFAreaInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &RouterOSPFVRFAreaInterfaceDataSource{}
)

func NewRouterOSPFVRFAreaInterfaceDataSource() datasource.DataSource {
	return &RouterOSPFVRFAreaInterfaceDataSource{}
}

type RouterOSPFVRFAreaInterfaceDataSource struct {
	client *client.Client
}

func (d *RouterOSPFVRFAreaInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_router_ospf_vrf_area_interface"
}

func (d *RouterOSPFVRFAreaInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Router OSPF VRF Area Interface configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"process_name": schema.StringAttribute{
				MarkdownDescription: "Name for this OSPF process",
				Required:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "Name for this OSPF vrf",
				Required:            true,
			},
			"area_id": schema.StringAttribute{
				MarkdownDescription: "Enter the OSPF area configuration submode",
				Required:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Enable routing on an interface ",
				Required:            true,
			},
			"network_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF broadcast multi-access network",
				Computed:            true,
			},
			"network_non_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF NBMA network",
				Computed:            true,
			},
			"network_point_to_point": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF point-to-point network",
				Computed:            true,
			},
			"network_point_to_multipoint": schema.BoolAttribute{
				MarkdownDescription: "Specify OSPF point-to-multipoint network",
				Computed:            true,
			},
			"cost": schema.Int64Attribute{
				MarkdownDescription: "Interface cost",
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Router priority",
				Computed:            true,
			},
			"passive_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable passive",
				Computed:            true,
			},
			"passive_disable": schema.BoolAttribute{
				MarkdownDescription: "Disable passive",
				Computed:            true,
			},
		},
	}
}

func (d *RouterOSPFVRFAreaInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RouterOSPFVRFAreaInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RouterOSPFVRFAreaInterfaceData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, err := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Unable to apply gNMI Get operation", err.Error())
		return
	}

	config.fromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
