# Setting up filebeat
In order to setup filebeat, we use the official Helm chart maintained by Elastic which can be found [here](https://github.com/elastic/helm-charts/tree/master/filebeat)


## Installation

* Configure the `filebeat-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install filebeat elastic/filebeat -f helm/filebeat-vals.yaml -n elastic --version 7.7.1
  ```

* Verify that the service was correctly initiated:
  ```console
  kubectl get pods -n elastic -l chart=filebeat
  ```
  
