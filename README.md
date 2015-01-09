Potty Time
==========

IP, UDP, TCP?

Monitor bandwidth on an interface with some Go code, breaking it down by IP blocks.

```
docker run -it rgbkrk/pottytime 
```

Customizing the interface to sniff on

```
docker run -it --net=host rgbkrk/pottytime -device bond0.101
```
