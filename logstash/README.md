# Setting up logstash 
In order to setup logstash, we use the official Helm chart maintained by Elastic which can be found [here](https://github.com/elastic/helm-charts/tree/master/logstash)


## Installation

* Configure the `logstash-vals.yaml` file according to your settings and install the chart.
  ```console
  helm install logstash elastic/logstash -f helm/logstash-vals.yaml -n elastic --version 7.7.1
  ```

* Verify that the service was correctly initiated:
  ```console
  kubectl get pods -n elastic -l chart=logstash
  ```
  
