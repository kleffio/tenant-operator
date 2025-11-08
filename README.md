# Kubernetes Tenant Operator

## Overview

The Kubernetes Tenant Operator is a specialized controller designed for a multi-tenant Kubernetes-based web hosting Platform as a Service (PaaS). It streamlines the management of tenants by automating the provisioning and de-provisioning of their dedicated namespaces.

This operator introduces a `Tenant` Custom Resource Definition (CRD), which represents a single tenant on the platform. When a new `Tenant` resource is created, the operator's controller springs into action, creating a unique namespace for that tenant. This namespace is appropriately labeled to identify the tenant and enable further configurations, such as Istio sidecar injection.

Conversely, when a `Tenant` resource is deleted, the operator ensures a clean teardown by removing the corresponding namespace and all its associated resources. This is managed gracefully through the use of a finalizer, preventing the `Tenant` resource from being fully deleted until its namespace has been successfully removed.

## Features

*   **Automated Namespace Provisioning**: Automatically creates a dedicated namespace for each new `Tenant` custom resource.
*   **Tenant-Specific Labeling**: Applies labels to tenant namespaces for easy identification and policy enforcement, including:
    *   `istio-injection`
    *   `kleff.io/tenant-username`
    *   `kleff.io/tenant-plan`
*   **Graceful Deletion with Finalizers**: Utilizes a finalizer (`kleff.io/namespace-finalizer`) to ensure that a tenant's namespace is deleted before the `Tenant` resource is removed from the cluster.
*   **Status Conditions**: Updates the `status` of the `Tenant` resource with conditions like `NamespaceReady` and `TenantReady` to reflect the state of the provisioning process.

## Custom Resource Definition (CRD)

The core of the Tenant Operator is the `Tenant` custom resource. Below is an example of a `Tenant` resource definition:

```yaml
apiVersion: kleff.kleff.io/v1
kind: Tenant
metadata:
  name: example-tenant
spec:
  plan: "basic"
  userId: "user-12345"
  username: "exampleuser"
```



## Roadmap: Planned Enhancements

This operator is designed to be extensible. Future versions will deepen tenant isolation and automate more aspects of the PaaS environment:

*   **OPA Gatekeeper Integration**: Automatically apply custom security and governance policies to each tenant's namespace to enforce rules like trusted image registries or pod security standards.
*   **Automated Istio Configuration**: Generate tenant-specific Istio resources, such as `Gateways` and `AuthorizationPolicies`, to manage ingress traffic and secure service-to-service communication within the mesh.
*   **Dynamic Resource Management**: Create `ResourceQuota` and `LimitRange` objects based on the tenant's selected `plan` to ensure fair resource allocation and maintain cluster stability.