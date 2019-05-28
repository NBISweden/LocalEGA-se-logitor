# Backing up Elasticsearch snapshots in a S3 data lake 
In order to setup Elasticsearch, we use the official Helm chart maintained by Elastic which can be found [here](https://github.com/elastic/helm-charts/tree/master/elasticsearch)

## Installation

* Add the Elastic Helm charts repo.
  ```console
  helm repo add elastic https://helm.elastic.co
  ```
* Configure the `elastic-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install --name elasticsearch -f helm/elastic-vals.yaml --namespace=myns elastic/elasticsearch --version 7.1.0
  ```
* Configure a log exporter daemon like `fluentd`, `fluentbit` or `logstash` to talk to the newly deployed Elasticsearch service. If you are using Rancher, you can easily set this up in the logging settings of your project.

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
    
