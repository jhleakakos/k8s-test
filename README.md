Small test for Kubernetes.

# Command Reminders

Can test images and containers with a `run` command or with `compose`

- `docker compose up -d` and `docker compose down` to use compose to run individual containers

Prep for k8s

- `docker build -t <image_name> .` from within directory with Dockerfile to build the image
- `kubectl apply -f <deployment_file_yaml>` to deploy with k8s

K8s basics

- `kubectl get deployments` to check deployments
- `kubectl get services` to check services (in particular to get the nodeport values)
- `kubectl delete deployments <deployment_name>` to delete a deployment
- `kubectl delete services <service_name>` to delete a service

We are using a NodePort service to expose the deployment outside of the cluster. To connect

- `kubectl get services` to get the NodePort value (xxxx:NodePort)
- you can then curl http://localhost:NodePort or reach it from a browser

# Updates

To test an update, change something in a source file and rebuild the image with a new version tag in the name. Then update the k8s yaml file to match the new image name. Then you can run a new `kubectl apply` to push out the deployment. K8s will create a new replica set and start to transfer pods over to it.

If you simulate a failed update, you can use `kubectl rollout <subcommand>` to investigate and initiate a rollback.