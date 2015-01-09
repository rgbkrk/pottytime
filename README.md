Potty Time
==========

IP, UDP, TCP?

Monitor bandwidth on an interface with some Go code. This version has a default BPF to ignore Rackspace's ServiceNet.

```
docker run -it rgbkrk/pottytime 
```

Customizing the interface to sniff on

```
docker run -it --net=host rgbkrk/pottytime -device bond0.101
```
