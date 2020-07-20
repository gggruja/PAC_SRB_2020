resource "kubernetes_namespace" "monitoring" {
  metadata {
    name = "monitoring"
  }
}

resource "helm_release" "prometheus-operator" {
  repository = local.helm_repository_stable
  chart = "prometheus-operator"
  namespace = kubernetes_namespace.monitoring.metadata[0].name
  name = "prometheus-operator"

  values = [
    file("helm/prometheus-operator.yaml")
  ]
}

resource "kubernetes_secret" "prometheus-basic-auth" {
  metadata {
    name = "prometheus-basic-auth"
    namespace = kubernetes_namespace.monitoring.metadata[0].name
  }

  data = {
    "auth" = "prometheus:$apr1$BW1Uu32h$Ru5z2hKVnfHNSPa4y2CT41"
  }
}
