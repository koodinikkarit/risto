IMAGE_TAG="$CI_COMMIT_REF_NAME-$CI_COMMIT_SHA"

RELEASE_NAME=$(echo "$CI_COMMIT_REF_NAME" | sed 's/\./-/g')

helm upgrade \
--wait \
--set ristoImage=jaska/risto:$IMAGE_TAG \
--install risto-$RELEASE_NAME ./deployment
