# afxdp-network-operator

The AF_XDP Network Operator is designed to help the user to provision
and configure AF_XDP CNI plugin and Device plugin in a Kubernetes cluster.

## Motivation

AF_XDP requires a number of components in order to be provisioned and
configured appropriately. An operator is a natural choice for a central
point to coordinate the relevant components in one place.

## Goals

- To provision, advertise and manage a set of interfaces for Pods that want to use AF_XDP.
- Pods should be able to run without any special privileges.
- Pods should be able to take advantage of PFs, PFs/SFs and SFs.
- To interwork with BPFd (eBPF lifecycle management daemon).

## Targetted Functionality

- Initialize the supported AF_XDP NIC types on selected nodes.
- Provision/upgrade AF_XDP device plugin executable on selected node.
- Provision/upgrade AF_XDP CNI plugin executable on selected nodes.
- Manage configuration of AF_XDP device plugin on host.
- Generate net-att-def CRs for AF_XDP CNI plugin
- Work with BPFd to load and unload BPF progs on the netdevs
- Supports operation in a virtualized Kubernetes deployment - Needs more investigation

## AF_XDP CNI and Device plugin

Existing plugins to integrate with https://github.com/intel/afxdp-plugins-for-kubernetes

More information will be added as the operator is developed.

## API

Similar to the SR-IOV network operator, the AF_XDP network operator introduces following CRDs:

- AfxdpNetwork

- AfxdpNetworkNodeState

- AfxdpNetworkNodePolicy

### AfxdpNetwork

A custom resource of AfxdpNetwork could represent the a layer-2 broadcast domain where
some AF_XDP netdevices are attach to. It is primarily used to generate a NetworkAttachmentDefinition
CR with an AF_XDP CNI plugin configuration.

This AfxdpNetwork CR also contains the `poolName` which is aligned with the `poolName` of
AF_XDP device plugin. One AfxdpNetwork obj maps to one `poolName`, but one `poolName` can
be shared by different AfxdpNetwork CRs.

This CR should be managed by cluster admin. Here is an example:

```yaml
apiVersion: afxdpnetwork.hns.dev/v1aplha1
kind: AfxdpNetwork
metadata:
  name: example-network
  namespace: example-namespace
spec:
  ipam: |
    {
      "type": "afxdp",
      "subnet": "10.56.217.0/24",
      "rangeStart": "10.56.217.171",
      "rangeEnd": "10.56.217.181",
      "routes": [{
        "dst": "0.0.0.0/0"
      }],
      "gateway": "10.56.217.1"
    }
  poolName: mypool
```

### AfxdpNetworkNodeState

The custom resource to represent the AF_XDP interface states of each host, which should only
be managed by the operator itself.

- The `spec` of this CR represents the desired configuration which should be applied to the
  interfaces and AF_XDP device plugin.
- The `status` contains current states of those PFs/SFs (baremetal only), and the states of
  the PFs/SFs. It helps user to discover AF_XDP network hardware on node, or attached PFs/SFs
  in the case of a virtual deployment.

The spec is rendered by afxdp-policy-controller, and consumed by afxdp-config-daemon. afxdp-config-daemon
is responsible for updating the `status` field to reflect the latest status, this information
can be used as input to create AfxdpNetworkNodePolicy CR.

An example of AfxdpNetworkNodeState CR:

```yaml
apiVersion: afxdpnetwork.hns.dev/v1aplha1
kind: AfxdpNetworkNodeState
metadata:
  name: worker-node-1
  namespace: afxdp-network-operator
spec:
  interfaces:
  - deviceType: netdev
    mtu: 1500
    numSfs: 4
    pciAddress: 0000:86:00.0
status:
  interfaces:
  - deviceID: "1583"
    driver: ice
    mtu: 1500
    numSfs: 4
    pciAddress: 0000:86:00.0
    maxVfs: 64
    vendor: "8086"
    Sfs:
      - deviceID: 154c
      driver: ice
      pciAddress: 0000:86:02.0
      vendor: "8086"
      - deviceID: 154c
      driver: ice
      pciAddress: 0000:86:02.1
      vendor: "8086"
      - deviceID: 154c
      driver: vfio-pci
      pciAddress: 0000:86:02.2
      vendor: "8086"
      - deviceID: 154c
      driver: vfio-pci
      pciAddress: 0000:86:02.3
      vendor: "8086"
```

