locals {
  keycloak_username = "keycloak"
  keycloak_database = "keycloak"
}

resource "kubernetes_namespace" "keycloak" {
  metadata {
    name = "keycloak"
    labels = {
      "app" = "keycloak"
    }
  }
}
