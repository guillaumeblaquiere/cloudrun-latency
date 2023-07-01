# Overview

Cloud Run latency is a simple HTTP service that measure the service to service latency in Cloud Run, but also
in any environment that run container, like Kubernetes and GKE on Google Cloud.

*This [article](https://medium.com/google-cloud/cloud-run-gke-istio-network-latency-comparison-60f9599a50bb) 
presents that the concept and how to test the latency*

# General architecture

There is only 1 container with 2 endpoints
* `/ping` logs the latency when it calls another service. For that, it receives 
  * The target URL: the service to call and to measure the latency
  * The number of call to perform to the target URL
* `/pong` which do nothing, only return a HTTP 200 response code.

# How to use

Build the container and deploy it on the target environment (see the article for more details)

# Licence

This library is licensed under Apache 2.0. Full license text is available in
[LICENSE](https://github.com/guillaumeblaquiere/cloudrun-latency/tree/master/LICENSE).