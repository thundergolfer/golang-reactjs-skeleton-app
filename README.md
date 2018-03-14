# 12 Factor (App)
A basic twelve-factor app for use as a best-practices project skeleton

-----

## How To Use This

* Learn from it. There's notes in this `README.md` and throughout the "notes" `git` branch as to how this app fulfils the twelve factors.
* Use it as a project base/skeleton. This app is set up as to enable quick development and deployment of web applications.

## Installation

1. Install the back-end with [`dep`](https://github.com/golang/dep) by running `dep ensure`
2. Install the front-end with [`yarn`](https://yarnpkg.com/lang/en/) by running `yarn`.

## Running The App

### Back-End

`go run ./backend/*.go`

## The Twelve (XII) Factors

#### **I. Codebase** - One codebase tracked in revision control, many deploys

All code for this app (front-end, back-end, deployment) is contained in this repository, thus fulfilling factor `1`.

#### **II. Dependencies** - Explicitly declare and isolate dependencies

* Back-end dependencies are explicitly declared by `Gopkg.lock` and isolated in `vendor/`.
* Front-end dependencies are explicitly declared in `yarn.lock` and isolated in `frontend/node_modules/`
* The `curl` command, which is used during app build, is vendored into `vendor/` and not assumed to be present on the app's system.


#### **III. Config** - Store config in the environment

As stipulated, all configuration that varies between deployments of the app is picked up by the app from the environment. See: `backend/config.go`.

#### **IV. Backing services** - Treat backing services as attached resources


#### **V. Build, release, run** - Strictly separate build and run stages

The build stage is strictly separated into the building of a single Docker image, done with `scripts/build_docker_image.sh`.

**TODO:** Think about how config get packaged into releases


#### **VI. Processes** - Execute the app as one or more stateless processes

The back-end web app operates statelessly if any datastore but the `InMemoryStorer` is used. State is kept in an attached datastore resource, like the `GoogleCloudStorer`.

#### **VII. Port binding** - Export services via port binding

Our back-end web server exposes itself on port `8080`.

#### **VIII. Concurrency** - Scale out via the process model

This example app currently only has 1 *process-type*, the web-server, but that process-type can be instantiated into many processes for horizontal scale. **NOTE:** What about race conditions?

#### **IX. Disposability** - Maximize robustness with fast startup and graceful shutdown

This is a minimal toy example app, so it's going to startup quickly and without fuss. It will also shutdown cleanly, calling no mandatory 'cleanup' tasks that could fail in a crash scenario.

#### **X. Dev/prod parity** - Keep development, staging, and production as similar as possible

[*Google Cloud Datastore Emulator*](https://cloud.google.com/datastore/docs/tools/datastore-emulator) and [*Minio*](https://github.com/minio/minio) allow for local development with production-like backing services. *Docker* facilitates OS-level virtualisation both locally and in production (ie. in Google Cloud, AWS).

#### **XI. Logs** - Treat logs as event streams

Fulfilled, as application logs are dumped out to `stdout` with no concern for storage or routing, and they have nice structure granted by the [logrus](https://github.com/sirupsen/logrus) library.

#### **XII. Admin processes** - Run admin/management tasks as one-off processes

One-time maintenance scripts are committed in the repo in `scripts/` alongside regular application code. The back-end is in `golang` though, so we don't get the REPL the *12 Factor Methodology* strongly encourages.
