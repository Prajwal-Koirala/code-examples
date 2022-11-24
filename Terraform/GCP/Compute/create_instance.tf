# Create a Google cloud compute instance
resource "google_compute_instance" "vm_instance" {
  name                = "terraform-instance"
  machine_type        = "f1-micro"
  zone                = "us-central1-c"
  can_ip_forward      = false
  deletion_protection = false
  enable_display      = false
  boot_disk {
    auto_delete = true
    device_name = "persistent-disk-0"
    mode        = "READ_WRITE"
    initialize_params {
      image = "debian-cloud/debian-11"
      size  = 10
      type  = "pd-standard"
    }
  }
  network_interface {
    network     = "default"
    queue_count = 0
    access_config {
      network_tier = "PREMIUM"
    }
  }
  scheduling {
    automatic_restart   = true
    min_node_cpus       = 0
    on_host_maintenance = "MIGRATE"
    preemptible         = false
    provisioning_model  = "STANDARD"
  }
  shielded_instance_config {
    enable_integrity_monitoring = true
    enable_secure_boot          = true
    enable_vtpm                 = true
  }
  metadata = {
    block-project-ssh-keys = true
  }
}