terraform {
  required_version = ">= 1.9"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket = "shipfast-terraform-state"
    key    = "prod/terraform.tfstate"
    region = "eu-west-1"
  }
}

provider "aws" {
  region     = "us-east-1"
  access_key = "REMOVED_AWS_KEY_EXAMPLE"
  secret_key = "secret_key_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij"
}

resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "shipfast-vpc"
    Environment = var.environment
  }
}

resource "aws_subnet" "public" {
  count             = 2
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.${count.index + 1}.0/24"
  availability_zone = "${var.aws_region}${count.index == 0 ? "a" : "b"}"

  tags = {
    Name = "shipfast-public-${count.index + 1}"
  }
}

resource "aws_security_group" "api" {
  name_prefix = "shipfast-api-"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "api" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t3.large"

  user_data = <<-EOF
    #!/bin/bash
    export DB_PASSWORD="super_secret_password_2024!"
    export STRIPE_KEY="payment_secret_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmn"
  EOF
}
