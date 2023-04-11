---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_evpn Resource - terraform-provider-iosxr"
subcategory: "EVPN"
description: |-
  This resource can manage the EVPN configuration.
---

# iosxr_evpn (Resource)

This resource can manage the EVPN configuration.

## Example Usage

```terraform
resource "iosxr_evpn" "example" {
  source_interface = "Loopback0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.
- `source_interface` (String) Configure EVPN router-id implicitly through Loopback Interface

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_evpn.example "Cisco-IOS-XR-um-l2vpn-cfg:evpn"
```