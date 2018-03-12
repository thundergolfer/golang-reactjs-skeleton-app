DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/../../.deployment_env_vars

kubectl run ${DEPLOYMENT_NAME} --image=gcr.io/${PROJECT_ID}/${IMAGE_NAME}:v${TAG} --port $PORT
