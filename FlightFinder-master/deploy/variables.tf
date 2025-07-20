variable "region" {
  type    = string
  default = "us-east-1" # cheapest region
}

variable "ami" {
  description = "AMI found on https://cloud-images.ubuntu.com/locator/ec2/"
  type        = string
  default     = "ami-03ff931c79d0e2c80" # us-east-1	impish	21.10	amd64	hvm:ebs-ssd	20220201	ami-03ff931c79d0e2c80	hvm"
}

variable "key_pair_name" {
  description = "EC2 KeyPair name found on AWS Portal under EC2/Network & Security/Key Pairs"
  type        = string
  default     = ""
}
