# Mora

[![Build Status](https://travis-ci.org/bbengfort/mora.svg?branch=master)](https://travis-ci.org/bbengfort/mora)
[![Coverage Status](https://coveralls.io/repos/github/bbengfort/mora/badge.svg?branch=master)](https://coveralls.io/github/bbengfort/mora?branch=master)
[![GoDoc Reference](https://godoc.org/github.com/bbengfort/mora?status.svg)](https://godoc.org/github.com/bbengfort/mora)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbengfort/mora)](https://goreportcard.com/report/github.com/bbengfort/mora)
[![Stories in Ready](https://badge.waffle.io/bbengfort/mora.png?label=ready&title=Ready)](https://waffle.io/bbengfort/mora)

**A study of  the uptime and latency of user-oriented distributed systems.**

![Mora Architecture Diagram](http://bbengfort.github.io/assets/images/2016-05-10-mora-architecture.png)

Mora is a distributed systems study that measures the uptime and latency of heterogenous, partition-prone, user-oriented networks. The project contains three pieces: Oro and Scio which ping each other to measure latency inside of the network, then report those pings to [Scribo](https://github.com/bbengfort/scribo), which is simply a RESTful API designed to record experimental data.

This package provides library functionality for the Oro and Scio applications, as well as the command to build a Scio daemon.

## About

Mora (delay, waiting) observes ping latencies between nodes in a wide area, heterogenous, user-oriented network by running a local service that pings other nodes in the network. Oro (speak) is the name of the mobile application, and Scio (understand) is the name of the desktop client. The ping data is collected by a centralized RESTful microservice called Scribo (record). This data will be used for scientific research concerning distributed systems.

### Documentation

[Documenting go code correctly](https://blog.golang.org/godoc-documenting-go-code) is vital, because [godoc.org](https://godoc.org/) will automatically [pull package documentation](https://godoc.org/-/about) from GitHub.

To view the documentation locally before pushing:

    godoc -http=:6060

This will provide a dashboard for the Go code on your local machine. You can find documentation for this package here: [godoc.org: package mora](https://godoc.org/github.com/bbengfort/mora).

### Contributing

Mora is open source, and I'd love your help, particularly if you are a student at the University of Maryland and are interested in studying distributed systems. If you would like to contribute, you can do so in the following ways:

1. Add issues or bugs to the bug tracker: [https://github.com/bbengfort/scribo/issues](https://github.com/bbengfort/mora/issues)
2. Work on a card on the dev board: [https://waffle.io/bbengfort/scribo](https://waffle.io/bbengfort/mora)
3. Create a pull request in Github: [https://github.com/bbengfort/scribo/pulls](https://github.com/bbengfort/mora/pulls)

Note that labels in the Github issues are defined in the blog post: [How we use labels on GitHub Issues at Mediocre Laboratories](https://mediocre.com/forum/topics/how-we-use-labels-on-github-issues-at-mediocre-laboratories).

When doing a pull request, keep in mind that the project is set up in a typical production/release/development cycle as described in _[A Successful Git Branching Model](http://nvie.com/posts/a-successful-git-branching-model/)_. A typical workflow is as follows:

1. Select a card from the [dev board](https://waffle.io/bbengfort/mora) - preferably one that is "ready" then move it to "in-progress".

2. Create a branch off of develop called "feature-[feature name]", work and commit into that branch.

        ~$ git checkout -b feature-myfeature develop

3. Once you are done working (and everything is tested) merge your feature into develop.

        ~$ git checkout develop
        ~$ git merge --no-ff feature-myfeature
        ~$ git branch -d feature-myfeature
        ~$ git push origin develop

4. Repeat. Releases will be routinely pushed into master via release branches, then deployed to the server.

Note that no pull requests into master will be considered; only those that pull into develop.

### Throughput

[![Throughput Graph](https://graphs.waffle.io/bbengfort/mora/throughput.svg)](https://waffle.io/bbengfort/mora/metrics/throughput)

## Contributors

Thank you for all your help contributing to make Mora a great project!

### Maintainers

- Benjamin Bengfort: [@bbengfort](https://github.com/bbengfort/)

### Contributors

- Your name here!

## Changelog

The release versions that are tagged in Git. You can see the tags through the GitHub web application and download the tarball of the version you'd like.

The versioning uses a three part version system, "a.b.c" - "a" represents a major release that may not be backwards compatible. "b" is incremented on minor releases that may contain extra features, but are backwards compatible. "c" releases are bug fixes or other micro changes that developers should feel free to immediately update to.
