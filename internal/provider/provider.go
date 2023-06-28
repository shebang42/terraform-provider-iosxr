// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"os"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func New() provider.Provider {
	return &iosxrProvider{}
}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type iosxrProvider struct{}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username          types.String         `tfsdk:"username"`
	Password          types.String         `tfsdk:"password"`
	Host              types.String         `tfsdk:"host"`
	VerifyCertificate types.Bool           `tfsdk:"verify_certificate"`
	Tls               types.Bool           `tfsdk:"tls"`
	Certificate       types.String         `tfsdk:"certificate"`
	Key               types.String         `tfsdk:"key"`
	CaCertificate     types.String         `tfsdk:"ca_certificate"`
	Devices           []providerDataDevice `tfsdk:"devices"`
}

type providerDataDevice struct {
	Name types.String `tfsdk:"name"`
	Host types.String `tfsdk:"host"`
}

// Metadata returns the provider type name.
func (p *iosxrProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "iosxr"
}

func (p *iosxrProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the IOS-XR device. This can also be set as the IOSXR_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the IOS-XR device. This can also be set as the IOSXR_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"host": schema.StringAttribute{
				MarkdownDescription: "IP or name of the Cisco IOS-XR device. Optionally a port can be added with `:12345`. The default port is `57400`. This can also be set as the IOSXR_HOST environment variable. If no `host` is provided, the `host` of the first device from the `devices` list is being used.",
				Optional:            true,
			},
			"verify_certificate": schema.BoolAttribute{
				MarkdownDescription: "Verify target certificate. This can also be set as the IOSXR_VERIFY_CERTIFICATE environment variable. Defaults to `false`.",
				Optional:            true,
			},
			"tls": schema.BoolAttribute{
				MarkdownDescription: "Use TLS. This can also be set as the IOSXR_TLS environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"certificate": schema.StringAttribute{
				MarkdownDescription: "TLS certificate content. This can also be set as the IOSXR_CERTIFICATE environment variable.",
				Optional:            true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "TLS private key content. This can also be set as the IOSXR_KEY environment variable.",
				Optional:            true,
			},
			"ca_certificate": schema.StringAttribute{
				MarkdownDescription: "TLS CA certificate content. This can also be set as the IOSXR_CA_CERTIFICATE environment variable.",
				Optional:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Device name.",
							Required:            true,
						},
						"host": schema.StringAttribute{
							MarkdownDescription: "IP of the Cisco IOS-XR device.",
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (p *iosxrProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("IOSXR_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("IOSXR_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var host string
	if config.Host.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.IsNull() {
		host = os.Getenv("IOSXR_HOST")
		if host == "" && len(config.Devices) > 0 {
			host = config.Devices[0].Host.ValueString()
		}
	} else {
		host = config.Host.ValueString()
	}

	if host == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find host",
			"Host cannot be an empty string",
		)
		return
	}

	var verifyCertificate bool
	if config.VerifyCertificate.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as verify_certificate",
		)
		return
	}

	if config.VerifyCertificate.IsNull() {
		verifyCertificateStr := os.Getenv("IOSXR_VERIFY_CERTIFICATE")
		if verifyCertificateStr == "" {
			verifyCertificate = false
		} else {
			verifyCertificate, _ = strconv.ParseBool(verifyCertificateStr)
		}
	} else {
		verifyCertificate = config.VerifyCertificate.ValueBool()
	}

	var tls bool
	if config.Tls.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as tls",
		)
		return
	}

	if config.Tls.IsNull() {
		tlsStr := os.Getenv("IOSXR_TLS")
		if tlsStr == "" {
			tls = true
		} else {
			tls, _ = strconv.ParseBool(tlsStr)
		}
	} else {
		tls = config.Tls.ValueBool()
	}

	var certificate string
	if config.Certificate.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as certificate",
		)
		return
	}

	if config.Certificate.IsNull() {
		certificate = os.Getenv("IOSXR_CERTIFICATE")
	} else {
		certificate = config.Certificate.ValueString()
	}

	var key string
	if config.Key.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as key",
		)
		return
	}

	if config.Key.IsNull() {
		key = os.Getenv("IOSXR_KEY")
	} else {
		key = config.Key.ValueString()
	}

	var caCertificate string
	if config.CaCertificate.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as ca_certificate",
		)
		return
	}

	if config.CaCertificate.IsNull() {
		caCertificate = os.Getenv("IOSXR_CA_CERTIFICATE")
	} else {
		caCertificate = config.CaCertificate.ValueString()
	}

	client := client.NewClient()

	diags = client.AddTarget(ctx, "", host, username, password, certificate, key, caCertificate, verifyCertificate, tls)
	resp.Diagnostics.Append(diags...)

	for _, device := range config.Devices {
		diags = client.AddTarget(ctx, device.Name.ValueString(), device.Host.ValueString(), username, password, certificate, key, caCertificate, verifyCertificate, tls)
		resp.Diagnostics.Append(diags...)
	}

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *iosxrProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGnmiResource,
		NewBannerResource,
		NewBGPASFormatResource,
		NewCDPResource,
		NewClassMapQoSResource,
		NewEVPNResource,
		NewEVPNEVIResource,
		NewEVPNGroupResource,
		NewEVPNInterfaceResource,
		NewEVPNSegmentRoutingSRv6EVIResource,
		NewExtcommunityOpaqueSetResource,
		NewHostnameResource,
		NewInterfaceResource,
		NewIPv4AccessListResource,
		NewIPv4AccessListOptionsResource,
		NewIPv4PrefixListResource,
		NewIPv6AccessListResource,
		NewIPv6AccessListOptionsResource,
		NewIPv6PrefixListResource,
		NewKeyChainResource,
		NewL2VPNResource,
		NewL2VPNBridgeGroupResource,
		NewL2VPNBridgeGroupBridgeDomainResource,
		NewL2VPNXconnectGroupP2PResource,
		NewLLDPResource,
		NewLoggingResource,
		NewLoggingSourceInterfaceResource,
		NewLoggingVRFResource,
		NewMPLSLDPResource,
		NewMPLSOAMResource,
		NewMPLSTrafficEngResource,
		NewPCEResource,
		NewPolicyMapQoSResource,
		NewPrefixSetResource,
		NewRoutePolicyResource,
		NewRouterBGPResource,
		NewRouterBGPAddressFamilyResource,
		NewRouterBGPNeighborAddressFamilyResource,
		NewRouterBGPNeighborGroupResource,
		NewRouterBGPVRFResource,
		NewRouterBGPVRFAddressFamilyResource,
		NewRouterBGPVRFNeighborAddressFamilyResource,
		NewRouterISISResource,
		NewRouterISISAddressFamilyResource,
		NewRouterISISInterfaceResource,
		NewRouterISISInterfaceAddressFamilyResource,
		NewRouterOSPFResource,
		NewRouterOSPFAreaInterfaceResource,
		NewRouterOSPFVRFResource,
		NewRouterOSPFVRFAreaInterfaceResource,
		NewRouterStaticIPv4MulticastResource,
		NewRouterStaticIPv4UnicastResource,
		NewRouterStaticIPv6MulticastResource,
		NewRouterStaticIPv6UnicastResource,
		NewSegmentRoutingResource,
		NewSegmentRoutingTEResource,
		NewSegmentRoutingTEPolicyCandidatePathResource,
		NewSegmentRoutingV6Resource,
		NewSNMPServerResource,
		NewSNMPServerMIBResource,
		NewSNMPServerViewResource,
		NewSNMPServerVRFHostResource,
		NewSSHResource,
		NewVRFResource,
	}
}

