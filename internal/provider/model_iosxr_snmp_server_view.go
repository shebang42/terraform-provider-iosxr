// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SNMPServerView struct {
	Device          types.String                    `tfsdk:"device"`
	Id              types.String                    `tfsdk:"id"`
	DeleteMode      types.String                    `tfsdk:"delete_mode"`
	ViewName        types.String                    `tfsdk:"view_name"`
	MibViewFamilies []SNMPServerViewMibViewFamilies `tfsdk:"mib_view_families"`
}
type SNMPServerViewData struct {
	Device          types.String                    `tfsdk:"device"`
	Id              types.String                    `tfsdk:"id"`
	ViewName        types.String                    `tfsdk:"view_name"`
	MibViewFamilies []SNMPServerViewMibViewFamilies `tfsdk:"mib_view_families"`
}
type SNMPServerViewMibViewFamilies struct {
	Name     types.String `tfsdk:"name"`
	Included types.Bool   `tfsdk:"included"`
	Excluded types.Bool   `tfsdk:"excluded"`
}

func (data SNMPServerView) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server/views/view[view-name=%s]", data.ViewName.ValueString())
}

func (data SNMPServerViewData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server/views/view[view-name=%s]", data.ViewName.ValueString())
}

func (data SNMPServerView) toBody(ctx context.Context) string {
	body := "{}"
	if !data.ViewName.IsNull() && !data.ViewName.IsUnknown() {
		body, _ = sjson.Set(body, "view-name", data.ViewName.ValueString())
	}
	if len(data.MibViewFamilies) > 0 {
		body, _ = sjson.Set(body, "mib-view-families.mib-view-family", []interface{}{})
		for index, item := range data.MibViewFamilies {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, "mib-view-families.mib-view-family"+"."+strconv.Itoa(index)+"."+"mib-view-family-name", item.Name.ValueString())
			}
			if !item.Included.IsNull() && !item.Included.IsUnknown() {
				if item.Included.ValueBool() {
					body, _ = sjson.Set(body, "mib-view-families.mib-view-family"+"."+strconv.Itoa(index)+"."+"included", map[string]string{})
				}
			}
			if !item.Excluded.IsNull() && !item.Excluded.IsUnknown() {
				if item.Excluded.ValueBool() {
					body, _ = sjson.Set(body, "mib-view-families.mib-view-family"+"."+strconv.Itoa(index)+"."+"excluded", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *SNMPServerView) updateFromBody(ctx context.Context, res []byte) {
	for i := range data.MibViewFamilies {
		keys := [...]string{"mib-view-family-name"}
		keyValues := [...]string{data.MibViewFamilies[i].Name.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "mib-view-families.mib-view-family").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("mib-view-family-name"); value.Exists() && !data.MibViewFamilies[i].Name.IsNull() {
			data.MibViewFamilies[i].Name = types.StringValue(value.String())
		} else {
			data.MibViewFamilies[i].Name = types.StringNull()
		}
		if value := r.Get("included"); !data.MibViewFamilies[i].Included.IsNull() {
			if value.Exists() {
				data.MibViewFamilies[i].Included = types.BoolValue(true)
			} else {
				data.MibViewFamilies[i].Included = types.BoolValue(false)
			}
		} else {
			data.MibViewFamilies[i].Included = types.BoolNull()
		}
		if value := r.Get("excluded"); !data.MibViewFamilies[i].Excluded.IsNull() {
			if value.Exists() {
				data.MibViewFamilies[i].Excluded = types.BoolValue(true)
			} else {
				data.MibViewFamilies[i].Excluded = types.BoolValue(false)
			}
		} else {
			data.MibViewFamilies[i].Excluded = types.BoolNull()
		}
	}
}

func (data *SNMPServerViewData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "mib-view-families.mib-view-family"); value.Exists() {
		data.MibViewFamilies = make([]SNMPServerViewMibViewFamilies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SNMPServerViewMibViewFamilies{}
			if cValue := v.Get("mib-view-family-name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("included"); cValue.Exists() {
				item.Included = types.BoolValue(true)
			} else {
				item.Included = types.BoolValue(false)
			}
			if cValue := v.Get("excluded"); cValue.Exists() {
				item.Excluded = types.BoolValue(true)
			} else {
				item.Excluded = types.BoolValue(false)
			}
			data.MibViewFamilies = append(data.MibViewFamilies, item)
			return true
		})
	}
}

func (data *SNMPServerView) getDeletedListItems(ctx context.Context, state SNMPServerView) []string {
	deletedListItems := make([]string, 0)
	for i := range state.MibViewFamilies {
		keys := [...]string{"mib-view-family-name"}
		stateKeyValues := [...]string{state.MibViewFamilies[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.MibViewFamilies[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MibViewFamilies {
			found = true
			if state.MibViewFamilies[i].Name.ValueString() != data.MibViewFamilies[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/mib-view-families/mib-view-family%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *SNMPServerView) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.MibViewFamilies {
		keys := [...]string{"mib-view-family-name"}
		keyValues := [...]string{data.MibViewFamilies[i].Name.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		if !data.MibViewFamilies[i].Included.IsNull() && !data.MibViewFamilies[i].Included.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/mib-view-families/mib-view-family%v/included", data.getPath(), keyString))
		}
		if !data.MibViewFamilies[i].Excluded.IsNull() && !data.MibViewFamilies[i].Excluded.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/mib-view-families/mib-view-family%v/excluded", data.getPath(), keyString))
		}
	}
	return emptyLeafsDelete
}

func (data *SNMPServerView) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.MibViewFamilies {
		keys := [...]string{"mib-view-family-name"}
		keyValues := [...]string{data.MibViewFamilies[i].Name.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/mib-view-families/mib-view-family%v", data.getPath(), keyString))
	}
	return deletePaths
}
