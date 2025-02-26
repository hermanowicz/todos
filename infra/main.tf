terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.45"
    }
  }
}

# Configure the Hetzner Cloud Provider
provider "hcloud" {}

resource "hcloud_network" "labs-net" {
  name     = "labs-net"
  ip_range = "10.0.0.0/16"
}

resource "hcloud_network_subnet" "network-subnet" {
  type         = "cloud"
  network_id   = hcloud_network.main-net.id
  network_zone = "eu-central"
  ip_range     = "10.0.1.0/24"
}


