variable "domain_name" {
  type    = string
  default = "os4gophers"
}

variable "opensearch_username" {
  type    = string
  default = "opensearch_user"
}

variable "opensearch_password" {
  type      = string
  default   = "W&lcome123"
  sensitive = true
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "aws_opensearch_domain" "opensearch" {

  domain_name    = var.domain_name
  engine_version = "OpenSearch_2.13"

  cluster_config {
    dedicated_master_enabled = true
    dedicated_master_type    = "m6g.large.search"
    dedicated_master_count   = 3
    instance_type            = "r6g.large.search"
    instance_count           = 3
    zone_awareness_enabled   = true
    zone_awareness_config {
      availability_zone_count = 3
    }
    warm_enabled = true
    warm_type    = "ultrawarm1.large.search"
    warm_count   = 2
  }

  advanced_security_options {
    enabled                        = true
    anonymous_auth_enabled         = false
    internal_user_database_enabled = true
    master_user_options {
      master_user_name     = var.opensearch_username
      master_user_password = var.opensearch_password
    }
  }

  domain_endpoint_options {
    enforce_https       = true
    tls_security_policy = "Policy-Min-TLS-1-2-2019-07"
  }

  encrypt_at_rest {
    enabled = true
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 300
    volume_type = "gp3"
    throughput  = 250
  }

  node_to_node_encryption {
    enabled = true
  }

  access_policies = <<CONFIG
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "es:*",
            "Principal": "*",
            "Effect": "Allow",
            "Resource": "arn:aws:es:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:domain/${var.domain_name}/*"
        }
    ]
}
CONFIG
}

output "domain_endpoint" {
  value = "https://${aws_opensearch_domain.opensearch.endpoint}"
}

output "opensearch_dashboards" {
  value = "https://${aws_opensearch_domain.opensearch.dashboard_endpoint}"
}
