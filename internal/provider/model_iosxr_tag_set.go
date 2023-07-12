// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TagSet struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	SetName   types.String `tfsdk:"set_name"`
	RplTagSet types.String `tfsdk:"rpl_tag_set"`
}

type TagSetData struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	SetName   types.String `tfsdk:"set_name"`
	RplTagSet types.String `tfsdk:"rpl_tag_set"`
}

func (data TagSet) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/tag-sets/tag-set[set-name=%s]", data.SetName.ValueString())
}

func (data TagSetData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/tag-sets/tag-set[set-name=%s]", data.SetName.ValueString())
}

func (data TagSet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.SetName.IsNull() && !data.SetName.IsUnknown() {
		body, _ = sjson.Set(body, "set-name", data.SetName.ValueString())
	}
	if !data.RplTagSet.IsNull() && !data.RplTagSet.IsUnknown() {
		body, _ = sjson.Set(body, "rpl-tag-set", data.RplTagSet.ValueString())
	}
	return body
}

func (data *TagSet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-tag-set"); value.Exists() && !data.RplTagSet.IsNull() {
		data.RplTagSet = types.StringValue(value.String())
	} else {
		data.RplTagSet = types.StringNull()
	}
}

func (data *TagSetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "rpl-tag-set"); value.Exists() {
		data.RplTagSet = types.StringValue(value.String())
	}
}

func (data *TagSet) getDeletedListItems(ctx context.Context, state TagSet) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *TagSet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *TagSet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.RplTagSet.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/rpl-tag-set", data.getPath()))
	}
	return deletePaths
}