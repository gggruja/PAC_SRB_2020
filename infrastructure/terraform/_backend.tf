locals {
  backend_username = "backend"
  backend_database = "backend"
}

resource "kubernetes_namespace" "backend" {
  metadata {
    name = "backend"
  }
}


resource "helm_release" "backend" {
  name = "backend"
  chart = "../chart/backend"
  namespace = kubernetes_namespace.backend.metadata[0].name

  depends_on = [
    helm_release.backend-mariadb,
    kubernetes_config_map.backend-config,
    kubernetes_secret.backend-mariadb-access
  ]

  values = [
    file("helm/backend.yaml")
  ]
}

resource "kubernetes_config_map" "backend-config" {
  metadata {
    name = "backend-config"
    namespace = kubernetes_namespace.backend.metadata[0].name
  }

  data = {
    BIND_ADDRESS = ":80"
    DB_DRIVER = "mysql"
    DB_HOST = helm_release.backend-mariadb.metadata[0].name
    DB_PORT = "3306"
    DB_NAME = "backend"
  }
}

resource "kubernetes_cron_job" "backend-init-db" {
  metadata {
    name = "backend-init-db"
    namespace = kubernetes_namespace.backend.metadata[0].name
  }

  depends_on = [
    helm_release.backend
  ]


  spec {
    concurrency_policy            = "Replace"
    failed_jobs_history_limit     = 5
    schedule                      = "*/1 * * * *"
    starting_deadline_seconds     = 10
    successful_jobs_history_limit = 10
    suspend                       = true
    job_template {
      metadata {
      }
      spec {
        backoff_limit = 2
        ttl_seconds_after_finished    = 10
        template {
          metadata {}
          spec {
            container {
              name    = "backend-init-db"
              image   = "busybox"
              command = ["/bin/sh", "-c", "wget backend/init"]
            }
            restart_policy = "Never"
          }
        }
      }
    }
  }
}


resource "random_password" "backend-mariadb-password" {
  length = 16
  special = false
}

resource "kubernetes_secret" "backend-mariadb-access" {
  metadata {
    name = "backend-mariadb-access"
    namespace = kubernetes_namespace.backend.metadata[0].name
  }

  data = {
    DB_USER = "backend"
    DB_PASSWORD = random_password.backend-mariadb-password.result
    "username" = local.backend_username
    "database" = local.backend_database
    "password" = random_password.backend-mariadb-password.result
  }
}

resource "helm_release" "backend-mariadb" {
  name = "backend-mariadb"
  chart = "mariadb"
  repository = local.helm_repository_bitnami
  namespace = kubernetes_namespace.backend.metadata[0].name

  depends_on = [
    helm_release.prometheus-operator
  ]

  values = [
    file("helm/backend-database.yaml")
  ]

  set {
    name = "db.user"
    value = local.backend_username
  }

  set {
    name = "db.name"
    value = local.backend_database
  }

  set {
    name = "db.password"
    value = random_password.backend-mariadb-password.result
  }
}
