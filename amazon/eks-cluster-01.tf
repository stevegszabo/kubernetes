terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "ca-central-1"
}

data "aws_iam_role" "role-cluster" {
  name = "eksClusterRole"
}

data "aws_iam_role" "role-node" {
  name = "eksNodeRole"
}

data "aws_security_group" "this" {
  name = "eksClusterSecurityGroup"
}

data "aws_subnet" "subnet01" {
  id = "subnet-0adfd96b51466392c"
}

data "aws_subnet" "subnet02" {
  id = "subnet-0239cb207ee319ea2"
}

data "aws_subnet" "subnet03" {
  id = "subnet-079b5948f9c8a5765"
}

resource "aws_eks_cluster" "this" {
  name                      = "eks-cluster-01"
  version                   = "1.22"
  role_arn                  = data.aws_iam_role.role-cluster.arn
  enabled_cluster_log_types = []

  vpc_config {
    endpoint_private_access = false
    endpoint_public_access  = true
    security_group_ids      = [data.aws_security_group.this.id]
    public_access_cidrs     = ["0.0.0.0/0"]
    subnet_ids              = [
      data.aws_subnet.subnet01.id,
      data.aws_subnet.subnet02.id,
      data.aws_subnet.subnet03.id
    ]
  }

  kubernetes_network_config {
    ip_family         = "ipv4"
    service_ipv4_cidr = "172.20.0.0/16"
  }

  lifecycle {
    ignore_changes = [
      enabled_cluster_log_types
    ]
  }
}

resource "aws_eks_node_group" "this" {
  node_group_name = "eks-group-01"
  cluster_name    = aws_eks_cluster.this.name
  node_role_arn   = data.aws_iam_role.role-node.arn
  ami_type        = "AL2_x86_64"
  capacity_type   = "ON_DEMAND"
  disk_size       = "20"
  instance_types  = ["t3.medium"]
  subnet_ids      = [
    data.aws_subnet.subnet01.id,
    data.aws_subnet.subnet02.id,
    data.aws_subnet.subnet03.id
  ]

  remote_access {
    ec2_ssh_key = "gamepc"
  }

  scaling_config {
    desired_size = 1
    max_size     = 3
    min_size     = 1
  }

  update_config {
    max_unavailable = 1
  }

  lifecycle {
    ignore_changes = [scaling_config[0].desired_size]
  }
}

output "eks_endpoint" {
  value = aws_eks_cluster.this.endpoint
}

output "aws_eks_node_group" {
  value = aws_eks_node_group.this.id
}
