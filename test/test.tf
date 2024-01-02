provider "google" {}
provider "aws" {}

variable "foo" {
  default = "bar"
}

output "bar" {
  value = var.foo
}
