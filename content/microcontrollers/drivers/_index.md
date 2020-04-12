---
title: "Drivers"
chapter: true
weight: 90
---

# Drivers

TinyGo has driver support for 44 different sensors and devices such as digital accelerometers and multicolor LEDs.

All of the drivers code is in the TinyGo Drivers repository located at [https://github.com/tinygo-org/drivers/](https://github.com/tinygo-org/drivers/).

The following 48 devices are supported.

| Device Name | Interface Type |
|----------|-------------|
| [ADT7410 I2C Temperature Sensor](https://www.analog.com/media/en/technical-documentation/data-sheets/ADT7410.pdf) | I2C |
| [ADXL345 accelerometer](http://www.analog.com/media/en/technical-documentation/data-sheets/ADXL345.pdf) | I2C |
| [AMG88xx 8x8 Thermal camera sensor](https://cdn-learn.adafruit.com/assets/assets/000/043/261/original/Grid-EYE_SPECIFICATIONS%28Reference%29.pdf) | I2C |
| [APA102 RGB LED](https://cdn-shop.adafruit.com/product-files/2343/APA102C.pdf) | SPI |
| [AT24CX 2-wire serial EEPROM](https://www.openimpulse.com/blog/wp-content/uploads/wpsc/downloadables/24C32-Datasheet.pdf) | I2C |
| [BBC micro:bit LED matrix](https://github.com/bbcmicrobit/hardware/blob/master/SCH_BBC-Microbit_V1.3B.pdf) | GPIO |
| [BH1750 ambient light sensor](https://www.mouser.com/ds/2/348/bh1750fvi-e-186247.pdf) | I2C |
| [BlinkM RGB LED](http://thingm.com/fileadmin/thingm/downloads/BlinkM_datasheet.pdf) | I2C |
| [BME280 humidity/pressure sensor](https://cdn-shop.adafruit.com/datasheets/BST-BME280_DS001-10.pdf) | I2C |
| [BMP180 barometer](https://cdn-shop.adafruit.com/datasheets/BST-BMP180-DS000-09.pdf) | I2C |
| [Buzzer](https://en.wikipedia.org/wiki/Buzzer#Piezoelectric) | GPIO |
| [DS1307 real time clock](https://datasheets.maximintegrated.com/en/ds/DS1307.pdf) | I2C |
| [DS3231 real time clock](https://datasheets.maximintegrated.com/en/ds/DS3231.pdf) | I2C |
| [ESP32 as WiFi Coprocessor with Arduino nina-fw](https://github.com/arduino/nina-fw) | SPI |
| [ESP8266/ESP32 AT Command set for WiFi/TCP/UDP](https://github.com/espressif/esp32-at) | UART |
| [GPS module](https://www.u-blox.com/en/product/neo-6-series) | I2C/UART |
| [HC-SR04 Ultrasonic distance sensor](https://cdn.sparkfun.com/datasheets/Sensors/Proximity/HCSR04.pdf) | GPIO |
| [HD44780 LCD controller](https://www.sparkfun.com/datasheets/LCD/HD44780.pdf) | GPIO |
| [HUB75 RGB led matrix](https://cdn-learn.adafruit.com/downloads/pdf/32x16-32x32-rgb-led-matrix.pdf) | SPI |
| [ILI9341 TFT color display](https://cdn-shop.adafruit.com/datasheets/ILI9341.pdf) | SPI |
| [L293x motor driver](https://www.ti.com/lit/ds/symlink/l293d.pdf) | GPIO/PWM |
| [L9110x motor driver](https://www.elecrow.com/download/datasheet-l9110.pdf) | GPIO/PWM |
| [LIS3DH accelerometer](https://www.st.com/resource/en/datasheet/lis3dh.pdf) | I2C |
| [LSM6DS3 accelerometer](https://www.st.com/resource/en/datasheet/lsm6ds3.pdf) | I2C |
| [MAG3110 magnetometer](https://www.nxp.com/docs/en/data-sheet/MAG3110.pdf) | I2C |
| [MCP3008 analog to digital converter (ADC)](http://ww1.microchip.com/downloads/en/DeviceDoc/21295d.pdf) | SPI |
| [Microphone - PDM](https://cdn-learn.adafruit.com/assets/assets/000/049/977/original/MP34DT01-M.pdf) | I2S/PDM |
| [MMA8653 accelerometer](https://www.nxp.com/docs/en/data-sheet/MMA8653FC.pdf) | I2C |
| [MPU6050 accelerometer/gyroscope](https://store.invensense.com/datasheets/invensense/MPU-6050_DataSheet_V3%204.pdf) | I2C |
| [PCD8544 display](http://eia.udg.edu/~forest/PCD8544_1.pdf) | SPI |
| [Resistive Touchscreen (4-wire)](http://ww1.microchip.com/downloads/en/Appnotes/doc8091.pdf) | GPIO |
| [Semihosting](https://wiki.segger.com/Semihosting) | Debug |
| [Shift register (PISO)](https://en.wikipedia.org/wiki/Shift_register#Parallel-in_serial-out_\(PISO\)) | GPIO |
| [Shift registers (SIPO)](https://en.wikipedia.org/wiki/Shift_register#Serial-in_parallel-out_(SIPO)) | GPIO |
| [SHT3x Digital Humidity Sensor](https://www.sensirion.com/fileadmin/user_upload/customers/sensirion/Dokumente/0_Datasheets/Humidity/Sensirion_Humidity_Sensors_SHT3x_Datasheet_digital.pdf) | I2C |
| [SPI NOR Flash Memory](https://en.wikipedia.org/wiki/Flash_memory#NOR_flash) | SPI/QSPI |
| [SSD1306 OLED display](https://cdn-shop.adafruit.com/datasheets/SSD1306.pdf) | I2C / SPI |
| [SSD1331 TFT color display](https://www.crystalfontz.com/controllers/SolomonSystech/SSD1331/381/) | SPI |
| [ST7735 TFT color display](https://www.crystalfontz.com/controllers/Sitronix/ST7735R/319/) | SPI |
| [ST7789 TFT color display](https://cdn-shop.adafruit.com/product-files/3787/3787_tft_QT154H2201__________20190228182902.pdf) | SPI |
| [Stepper motor "Easystepper" controller](https://en.wikipedia.org/wiki/Stepper_motor) | GPIO |
| [Thermistor](https://www.farnell.com/datasheets/33552.pdf) | ADC |
| [TMP102 I2C Temperature Sensor](https://download.mikroe.com/documents/datasheets/tmp102-data-sheet.pdf) | I2C |
| [VEML6070 UV light sensor](https://www.vishay.com/docs/84277/veml6070.pdf) | I2C |
| [VL53L1X time-of-flight distance sensor](https://www.st.com/resource/en/datasheet/vl53l1x.pdf) | I2C |
| [Waveshare 2.13" (B & C) e-paper display](https://www.waveshare.com/w/upload/d/d3/2.13inch-e-paper-b-Specification.pdf) | SPI |
| [Waveshare 2.13" e-paper display](https://www.waveshare.com/w/upload/e/e6/2.13inch_e-Paper_Datasheet.pdf) | SPI |
| [WS2812 RGB LED](https://cdn-shop.adafruit.com/datasheets/WS2812.pdf) | GPIO |

We also give you the ability to add new drivers. If your device isn't listed here, please raise an issue in the [issue tracker](https://github.com/tinygo-org/drivers/issues).