From this example, in status field, the user can find out there are 2 SRIOV capable NICs
on node 'work-node-1'; in spec field, user can learn what the expected configure is
generated from the combination of AfxdpNetworkNodePolicy CRs. **In the virtual deployment
case, a single VF will be associated with each device.**

### AfxdpNetworkNodePolicy

This CRD is the key of AF_XDP network operator. This custom resource should be managed
by cluster admin, to instruct the operator to:

1. Render the spec of AfxdpNetworkNodeState CR for selected node, to configure the
   AF_XDP interfaces.
2. Deploy AF_XDP CNI plugin and device plugin on selected node.
3. Generate the configuration of AF_XDP device plugin.

An example of AfxdpNetworkNodePolicy CR:

```yaml
apiVersion: afxdpnetwork.hns.dev/v1aplha1
kind: AfxdpNetworkNodePolicy
metadata:
  name: policy-1
  namespace: afxdp-network-operator
spec:
  deviceType: vfio-pci
  mtu: 1500
  nicSelector:
    deviceID: "1583"
    rootDevices:
    - 0000:86:00.0
    vendor: "8086"
  nodeSelector:
    feature.node.kubernetes.io/network-afxdp.capable: "true"
  numSfs: 4
  priority: 90
  poolName: intelnics
```

In this example, user selected the nic from vendor `8086` which is intel, device
module is `1583` which is XL710 for 40GbE, on nodes labeled with `network-afxdp.capable`
equals 'true'. Then for those PFs, create 4 PFs/SFs each, set mtu to 1500 and the
load the vfio-pci driver to those virtual functions.

In a virtual deployment:
- TODO...

#### Multiple policies

When multiple AfxdpNetworkNodeConfigPolicy CRs are present, the `priority` field
(0 is the highest priority) is used to resolve any conflicts. Conflicts occur
only when same PF is referenced by multiple policies. The final desired
configuration is saved in `AfxdpNetworkNodeState.spec.interfaces`.

Policies processing order is based on priority (lowest first), followed by `name`
field (starting from `a`). Policies with same **priority** or **non-overlapping
VF groups** (when #-notation is used in pfName field) are merged, otherwise only
the highest priority policy is applied. In case of same-priority policies and
overlapping VF groups, only the last processed policy is applied.

#### Externally Manage virtual functions

When `ExternallyManage` is request on a policy the operator will only skip the virtual
function creation. The operator will only bind the virtual functions to the requested
driver and expose them via the device plugin. Another difference when this field is
requested in the policy is that when this policy is removed the operator will not
remove the virtual functions from the policy.

> **Note:** This means the user must create the virtual functions before they apply the
   policy or the webhook will reject the policy creation.

It's possible to use something like nmstate kubernetes-nmstate or just a simple systemd
file to create the virtual functions on boot.

This feature was created to support deployments where the user want to use some of the
virtual funtions for the host communication like storage network or out of band managment
and the virtual functions must exist on boot and not only after the operator and config-daemon
are running.

## Components and design

This operator is split into 2 components:

- controller
- afxdp-config-daemon

The controller is responsible for:

1. Read the AfxdpNetworkNodePolicy CRs and AfxdpNetwork CRs as input.
2. Render the manifests for AF_XDP CNI plugin and device plugin daemons.
3. Render the spec of AfxdpNetworkNodeState CR for each node.

The afxdp-config-daemon is responsible for:

1. Discover the AF_XDP NICs on each node, then sync the status of AfxdpNetworkNodeState CR.
2. Take the spec of AfxdpNetworkNodeState CR as input to configure those NICs.

## Workflow

TODO
