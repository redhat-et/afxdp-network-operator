# afxdp-network-operator

The AF_XDP Network Operator is designed to help the user to provision
and configure AF_XDP CNI plugin and Device plugin in a K8s cluster.

## Motivation

AF_XDP requires a number of components in order to be provisioned and
configured appropriately. An operator is a natural choice for a central
point to coordinate the relevant components in one place.

## Goals

- To provision, advertise and manage a set of interfaces for Pods that want to use AF_XDP.
- Pods should be able to run without any special privileges.
- Pods should be able to take advantage of PFs, VFs and SFs.
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