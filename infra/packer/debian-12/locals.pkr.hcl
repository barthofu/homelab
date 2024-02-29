locals {

  # Use local ISO file if provided
  use_iso_file = var.iso_file != null ? true : false

  # Useful when behind NAT or port forwarding scenarios
  http_url = join("", ["http://", coalesce(var.http_server_host, "{{ .HTTPIP }}"), ":", coalesce(var.http_server_port, "{{ .HTTPPort }}")])

}