## To-Do

Here I keep track of the various projects I'm working on for my cluster

### Prometheus-Grafana Monitoring
- Create a dashboard for cluster resource metrics (CPU, Memory, Disk)
- Create a dashboard for CronJobs (starts, failures, coming up)

### Self Hosting
- Configure Flux to track personal-website repo for CI/CD
  - TODO: Double check ImagePolicy, ImageUpdateAutomation, and ImageRegistry objects to configure CD for personal website
  - Site currently down due to deployment throwing "Authentiaction token invalid" error on image pull from ECR

## Remote Access to Cluster
- Investigate using a WireGuard VPN to access home network remotely
- Investigate using CloudFlare to allow kubectl securely from a remote network to home cluster

## Data Engineering
- Add Airflow to cluster
- Add Spark on cluster