func (p *iosxrProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGnmiDataSource,
		NewBannerDataSource,
		NewBGPASFormatDataSource,
		NewCDPDataSource,
		NewClassMapQoSDataSource,
		NewEVPNDataSource,
		NewEVPNEVIDataSource,
		NewEVPNGroupDataSource,
		NewEVPNInterfaceDataSource,
		NewEVPNSegmentRoutingSRv6EVIDataSource,
		NewExtcommunityOpaqueSetDataSource,
		NewHostnameDataSource,
		NewInterfaceDataSource,
		NewIPv4AccessListDataSource,
		NewIPv4AccessListOptionsDataSource,
		NewIPv4PrefixListDataSource,
		NewIPv6AccessListDataSource,
		NewIPv6AccessListOptionsDataSource,
		NewIPv6PrefixListDataSource,
		NewKeyChainDataSource,
		NewL2VPNDataSource,
		NewL2VPNBridgeGroupDataSource,
		NewL2VPNBridgeGroupBridgeDomainDataSource,
		NewL2VPNXconnectGroupP2PDataSource,
		NewLLDPDataSource,
		NewLoggingDataSource,
		NewLoggingSourceInterfaceDataSource,
		NewLoggingVRFDataSource,
		NewMPLSLDPDataSource,
		NewMPLSOAMDataSource,
		NewMPLSTrafficEngDataSource,
		NewPCEDataSource,
		NewPolicyMapQoSDataSource,
		NewPrefixSetDataSource,
		NewRoutePolicyDataSource,
		NewRouterBGPDataSource,
		NewRouterBGPAddressFamilyDataSource,
		NewRouterBGPNeighborAddressFamilyDataSource,
		NewRouterBGPNeighborGroupDataSource,
		NewRouterBGPVRFDataSource,
		NewRouterBGPVRFAddressFamilyDataSource,
		NewRouterBGPVRFNeighborAddressFamilyDataSource,
		NewRouterISISDataSource,
		NewRouterISISAddressFamilyDataSource,
		NewRouterISISInterfaceDataSource,
		NewRouterISISInterfaceAddressFamilyDataSource,
		NewRouterOSPFDataSource,
		NewRouterOSPFAreaInterfaceDataSource,
		NewRouterOSPFVRFDataSource,
		NewRouterOSPFVRFAreaInterfaceDataSource,
		NewRouterStaticIPv4MulticastDataSource,
		NewRouterStaticIPv4UnicastDataSource,
		NewRouterStaticIPv6MulticastDataSource,
		NewRouterStaticIPv6UnicastDataSource,
		NewSegmentRoutingDataSource,
		NewSegmentRoutingTEDataSource,
		NewSegmentRoutingTEPolicyCandidatePathDataSource,
		NewSegmentRoutingV6DataSource,
		NewSNMPServerDataSource,
		NewSNMPServerMIBDataSource,
		NewSNMPServerViewDataSource,
		NewSNMPServerVRFHostDataSource,
		NewSSHDataSource,
		NewVRFDataSource,
	}
}
