# loglyzer
## Feature
Functions are below.
- Ingest log data from given file
- Data ingested send to Elasticsearch
- Discover data in Kibana

## Prerequisite
Install docker and docker compose
If use Windows, need to enable WSL2
## Getting Started



## Development
using below stacks.

|tool|purpose|link|
|-|-|-|
|Docker|launch app on container||
|VSCode|editor||
|Air|hot reload of go server on docker|https://github.com/cosmtrek/air|
|Delve|debug golang|https://github.com/go-delve/delve|
|Elasticsearch|DB||
|Kibana|visualize||


## Trouble Shooting

### Error of Elasticsearch

#### vm.max_map_count is too low

see docker logs of failed container
```shell
docker logs loglyzer-es01-1 -f --tail=100
```

if display error as below, you should increase value of vm.max_map_count by using sysctl.
```json
{
    "@timestamp": "2023-10-14T08:34:01.103Z",
    "log.level": "ERROR",
    "message": "node validation exception\n[1] bootstrap checks failed. You must address the points described in the following [1] lines before starting Elasticsearch.\nbootstrap check failure [1] of [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]",
    "ecs.version": "1.2.0",
    "service.name": "ES_ECS",
    "event.dataset": "elasticsearch.server",
    "process.thread.name": "main",
    "log.logger": "org.elasticsearch.bootstrap.Elasticsearch",
    "elasticsearch.node.name": "es01",
    "elasticsearch.cluster.name": "es-cluster"
}
```

how to increase vm.max_map_count is as your enviroment
if use Rancher Desktop, you may read [this](https://docs.rancherdesktop.io/how-to-guides/increasing-open-file-limit)

#### suddenly stop Elasticsearch

see docker logs of failed container
```shell
docker logs loglyzer-es01-1 -f --tail=100
```

if display error as below at last line, you should increase value of ES_MEM_LIMIT in loglyzer/docker/.env
```
 es          | ERROR: Elasticsearch exited unexpectedly
```
if it doesn't improve, you should check memory of Docker allocated.