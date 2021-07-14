---
title: "Tinygo flash errors"
weight: 11
description: |
  Solving common problems when flashing your first TinyGo program.
---

Flashing to a port may be a source of pain with a unconfigured computer. Here are some common problems when trying to get your first program on a device.

* [Serial port permission problems (Linux)](#permission-denied-linux)

### Permission Denied (Linux)

This error is usually encountered when flashing your first program, here's a typical error:
```
avrdude: ser_open(): can't open device "/dev/ttyACM0": Permission denied
```

The above error may be followed by `ioctl` errors too.

To fix this we must make sure our user is included in the group with access to serial ports. On ubuntu and Debian this is the **dialout** group.

First we can check if we are not already part of the `dialout` group:

```shell
groups
```
The output should look something like `youruser adm cdrom sudo dip plugdev lpadmin lxd sambashare`. If dialout is not present in the groups we can add ourselves to it with the following command:

```shell
sudo usermod -aG dialout $(whoami)
```
This will add our current user to dialout group. **Log out and log in again to complete the procedure**. You should now be able to flash to your device. 

As a last resort if this does not work you can try to modify the permissions with `chmod`. For the error above we may try `sudo chmod a+rw /dev/ttyACM0` where `/dev/ttyACM0` is the serial port we are trying to flash to.
