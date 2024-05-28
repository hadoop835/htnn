---
title: Observability
---

HTNN is 100% compatible with Istio and Envoy, which allows us to use Istio's [observability features](https://istio.io/latest/docs/concepts/observability/).

In addition to this, HTNN also provides some extra observability features.

Note: The following features depend on Istio's debug interface and Prometheus metrics being enabled. By default, they are both turned on.

## Log

HTNN control plane's additional functionalities all use the logger named `htnn`. You can dynamically adjust the log level through [ControlZ](https://istio.io/latest/docs/ops/diagnostic-tools/controlz/). By setting the log level of `htnn` to `debug`, you can view the entire reconciliation process.

HTNN data plane features developed in Go use the logger named `golang`. You can dynamically adjust the log level through [Envoy Admin API](https://www.envoyproxy.io/docs/envoy/latest/operations/admin#post--logging) or with `istioctl pc log $pod_name --level golang:debug`.

## Metrics

The HTNN control plane adds the following metrics:

| Name                                             | Type      | Description                                                      |
|--------------------------------------------------|-----------|------------------------------------------------------------------|
| htnn_httpfilterpolicy_reconcile_duration_seconds | histogram | How long in seconds HTNN reconciles HTTPFilterPolicy.            |
| htnn_httpfilterpolicy_translate_duration_seconds | histogram | How long in seconds HTNN translates HTTPFilterPolicy in a batch. |
| htnn_consumer_reconcile_duration_seconds         | histogram | How long in seconds HTNN reconciles Consumer.                    |
| htnn_serviceregistry_reconcile_duration_seconds  | histogram | How long in seconds HTNN reconciles ServiceRegistry.             |

You can access these metrics by default via Istio's Prometheus port `127.0.0.1:15014/metrics`. Note that if a metric has no data, it will not appear.

## Debug

The EnvoyFilter and ServiceEntry generated by the HTNN control plane can be obtained through Istio's own `configz` interface. For example, by running `kubectl exec -it istiod-xxx -- curl 127.0.0.1:8080/debug/configz | jq`, you can see:

```
{
  "kind": "EnvoyFilter",
  "apiVersion": "networking.istio.io/v1alpha3",
  "metadata": {
    "name": "htnn-http-filter",
    "namespace": "istio-system",
    "resourceVersion": "52795",
    "creationTimestamp": "2024-05-10T10:38:02Z",
    "labels": {
      "htnn.mosn.io/created-by": "HTTPFilterPolicy"
    },
    "annotations": {
      "htnn.mosn.io/info": "{\"httpfilterpolicies\":[\"nodesentry/policy\"]}"
    }
  },
  ...
},
...
```

The generated EnvoyFilter is tagged with the label "htnn.mosn.io/created-by", marking what kind of resource it was generated from. There is also an annotation "htnn.mosn.io/info" which contains the following fields:

* `httpfilterpolicies`: Policies for generating this EnvoyFilter, named `$namespace/$name`.