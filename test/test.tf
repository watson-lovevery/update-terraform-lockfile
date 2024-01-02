provider "google" {}

variable "foo" {
  default = "bar"
}

output "bar" {
  value = var.foo
}
