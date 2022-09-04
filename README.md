Small test for Kubernetes.

# Command Reminders

- `docker build -t <image_name> .` from within directory with Dockerfile to build the image
- `docker compose up -d` and `docker compose down` to use compose to run individual containers
- `kubectl apply -f <deployment_file_yaml>` to deploy with k8s

We are using a NodePort service to expose the deployment outside of the cluster. To connect:

- `kubectl get services` to get the NodePort value (xxxx:NodePort)
- you can then curl http://localhost:NodePort or reach it from a browser
