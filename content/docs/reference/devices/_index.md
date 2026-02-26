---
title: "Devices"
chapter: true
weight: 4
description: |
  Sensors and displays that are supported by TinyGo.
---

TinyGo has support for many different devices and sensors such as digital accelerometers, OLED displays, WiFi adaptors, and more.

Drivers are packages designed to make it easier to use these devices from your own TinyGo programs.

All of these drivers can be found in the TinyGo Drivers repository located at [https://github.com/tinygo-org/drivers/](https://github.com/tinygo-org/drivers/)

The following devices are supported.

| Driver                        | Device Name                                                             | Datasheet   | Interface Type |
|-------------------------------|-------------------------------------------------------------------------|-------------|----------------|
| adafruit4650                  | [Adafruit FeatherWing OLED - 128x64 OLED display](https://pkg.go.dev/tinygo.org/x/drivers/adafruit4650)  | [docs](https://learn.adafruit.com/adafruit-128x64-oled-featherwing)            | I2C               |
| adt7410                       | [ADT7410 I2C Temperature Sensor](https://pkg.go.dev/tinygo.org/x/drivers/adt7410) | [datasheet](https://www.analog.com/media/en/technical-documentation/data-sheets/ADT7410.pdf) | I2C |
| adxl345                       | [ADXL345 accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/adxl345) | [datasheet](http://www.analog.com/media/en/technical-documentation/data-sheets/ADXL345.pdf) | I2C |
| aht20                         | [AHT20 I2C Temperature and Humidity Sensor](https://pkg.go.dev/tinygo.org/x/drivers/aht20) | [datasheet](http://www.aosong.com/userfiles/files/media/AHT20%20%E8%8B%B1%E6%96%87%E7%89%88%E8%AF%B4%E6%98%8E%E4%B9%A6%20A0%2020201222.pdf) | I2C |
| amg88xx                       | [AMG88xx 8x8 Thermal camera sensor](https://pkg.go.dev/tinygo.org/x/drivers/amg88xx) | [datasheet](https://cdn-learn.adafruit.com/assets/assets/000/043/261/original/Grid-EYE_SPECIFICATIONS%28Reference%29.pdf) | I2C |
| apa102                        | [APA102 RGB LED](https://pkg.go.dev/tinygo.org/x/drivers/apa102) | [datasheet](https://cdn-shop.adafruit.com/product-files/2343/APA102C.pdf) | SPI |
| apds9960                      | [APDS9960 Digital proximity, ambient light, RGB and gesture sensor](https://pkg.go.dev/tinygo.org/x/drivers/apds9960) | [datasheet](https://cdn.sparkfun.com/assets/learn_tutorials/3/2/1/Avago-APDS-9960-datasheet.pdf) | I2C |
| as560x                        | [AS5600 / AS5601](https://pkg.go.dev/tinygo.org/x/drivers/as560x) ([on-axis magnetic rotary position sensors](https://ams.com/angle-position-on-axis)) | [datasheet](https://ams.com/documents/20143/36005/AS5600_DS000365_5-00.pdf) | I2C |
| at24cx                        | [AT24CX 2-wire serial EEPROM](https://pkg.go.dev/tinygo.org/x/drivers/at24cx) | [datasheet](https://www.openimpulse.com/blog/wp-content/uploads/wpsc/downloadables/24C32-Datasheet.pdf) | I2C |
| axp192                        | [AXP192 single Cell Li-Battery and Power System Management](https://pkg.go.dev/tinygo.org/x/drivers/axp192) | [datasheet](https://github.com/m5stack/M5-Schematic/blob/master/Core/AXP192%20Datasheet_v1.1_en_draft_2211.pdf) | I2C |
| bh1750                        | [BH1750 ambient light sensor](https://pkg.go.dev/tinygo.org/x/drivers/bh1750) | [datasheet](https://www.mouser.com/ds/2/348/bh1750fvi-e-186247.pdf) | I2C |
| blinkm                        | [BlinkM RGB LED](https://pkg.go.dev/tinygo.org/x/drivers/blinkm) | [datasheet](http://thingm.com/fileadmin/thingm/downloads/BlinkM_datasheet.pdf) | I2C |
| bma42x                        | [BMA42X triaxial acceleration sensor](https://pkg.go.dev/tinygo.org/x/drivers/bma427) | [datasheet](https://files.pine64.org/doc/datasheet/pinetime/BST-BMA421-FL000.pdf) | I2C |
| bme280                        | [BME280 humidity/pressure sensor](https://pkg.go.dev/tinygo.org/x/drivers/bme280) | [datasheet](https://cdn-shop.adafruit.com/datasheets/BST-BME280_DS001-10.pdf) | I2C |
| bmi160                        | [BMI160 accelerometer/gyroscope](https://pkg.go.dev/tinygo.org/x/drivers/bmi160) | [datasheet](https://www.bosch-sensortec.com/media/boschsensortec/downloads/datasheets/bst-bmi160-ds000.pdf) | SPI |
| bmp180                        | [BMP180 barometer](https://pkg.go.dev/tinygo.org/x/drivers/bmp180) | [datasheet](https://cdn-shop.adafruit.com/datasheets/BST-BMP180-DS000-09.pdf) | I2C |
| bmp280                        | [BMP280 temperature/barometer](https://pkg.go.dev/tinygo.org/x/drivers/bmp280) | [datasheet](https://www.bosch-sensortec.com/media/boschsensortec/downloads/datasheets/bst-bmp280-ds001.pdf) | I2C |
| bmp388                        | [BMP388 pressure sensor](https://pkg.go.dev/tinygo.org/x/drivers/bmp388) | [datasheet](https://www.bosch-sensortec.com/media/boschsensortec/downloads/datasheets/bst-bmp388-ds001.pdf) | I2C |
| bno08x                        | [BNO08x 9-DOF IMUsensor](https://pkg.go.dev/tinygo.org/x/drivers/bno08x) | [datasheet](https://www.ceva-ip.com/wp-content/uploads/BNO080_085-Datasheet.pdf) | I2C |
| buzzer                        | [Buzzer](https://pkg.go.dev/tinygo.org/x/drivers/buzzer) | [datasheet](https://en.wikipedia.org/wiki/Buzzer#Piezoelectric) | GPIO |
| comboat                       | [Aithinker-Combo-AT WiFi](https://pkg.go.dev/tinygo.org/x/drivers/comboat) | [docs](https://aithinker-combo-guide.readthedocs.io/en/latest/docs/instruction/index.html) | UART |
| dht                           | [DHTXX thermometer and humidity sensor](https://pkg.go.dev/tinygo.org/x/drivers/dht) | [datasheet](https://cdn-shop.adafruit.com/datasheets/Digital+humidity+and+temperature+sensor+AM2302.pdf) | GPIO |
| ds1307                        | [DS1307 real time clock](https://pkg.go.dev/tinygo.org/x/drivers/ds1307) | [datasheet](https://datasheets.maximintegrated.com/en/ds/DS1307.pdf) | I2C |
| ds18b20                       | [DS18B20 digital thermometer](https://pkg.go.dev/tinygo.org/x/drivers/ds18b20) | [datasheet](https://datasheets.maximintegrated.com/en/ds/DS1307.pdf) | I2C |
| ds3231                        | [DS3231 real time clock](https://pkg.go.dev/tinygo.org/x/drivers/ds3231) | [datasheet](https://datasheets.maximintegrated.com/en/ds/DS3231.pdf) | I2C |
| easystepper                   | [Stepper motor "Easystepper" controller](https://pkg.go.dev/tinygo.org/x/drivers/easystepper) | [datasheet](https://en.wikipedia.org/wiki/Stepper_motor) | GPIO |
| encoders                      | [Rotary Encoder](https://pkg.go.dev/tinygo.org/x/drivers/encoders) | [datasheet](https://www.mouser.com/datasheet/2/414/TTRB_S_A0002793947_1-2565369.pdf) | GPIO |
| ens160                        | [ScioSense ENS160 digital gas sensor](https://pkg.go.dev/tinygo.org/x/drivers/ens160) | [datasheet](https://www.sciosense.com/wp-content/uploads/2023/12/ENS160-Datasheet.pdf) | I2C |
| espat                         | [ESP8266/ESP32 AT Command set for WiFi/TCP/UDP](https://pkg.go.dev/tinygo.org/x/drivers/espat) | [datasheet](https://github.com/espressif/esp32-at) | UART |
| flash                         | [SPI NOR Flash Memory](https://pkg.go.dev/tinygo.org/x/drivers/flash) | [datasheet](https://en.wikipedia.org/wiki/Flash_memory#NOR_flash) | SPI/QSPI |
| ft6336                        | [FT6336 touch controller](https://pkg.go.dev/tinygo.org/x/drivers/ft6336) | [datasheet](https://focuslcds.com/content/FT6236.pdf) | I2C |
| gc9a01                        | [Waveshare GC9A01 TFT round display](https://pkg.go.dev/tinygo.org/x/drivers/gc9a01) | [datasheet](https://www.waveshare.com/w/upload/5/5e/GC9A01A.pdf) | SPI |
| gps                           | [GPS module](https://pkg.go.dev/tinygo.org/x/drivers/gps) | [datasheet](https://www.u-blox.com/en/product/neo-6-series) | I2C/UART |
| hcsr04                        | [HC-SR04 Ultrasonic distance sensor](https://pkg.go.dev/tinygo.org/x/drivers/hcsr04) | [datasheet](https://cdn.sparkfun.com/datasheets/Sensors/Proximity/HCSR04.pdf) | GPIO |
| hd44780                       | [HD44780 LCD controller](https://pkg.go.dev/tinygo.org/x/drivers/hd44780) | [datasheet](https://www.sparkfun.com/datasheets/LCD/HD44780.pdf) | GPIO |
| hd44780i2c                    | [HD44780 LCD controller with I2C interface](https://pkg.go.dev/tinygo.org/x/drivers/hd44780i2c) |[datasheet](https://www.sparkfun.com/datasheets/LCD/HD44780.pdf) | I2C |
| honeyhsc                      | [TruStability® pressure sensor](https://pkg.go.dev/tinygo.org/x/drivers/honeyhsc) | [datahseet](https://prod-edam.honeywell.com/content/dam/honeywell-edam/sps/siot/en-gb/products/sensors/pressure-sensors/board-mount-pressure-sensors/common/documents/sps-siot-hsc-series-ssc-series-installation-instructions-50044171-f-en-ciid-151614.pdf)| I2C/SPI |
| hts221                        | [HTS221 digital humidity and temperature sensor](https://pkg.go.dev/tinygo.org/x/drivers/hts221) | [datasheet](https://www.st.com/resource/en/datasheet/hts221.pdf) | I2C |
| hub75                         | [HUB75 RGB led matrix](https://pkg.go.dev/tinygo.org/x/drivers/hub75) | [datasheet](https://cdn-learn.adafruit.com/downloads/pdf/32x16-32x32-rgb-led-matrix.pdf) | SPI |
| i2csoft                       | [software I2C driver](https://pkg.go.dev/tinygo.org/x/drivers/i2csoft) | [datasheet](https://www.ti.com/lit/an/slva704/slva704.pdf) | GPIO |
| ili9341                       | [ILI9341 TFT color display](https://pkg.go.dev/tinygo.org/x/drivers/ili9341) | [datasheet](https://cdn-shop.adafruit.com/datasheets/ILI9341.pdf) | SPI |
| ina219                        | [INA219 current and power monitor](https://pkg.go.dev/tinygo.org/x/drivers/ina219) | [datasheet](https://www.ti.com/lit/ds/symlink/ina219.pdf) | I2C |
| ina260                        | [INA260 Volt/Amp/Power meter](https://pkg.go.dev/tinygo.org/x/drivers/ina260) | [datasheet](https://www.ti.com/lit/ds/symlink/ina260.pdf) | I2C |
| irremote                      | [Infrared remote control](https://pkg.go.dev/tinygo.org/x/drivers/irremote) | [datasheet](https://en.wikipedia.org/wiki/Consumer_IR) | GPIO |
| is31fl3731                    | [IS31FL3731 matrix LED driver](https://pkg.go.dev/tinygo.org/x/drivers/is31fl3731) | [datasheet](https://www.lumissil.com/assets/pdf/core/IS31FL3731_DS.pdf) | I2C |
| keypad4x4                     | [4x4 Membrane Keypad](https://pkg.go.dev/tinygo.org/x/drivers/keypad4x4) | [datasheet](https://cdn.sparkfun.com/assets/f/f/a/5/0/DS-16038.pdf) | GPIO |
| l293x                         | [L293x motor driver](https://pkg.go.dev/tinygo.org/x/drivers/l293x) | [datasheet](https://www.ti.com/lit/ds/symlink/l293d.pdf) | GPIO/PWM |
| l3gd20                        | [L3GD20 3-axis angular rate sensor](https://pkg.go.dev/tinygo.org/x/drivers/l3gd20) | [datasheet](https://www.st.com/resource/en/datasheet/l3gd20.pdf) | I2C |
| l9110x                        | [L9110x motor driver](https://pkg.go.dev/tinygo.org/x/drivers/l9110x) | [datasheet](https://www.elecrow.com/download/datasheet-l9110.pdf) | GPIO/PWM |
| lis2mdl                       | [LIS2MDL magnetometer](https://pkg.go.dev/tinygo.org/x/drivers/lis2mdl) | [datasheet](https://www.st.com/resource/en/datasheet/lis2mdl.pdf) | I2C |
| lis3dh                        | [LIS3DH accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lis3dh) | [datasheet](https://www.st.com/resource/en/datasheet/lis3dh.pdf) | I2C |
| lps22hb                       | [LPS22HB MEMS nano pressure sensor](https://pkg.go.dev/tinygo.org/x/drivers/lps22hb) | [datasheet](https://www.st.com/resource/en/datasheet/dm00140895.pdf) | I2C |
| lsm303agr                     | [LSM303AGR accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm303agr) | [datasheet](https://www.st.com/resource/en/datasheet/lsm303agr.pdf) | I2C |
| lsm303dlhc                    | [LSM303DLHC accelerometer/magnetometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm303dlhc) | [datasheet](https://www.st.com/resource/en/datasheet/lsm303dlhc.pdf) | I2C |
| lsm6ds3                       | [LSM6DS3 accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm6ds3) | [datasheet](https://www.st.com/resource/en/datasheet/lsm6ds3.pdf) | I2C |
| lsm6ds3tr                     | [LSM6DS3TR accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm6ds3tr) | [datasheet](https://www.st.com/resource/en/datasheet/lsm6ds3tr.pdf)| I2C |
| lsm6dsox                      | [LSM6DSOX accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm6ds0x) | [datasheet](https://www.st.com/resource/en/datasheet/lsm6dsox.pdf) | I2C |
| lsm9ds1                       | [LSM9DS1 accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/lsm9ds1) | [datasheet](https://www.st.com/resource/en/datasheet/lsm9ds1.pdf)| I2C |
| mag3110                       | [MAG3110 magnetometer](https://pkg.go.dev/tinygo.org/x/drivers/mag3110) | [datasheet](https://www.nxp.com/docs/en/data-sheet/MAG3110.pdf) | I2C |
| makeybutton                   | [Makey Button](https://pkg.go.dev/tinygo.org/x/drivers/makeybutton) | [datasheet](https://makeymakey.com/) | GPIO |
| max6675                       | [MAX6675 thermocouple-to-digital converter](https://pkg.go.dev/tinygo.org/x/drivers/max6675) | [datasheet](https://www.analog.com/media/en/technical-documentation/data-sheets/max6675.pdf) | SPI |
| max72xx                       | [MAX7219 & MAX7221 display driver](https://pkg.go.dev/tinygo.org/x/drivers/max72xx) | [datasheet](https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf) | SPI |
| mcp23017                      | [MCP23017 port expander](https://pkg.go.dev/tinygo.org/x/drivers/mcp23017) | [datasheet](https://ww1.microchip.com/downloads/en/DeviceDoc/20001952C.pdf) | I2C |
| mcp2515                       | [MCP2515 Stand-Alone CAN Controller with SPI Interface](https://pkg.go.dev/tinygo.org/x/drivers/mcp2515) | [datasheet](https://ww1.microchip.com/downloads/en/DeviceDoc/MCP2515-Family-Data-Sheet-DS20001801K.pdf) | SPI |
| mcp3008                       | [MCP3008 analog to digital converter (ADC)](https://pkg.go.dev/tinygo.org/x/drivers/mcp3008) | [datasheet](http://ww1.microchip.com/downloads/en/DeviceDoc/21295d.pdf) | SPI |
| mcp9808                       | [MCP9808 High Accuracy I2C Temperature Sensor](https://pkg.go.dev/tinygo.org/x/drivers/mcp9808) | [datasheet](https://cdn-shop.adafruit.com/datasheets/MCP9808.pdf) | I2C |
| microbitmatrix                 | [BBC micro:bit LED matrix](https://pkg.go.dev/tinygo.org/x/drivers/microbitmatrix) | [datasheet](https://github.com/bbcmicrobit/hardware/blob/master/SCH_BBC-Microbit_V1.3B.pdf) | GPIO |
| microphone                    | [Microphone - PDM](https://pkg.go.dev/tinygo.org/x/drivers/microphone) | [datasheet](https://cdn-learn.adafruit.com/assets/assets/000/049/977/original/MP34DT01-M.pdf) | I2S/PDM |
| mma8653                       | [MMA8653 accelerometer](https://pkg.go.dev/tinygo.org/x/drivers/mma8653) | [datasheet](https://www.nxp.com/docs/en/data-sheet/MMA8653FC.pdf) | I2C |
| mpu6050                       | [MPU6050 accelerometer/gyroscope](https://pkg.go.dev/tinygo.org/x/drivers/mpu6050) | [datasheet](https://store.invensense.com/datasheets/invensense/MPU-6050_DataSheet_V3%204.pdf) | I2C |
| mpu6886                       | [MPU6886 accelerometer/gyroscope](https://pkg.go.dev/tinygo.org/x/drivers/mpu6886) | [datasheet](https://github.com/m5stack/M5-Schematic/blob/master/datasheet/MPU-6886-000193%2Bv1.1_GHIC.PDF.pdf) | I2C |
| mpu9150                       | [MPU9150 accelerometer/gyroscope]() |[datahseet](https://invensense.tdk.com/wp-content/uploads/2015/02/MPU-9150-Datasheet.pdf) | I2C |
| ndir                          | [NDIR CO2 Sensor](https://pkg.go.dev/tinygo.org/x/drivers/ndir) | [datasheet](http://sandboxelectronics.com/?p=1126) | I2C |
| onewire                       | [One Wire bus system](https://pkg.go.dev/tinygo.org/x/drivers/onewire) | [datasheet](https://en.wikipedia.org/wiki/1-Wire) | 1-wire |
| p1am                          | [P1AM-100 base controller](https://pkg.go.dev/tinygo.org/x/drivers/p1am) | [docs](https://github.com/facts-engineering/P1AM/tree/1.0.1) | SPI |
| pca9685                       | [PCA9685 16-channel, 12-bit PWM controller](https://pkg.go.dev/tinygo.org/x/drivers/pca9685) | [datasheet](https://www.nxp.com/docs/en/data-sheet/PCA9685.pdf) | I2C |
| pcd8544                       | [PCD8544 display](https://pkg.go.dev/tinygo.org/x/drivers/pcd8544) | [datasheet](http://eia.udg.edu/~forest/PCD8544_1.pdf) | SPI |
| pcf8523                       | [PCF8523 real time clock](https://pkg.go.dev/tinygo.org/x/drivers/pcf8523) | [datasheet](https://www.nxp.com/docs/en/data-sheet/PCF8523.pdf) | I2C |
| pcf8563                       | [PCF8563 real time clock](https://pkg.go.dev/tinygo.org/x/drivers/pcf8563) | [datasheet](https://www.nxp.com/docs/en/data-sheet/PCF8563.pdf) | I2C |
| pcf8591                       | [PCF8591 Analog to Digital/Digital to Analog Converter](https://pkg.go.dev/tinygo.org/x/drivers/pcf8591) | [datasheet](https://www.nxp.com/docs/en/data-sheet/PCF8591.pdf) | I2C |
| qmi8658c                      | [QMI8658C accelerometer/gyroscope](https://pkg.go.dev/tinygo.org/x/drivers/qmi8658c) | [datasheet](https://www.qstcorp.com/upload/pdf/202202/%EF%BC%88%E5%B7%B2%E4%BC%A0%EF%BC%89QMI8658C%20datasheet%20rev%200.9.pdf) | I2C |
| rtl8720dn                     | [RTL8720DN 2.4G/5G Dual Bands Wireless and BLE5.0](https://pkg.go.dev/tinygo.org/x/drivers/lsm6ds0xhttps://pkg.go.dev/tinygo.org/x/drivers/rtl8720dn) | [datasheet](https://www.seeedstudio.com/Realtek8720DN-2-4G-5G-Dual-Bands-Wireless-and-BLE5-0-Combo-Module-p-4442.html) | UART |
| scd4x                         | [SCD4x CO2 Sensor](https://pkg.go.dev/tinygo.org/x/drivers/scd4x) | [datasheet](https://sensirion.com/media/documents/C4B87CE6/627C2DCD/CD_DS_SCD40_SCD41_Datasheet_D1.pdf) | I2C |
| sdcard                        | [SPI SDCARD/MMC](https://pkg.go.dev/tinygo.org/x/drivers/sdcard) | [datasheet](https://en.wikipedia.org/wiki/SD_card) | SPI |
| seesaw                        | [Adafruit Seesaw](https://pkg.go.dev/tinygo.org/x/drivers/seesaw) | [docs](https://learn.adafruit.com/adafruit-seesaw-atsamd09-breakout/overview) | I2C |
| semihosting                   | [Semihosting](https://pkg.go.dev/tinygo.org/x/drivers/semihosting) | [datasheet](https://wiki.segger.com/Semihosting) | Debug |
| servo                         | [Servo](https://pkg.go.dev/tinygo.org/x/drivers/servo) | [datasheet](https://learn.sparkfun.com/tutorials/hobby-servo-tutorial/all) | PWM |
| sgp30                         | [SGP30 VOC sensor](https://pkg.go.dev/tinygo.org/x/drivers/sgp30) | [datasheet](https://sensirion.com/media/documents/984E0DD5/61644B8B/Sensirion_Gas_Sensors_Datasheet_SGP30.pdf) | I2C |
| sh1106                        | [SH1106 OLED display](https://pkg.go.dev/tinygo.org/x/drivers/sh1106) | [datasheet](https://www.velleman.eu/downloads/29/infosheets/sh1106_datasheet.pdf) | I2C / SPI |
| sharpmem                      | [Sharp Memory Displays](https://pkg.go.dev/tinygo.org/x/drivers/sharpmem) | [docs](https://sharpdevices.com/memory-lcd/) | SPI |
| shifter                       | [8bit shift registers (74HC165, 74165, etc)]() |             | GPIO |
| shiftregister                 | [Shift register (PISO)](https://pkg.go.dev/tinygo.org/x/drivers/shiftregister) | [datasheet](https://en.wikipedia.org/wiki/Shift_register#Parallel-in_serial-out_\(PISO\)) | GPIO |
| sht3x                         | [SHT3x Digital Humidity Sensor](https://pkg.go.dev/tinygo.org/x/drivers/sht3x) | [datasheet](https://sensirion.com/media/documents/213E6A3B/63A5A569/Datasheet_SHT3x_DIS.pdf) | I2C |
| sht4x                         | [SHT4x Digital Humidity Sensor](https://pkg.go.dev/tinygo.org/x/drivers/sht4x) | [datasheet](https://sensirion.com/media/documents/33FD6951/63E1087C/Datasheet_SHT4x_1.pdf) | I2C |
| shtc3                         | [SHTC3 Digital Humidity Sensor (RH/T)](https://pkg.go.dev/tinygo.org/x/drivers/shtc3) | [datasheet](https://www.sensirion.com/fileadmin/user_upload/customers/sensirion/Dokumente/2_Humidity_Sensors/Datasheets/Sensirion_Humidity_Sensors_SHTC3_Datasheet.pdf)  | I2C |
| si5351                        | [SI5351 clock generator](https://pkg.go.dev/tinygo.org/x/drivers/si5351) | [datasheet](https://cdn-shop.adafruit.com/datasheets/Si5351.pdf) | I2C |
| ssd1289                       | [SSD1289 TFT color display](https://pkg.go.dev/tinygo.org/x/drivers/ssd1289) | [datasheet](http://aitendo3.sakura.ne.jp/aitendo_data/product_img/lcd/tft2/M032C1289TP/3.2-SSD1289.pdf) | GPIO |
| ssd1306                       | [SSD1306 OLED display](https://pkg.go.dev/tinygo.org/x/drivers/ssd1306) | [datasheet](https://cdn-shop.adafruit.com/datasheets/SSD1306.pdf) | I2C / SPI |
| ssd1331                       | [SSD1331 TFT color display](https://pkg.go.dev/tinygo.org/x/drivers/ssd1331) | [datasheet](https://www.crystalfontz.com/controllers/SolomonSystech/SSD1331/381/) | SPI |
| ssd1351                       | [SSD1351 OLED display](https://pkg.go.dev/tinygo.org/x/drivers/ssd1351) | [datasheet](https://download.mikroe.com/documents/datasheets/ssd1351-revision-1.3.pdf) | SPI |
| st7735                        | [ST7735 TFT color display](https://pkg.go.dev/tinygo.org/x/drivers/st7735) | [datasheet](https://www.crystalfontz.com/controllers/Sitronix/ST7735R/319/) | SPI |
| st7789                        | [ST7789 TFT color display](https://pkg.go.dev/tinygo.org/x/drivers/st7789) | [datasheet](https://cdn-shop.adafruit.com/product-files/3787/3787_tft_QT154H2201__________20190228182902.pdf) | SPI |
| sx126x                        | [Semtech SX126x Lora](https://pkg.go.dev/tinygo.org/x/drivers/sx126x) | [datasheet](https://www.semtech.com/products/wireless-rf/lora-connect/sx1261) | SPI |
| sx127x                        | [Semtech SX127x Lora](https://pkg.go.dev/tinygo.org/x/drivers/sx127x) | [datasheet](https://www.semtech.com/products/wireless-rf/lora-connect/sx1276) | SPI |
| thermistor                    | [Thermistor](https://pkg.go.dev/tinygo.org/x/drivers/thermistor) | [datasheet](https://www.farnell.com/datasheets/33552.pdf) | ADC |
| tm1637                        | [TM1637 7-segment LED display](https://pkg.go.dev/tinygo.org/x/drivers/tm1637) | [datasheet](https://www.mcielectronics.cl/website_MCI/static/documents/Datasheet_TM1637.pdf) | I2C |
| tmc2209                       | [TMC2209 stepper motor](https://pkg.go.dev/tinygo.org/x/drivers/tmc2209) | [datasheet](https://www.analog.com/media/en/technical-documentation/data-sheets/tmc2209_datasheet_rev1.09.pdf) | UART |
| tmc5160                       | [TMC5160 stepper motor](https://pkg.go.dev/tinygo.org/x/drivers/tmc5160) | [datasheet](https://www.analog.com/media/en/technical-documentation/data-sheets/tmc5160a_datasheet_rev1.17.pdf) | UART |
| tmp102                        | [TMP102 I2C Temperature Sensor](https://pkg.go.dev/tinygo.org/x/drivers/tmp102) | [datasheet](https://download.mikroe.com/documents/datasheets/tmp102-data-sheet.pdf) | I2C |
| touch                         | [Resistive Touchscreen (4-wire)](https://pkg.go.dev/tinygo.org/x/drivers/touch/resistive) | [datasheet](http://ww1.microchip.com/downloads/en/Appnotes/doc8091.pdf) | GPIO |
| ttp229                        | [TTP229 (BSF version) 16 keys or 8 keys touch pad detector](https://pkg.go.dev/tinygo.org/x/drivers/ttp229) | [datasheet](https://www.sunrom.com/download/SUNROM-TTP229-BSF_V1.1_EN.pdf) | GPIO |
| uc8151                        | [UC8151 All-in-one driver IC for ESL](https://pkg.go.dev/tinygo.org/x/drivers/uc8151) | [datasheet](https://www.buydisplay.com/download/ic/UC8151C.pdf) | I2C |
| veml6070                      | [VEML6070 UV light sensor](https://pkg.go.dev/tinygo.org/x/drivers/veml6070) | [datasheet](https://www.vishay.com/docs/84277/veml6070.pdf) | I2C |
| vl53l1x                       | [VL53L1X time-of-flight distance sensor](https://pkg.go.dev/tinygo.org/x/drivers/vl53l1x) | [datasheet](https://www.st.com/resource/en/datasheet/vl53l1x.pdf) | I2C |
| vl6180x                       | [VL6180X time-of-flight distance sensor](https://pkg.go.dev/tinygo.org/x/drivers/vl6180x) | [datasheet](https://www.st.com/resource/en/datasheet/vl6180x.pdf) | I2C |
| w5500                         | [W5500 Ethernet controller](https://pkg.go.dev/tinygo.org/x/drivers/w5500) | [datasheet](https://docs.wiznet.io/img/products/w5500/W5500_ds_v110e.pdf) | SPI |
| waveshare-epd/epd1in54        | [Waveshare 1.54" black and white e-paper device](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd1in54) | [datasheet](https://www.waveshare.com/w/upload/e/e5/1.54inch_e-paper_V2_Datasheet.pdf) | SPI |
| waveshare-epd/epd2in13        | [Waveshare 2.13" e-paper display](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd2in13) | [datasheet](https://www.waveshare.com/w/upload/e/e6/2.13inch_e-Paper_Datasheet.pdf) | SPI |
| waveshare-epd/epd2in13x       | [Waveshare 2.13" (B & C) e-paper display](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd2in13x) | [datasheet](https://www.waveshare.com/w/upload/d/d3/2.13inch-e-paper-b-Specification.pdf) | SPI |
| waveshare-epd/epd2in66b       | [Waveshare 2.66" e-Paper display](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd2in66b) | [datasheet](https://files.waveshare.com/upload/e/ec/2.66inch-e-paper-b-specification.pdf) | SPI |
| waveshare-epd/epd2in9         | [Waveshare 2.9" e-paper display (V1)](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd2in9) | [datasheet](https://www.waveshare.com/w/upload/e/e6/2.9inch_e-Paper_Datasheet.pdf) | SPI |
| waveshare-epd/epd4in2         | [Waveshare 4.2" e-paper B/W display](https://pkg.go.dev/tinygo.org/x/drivers/waveshare-epd/epd4in2) | [datasheet](https://www.waveshare.com/w/upload/6/6a/4.2inch-e-paper-specification.pdf) | SPI |
| wifinina                      | [ESP32 as WiFi Coprocessor with Arduino nina-fw](https://pkg.go.dev/tinygo.org/x/drivers/wifinina) | [datasheet](https://github.com/arduino/nina-fw) | SPI |
| ws2812                        | [WS2812 RGB LED](https://pkg.go.dev/tinygo.org/x/drivers/ws2812) | [datasheet](https://cdn-shop.adafruit.com/datasheets/WS2812.pdf) | GPIO |
| xpt2046                       | [XPT2046 touch controller](https://pkg.go.dev/tinygo.org/x/drivers/xpt2046) | [datasheet](http://grobotronics.com/images/datasheets/xpt2046-datasheet.pdf) | GPIO |

We also give you the ability to add new drivers. If your device isn't listed here, please raise an issue in the [issue tracker](https://github.com/tinygo-org/drivers/issues).

If you want to know more about how drivers are implemented please see the [Drivers page](../../concepts/drivers) under "Concepts".
