source "file" "answers" {
  content = templatefile("${path.root}/templates/answers.template", {
    dns_servers      = var.dns_servers
  })
  target = "${path.root}/http/answers"
}

source "file" "setup" {
  content = templatefile("${path.root}/templates/setup.sh.template", {
    username      = var.username
    password      = var.password
  })
  target = "${path.root}/http/setup.sh"
}
