---
title: "Tinygo flash errors"
weight: 11
description: |
  Solving common problems when flashing your first TinyGo program.
---

Flashing to a port may be a source of pain with a unconfigured computer. Here are some common problems when trying to get your first program on a device.

* [Unable to locate any volume (Linux)](#unable-to-locate-any-volume-inux)
* [Serial port permission problems (Linux)](#permission-denied-linux)

### Unable to locate any volume (Linux)

This error happens with devices such as the Pi Pico (and the Pico-W) which use the UF2 protocol for flashing programs.  With this protocol, the device 
emulates a USB block storage device that can be mounted. When you copy a properly formatted flash image onto this virtual disk drive, the side effect
is to flash that image onto the device.

For the pico with `target=pico`, `tinygo` expects to see this device in the current directory under the name `RPI-RP2`. If this volume is is not found,
you will get an error message like this:
```bash
failed to flash /tmp/tinygo173718394/main.uf2: unable to locate any volume: [RPI-RP2]
```

To fix this problem, you need to mount the volume after restarting the Pico while holding down the `BOOTSEL` button. You can do this manually, but it
is better to edit the `/etc/fstab` file to allow users to do this themselves. You can do this by adding a line like this to `/etc/fstab`:
```bash
LABEL=RPI-RP2  /path/to/volume/RPI-RP2  vfat	defaults,umask=000,users
```
where you replace `/path/to/volume` with a pathname of your choosing. It is common for this to be put in your home directory or in a system standard
place such as `/mnt/RPI-RP2`, but it can also be in the directory where you expect to run `tinygo`. You will also need to create a directory with the path 
`/path/to/volume/RPI-RP2`. Once you have done this and after restarting the Pico in
`BOOTSEL` mode, you should be able to see the Pico's block device using this command
```bash
lsblock -f
```
On a typical system, the Pico will show up like this
```bash
sda                                                                                  
└─sda1      vfat     RPI-RP2     000A-9BF3
```
you can mount the Pico using this command:
```bash
mount /path/to/volume/RPI-RP2
```
You can verify that this has worked using this command:
```bash
mount | grep RPI
```
If you get output that looks like this, you have succeeded in mounting the volume
```bash
/dev/sda1 on /mnt/RPI-RP2 type vfat (rw,nosuid,...lots more stuff...)
```
If you mounted your Pico someplace other than in your development directory, you should create a symbolic
link to the mounted volume in the directory where you plan to run `tinygo`. This can be done using
```bash
ln -s /mnt/RPI-RP2/
```
Note that editing `/etc/fstab`, creating the mount point directory with `mkdir` and creating your symbolic
link all only need to be done once. During development, you will only need to do the `mount` command
after restarting the Pico. After flashing the Pico, it is good practice to unmount the drive using the command
```bash
umount /path/to/volume/RPI-RP2
```
### Permission Denied (Linux)

This error is usually encountered when flashing your first program, here's a typical error:
```
avrdude: ser_open(): can't open device "/dev/ttyACM0": Permission denied
```

The above error may be followed by `ioctl` errors too.

To fix this we must make sure our user is included in the group with access to serial ports. On ubuntu and Debian this is the **dialout** group. The equivalent group on Arch Linux based distributions is **uucp**.

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

### Disable "not ejected safely" notification (Mac)

Some boards are flashed by copying over `.uf2` file to the board mounted as Mass Storage Device.  
Upon successful flashing, the board reboots and this looks to the OS as the board was _unsafely_ ejected (yanked out by user), so OS warns with a notification that has to be acknowledged (it does not disappear itself).  
This can become annoying very quickly during intense flash-run-debug sessions. Thankfully, it is possible to disable this behaviour.  

In terminal, execute following and **restart the system**:
```
sudo defaults write /Library/Preferences/SystemConfiguration/com.apple.DiskArbitration.diskarbitrationd.plist DADisableEjectNotification -bool YES && sudo pkill diskarbitrationd
```

Revert to the default behaviour:
```
sudo defaults delete /Library/Preferences/SystemConfiguration/com.apple.DiskArbitration.diskarbitrationd.plist DADisableEjectNotification && sudo pkill diskarbitrationd
```
