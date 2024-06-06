# htnn-controller

![Version: 0.1.3](https://img.shields.io/badge/Version-0.1.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.2.1](https://img.shields.io/badge/AppVersion-0.2.1-informational?style=flat-square)

A Helm chart for HTNN controller

## Install

To install the chart with the release `htnn-controller`:

```shell
helm repo add htnn https://mosn.github.io/htnn
helm repo update

helm install htnn-controller htnn/htnn-controller --namespace istio-system --create-namespace
```

For more information like how to configure and troubleshoot, please refer to the [Installation Guide](https://github.com/mosn/htnn/blob/main/site/content/en/docs/getting-started/installation.md).

## Uninstall

To uninstall the Helm release `htnn-controller`:

```shell
helm uninstall htnn-controller -n istio-system
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| spacewander |  |  |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://istio-release.storage.googleapis.com/charts | istio-base(base) | 1.21.2 |
| https://istio-release.storage.googleapis.com/charts | istiod | 1.21.2 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| istiod.global.istioNamespace | string | `"istio-system"` |  |
| istiod.global.proxy.image | string | `"ghcr.io/mosn/htnn-proxy:dev"` |  |
| istiod.pilot.env.HTNN_ENABLE_LDS_PLUGIN_VIA_ECDS | string | `"false"` |  |
| istiod.pilot.env.PILOT_ENABLE_HTNN | string | `"true"` |  |
| istiod.pilot.env.PILOT_ENABLE_HTNN_STATUS | string | `"true"` |  |
| istiod.pilot.env.PILOT_SCOPE_GATEWAY_TO_NAMESPACE | string | `"true"` |  |
| istiod.pilot.hub | string | `"ghcr.io/mosn"` |  |
| istiod.pilot.image | string | `"htnn-controller"` |  |
| istiod.pilot.tag | string | `"dev"` |  |
| istiod.revision | string | `""` |  |

