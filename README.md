
Daton
=====

Dev environment
----------------

For the dev environment we use Docker and Makefile

### Makefile commands

    $ make build_docker
    $ make cmd ARGS="env"
    $ make shell
    $ make run
    $ make stop
    $ make rm
    $ make test


### Library management

If you update or add a new dependency you need to build docker image again

    $ make build_docker

#### Add

    $ make shell
    $ go get package
    $ godep save package

#### Update

    $ make shell
    $ go get -u package
    $ godep update package

Licensing
---------

See [LICENSE](https://github.com/slok/daton/blob/master/LICENSE)