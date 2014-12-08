Local statsd
============

Statsd is cute, but it doesn't handle authentication nor security.
DTLS is hard, and out of the statsd scope.
VLAN or VPN is not always an option.

You can use on statsd server per node, but the result will not be consolidated, you have to use some centralized daemon.

Here is a fake statsd agent, listening only localhost, and sending values, via TCP to one statsd (to rules them all).
The TCP connection can be TLS with certificates authentication, talking to a stunnel proxy.


    +--------------------+            +-----------------+
    | node A             |    TLS     | node B          |
    | agent->localstatsd-+----------->+-stunnel->statsd |
    |                    |            |                 |
    +--------------------+            +-----------------+

Build it
--------

    docker build

Status
------

 * √ listening UDP
 * √ connecting an reconnecting TCP
 * _ CLI options
 * _ TLS authentication
 * _ counting its own latencies

Licence
-------

3 terms BSD licence © Mathieu Lecarme 2014
