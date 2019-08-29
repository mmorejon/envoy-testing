# Envoy Testing

## Commands

* start services
  ```
  docker-compose --project-name dynamic-files --file docker-compose-files.yaml up --build
  ```
* remove services
  ```
  docker-compose --project-name dynamic-files --file docker-compose-files.yaml rm
  ```
* update envoy
  ```
  docker container exec -it  dynamic-files_proxy-api-eds_1 /tmp-files/update.sh lds cds eds
  ```
* test envoy
  ```
  curl -i --cacert cert.pem https://localhost
  ```