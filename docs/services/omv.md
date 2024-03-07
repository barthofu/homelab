# OMV

OpenMediaVault (OMV) is installed on the VM `nas` using Ansible.

## Configuraton

1. Connects to the web UI
2. Change the `admin` password
3. Go to `Network` > `Interfaces`
4. Add a new **ethernet** interface with those parameters:
   - **Device**: `eth0`
   - **IPv4**:
     - **Method**: `static`
     - **Adress**: *the IP adress of the nas VM (e.g: 192.168.1.110)*
     - **Gateway**: *the gateway adress (e.g: 192.168.1.254)*
5. Click on the save button

Now you can setup *users* and *disks* how you want! See the [storage doc](../storage.md) to know the disks layouts and storage organization.