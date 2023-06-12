io_mode = "async"

service "http" "web_proxy" {
  listen_addr = "0.0.0.0:8080"

  process "main" {
    command = "server"
  }

}