# Network Upgrade: Uluwatu

This guide describes the process to participate in the *critical coordinated* Omni Omega  “Uluwatu” network upgrade (hard fork).  (Omni network upgrades are named after iconic surf spots; Uluwatu is in Bali).

## TL;DR

- Simply ensure the `omniops/halovisor:v0.9.0` docker image is running **BEFORE** the upgrade height.
  - `halovisor:v0.9.0`: wraps cosmovisor with `halo:v0.8.1` and `halo:v0.9.0`
  - It will perform the binary switch automatically at the required block.
- Omega upgrade height: TBD
- Approximate upgrade date: 7~11 Oct 2024
- Version(s) supported before upgrade: `halo:v0.4.0 .. v0.8.1`
- Version required after upgrade: `halo:v0.9.0`

> 🚧 Like any blockchain network upgrade (hard fork), nodes that do not upgrade will crash or stall.

## Details

The “uluwatu” upgrade is the first network upgrade (hard fork) planned for the Omni Omega network and is included in the `halo:v0.9.0` release.

The upgrade contains changes to `halo`’s `attest` module logic ensuring that attestations are only deleted when they exit the modified vote window. See [issue](https://github.com/omni-network/omni/issues/1787) and [PR](https://github.com/omni-network/omni/pull/1983) for details.

No changes to `geth` is required, this version is compatible with `v1.14.11`.

See [Run a Full Node](./1-run-full-node.md#halo-deployment-instructions) for more details on running `halo` with `cosmovisor`.

See the [Operator FAQ](./5-faq.md)  for details on `halovisor vs halo` and `docker vs binaries`
