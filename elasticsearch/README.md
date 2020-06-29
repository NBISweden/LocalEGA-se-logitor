# Backing up Elasticsearch snapshots in a S3 data lake 
In order to setup Elasticsearch, we use the official Helm chart maintained by Elastic which can be found [here](https://github.com/elastic/helm-charts/tree/master/elasticsearch)


## Prerequisites


* Set up an S3 endpoint as follows:

```
kubectl create ns elastic
```

```
helm install myminio stable/minio -n elastic --set accessKey=myaccesskey,secretKey=mysecretkey --version 5.0.30
```

## Installation

* Add the Elastic Helm charts repo.
  ```console
  helm repo add elastic https://helm.elastic.co
  ```
* Configure the `elastic-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install elasticsearch elastic/elasticsearch -f helm/elastic-vals.yaml -n elastic --version 7.8.0
  ```

* Configure a log exporter daemon like `filebeat`. See the filebeat setup in this repository.

* Verify that indexes were actually recorded by running:
  ```console
  curl -X GET "$ELASTICSEARCH_SERVICE_HOST:$ELASTICSEARCH_SERVICE_PORT/_cat/indices?v&pretty"
  ```
  
  ## Automating snapshots

 * First, a S3 repository needs to be created for your snapshots. Please run:
 
   ```console
   curl -X PUT http://$ELASTICSEARCH_SERVICE_HOST:$ELASTICSEARCH_SERVICE_PORT/_snapshot/$REPOSITORY_NAME -H 'Content-Type: application/json' -d '{ "type": "s3", "settings": { "bucket": "bucket123" } }'
   ```
  
 * To automate the snapshot upload to the S3 backend, you can easily setup a cron job on Kubernetes. Please edit the configuration file to adapt it to your environment.
  
    ```console
    kubectl apply -f cron/s3-snap-cron.yaml
    ```
    
