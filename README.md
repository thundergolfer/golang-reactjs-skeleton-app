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


#### **IV. Backing services** - Treat backing services as attached resources


#### **V. Build, release, run** - Strictly separate build and run stages


#### **VI. Processes** - Execute the app as one or more stateless processes


#### **VII. Port binding** - Export services via port binding


#### **VIII. Concurrency** - Scale out via the process model


#### **IX. Disposability** - Maximize robustness with fast startup and graceful shutdown


#### **X. Dev/prod parity** - Keep development, staging, and production as similar as possible


#### **XI. Logs** - Treat logs as event streams


#### **XII. Admin processes** - Run admin/management tasks as one-off processes
