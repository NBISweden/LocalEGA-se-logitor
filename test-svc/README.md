# Setting up the test-svc

## Installation

* Set up the config file according to your settings and run:
  ```console
  kubectl apply -n test -f svc-manifesto.yaml
  ```

* Verify that the service was correctly initiated:
  ```console
  kubectl get pods -n test
  ```
  
