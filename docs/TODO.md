# TODOs

1. Stage

- Installation of CamBPM into local environment
    - Tomcat
    - H2

Requirements:
- File download with basic auth
- XML manipulation -> XML library - https://github.com/beevik/etree
- File operations (cp,mv, ...)
- Process starting
- Download caching

1.1 Provisioning steps

- Download distro
- Download JDBC driver
- Configure distro

2. Stage

- Wildfly support (JBoss)
- WebLogic
- WebSphere

3. Stage

- Support docker backend

4. Stage

- Usage of Docker containers, including database container

5. Stage

- Support spin up of environment CamBPM + Server + Database
- Support more backends like Docker-Compose / Swarm / Kubernetes

6. Stage

- Support installation with downloaded servers like JBossEAP etc.
- Support installation with non-camunda Docker images like official Tomcat runtime etc.

7. Stage

- Replace the official docker-camunda-bpm-platform images with it


- Support to manage already created configurations locally?
- Checksum support
- Download / manage a custom Dockerfile (from GitHub, with scripts for installation)

- interactive mode for manually specifying stuff
- load cvm configuration (databases, jdks, servers, camundas)
- support env variables for installation with prefix CVM_, eg.
  - CAMUNDA_VERSION
  - JDK_VERSION
  - DATABASE_VERSION
  - JDBC_DRIVER_VERSION
  - CONFIG_URL

# Features

## Usability

- Outputs for ports / state / application

## Version matrix

- Camunda
- Application server
- Database

## Development

- Debugging
- local repositories
- run integration tests?

## QA

- access for file deployments


https://google.github.io/styleguide/shell.xml
https://github.com/jpetazzo/pipework/blob/master/pipework
https://github.com/kevinburke/go-bindata
https://www.appveyor.com/
