
terraform {
  required_version = "~> 1.0"
  required_providers {
    aws = {
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = var.region
  # the below keys are needed if you dont have aws cli configured ("aws configure")
  # access_key = "ACCESS_KEY_HERE"
  # secret_key = "SECRET_KEY_HERE"
}

