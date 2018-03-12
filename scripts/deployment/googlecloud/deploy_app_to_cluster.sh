DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/../../.deployment_env_vars

kubectl expose deployment ${DEPLOYMENT_NAME} --type=LoadBalancer --port 80 --target-port $PORT
