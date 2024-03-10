source "file" "preseed" {
  content = templatefile("${path.root}/templates/cloud.yml.template", {
    ssh_password = var.ssh_password
    ssh_public_key = var.ssh_public_key
    ssh_ansible_public_key = var.ssh_ansible_public_key
  })
  target = "${path.root}/http/cloud.yml"
}