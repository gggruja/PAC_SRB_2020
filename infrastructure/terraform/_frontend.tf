resource "kubernetes_namespace" "frontend" {
  metadata {
    name = "frontend"
  }
}

resource "helm_release" "frontend" {
  name = "frontend"
  chart = "../chart/frontend"
  namespace = kubernetes_namespace.frontend.metadata[0].name

  depends_on = [
    helm_release.backend
  ]

  values = [
    file("helm/frontend.yaml")
  ]
}
