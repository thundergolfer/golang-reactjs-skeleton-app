CONTAINER_CLUSTER_NAME="twelve-factor-cluster"

gcloud container clusters create ${CONTAINER_CLUSTER_NAME} --num-nodes=1
