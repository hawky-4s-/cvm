# Design

CMV is a biased CamBPM installation manager for QA and Demo purposes.
It manages the creation and configuration of Camunda BPM servers. Multiple backends allow the usage of

## Entities

- Camunda Distro
- Server
- Database
- JDK

## Use case sequences

### Simple CamBPM installation

1. Fetch CVM configuration from Remote, Local, ENV etc.
2.
```
cvm create <name> \
    -c <camunda_version> \
    -s <server> \
    -db <database> -dbhost <database_host> -dbport <database_port> -dbname <database_name>
```

Simple configuration always goes with default values:
- camunda_version: latest
- server: tomcat
- db: h2
- dbhost, dbport, dbname will be default values for the databases
    - h2
    - mysql
    - mariadb
    - postgresql
    - db2
    - oracle
