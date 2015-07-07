
Daton
=====

Dev environment
----------------

For the dev environment we use Docker and Makefile

### Makefile commands

build_docker:
    $ make build_docker

cmd:
    $ make cmd ARGS="env"

shell:
    $ make shell

run:
    $ make run

stop:
    $ make stop

rm:
    $ make rm

test:
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