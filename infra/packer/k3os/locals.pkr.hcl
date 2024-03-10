locals {
  boot_command_pre = ["<wait>", "<tab>", "<down>", "<wait>", "e", "<down>", "<down>", "<down>", "<down>", "<down>", "<down>", "<end>"]
  boot_command_args = [
    " ", "k3os.install.device=/dev/vda",
    " ", "k3os.mode=install",
    " ", "k3os.install.silent=true",
    " ", "k3os.install.debug=true",
  ]
  boot_command_args_proxmox = [
    " ", "k3os.install.power_off=true",
    " ", "k3os.install.config_url=${var.config_url}",
    " ", "k3os.install.tty=ttyS0"
  ]
  boot_command_post = ["<F10>"]
}