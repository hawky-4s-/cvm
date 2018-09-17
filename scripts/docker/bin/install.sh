#!/bin/bash -eux

VERSION=${VERSION:-latest} # numerical value or latest (autodiscovery)
DISTRO=${DISTRO:-tomcat} # tomcat, jboss, wildfly, wildfly8, wildfly10, wildfly11, jbosseap(6.1/6.2/6.3/6.4/7.0/7.1), weblogic(12.1/12R2), websphere(8.0/8.5/9.0)
EE=${EE:-true} # true / false
SNAPSHOT=${SNAPSHOT:-false} # true / false
DATABASE=${DATABASE:-postgresql} # mysql, mariadb, postgresql, oracle 10/11/12, db2 9.7/10.1/10.5/11.1
DATABASE_VERSION=${DATABASE_VERSION:-latest} # numerical value or latest (autodiscovery)

function cambpmToApplicationServerMatrix {
    7.6 tomcat, jboss, wildfly ...
    7.7 tomcat, jboss, wildfly ...
    7.8 tomcat, jboss, wildfly ...
    7.9 tomcat, jboss, wildfly ...
    7.10 tomcat, jboss, wildfly ...
}

function applicationServerJDKLookup {
    # only exceptions in map... if not found, default to jdk8
    # lookup values in job generator config
    declare -A applicationServerToJDK

    applicationServerToJDK=( ["tomcat"] )


}

function main() {

}

main
