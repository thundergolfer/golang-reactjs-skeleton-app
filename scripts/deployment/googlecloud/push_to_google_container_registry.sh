DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $DIR/../../.deployment_env_vars

gcloud docker -- push gcr.io/${PROJECT_ID}/${IMAGE_NAME}:v${TAG}
