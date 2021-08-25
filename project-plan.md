# Kubernetes Cluster Homelab Project Plan

The long-term goal of this project is to build a homelab containing a Kubernetes cluster with multiple physical nodes for the purpose of running and hosting containerized workloads. These workloads might include web hosting, databases, and running ML or scientific computing jobs.

Below is a project plan defining specific intermediate goals to achieve on the way to a fully featured and functional self-hosted kubernetes cluster.

## Project Plan
- Install Ubuntu server on nodes
- Configure ssh access into nodes from local network
- Install kubernetes on all nodes
- Configure kubernetes master and worker nodes
- Deploy test workload to kubernetes with kubectl
- Deploy personal website to cluster
- Expose personal website to internet from cluster
  - Will require setting up other security measures to be determined
- Implement Prometheus monitoring
- Implement a Grafana dashboard