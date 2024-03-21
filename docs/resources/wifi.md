# routeros_wifi (Resource)
*<span style="color:red">This resource requires a minimum version of RouterOS 7.13.</span>*

## Example Usage
```terraform
resource "routeros_wifi" "wifi1" {
  configuration = {
    manager = "capsman"
  }
  name = "wifi1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the interface.

### Optional

- `aaa` (Map of String) AAA inline settings.
- `arp` (String) Address Resolution Protocol mode:
	* disabled - the interface will not use ARP
	* enabled - the interface will use ARP
	* local-proxy-arp - the router performs proxy ARP on the interface and sends replies to the same interface
	* proxy-arp - the router performs proxy ARP on the interface and sends replies to other interfaces
	* reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the ARP table. No dynamic entries will be automatically stored in the ARP table. Therefore for communications to be successful, a valid static entry must already exist.
- `arp_timeout` (String) ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in IP/Settings, default is 30s. Can use postfix ms, s, M, h, d for milliseconds, seconds, minutes, hours or days. If no postfix is set then seconds (s) is used.
- `channel` (Map of String) Channel inline settings.
- `configuration` (Map of String) Configuration inline settings.
- `datapath` (Map of String) Datapath inline settings.
- `disable_running_check` (Boolean) An option to set the running property to true if it is not disabled.
- `disabled` (Boolean)
- `interworking` (Map of String) Interworking inline settings.
- `l2mtu` (Number) Layer2 Maximum transmission unit. [See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).
- `mac_address` (String) MAC address (BSSID) to use for the interface.
- `master_interface` (String) The corresponding master interface of the virtual one.
- `mtu` (Number) Layer3 maximum transmission unit
- `security` (Map of String) Security inline settings.
- `steering` (Map of String) Steering inline settings.

### Read-Only

- `bound` (Boolean) A flag whether the interface is currently available for the WiFi manager.
- `default_name` (String) The interface's default name.
- `id` (String) The ID of this resource.
- `inactive` (Boolean) A flag whether the interface is currently inactive.
- `master` (Boolean) A flag whether the interface is not a virtual one.
- `radio_mac` (String) The MAC address of the associated radio.
- `running` (Boolean) A flag whether the interface has established a link to another device.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/interface/wifi get [print show-ids]]
terraform import routeros_wifi.wifi1 '*1'
```