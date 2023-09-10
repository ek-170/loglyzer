#!/bin/sh
script_dir="$(cd "$(dirname "$0")" && pwd)"
data_dir="$script_dir/init_data"
es_dir="$data_dir/es"
ingest_dir="$es_dir/ingest"
elasticsearch_url="http://localhost:9200"

set -eu

printf '\e[33mInitialize Elasticsearch\e[m\n'
printf 'init start Ingest Pipeline\n'

for json_file in "$ingest_dir"/*.json; do
    if [ -e "$json_file" ]; then
        filename=$(basename "$json_file")
        # replace - to /
        path="${filename//-//}"
        # remove extension
        path="${path//.json//}"
        # remove last /
        path="${path/%?/}"

        curl_command="curl -XPUT -s \"$elasticsearch_url/$path\" -H \"Content-Type: application/json\" -d @$json_file"
        echo "Executing: $curl_command"
        eval "$curl_command" | jq
        printf "Completed: %s\n" "$filename"
    fi
done

printf "\e[33mElasticsearch Initialization completed.\e[m\n"