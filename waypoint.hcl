project = "cats"

app "cats-docker" {
  labels = {
    "service" = "cats-docker"
  }
  build {
    use "pack" {}
  }

  deploy {
    use "docker" {}
  }
}
