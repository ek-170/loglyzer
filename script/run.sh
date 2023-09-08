#!/bin/bash

help(){
    echo "run.sh usage is \"run.sh [port] [git branch name]\""
    exit 1
}

if [ $# -ne 2 ]; then
    help
fi

set -eux

cd $(dirname $0)
pwd

echo DOCKER_HTTP_PORT="$1" > .env
cat .env

if [ -e ./loglyzer ]; then
    rm -rf ./loglyzer 
fi

cd ./loglyzer/apps/web
git branch --list --all 
yarn install
yarn run build
yarn run export
ls

cd ../../../
if [ -e ./httpd/loglyzer ]; then
    rm -rf ./httpd/loglyzer
fi
mv ./loglyzer/apps/web/out ./loglyzer/apps/web/loglyzer
mv ./loglyzer/apps/web/loglyzer ./httpd/.

if [ -e ./tomcat/loglyzer ]; then
    rm -rf ./tomcat/loglyzer
fi
mv ./loglyzer ./tomcat/.

docker compose down

docker compose up -d --build

docker compose exec tomcat /bin/bash -c "cd /opt/loglyzer/migrations && ./gradlew run_pg"
