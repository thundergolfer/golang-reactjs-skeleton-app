DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/../../.deployment_env_vars

PORT="8080"

kubectl expose deployment ${DEPLOYMENT_NAME} --type=LoadBalancer --port 80 --target-port $PORT

echo "Deployment IP details: "
echo "-----------------------"
kubectl get service
