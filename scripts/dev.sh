ENV_FILE=".env"
if [ -f $CREDS_FILE ]; then
    set -a
    . $PWD/$ENV_FILE
    set +a
else
    echo "'$ENV_FILE' could not be found. See Readme on how to set up."
    exit 1
fi

# Run the application
./bin/${APP:-"app"}
#time /$GOPATH/bin/realize start run
