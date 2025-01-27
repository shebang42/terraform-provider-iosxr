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
	_ datasource.DataSource              = &SNMPServerDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPServerDataSource{}
)

func NewSNMPServerDataSource() datasource.DataSource {
	return &SNMPServerDataSource{}
}

type SNMPServerDataSource struct {
	client *client.Client
}

func (d *SNMPServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_server"
}

func (d *SNMPServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the SNMP Server configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Text for mib Object sysLocation",
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: "Text for mib Object sysContact",
				Computed:            true,
			},
			"traps_rf": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP RF-MIB traps",
				Computed:            true,
			},
			"traps_bfd": schema.BoolAttribute{
				MarkdownDescription: "Enable BFD traps",
				Computed:            true,
			},
			"traps_ntp": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Cisco Ntp traps",
				Computed:            true,
			},
			"traps_ethernet_oam_events": schema.BoolAttribute{
				MarkdownDescription: "Enable all OAM event traps",
				Computed:            true,
			},
			"traps_copy_complete": schema.BoolAttribute{
				MarkdownDescription: "Enable CISCO-CONFIG-COPY-MIB ccCopyCompletion traps",
				Computed:            true,
			},
			"traps_snmp_linkup": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMPv2-MIB linkUp traps",
				Computed:            true,
			},
			"traps_snmp_linkdown": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMPv2-MIB linDownp traps",
				Computed:            true,
			},
			"traps_power": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity power traps",
				Computed:            true,
			},
			"traps_config": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP config traps",
				Computed:            true,
			},
			"traps_entity": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity traps",
				Computed:            true,
			},
			"traps_system": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP SYSTEMMIB-MIB traps",
				Computed:            true,
			},
			"traps_bridgemib": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP Trap for Bridge MIB",
				Computed:            true,
			},
			"traps_entity_state_operstatus": schema.BoolAttribute{
				MarkdownDescription: "Enable entity oper status enable notification",
				Computed:            true,
			},
			"traps_entity_redundancy_all": schema.BoolAttribute{
				MarkdownDescription: "Enable all CISCO-ENTITY-REDUNDANCY-MIB traps",
				Computed:            true,
			},
			"trap_source_both": schema.StringAttribute{
				MarkdownDescription: "Assign an interface for the source address of all traps",
				Computed:            true,
			},
			"traps_l2vpn_all": schema.BoolAttribute{
				MarkdownDescription: "Enable all L2VPN traps",
				Computed:            true,
			},
			"traps_l2vpn_vc_up": schema.BoolAttribute{
				MarkdownDescription: "Enable VC up traps",
				Computed:            true,
			},
			"traps_l2vpn_vc_down": schema.BoolAttribute{
				MarkdownDescription: "Enable VC down traps",
				Computed:            true,
			},
			"traps_sensor": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity sensor traps",
				Computed:            true,
			},
			"traps_fru_ctrl": schema.BoolAttribute{
				MarkdownDescription: "Enable SNMP entity FRU control traps",
				Computed:            true,
			},
			"traps_isis_all": schema.StringAttribute{
				MarkdownDescription: "Enable all IS-IS traps",
				Computed:            true,
			},
			"traps_isis_database_overload": schema.StringAttribute{
				MarkdownDescription: "isisDatabaseOverload",
				Computed:            true,
			},
			"traps_isis_manual_address_drops": schema.StringAttribute{
				MarkdownDescription: "isisManualAddressDrops",
				Computed:            true,
			},
			"traps_isis_corrupted_lsp_detected": schema.StringAttribute{
				MarkdownDescription: "isisCorruptedLSPDetected",
				Computed:            true,
			},
			"traps_isis_attempt_to_exceed_max_sequence": schema.StringAttribute{
				MarkdownDescription: "isisAttemptToExceedMaxSequence",
				Computed:            true,
			},
			"traps_isis_id_len_mismatch": schema.StringAttribute{
				MarkdownDescription: "isisIDLenMismatch",
				Computed:            true,
			},
			"traps_isis_max_area_addresses_mismatch": schema.StringAttribute{
				MarkdownDescription: "isisMaxAreaAddressesMismatch",
				Computed:            true,
			},
			"traps_isis_own_lsp_purge": schema.StringAttribute{
				MarkdownDescription: "isisOwnLSPPurge",
				Computed:            true,
			},
			"traps_isis_sequence_number_skip": schema.StringAttribute{
				MarkdownDescription: "isisSequenceNumberSkip",
				Computed:            true,
			},
			"traps_isis_authentication_type_failure": schema.StringAttribute{
				MarkdownDescription: "isisAuthenticationTypeFailure",
				Computed:            true,
			},
			"traps_isis_authentication_failure": schema.StringAttribute{
				MarkdownDescription: "isisAuthenticationFailure",
				Computed:            true,
			},
			"traps_isis_version_skew": schema.StringAttribute{
				MarkdownDescription: "isisVersionSkew",
				Computed:            true,
			},
			"traps_isis_area_mismatch": schema.StringAttribute{
				MarkdownDescription: "isisAreaMismatch",
				Computed:            true,
			},
			"traps_isis_rejected_adjacency": schema.StringAttribute{
				MarkdownDescription: "isisRejectedAdjacency",
				Computed:            true,
			},
			"traps_isis_lsp_too_large_to_propagate": schema.StringAttribute{
				MarkdownDescription: "isisLSPTooLargeToPropagate",
				Computed:            true,
			},
			"traps_isis_orig_lsp_buff_size_mismatch": schema.StringAttribute{
				MarkdownDescription: "isisOrigLSPBuffSizeMismatch",
				Computed:            true,
			},
			"traps_isis_protocols_supported_mismatch": schema.StringAttribute{
				MarkdownDescription: "isisProtocolsSupportedMismatch",
				Computed:            true,
			},
			"traps_isis_adjacency_change": schema.StringAttribute{
				MarkdownDescription: "isisAdjacencyChange",
				Computed:            true,
			},
			"traps_isis_lsp_error_detected": schema.StringAttribute{
				MarkdownDescription: "isisLSPErrorDetected",
				Computed:            true,
			},
			"traps_bgp_cbgp2_updown": schema.BoolAttribute{
				MarkdownDescription: "Enable CISCO-BGP4-MIB v2 up/down traps",
				Computed:            true,
			},
			"traps_bgp_bgp4_mib_updown": schema.BoolAttribute{
				MarkdownDescription: "Enable CISCO-BGP4-MIB v2 up/down traps",
				Computed:            true,
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: "Context Name",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"context_name": schema.StringAttribute{
							MarkdownDescription: "Context Name",
							Computed:            true,
						},
					},
				},
			},
			"vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "VRF name",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "VRF name",
							Computed:            true,
						},
						"context": schema.StringAttribute{
							MarkdownDescription: "SNMP Context Name",
							Computed:            true,
						},
					},
				},
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: "Name of the user",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Name of the user",
							Computed:            true,
						},
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Group to which the user belongs",
							Computed:            true,
						},
						"v3_auth_md5_encryption_aes": schema.StringAttribute{
							MarkdownDescription: "Specifies an aes-128 ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_auth_md5_encryption_default": schema.StringAttribute{
							MarkdownDescription: "Specifies an default ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_auth_sha_encryption_aes": schema.StringAttribute{
							MarkdownDescription: "Specifies an aes-128 ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_auth_sha_encryption_default": schema.StringAttribute{
							MarkdownDescription: "Specifies an default ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_priv_aes_aes_128_encryption_default": schema.StringAttribute{
							MarkdownDescription: "Specifies an default ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_priv_aes_aes_128_encryption_aes": schema.StringAttribute{
							MarkdownDescription: "Specifies an aes-128 ENCRYPTED authentication password",
							Computed:            true,
						},
						"v3_ipv4": schema.StringAttribute{
							MarkdownDescription: "Type of Access-list",
							Computed:            true,
						},
						"v3_systemowner": schema.BoolAttribute{
							MarkdownDescription: "System Owner permissions for MIB objects",
							Computed:            true,
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "Name of the group",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Name of the group",
							Computed:            true,
						},
						"v3_priv": schema.BoolAttribute{
							MarkdownDescription: "group using authPriv Security Level",
							Computed:            true,
						},
						"v3_read": schema.StringAttribute{
							MarkdownDescription: "specify a read view for this group",
							Computed:            true,
						},
						"v3_write": schema.StringAttribute{
							MarkdownDescription: "specify a write view for this group",
							Computed:            true,
						},
						"v3_context": schema.StringAttribute{
							MarkdownDescription: "Attach a SNMP context",
							Computed:            true,
						},
						"v3_notify": schema.StringAttribute{
							MarkdownDescription: "specify a notify view for the group",
							Computed:            true,
						},
						"v3_ipv4": schema.StringAttribute{
							MarkdownDescription: "Type of Access-list",
							Computed:            true,
						},
						"v3_ipv6": schema.StringAttribute{
							MarkdownDescription: "Type of Access-list",
							Computed:            true,
						},
					},
				},
			},
			"communities": schema.ListNestedAttribute{
				MarkdownDescription: "The UNENCRYPTED (cleartext) community string",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"community": schema.StringAttribute{
							MarkdownDescription: "The UNENCRYPTED (cleartext) community string",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Restrict this community to a named view",
							Computed:            true,
						},
						"ro": schema.BoolAttribute{
							MarkdownDescription: "Read-only community",
							Computed:            true,
						},
						"rw": schema.BoolAttribute{
							MarkdownDescription: "Read-write community",
							Computed:            true,
						},
						"sdrowner": schema.BoolAttribute{
							MarkdownDescription: "SDR Owner permissions for MIB Objects",
							Computed:            true,
						},
						"systemowner": schema.BoolAttribute{
							MarkdownDescription: "System Owner permissions for MIB objects",
							Computed:            true,
						},
						"ipv4": schema.StringAttribute{
							MarkdownDescription: "Type of Access-list",
							Computed:            true,
						},
						"ipv6": schema.StringAttribute{
							MarkdownDescription: "Type of Access-list",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SNMPServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *SNMPServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SNMPServerData

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
