# Setting up filebeat
In order to setup filebeat, we use the official Helm chart maintained by Elastic which can be found [here](https://github.com/elastic/helm-charts/tree/master/filebeat)


## Installation

* Configure the `filebeat-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install filebeat elastic/filebeat -f helm/filebeat-vals.yaml -n elastic --version 7.8.0
  ```

* Verify that the service was correctly initiated:
  ```console
  kubectl get pods -n elastic -l chart=filebeat
  ```

* Explore the ES API in search for a namespace's logs:
```console
curl --location --request GET '$ELASTICSEARCH_HOST/$ES_INDEX/_search?pretty=true&size=100' --header 'Content-Type: application/json' --data-raw '{
    "query" : {
        "term": { "kubernetes.namespace" : "default"}
    }
}'
```
* Or get the logs from a specific container tag in a given namespace over a specific time range:

```console
curl --location --request  GET '$ELASTICSEARCH_HOST/$ES_INDEX/_search?pretty=true&size=100' --header 'Content-Type: application/json' --data-raw '{
    "query" : {
        "bool": {
        "must": [
        { "match": { "kubernetes.container.image": "novella/svc-test:003" }}
      ],
        "filter": [
            {"term": {"kubernetes.namespace": "default"}},
            {"range": { "@timestamp": { "gte": "2020-06-28" }}}
        ]
    }
    }
}'
```
  
