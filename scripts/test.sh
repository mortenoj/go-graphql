ENV_FILE=".env.test"
if [ -f $CREDS_FILE ]; then
    set -a
    . $PWD/$ENV_FILE
    set +a
else
    echo "'$ENV_FILE' could not be found. See Readme on how to set up."
    exit 1
fi

echo "====>Running tests"

if [[ -n "${COVER}" ]]; then
    go test  -c -coverpkg ./...
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out
    rm coverage.out
    #go list ./... | xargs -t -n4 go test $TESTARGS -coverpkg
else
    go list ./... | xargs -t -n4 go test $TESTARGS -timeout=2m -parallel=4
fi
