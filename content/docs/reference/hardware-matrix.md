---
title: " Supported hardware feature matrix"
weight: 10
description: >
  Matrix between Microcontrollers and Features
---

Notes:

* '-': Not supported by Hardware
* '?': Supported status is currently unknown
* 'n': Not yet supported by TinyGo
* 'p': Partially supported by TinyGo
* 'x': Supported

see TinyGo documentation of the microcontroller for details to:

* WiFi
* BT (Bluetooth)
* IMU (Inertial measurement unit, e.g. acceleration, rotation, magnetometer)
* NePx (NeoPixel, WS2812)
* other (other peripheral or built-in devices, e.g. temperature, GSM)

| Microcontroller                             |GPIO|UART|SPI|I2C|ADC|PWM|USBDev|BT |WiFi|IMU|NePx|other|
|:-------------------------------------------:|:--:|:--:|:-:|:-:|:-:|:-:|:----:|:-:|:--:|:-:|:--:|:---:|
| [Adafruit Circuit Playground Bluefruit](https://www.adafruit.com/product/4333)| x  | x  | x | x | x | x | x    | x | -  | x | x  | p   |
| [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333)  | x  | x  | x | x | x | x | x    | - | -  | x | x  | p   |
| [Adafruit CLUE](https://www.adafruit.com/product/4500)                        | x  | x  | x | x | x | x | x    | x | -  | x | x  | p   |
| [Adafruit Feather M0](https://www.adafruit.com/product/3403)                  | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Adafruit Feather M4](https://www.adafruit.com/product/3857)                  | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit Feather M4 CAN](https://www.adafruit.com/product/4759)              | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit Feather nRF52840 Express](https://www.adafruit.com/product/4062)    | x  | x  | x | x | x | x | x    | x | -  | - | x  | -   |
| [Adafruit Feather nRF52840 Sense](https://www.adafruit.com/product/4516)      | x  | x  | x | x | x | x | x    | x | -  | x | x  | p   |
| [Adafruit Feather RP2040](https://www.adafruit.com/product/4884)              | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Adafruit Feather STM32F405 Express](https://www.adafruit.com/product/4382)   | x  | x  | x | x | x | n | n    | - | -  | - | x  | -   |
| [Adafruit Grand Central M4](https://www.adafruit.com/product/4064)            | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit ItsyBitsy M0](https://www.adafruit.com/product/3727)                | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Adafruit ItsyBitsy M4](https://www.adafruit.com/product/3800)                | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Adafruit ItsyBitsy nRF52840 Express](https://www.adafruit.com/product/4481)  | x  | x  | x | x | x | x | x    | x | -  | - | n  | -   |
| [Adafruit MacroPad RP2040](https://www.adafruit.com/product/5100)             | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit Matrix Portal M4](https://www.adafruit.com/product/4745)            | x  | x  | x | x | x | x | x    | - | n  | x | x  | p   |
| [Adafruit Metro M4 Express AirLift](https://www.adafruit.com/product/4000)    | x  | x  | x | x | x | x | x    | - | n  | - | x  | -   |
| [Adafruit PyBadge](https://www.adafruit.com/product/4200)                     | x  | x  | x | x | x | x | x    | - | -  | x | x  | p   |
| [Adafruit PyGamer](https://www.adafruit.com/product/4242)                     | x  | x  | x | x | x | x | x    | - | -  | x | x  | p   |
| [Adafruit PyPortal](https://www.adafruit.com/product/4116)                    | x  | x  | x | x | x | x | x    | - | n  | x | x  | p   |
| [Adafruit QT Py](https://www.adafruit.com/product/4600)                       | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit QT Py ESP32-C3](https://www.adafruit.com/product/5405)              | x  | x  | x | x | n | n | n    | n | n  | - | x  | -   |
| [Adafruit QT Py RP2040](https://www.adafruit.com/product/4900)                | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Adafruit Trinket M0](https://www.adafruit.com/product/3500)                  | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Adafruit Trinkey QT2040](https://www.adafruit.com/product/5056)              | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [Arduino Mega 1280](https://docs.arduino.cc/retired/other/arduino-older-boards#arduino-mega/)                    | x  | x  | x | x | x | x | -    | - | -  | - | -  | -   |
| [Arduino Mega 2560](https://store.arduino.cc/arduino-mega-2560-rev3)                    | x  | x  | x | x | x | n | -    | - | -  | - | -  | -   |
| [Arduino MKR WiFi 1010](https://store.arduino.cc/usa/mkr-wifi-1010)                | x  | x  | x | x | x | x | x    | x | x  | - | -  | -   |
| [Arduino MKR1000](https://store.arduino.cc/arduino-mkr1000-wifi)                      | x  | x  | x | x | x | x | x    | x | x  | - | -  | -   |
| [Arduino Nano](https://store.arduino.cc/arduino-nano)                         | x  | x  | x | x | x | x | -    | - | -  | - | -  | -   |
| [Arduino Nano 33 BLE (Sense)](https://store.arduino.cc/arduino-nano-33-ble)          | x  | x  | x | x | x | x | x    | x | -  | x | -  | p   |
| [Arduino Nano 33 IoT](https://store.arduino.cc/nano-33-iot)                  | x  | x  | x | x | x | x | x    | x | x  | x | -  | p   |
| [Arduino Nano RP2040 Connect](https://store.arduino.cc/nano-rp2040-connect)          | x  | x  | x | x | x | x | x    | x | x  | x | -  | p   |
| [Arduino Uno](https://store.arduino.cc/arduino-uno-rev3)                          | x  | x  | x | x | x | x | -    | - | -  | - | -  | -   |
| [Arduino Zero](https://store.arduino.cc/arduino-zero)                         | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [BBC micro:bit](https://microbit.org)                        | x  | x  | x | x | x | n | -    | x | -  | n | -  | p   |
| [Blues Wireless Swan](https://blues.io/products/swan/)                  | x  | x  | x | x | n | n | n    | - | -  | - | -  | p   |
| [Digispark](http://digistump.com/products/1)                            | x  | n  | n | n | x | n | -    | - | -  | - | -  | -   |
| [Dragino LoRaWAN GPS Tracker LGT-92](https://www.dragino.com/products/lora-lorawan-end-node/item/142-lgt-92.html)   | x  | x  | x | x | n | n | n    | - | -  | n | -  | -   |
| [ESP32 - mini32](https://www.lilygo.cc/en-pl/products/t7-v1-3-mini-32-esp32)                       | x  | x  | x | n | n | n | -    | n | n  | - | -  | -   |
| [ESP32 Core Board V2](https://docs.espressif.com/projects/esp-idf/en/release-v3.0/hw-reference/modules-and-boards.html#esp32-core-board-v2-esp32-devkitc)                  | x  | x  | x | n | n | n | -    | n | n  | - | -  | -   |
| [ESP8266 - d1mini](https://botland.store/withdrawn-products/6257-d1-mini-wifi-esp8266-iot-compatible-with-wemos-and-arduino.html)                     | x  | x  | n | n | n | n | -    | - | n  | - | -  | -   |
| [ESP8266 - NodeMCU](https://en.wikipedia.org/wiki/NodeMCU)                    | x  | x  | n | n | n | n | -    | - | n  | - | -  | -   |
| [Game Boy Advance](https://en.wikipedia.org/wiki/Game_Boy_Advance)                     | ?  | ?  | ? | ? | ? | ? | ?    | - | -  | - | -  | -   |
| [iLabs Challenger RP2040 LoRa](https://ilabs.se/product/challenger-rp2040-lora/)         | x  | x  | x | x | x | x | x    | - | -  | - | -  | p   |
| [M5Stack BASIC Kit](https://docs.m5stack.com/en/core/basic)                    | x  | x  | x | n | n | n | -    | n | n  | - | -  | -   |
| [M5Stack Core2](https://shop.m5stack.com/products/m5stack-core2-esp32-iot-development-kit)                        | x  | x  | x | n | n | n | -    | n | n  | - | -  | -   |
| [M5Stack M5Stamp-C3](https://docs.m5stack.com/en/core/stamp_c3)                   | x  | x  | x | n | n | n | n    | n | n  | - | x  | -   |
| [Makerdiary nRF52840-MDK](https://wiki.makerdiary.com/nrf52840-mdk/)              | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Makerdiary nRF52840-MDK USB Dongle](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/)   | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Makerfabs ESP32-C3 SPI](https://wiki.makerfabs.com/ESP32_C3_SPI_3.5_TFT_with_Touch.html)               | x  | x  | x | x | n | n | n    | n | n  | - | -  | p   |
| [Microchip SAM E54 Xplained Pro](https://www.microchip.com/developmenttools/productdetails/atsame54-xpro)       | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Nice Keyboards nice!nano](https://nicekeyboards.com/products/nice-nano-v1-0)             | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Nintendo Switch](https://en.wikipedia.org/wiki/Nintendo_Switch)                      | ?  | ?  | ? | ? | ? | ? | ?    | - | -  | - | -  | -   |
| [Nordic Semiconductor PCA10031](https://www.nordicsemi.com/eng/Products/nRF51-Dongle)        | x  | x  | x | x | x | n | -    | x | -  | - | -  | -   |
| [Nordic Semiconductor PCA10040](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF52-DK)        | x  | x  | x | x | x | x | -    | x | -  | - | -  | -   |
| [Nordic Semiconductor PCA10056](https://www.nordicsemi.com/Software-and-Tools/Development-Kits/nRF52840-DK)        | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Nordic Semiconductor PCA10059](https://www.nordicsemi.com/Software-and-tools/Development-Kits/nRF52840-Dongle)        | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Particle Argon](https://docs.particle.io/datasheets/wi-fi/argon-datasheet/)                       | x  | x  | x | x | x | x | x    | x | n  | - | -  | -   |
| [Particle Boron](https://docs.particle.io/datasheets/cellular/boron-datasheet/)                       | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Particle Xenon](https://docs.particle.io/datasheets/discontinued/xenon-datasheet/)                       | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Phytec reel board](https://www.phytec.eu/product-eu/internet-of-things/reelboard/)                    | x  | x  | x | x | x | x | x    | x | -  | n | -  | p   |
| [Pimoroni Badger2040](https://shop.pimoroni.com/products/badger-2040)                  | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Pimoroni Tufty2040](https://shop.pimoroni.com/products/tufty-2040)                   | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Pine64 PineTime](https://wiki.pine64.org/index.php/PineTime)                      | x  | x  | x | x | x | x | -    | x | -  | - | -  | -   |
| [PJRC Teensy 3.6](https://www.pjrc.com/store/teensy36.html)                      | x  | x  | n | n | n | n | n    | - | -  | - | -  | -   |
| [PJRC Teensy 4.0](https://www.pjrc.com/store/teensy40.html)                      | x  | x  | x | x | x | n | n    | - | -  | - | -  | -   |
| [ProductivityOpen P1AM-100](https://facts-engineering.github.io/modules/P1AM-100/P1AM-100.html)            | x  | x  | x | x | x | x | x    | x | x  | - | -  | -   |
| [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/)                    | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Seeed LoRa-E5 Development Kit](https://www.seeedstudio.com/LoRa-E5-Dev-Kit-p-4868.html)        | x  | x  | x | x | n | n | -    | - | -  | - | -  | p   |
| [Seeed Seeeduino XIAO](https://www.seeedstudio.com/Seeeduino-XIAO-Arduino-Microcontroller-SAMD21-Cortex-M0+-p-4426.html)                 | x  | x  | x | x | x | x | x    | - | -  | - | -  | -   |
| [Seeed Sipeed MAix Bit](https://www.seeedstudio.com/Sipeed-MAix-BiT-for-RISC-V-AI-IoT-p-2872.html)                | x  | x  | x | x | - | n | ?    | - | -  | - | -  | -   |
| [Seeed Wio Terminal](https://www.seeedstudio.com/Wio-Terminal-p-4509.html)                   | x  | x  | x | x | x | x | x    | n | n  | n | -  | p   |
| [Seeed XIAO BLE](https://www.seeedstudio.com/Seeed-XIAO-BLE-nRF52840-p-5201.html)                       | x  | x  | x | x | x | x | x    | x | -  | - | -  | -   |
| [Seeed XIAO ESP32 C3](https://www.seeedstudio.com/Seeed-XIAO-ESP32C3-p-5431.html)                  | x  | x  | x | x | n | n | n    | n | n  | - | -  | -   |
| [Seeed XIAO RP2040](https://www.seeedstudio.com/XIAO-RP2040-v1-0-p-5026.html)                    | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [SiFive HiFive1 Rev B](https://www.sifive.com/boards/hifive1-rev-b)                 | x  | x  | x | x | - | n | ?    | - | -  | - | -  | -   |
| [Sparkfun Thing Plus RP2040](https://www.sparkfun.com/products/17745)           | x  | x  | x | x | x | x | x    | - | -  | - | n  | p   |
| [ST Micro STM32 "Blue Pill" F103XX](https://stm32-base.org/boards/STM32F103C8T6-Blue-Pill.html)    | x  | x  | x | x | x | n | n    | - | -  | - | -  | -   |
| [ST Micro STM32 Nucleo-144 F722ZE](https://www.st.com/en/evaluation-tools/nucleo-f722ze.html)     | x  | x  | n | x | n | n | n    | - | -  | - | -  | p   |
| [ST Micro STM32 Nucleo-144 L552ZE](https://www.st.com/en/evaluation-tools/nucleo-l552ze-q.html)     | x  | x  | n | x | n | n | n    | - | -  | - | -  | p   |
| [ST Micro STM32 Nucleo-32 L031K6](https://www.st.com/en/evaluation-tools/nucleo-l031k6.html)      | x  | x  | x | x | n | n | -    | - | -  | - | -  | p   |
| [ST Micro STM32 Nucleo-32 L432KC](https://www.st.com/en/evaluation-tools/nucleo-l432kc.html)      | x  | x  | x | x | n | n | -    | - | -  | - | -  | p   |
| [ST Micro STM32 Nucleo-64 F103RB](https://www.st.com/en/evaluation-tools/nucleo-f103rb.html)      | x  | x  | x | x | x | n | -    | - | -  | - | -  | -   |
| [ST Micro STM32F4 Discovery](https://www.st.com/en/evaluation-tools/stm32f4discovery.html)           | x  | x  | x | x | x | n | n    | - | -  | n | -  | p   |
| [ST Micro STM32WL Nucleo-64 WL55JC](https://www.st.com/en/evaluation-tools/nucleo-wl55jc.html)    | x  | x  | x | x | n | n | -    | - | -  | - | -  | p   |
| [Waveshare RP2040-Zero](https://www.waveshare.com/wiki/RP2040-Zero)                | x  | x  | x | x | x | x | x    | - | -  | - | x  | -   |
| [X9 Pro Smartwatch](../microcontrollers/x9pro.md)                    | ?  | ?  | ? | ? | ? | ? | ?    | - | -  | - | -  | -   |
