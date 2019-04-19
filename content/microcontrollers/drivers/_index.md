---
title: "Drivers"
chapter: true
weight: 90
---

# Drivers

TinyGo has driver support for 14 different sensors and other devices, such as digital accelerometers and multicolor LEDs.

All of the drivers code is in the TinyGo Drivers repository located at [https://github.com/tinygo-org/drivers/](https://github.com/tinygo-org/drivers/).

| Device Name | Interface Type |
|----------|-------------|
| [ADXL345 accelerometer](http://www.analog.com/media/en/technical-documentation/data-sheets/ADXL345.pdf) | I2C |
| [APA102 RGB LED](https://cdn-shop.adafruit.com/product-files/2343/APA102C.pdf) | SPI |
| [BH1750 ambient light sensor](https://www.mouser.com/ds/2/348/bh1750fvi-e-186247.pdf) | I2C |
| [BlinkM RGB LED](http://thingm.com/fileadmin/thingm/downloads/BlinkM_datasheet.pdf) | I2C |
| [BMP180 barometer](https://cdn-shop.adafruit.com/datasheets/BST-BMP180-DS000-09.pdf) | I2C |
| [DS3231 real time clock](https://datasheets.maximintegrated.com/en/ds/DS3231.pdf) | I2C |
| ["Easystepper" stepper motor controller](https://en.wikipedia.org/wiki/Stepper_motor) | GPIO |
| [ESP8266/ESP32 AT Command set for WiFi/TCP/UDP](https://github.com/espressif/esp32-at) | UART |
| [LIS3DH accelerometer](https://www.st.com/resource/en/datasheet/lis3dh.pdf) | I2C |
| [MAG3110 magnetometer](https://www.nxp.com/docs/en/data-sheet/MAG3110.pdf) | I2C |
| [MMA8653 accelerometer](https://www.nxp.com/docs/en/data-sheet/MMA8653FC.pdf) | I2C |
| [MPU6050 accelerometer/gyroscope](https://store.invensense.com/datasheets/invensense/MPU-6050_DataSheet_V3%204.pdf) | I2C |
| [VL53L1X time-of-flight distance sensor](https://www.st.com/resource/en/datasheet/vl53l1x.pdf) | I2C |
| [WS2812 RGB LED](https://cdn-shop.adafruit.com/datasheets/WS2812.pdf) | GPIO |

We also give you the ability to add new drivers. If your device isn't listed here, please raise an issue in the [issue tracker](https://github.com/tinygo-org/drivers/issues).
