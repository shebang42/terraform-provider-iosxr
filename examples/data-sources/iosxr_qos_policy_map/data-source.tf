data "iosxr_qos_policy_map" "example" {
  policy_map_name = "core-ingress-classifier"
  name            = "class-default"
  type            = "qos"
}