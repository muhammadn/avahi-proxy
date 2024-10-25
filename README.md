# Multicast DNS Proxy written in Go

This was to solve the problem in trying to proxy all my local "home.lan" domain to Multicast DNS ".local" domain.

Problem was all existing solutions are eiher written in scripting languages such as Ruby and Python which are less performant.

So now, everything that is within Multicast DNS ".local" domain is resolvable with "home.lan" as well.

For example, i have a machine which is "muhammads-macbook-pro.local" which resolves with Multicast DNS (MDNS) but i want the host "muhammads-macbook-pro" to be resolvable as "muhammads-macbook-pro.home.lan" as well.

This helps with machines like Microsoft Windows,Android and Linux (without avahi installed) which does not have Multicast DNS built-in (thus cannot resolve ".local" TLD)

So you can actually run `dnsmasq` and then upstream any `.home.lan` requests to this multicast dns proxy software and return the IP from multicast `.local`.

Example in `dnsmasq.conf`:

```
server=/home.lan/127.0.0.1#5354

```

Run this program by specifying what is your LAN domain name. 

Example:

```
./avahi-proxy run --baseDomain home.lan
```

depending on what domain you configured for your LAN network.

Also you can specify the port which you want `avahi-proxy` to run on.

example:

```
./avahi-proxy run --baseDomain home.lan --port 5355
``` 

Contributions are welcomed.
