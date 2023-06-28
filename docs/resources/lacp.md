---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_lacp Resource - terraform-provider-iosxr"
subcategory: "System"
description: |-
  This resource can manage the LACP configuration.
---

# iosxr_lacp (Resource)

This resource can manage the LACP configuration.

## Example Usage

```terraform
resource "iosxr_lacp" "example" {
  mac      = "00:11:00:11:00:11"
  priority = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `mac` (String) The system ID to use in LACP negotiations.
- `priority` (Number) The system priority to use in LACP negotiations.
  - Range: `1`-`65535`

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_lacp.example "Cisco-IOS-XR-um-lacp-cfg:/lacp/system"
```