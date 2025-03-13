variable "project_id" {
  description = "The project ID to deploy resources to"
  type        = string
  default     = "rock-elevator-453623-f5"
}

variable "region" {
  description = "The region to deploy resources to (Asia/Seoul Region)"
  type        = string
  default     = "asia-northeast3"
}

variable "vpc_name" {
  description = "Name of the VPC Network"
  type        = string
  default     = "netflix-and-chill"
}

variable "subnet_cidr" {
  description = "CIDR block for the general subnet"
  type        = string
  default     = "10.0.0.0/24"
}

variable "connector_cidr" {
  description = "CIDR block for the Serverless VPC Connector subnet (/28 required)"
  type        = string
  default     = "10.0.2.0/28"
}

variable "redis_size" {
  description = "Size of the Redist Instance in GB"
  type        = number
  default     = 1
}

variable "artifact_registry_repo_name" {
  description = "Name of the Artifact Registry repository"
  type        = string
  default     = "short-url"
}
