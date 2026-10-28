---
title: "M5Stack StampFly"
weight: 3
---

The [M5Stack StampFly](https://docs.m5stack.com/en/app/StampFly_v1.1) is a small quadcopter built on the [M5Stamp S3A](https://docs.m5stack.com/en/core/StampS3A) module, which uses the Espressif [ESP32-S3](https://www.espressif.com/en/products/socs/esp32-s3) 32-bit Xtensa LX7 dual-core microcontroller. It has a BMI270 IMU, BMM150 magnetometer, BMP280 barometer, two VL53L3CX distance sensors, a PMW3901 optical flow sensor, and an INA3221 power monitor.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| WiFi      | YES | YES |
| Bluetooth | YES | YES |
| USBDevice | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `G0`              | `GPIO0`      | `BUTTON`          |
| `G1`              | `GPIO1`      | `GROVE_G1`, `ADC0` |
| `G2`              | `GPIO2`      | `GROVE_G2`, `ADC2` |
| `G3`              | `GPIO3`      | `SDA_PIN`, `ADC3` |
| `G4`              | `GPIO4`      | `SCL_PIN`, `ADC4` |
| `G5`              | `GPIO5`      | `MOTOR_FRONT_LEFT`, `ADC5` |
| `G6`              | `GPIO6`      | `TOF_BOTTOM_INT`, `ADC6` |
| `G7`              | `GPIO7`      | `TOF_BOTTOM_XSHUT`, `ADC7` |
| `G8`              | `GPIO8`      | `TOF_FRONT_INT`, `ADC8` |
| `G9`              | `GPIO9`      | `TOF_FRONT_XSHUT`, `ADC9` |
| `G10`             | `GPIO10`     | `MOTOR_REAR_LEFT`, `ADC10` |
| `G12`             | `GPIO12`     | `PMW3901_CS`, `ADC12` |
| `G13`             | `GPIO13`     | `GROVE_SDA`, `ADC13` |
| `G14`             | `GPIO14`     | `SPI1_MOSI_PIN`, `ADC14` |
| `G15`             | `GPIO15`     | `GROVE_SCL`, `ADC15` |
| `G21`             | `GPIO21`     | `WS2812_STAMP`    |
| `G39`             | `GPIO39`     | `WS2812`          |
| `G40`             | `GPIO40`     | `BUZZER`          |
| `G41`             | `GPIO41`     | `MOTOR_REAR_RIGHT` |
| `G42`             | `GPIO42`     | `MOTOR_FRONT_RIGHT` |
| `G43`             | `GPIO43`     | `UART_TX_PIN`, `SPI1_MISO_PIN` |
| `G44`             | `GPIO44`     | `UART_RX_PIN`, `SPI1_SCK_PIN` |
| `G46`             | `GPIO46`     | `BMI270_CS`       |
| `ADC11`           | `GPIO11`     |                   |
| `ADC16`           | `GPIO16`     |                   |
| `ADC17`           | `GPIO17`     |                   |
| `ADC18`           | `GPIO18`     |                   |
| `ADC19`           | `GPIO19`     |                   |
| `ADC20`           | `GPIO20`     |                   |

## Machine Package Docs

[Documentation for the machine package for the M5Stack StampFly](../../machine/m5stamp-fly)

## Flashing

### CLI Flashing

- Plug your StampFly into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command:

    ```shell
    tinygo flash -target=m5stamp-fly [PATH TO YOUR PROGRAM]
    ```

- The StampFly should restart and then begin running your program.

## Notes

The UART pins are shared with the SPI bus, so the serial console uses USB.

The board has an addressable RGB LED on `GPIO39` named `WS2812`. The Stamp-S3A module has another one on `GPIO21` named `WS2812_STAMP`.

The motors are driven by PWM on `MOTOR_FRONT_LEFT`, `MOTOR_FRONT_RIGHT`, `MOTOR_REAR_LEFT`, and `MOTOR_REAR_RIGHT`. Remove the propellers before testing motor code.

### WiFi

You can use the onboard wireless chip for WiFi using the `espradio` package.

For more information, see [https://github.com/tinygo-org/espradio](https://github.com/tinygo-org/espradio)

### Bluetooth

You can use the onboard wireless chip for Bluetooth Low Energy using the `bluetooth` package.

For more information, see [https://github.com/tinygo-org/bluetooth](https://github.com/tinygo-org/bluetooth)
