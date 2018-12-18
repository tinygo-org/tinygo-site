---
title: "Usage"
weight: 3
---

TinyGo should now be installed. Test it by running a test program:

    tinygo run examples/test

Before anything can be built for a bare-metal target, you need to generate some
files first:

    make gen-device

This will generate register descriptions, interrupt vectors, and linker scripts
for various devices. Also, you may need to re-run this command after updates,
as some updates cause changes to the generated files.

Now you can run a blinky example. For the PCA10040 [https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF52-DK](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF52-DK)
development board:

    tinygo flash -target=pca10040 examples/blinky2

Or for an Arduino Uno [https://store.arduino.cc/arduino-uno-rev3](https://store.arduino.cc/arduino-uno-rev3)

    tinygo flash -target=arduino examples/blinky1
