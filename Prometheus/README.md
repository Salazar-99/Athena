# Prometheus

This directory contains the files and configs I use for implementing Prometheus monitoring on my cluster

The Prometheus, AlertManager, and Grafana resources on my cluster are provided by the official `kube-prometheus-stack` helm chart.
This chart can be installed as follows:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack
```
