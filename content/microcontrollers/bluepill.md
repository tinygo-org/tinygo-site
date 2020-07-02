---
title: "Bluepill"
weight: 3
---

The [Bluepill](http://wiki.stm32duino.com/index.php?title=Blue_Pill) is a popular, ultra-cheap and compact ARM development board based on the ST Micro [STM32F103C8](https://www.st.com/en/microcontrollers/stm32f103c8.html) SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the Bluepill](../machine/bluepill)

## Flashing

### OpenOCD

Programs are loaded onto the Bluepill with a separate hardware programmer such as the [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) board, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Bluepill board with your TinyGo code.

- Plug your STLink v2 programmer into your computer's USB port.
- Plug your Bluepill into the STLink v2 programmer using the Bluepill `SWIO`, `SWCLK`, `3V3` and `GND` pins.
- Build and flash your TinyGo program using `tinygo flash -target=bluepill`

## OpenOCD Possible Modification

If you run the flasher, and you are getting the error **Warn : UNEXPECTED idcode: 0x2ba01477** you can mitigate the problem by editing the stm32f1x.cfg file within openOCD
```sh
> tinygo flash -target=bluepill -ocd-output main.go
Open On-Chip Debugger 0.10.0
Licensed under GNU GPL v2
For bug reports, read
        http://openocd.org/doc/doxygen/bugs.html
Info : auto-selecting first available session transport "hla_swd". To override use 'transport select <transport>'.
Info : The selected transport took over low-level target control. The results might differ compared to plain JTAG/SWD
adapter speed: 1000 kHz
adapter_nsrst_delay: 100
none separate
Info : Unable to match requested speed 1000 kHz, using 950 kHz
Info : Unable to match requested speed 1000 kHz, using 950 kHz
Info : clock speed 950 kHz
Info : STLINK v2 JTAG v29 API v2 SWIM v7 VID 0x0483 PID 0x3748
Info : using stlink api v2
Info : Target voltage: 3.160026
Warn : UNEXPECTED idcode: 0x2ba01477
Error: expected 1 of 1: 0x1ba01477
in procedure 'program' 
in procedure 'init' called at file "embedded:startup.tcl", line 506
in procedure 'ocd_bouncer'
** OpenOCD init failed **
shutdown command invoked

```


```sh
> sudo vim /usr/share/openocd/scripts/target/stm32f1x.cfg
```

Edit the line **set _CPUTAPID 0x1ba01477** and change the value to **set _CPUTAPID 0x2ba01477**.  This is usually a known problem with the chip on blue pills, since there are a lot of different chips out there.  This way, you can still write to those microprocessors.  

## Other possible problems

If you get an error on the line  **Info : STLINK v2 JTAG v29 API v2 SWIM v7 VID 0x0483 PID 0x3748**, you likely have the SWDIO/SWCLK lines in reverse.  This is also a known problem on some of the cheaper blue pills.  In some cases, the outer case of the st-link v2 is not labeled properly (and the pcb inside the case is labeled in reverse).  You wont hurt anything by switching the wires, and see if that fixes the problem.  It is better to try that than pry apart the st-link v2 casing.