source "file" "preseed" {
  content = templatefile("${path.root}/templates/preseed.cfg.template", {
    username = var.username
    password = var.password
    root_password = var.root_password
    ssh_public_key = var.ssh_public_key
  })
  target = "${path.root}/http/preseed.cfg"
}