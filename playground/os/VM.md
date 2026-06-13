# Virtual Machine

### HVM, Hardware Virtual Machine
Fully virtualized hardwares, taking advantages of hardware extensions that provide fast access. For example of hardware extensions, CPU virtualization extension, enhanced networking extension, GPU processing extension.

### PV, ParaVirtual
Boot with specific loader called `PV-GRUB`. Historically, this has better performance than HVM. However, after HVM can use hardware extensions, it is not longer true.

