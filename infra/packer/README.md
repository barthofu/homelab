# Homelab / Packer

This directory holds the configuration of [Packer](https://www.packer.io/) to build the base images for my homelab infrastructure.

## Table of contents

- [Table of contents](#table-of-contents)
- [Requirements](#requirements)
- [Usage](#usage)

## Requirements

- [Packer](https://www.packer.io/downloads)

## Usage

```bash
cd debian-12
packer build --var-file ../credentials.pkrvars.hcl .
```