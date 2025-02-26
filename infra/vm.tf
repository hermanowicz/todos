resource "hcloud_server" "server-1" {
  name        = "pyinfra-1"
  server_type = "cx22"
  image       = "ubuntu-24.04"
  location    = "fsn1"

  ssh_keys = [
    "hetzner_key"
  ]

  network {
    network_id = hcloud_network.main-net.id
    ip         = "10.0.1.5"
    # alias_ips = [
    #   "10.0.1.6",
    #   "10.0.1.7"
    # ]
  }

  depends_on = [
    hcloud_network_subnet.network-subnet
  ]
}
