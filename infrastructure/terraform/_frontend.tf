resource "kubernetes_namespace" "frontend" {
  metadata {
    name = "frontend"
  }
}
