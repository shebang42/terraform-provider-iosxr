// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RouterOSPFArea struct {
	Device      types.String `tfsdk:"device"`
	Id          types.String `tfsdk:"id"`
	ProcessName types.String `tfsdk:"process_name"`
	AreaId      types.String `tfsdk:"area_id"`
}

func (data RouterOSPFArea) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-ospf-cfg:/router/ospf/processes/process[process-name=%s]/areas/area[area-id=%s]", data.ProcessName.Value, data.AreaId.Value)
}

func (data RouterOSPFArea) toBody() string {
	body := "{}"

	return body
}

func (data *RouterOSPFArea) fromBody(res []byte) {
}

func (data *RouterOSPFArea) fromPlan(plan RouterOSPFArea) {
	data.Device = plan.Device
	data.ProcessName.Value = plan.ProcessName.Value
	data.AreaId.Value = plan.AreaId.Value
}
