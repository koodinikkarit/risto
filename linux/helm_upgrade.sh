IMAGE_TAG="$TRAVIS_BRANCH-$TRAVIS_COMMIT"

RELEASE_NAME=$(echo "$TRAVIS_BRANCH" | sed 's/\./-/g')

helm upgrade \
--wait \
--set ristoImage=jaska/risto:$IMAGE_TAG \
--set mysqlDatabase=$MYSQL_DATABASE \
--set mysqlPassword=$MYSQL_PASSWORD \
--set mysqlHost=$MYSQL_HOST \
--set mysqlPort=$MYSQL_PORT \
--install risto-$RELEASE_NAME ./deployment
