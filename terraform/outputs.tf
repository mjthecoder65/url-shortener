output "vpc_name" {
  value = google_compute_network.vpc.name
}

output "redis_host" {
  value = google_redis_instance.redis.host
}

output "redis_port" {
  value = google_redis_instance.redis.port
}

output "artifact_registry_repo_name" {
  value = google_artifact_registry_repository.repo.name
}

output "service_account_email" {
  value = google_service_account.github_deployer.email
}

output "service_account_key" {
  value     = google_service_account_key.github_deployer_key.private_key
  sensitive = true
}
