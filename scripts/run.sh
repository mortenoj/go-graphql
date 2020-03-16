ENV_FILE=".env"
if [ -f $CREDS_FILE ]; then
    set -a
    . $PWD/$ENV_FILE
    set +a
else
    echo "'$ENV_FILE' could not be found. See Readme on how to set up."
    exit 1
fi

if [[ "${DEV}" = true ]]; then
    # Run the application with realize
    time /$GOPATH/bin/realize start run
    #./bin/${APP:-"app"}
else
    # Run the application
    ./bin/${APP:-"app"}
fi
