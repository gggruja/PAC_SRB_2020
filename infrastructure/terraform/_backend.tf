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
