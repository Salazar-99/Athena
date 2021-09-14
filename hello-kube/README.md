## Hello-Kube

This is a simple Go application and it's associated Kubernetes deployment resources.

The Go application reads the ```GREETING``` environment variable and responds with "{GREETING}, World" to any requests at the "/" route.
The application is compiled and containerized via the Dockerfile.
The Kubernetes deployment runs a single replica of the application's container with the GREETING environment variable specified via a ConfigMap.
The ConfigMap is specified in k8s/configmap.yaml and defines a value for the GREETING environment variable via the greeting key's value.