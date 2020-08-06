resource "kubernetes_namespace" "backend" {
  metadata {
    name = "backend"
  }
}


resource "helm_release" "backend" {
  name = "backend"
  chart = "../chart/backend"
  namespace = kubernetes_namespace.backend.metadata[0].name

  values = [
    file("helm/backend.yaml")
  ]
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
