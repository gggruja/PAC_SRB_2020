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

resource "random_password" "keycloak-mariadb-password" {
  length = 16
  special = false
}

resource "kubernetes_secret" "mariadb-access" {
  metadata {
    name = "keycloak-mariadb-access"
    namespace = kubernetes_namespace.keycloak.metadata[0].name
  }

  data = {
    "username" = local.keycloak_username
    "database" = local.keycloak_database
    "password" = random_password.keycloak-mariadb-password.result
  }
}

resource "helm_release" "keycloak-mariadb" {
  name = "keycloak-mariadb"
  chart = "mariadb"
  repository = local.helm_repository_bitnami
  namespace = kubernetes_namespace.keycloak.metadata[0].name

  values = [
    file("helm/keycloak-database.yaml")
  ]

  set {
    name = "db.user"
    value = local.keycloak_username
  }

  set {
    name = "db.name"
    value = local.keycloak_database
  }

  set {
    name = "db.password"
    value = random_password.keycloak-mariadb-password.result
  }
}


resource "random_password" "keycloak-user" {
  length = 16
  special = false
}

resource "kubernetes_secret" "keycloak-user" {

  metadata {
    name = "keycloak-user"
    namespace = kubernetes_namespace.keycloak.metadata[0].name
  }

  data = {
    "password" = random_password.keycloak-user.result
  }

}

resource "helm_release" "keycloak" {
  depends_on = [
    helm_release.keycloak-mariadb
  ]

  name = "keycloak"
  namespace = kubernetes_namespace.keycloak.metadata[0].name
  chart = "keycloak"
  repository = local.helm_repository_codecentric

  values = [
    file("helm/keycloak.yaml")
  ]

  set {
    name = "keycloak.username"
    value = local.keycloak_username
  }

  set {
    name = "keycloak.persistence.dbName"
    value = local.keycloak_database
  }

  set {
    name = "keycloak.persistence.dbHost"
    value = helm_release.keycloak-mariadb.name
  }
}
