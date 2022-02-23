## Airflow Helm Chart

The official Airflow Helm Chart requires three avaiabale Perisistent Volumes for logs, redis, and postgresql. 
This folder contains definitions for these three PVs as well as a Storage Class definition for local storage on the
nodes of the cluster. Finally, a modified values.yaml file for my deployment of the Helm chart is included.