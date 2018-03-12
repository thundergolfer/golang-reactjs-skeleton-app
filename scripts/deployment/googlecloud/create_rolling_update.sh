set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

source $DIR/../../.deployment_env_vars

echo "Building docker image for ${TAG}"
./build_docker_image.sh

echo "Pushing build image to Google Container Registry"
./push_to_google_container_registry.sh

kubectl set image deployment/${DEPLOYMENT_NAME} ${DEPLOYMENT_NAME}=gcr.io/${PROJECT_ID}/${IMAGE_NAME}:v${TAG}
