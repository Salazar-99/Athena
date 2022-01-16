# Traefik

This directory contains the files and configs I use for implementing the Traefik Ingress Controller on my cluster.
The Traefik Ingress Controller can be installed via a Helm chart as follows

```bash
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
helm install traefik traefik/traefik -n traefik
```
