variable "private_ips" {
  default = []
}

variable "ssh_key_file" {
}

variable "ssh_user" {
  default = "centos"
}

variable "ssh_user_sudo_password" {
}

variable "sudo_cmd" {
  default = "sudo"
}

variable "chef_ips" {
  default = [] 
}

variable "automate-fqdn" {
  default = ""
}
