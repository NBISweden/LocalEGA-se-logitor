# Setting up Grafana 
In order to setup Grafana, we use the official Helm chart maintained by Kubernetes community which can be found [here](https://github.com/helm/charts/tree/master/stable/grafana)

## Installation

 * Configure the `grafana-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install --name grafana-ega -f grafana-vals.yaml stable/grafana
  ```

 * Extract your admin credentials from the grafana secret. The default user is `admin`. 
  ```console
  kubectl get secret grafana-ega -o json | jq -c '.data."admin-password"' | tr -d '"' | base64 -D
  ```
  
 * Now you should be able to login into the dashboard and add a data source in the configuration. If using Elasticsearch, you will need to specify the format of your indices. This can be consulted by:
  ```console
  curl "$ELASTICSEARCH_SERVICE_HOST:$ELASTICSEARCH_SERVICE_PORT/_all" | jq 'keys'
  ```
