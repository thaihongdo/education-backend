#!/bin/sh

export ENV="local"
export PORT="7171"

export CLIENT_URL="http://localhost:3000" 



echo "Setting env as $ENVIRONMENT_PREFIX\n"

GO111MODULE=on  go build -o cmd/entity-server/entity-server cmd/entity-server/main.go 
if [ $? -eq 0 ]
then
    cd cmd/entity-server/
    ./entity-server 
fi

