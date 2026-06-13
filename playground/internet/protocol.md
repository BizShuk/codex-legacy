# Protocol


### TTL , Time-to-Live
Tell router should be discarded if the packet has been too long in the network.

Value will be 1 ~ 255 and descrese each time going through a router. 

Prevent wrong packets stuck in routing loops. 

After it's discarded, it'll send a ICMP back to the originating host. Maybe trigger a resend from the original host.

`ping` and `traceroute` both use TTL 



In IP multicast, the TTL controls the scope or range in which a packet may be forwarded. By convention:

    0 is restricted to the same host
    1 is restricted to the same subnet
    32 is restricted to the same site
    64 is restricted to the same region
    128 is restricted to the same continent
    255 is unrestricted

For DNS server:
    
    describe a caching time for domain name

For HTTP:
    
    describe web content caching time

### ICMP Internet Control Message Protocol


### MPLS , Multiprotocol Label Switching