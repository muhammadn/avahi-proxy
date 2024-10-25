# Multicast DNS Proxy written in Go

This was to solve the problem in trying to proxy all my local "home.lan" domain to Multicast DNS ".local" domain.

Problem was all existing solutions are eiher written in scripting languages such as Ruby and Python which are less performant.

So now, everything that is within Multicast DNS ".local" domain is resolvable with "home.lan" as well.

For example, i have a machine which is "muhammads-macbook-pro.local" which resolves with Multicast DNS (MDNS) but i want the host "muhammads-macbook-pro" to be resolvable as "muhammads-macbook-pro.home.lan" as well.

This helps with machines like Microsoft Windows which does not have Multicast DNS built-in (thus cannot resolve ".local" TLD)

So you can actually run `dnsmasq` and then upstream any `.home.lan` requests to this multicast dns proxy software and return the IP from multicast `.local`.

Example in `dnsmasq.conf`:

```
server=/home.lan/127.0.0.1#5354

```

Note: this program runs on port 5354

Disclaimer: This is a prototype to get my network fixed. Just a quick and dirty code to get working, flags like changing `.home.lan` or port to something else is still in a work in progress. I intend to use cobra library for those functionalities.

Contributions are welcomed.
