---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_esi_set Data Source - terraform-provider-iosxr"
subcategory: "Route Policy"
description: |-
  This data source can read the ESI Set configuration.
---

# iosxr_esi_set (Data Source)

This data source can read the ESI Set configuration.

## Example Usage

```terraform
data "iosxr_esi_set" "example" {
  set_name = "POLICYSET"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `set_name` (String) Set name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `rpl` (String) Esi Set