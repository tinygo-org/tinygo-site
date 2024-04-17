
---
title: teensy40
---


## Constants

```go
const (
	//  = Pin  // [Pad]:       Alt Func 0       Alt Func 1       Alt Func 2       Alt Func 3      Alt Func 4            Alt Func 5  Alt Func 6            Alt Func 7             Alt Func 8             Alt Func 9
	//  = ----    -----------  ---------------  ---------------  ---------------  --------------  --------------------  ----------  --------------------  ---------------------  ---------------------  ----------------
	D0	= PA3	// [AD_B0_03]:  FLEXCAN2_RX      XBAR1_INOUT17    LPUART6_RX       USB_OTG1_OC     FLEXPWM1_PWMX01       GPIO1_IO03  REF_CLK_24M           LPSPI3_PCS0             ~                      ~
	D1	= PA2	// [AD_B0_02]:  FLEXCAN2_TX      XBAR1_INOUT16    LPUART6_TX       USB_OTG1_PWR    FLEXPWM1_PWMX00       GPIO1_IO02  LPI2C1_HREQ           LPSPI3_SDI              ~                      ~
	D2	= PD4	// [EMC_04]:    SEMC_DATA04      FLEXPWM4_PWMA02  SAI2_TX_DATA     XBAR1_INOUT06   FLEXIO1_FLEXIO04      GPIO4_IO04   ~                     ~                      ~                      ~
	D3	= PD5	// [EMC_05]:    SEMC_DATA05      FLEXPWM4_PWMB02  SAI2_TX_SYNC     XBAR1_INOUT07   FLEXIO1_FLEXIO05      GPIO4_IO05   ~                     ~                      ~                      ~
	D4	= PD6	// [EMC_06]:    SEMC_DATA06      FLEXPWM2_PWMA00  SAI2_TX_BCLK     XBAR1_INOUT08   FLEXIO1_FLEXIO06      GPIO4_IO06   ~                     ~                      ~                      ~
	D5	= PD8	// [EMC_08]:    SEMC_DM00        FLEXPWM2_PWMA01  SAI2_RX_DATA     XBAR1_INOUT17   FLEXIO1_FLEXIO08      GPIO4_IO08   ~                     ~                      ~                      ~
	D6	= PB10	// [B0_10]:     LCD_DATA06       QTIMER4_TIMER1   FLEXPWM2_PWMA02  SAI1_TX_DATA03  FLEXIO2_FLEXIO10      GPIO2_IO10  SRC_BOOT_CFG06        ENET2_CRS               ~                      ~
	D7	= PB17	// [B1_01]:     LCD_DATA13       XBAR1_INOUT15    LPUART4_RX       SAI1_TX_DATA00  FLEXIO2_FLEXIO17      GPIO2_IO17  FLEXPWM1_PWMB03       ENET2_RDATA00          FLEXIO3_FLEXIO17        ~
	D8	= PB16	// [B1_00]:     LCD_DATA12       XBAR1_INOUT14    LPUART4_TX       SAI1_RX_DATA00  FLEXIO2_FLEXIO16      GPIO2_IO16  FLEXPWM1_PWMA03       ENET2_RX_ER            FLEXIO3_FLEXIO16        ~
	D9	= PB11	// [B0_11]:     LCD_DATA07       QTIMER4_TIMER2   FLEXPWM2_PWMB02  SAI1_TX_DATA02  FLEXIO2_FLEXIO11      GPIO2_IO11  SRC_BOOT_CFG07        ENET2_COL               ~                      ~
	D10	= PB0	// [B0_00]:     LCD_CLK          QTIMER1_TIMER0   MQS_RIGHT        LPSPI4_PCS0     FLEXIO2_FLEXIO00      GPIO2_IO00  SEMC_CSX01            ENET2_MDC               ~                      ~
	D11	= PB2	// [B0_02]:     LCD_HSYNC        QTIMER1_TIMER2   FLEXCAN1_TX      LPSPI4_SDO      FLEXIO2_FLEXIO02      GPIO2_IO02  SEMC_CSX03            ENET2_1588_EVENT0_OUT   ~                      ~
	D12	= PB1	// [B0_01]:     LCD_ENABLE       QTIMER1_TIMER1   MQS_LEFT         LPSPI4_SDI      FLEXIO2_FLEXIO01      GPIO2_IO01  SEMC_CSX02            ENET2_MDIO              ~                      ~
	D13	= PB3	// [B0_03]:     LCD_VSYNC        QTIMER2_TIMER0   FLEXCAN1_RX      LPSPI4_SCK      FLEXIO2_FLEXIO03      GPIO2_IO03  WDOG2_RESET_B_DEB     ENET2_1588_EVENT0_IN    ~                      ~
	D14	= PA18	// [AD_B1_02]:  USB_OTG1_ID      QTIMER3_TIMER2   LPUART2_TX       SPDIF_OUT       ENET_1588_EVENT2_OUT  GPIO1_IO18  USDHC1_CD_B           KPP_ROW06              GPT2_CLK               FLEXIO3_FLEXIO02
	D15	= PA19	// [AD_B1_03]:  USB_OTG1_OC      QTIMER3_TIMER3   LPUART2_RX       SPDIF_IN        ENET_1588_EVENT2_IN   GPIO1_IO19  USDHC2_CD_B           KPP_COL06              GPT2_CAPTURE1          FLEXIO3_FLEXIO03
	D16	= PA23	// [AD_B1_07]:  FLEXSPIB_DATA00  LPI2C3_SCL       LPUART3_RX       SPDIF_EXT_CLK   CSI_HSYNC             GPIO1_IO23  USDHC2_DATA3          KPP_COL04              GPT2_COMPARE3          FLEXIO3_FLEXIO07
	D17	= PA22	// [AD_B1_06]:  FLEXSPIB_DATA01  LPI2C3_SDA       LPUART3_TX       SPDIF_LOCK      CSI_VSYNC             GPIO1_IO22  USDHC2_DATA2          KPP_ROW04              GPT2_COMPARE2          FLEXIO3_FLEXIO06
	D18	= PA17	// [AD_B1_01]:  USB_OTG1_PWR     QTIMER3_TIMER1   LPUART2_RTS_B    LPI2C1_SDA      CCM_PMIC_READY        GPIO1_IO17  USDHC1_VSELECT        KPP_COL07              ENET2_1588_EVENT0_IN   FLEXIO3_FLEXIO01
	D19	= PA16	// [AD_B1_00]:  USB_OTG2_ID      QTIMER3_TIMER0   LPUART2_CTS_B    LPI2C1_SCL      WDOG1_B               GPIO1_IO16  USDHC1_WP             KPP_ROW07              ENET2_1588_EVENT0_OUT  FLEXIO3_FLEXIO00
	D20	= PA26	// [AD_B1_10]:  FLEXSPIA_DATA03  WDOG1_B          LPUART8_TX       SAI1_RX_SYNC    CSI_DATA07            GPIO1_IO26  USDHC2_WP             KPP_ROW02              ENET2_1588_EVENT1_OUT  FLEXIO3_FLEXIO10
	D21	= PA27	// [AD_B1_11]:  FLEXSPIA_DATA02  EWM_OUT_B        LPUART8_RX       SAI1_RX_BCLK    CSI_DATA06            GPIO1_IO27  USDHC2_RESET_B        KPP_COL02              ENET2_1588_EVENT1_IN   FLEXIO3_FLEXIO11
	D22	= PA24	// [AD_B1_08]:  FLEXSPIA_SS1_B   FLEXPWM4_PWMA00  FLEXCAN1_TX      CCM_PMIC_READY  CSI_DATA09            GPIO1_IO24  USDHC2_CMD            KPP_ROW03              FLEXIO3_FLEXIO08        ~
	D23	= PA25	// [AD_B1_09]:  FLEXSPIA_DQS     FLEXPWM4_PWMA01  FLEXCAN1_RX      SAI1_MCLK       CSI_DATA08            GPIO1_IO25  USDHC2_CLK            KPP_COL03              FLEXIO3_FLEXIO09        ~
	D24	= PA12	// [AD_B0_12]:  LPI2C4_SCL       CCM_PMIC_READY   LPUART1_TX       WDOG2_WDOG_B    FLEXPWM1_PWMX02       GPIO1_IO12  ENET_1588_EVENT1_OUT  NMI_GLUE_NMI            ~                      ~
	D25	= PA13	// [AD_B0_13]:  LPI2C4_SDA       GPT1_CLK         LPUART1_RX       EWM_OUT_B       FLEXPWM1_PWMX03       GPIO1_IO13  ENET_1588_EVENT1_IN   REF_CLK_24M             ~                      ~
	D26	= PA30	// [AD_B1_14]:  FLEXSPIA_SCLK    ACMP_OUT02       LPSPI3_SDO       SAI1_TX_BCLK    CSI_DATA03            GPIO1_IO30  USDHC2_DATA6          KPP_ROW00              ENET2_1588_EVENT3_OUT  FLEXIO3_FLEXIO14
	D27	= PA31	// [AD_B1_15]:  FLEXSPIA_SS0_B   ACMP_OUT03       LPSPI3_SCK       SAI1_TX_SYNC    CSI_DATA02            GPIO1_IO31  USDHC2_DATA7          KPP_COL00              ENET2_1588_EVENT3_IN   FLEXIO3_FLEXIO15
	D28	= PC18	// [EMC_32]:    SEMC_DATA10      FLEXPWM3_PWMB01  LPUART7_RX       CCM_PMIC_RDY    CSI_DATA21            GPIO3_IO18  ENET2_TX_EN            ~                      ~                      ~
	D29	= PD31	// [EMC_31]:    SEMC_DATA09      FLEXPWM3_PWMA01  LPUART7_TX       LPSPI1_PCS1     CSI_DATA22            GPIO4_IO31  ENET2_TDATA01          ~                      ~                      ~
	D30	= PC23	// [EMC_37]:    SEMC_DATA15      XBAR1_IN23       GPT1_COMPARE3    SAI3_MCLK       CSI_DATA16            GPIO3_IO23  USDHC2_WP             ENET2_RX_EN            FLEXCAN3_RX             ~
	D31	= PC22	// [EMC_36]:    SEMC_DATA14      XBAR1_IN22       GPT1_COMPARE2    SAI3_TX_DATA    CSI_DATA17            GPIO3_IO22  USDHC1_WP             ENET2_RDATA01          FLEXCAN3_TX             ~
	D32	= PB12	// [B0_12]:     LCD_DATA08       XBAR1_INOUT10    ARM_TRACE_CLK    SAI1_TX_DATA01  FLEXIO2_FLEXIO12      GPIO2_IO12  SRC_BOOT_CFG08        ENET2_TDATA00           ~                      ~
	D33	= PD7	// [EMC_07]:    SEMC_DATA07      FLEXPWM2_PWMB00  SAI2_MCLK        XBAR1_INOUT09   FLEXIO1_FLEXIO07      GPIO4_IO07   ~                     ~                      ~                      ~
	D34	= PC15	// [SD_B0_03]:  USDHC1_DATA1     FLEXPWM1_PWMB01  LPUART8_RTS_B    XBAR1_INOUT07   LPSPI1_SDI            GPIO3_IO15  ENET2_RDATA00         SEMC_CLK6               ~                      ~
	D35	= PC14	// [SD_B0_02]:  USDHC1_DATA0     FLEXPWM1_PWMA01  LPUART8_CTS_B    XBAR1_INOUT06   LPSPI1_SDO            GPIO3_IO14  ENET2_RX_ER           SEMC_CLK5               ~                      ~
	D36	= PC13	// [SD_B0_01]:  USDHC1_CLK       FLEXPWM1_PWMB00  LPI2C3_SDA       XBAR1_INOUT05   LPSPI1_PCS0           GPIO3_IO13  FLEXSPIB_SS1_B        ENET2_TX_CLK           ENET2_REF_CLK2          ~
	D37	= PC12	// [SD_B0_00]:  USDHC1_CMD       FLEXPWM1_PWMA00  LPI2C3_SCL       XBAR1_INOUT04   LPSPI1_SCK            GPIO3_IO12  FLEXSPIA_SS1_B        ENET2_TX_EN            SEMC_DQS4               ~
	D38	= PC17	// [SD_B0_05]:  USDHC1_DATA3     FLEXPWM1_PWMB02  LPUART8_RX       XBAR1_INOUT09   FLEXSPIB_DQS          GPIO3_IO17  CCM_CLKO2             ENET2_RX_EN             ~                      ~
	D39	= PC16	// [SD_B0_04]:  USDHC1_DATA2     FLEXPWM1_PWMA02  LPUART8_TX       XBAR1_INOUT08   FLEXSPIB_SS0_B        GPIO3_IO16  CCM_CLKO1             ENET2_RDATA01           ~                      ~
)
```

Digital pins


```go
const (
	//  = Pin  // Dig | [Pad]      {ADC1/ADC2}
	A0	= PA18	// D14 | [AD_B1_02] {  7 / 7  }
	A1	= PA19	// D15 | [AD_B1_03] {  8 / 8  }
	A2	= PA23	// D16 | [AD_B1_07] { 12 / 12 }
	A3	= PA22	// D17 | [AD_B1_06] { 11 / 11 }
	A4	= PA17	// D18 | [AD_B1_01] {  6 / 6  }
	A5	= PA16	// D19 | [AD_B1_00] {  5 / 5  }
	A6	= PA26	// D20 | [AD_B1_10] { 15 / 15 }
	A7	= PA27	// D21 | [AD_B1_11] {  0 / 0  }
	A8	= PA24	// D22 | [AD_B1_08] { 13 / 13 }
	A9	= PA25	// D23 | [AD_B1_09] { 14 / 14 }
	A10	= PA12	// D24 | [AD_B0_12] {  1 / -  }
	A11	= PA13	// D25 | [AD_B0_13] {  2 / -  }
	A12	= PA30	// D26 | [AD_B1_14] {  - / 3  }
	A13	= PA31	// D27 | [AD_B1_15] {  - / 4  }
)
```

Analog pins


```go
const (
	LED	= D13

	UART_RX_PIN	= UART1_RX_PIN	// D0
	UART_TX_PIN	= UART1_TX_PIN	// D1

	SPI_SDI_PIN	= SPI1_SDI_PIN	// D12
	SPI_SDO_PIN	= SPI1_SDO_PIN	// D11
	SPI_SCK_PIN	= SPI1_SCK_PIN	// D13
	SPI_CS_PIN	= SPI1_CS_PIN	// D10

	I2C_SDA_PIN	= I2C1_SDA_PIN	// D18/A4
	I2C_SCL_PIN	= I2C1_SCL_PIN	// D19/A5
)
```

Default peripheral pins


```go
const (
	UART1_RX_PIN	= D0
	UART1_TX_PIN	= D1

	UART2_RX_PIN	= D7
	UART2_TX_PIN	= D8

	UART3_RX_PIN	= D15
	UART3_TX_PIN	= D14

	UART4_RX_PIN	= D16
	UART4_TX_PIN	= D17

	UART5_RX_PIN	= D21
	UART5_TX_PIN	= D20

	UART6_RX_PIN	= D25
	UART6_TX_PIN	= D24

	UART7_RX_PIN	= D28
	UART7_TX_PIN	= D29
)
```

#=====================================================#
|                        UART                         |
#===========#===========#=============#===============#
| Interface | Hardware  | Clock(Freq) |  RX/TX  : Alt |
#===========#===========#=============#=========-=====#
|   UART1   |  LPUART6  | OSC(24 MHz) |  D0/D1  : 2/2 |
|   UART2   |  LPUART4  | OSC(24 MHz) |  D7/D8  : 2/2 |
|   UART3   |  LPUART2  | OSC(24 MHz) | D15/D14 : 2/2 |
|   UART4   |  LPUART3  | OSC(24 MHz) | D16/D17 : 2/2 |
|   UART5   |  LPUART8  | OSC(24 MHz) | D21/D20 : 2/2 |
|   UART6   |  LPUART1  | OSC(24 MHz) | D25/D24 : 2/2 |
|   UART7   |  LPUART7  | OSC(24 MHz) | D28/D29 : 2/2 |
#===========#===========#=============#=========-=====#


```go
const (
	SPI1_SDI_PIN	= D12
	SPI1_SDO_PIN	= D11
	SPI1_SCK_PIN	= D13
	SPI1_CS_PIN	= D10

	SPI2_SDI_PIN	= D1
	SPI2_SDO_PIN	= D26
	SPI2_SCK_PIN	= D27
	SPI2_CS_PIN	= D0

	SPI3_SDI_PIN	= D34
	SPI3_SDO_PIN	= D35
	SPI3_SCK_PIN	= D37
	SPI3_CS_PIN	= D36
)
```

#==================================================================#
|                               SPI                                |
#===========#==========#===============#===========================#
| Interface | Hardware |  Clock(Freq)  | SDI/SDO/SCK/CS  :   Alt   |
#===========#==========#===============#=================-=========#
|   SPI1    |  LPSPI4  | PLL2(132 MHz) | D12/D11/D13/D10 : 3/3/3/3 |
|   SPI2    |  LPSPI3  | PLL2(132 MHz) |  D1/D26/D27/D0  : 7/2/2/7 |
|   SPI3    |  LPSPI1  | PLL2(132 MHz) | D34/D35/D37/D36 : 4/4/4/4 |
#===========#==========#===============#=================-=========#


```go
const (
	I2C1_SDA_PIN	= D18
	I2C1_SCL_PIN	= D19

	I2C2_SDA_PIN	= D17
	I2C2_SCL_PIN	= D16

	I2C3_SDA_PIN	= D25
	I2C3_SCL_PIN	= D24
)
```

#====================================================#
|                         I2C                        |
#===========#==========#=============#===============#
| Interface | Hardware | Clock(Freq) | SDA/SCL : Alt |
#===========#==========#=============#=========-=====#
|   I2C1    |  LPI2C1  | OSC(24 MHz) | D18/D19 : 3/3 |
|   I2C2    |  LPI2C3  | OSC(24 MHz) | D17/D16 : 1/1 |
|   I2C3    |  LPI2C4  | OSC(24 MHz) | D25/D24 : 0/0 |
#===========#==========#=============#=========-=====#


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.

Deprecated: use 100 * machine.KHz or 400 * machine.KHz instead.


```go
const (
	// I2CReceive indicates target has received a message from the controller.
	I2CReceive	I2CTargetEvent	= iota

	// I2CRequest indicates the controller is expecting a message from the target.
	I2CRequest

	// I2CFinish indicates the controller has ended the transaction.
	//
	// I2C controllers can chain multiple receive/request messages without
	// relinquishing the bus by doing 'restarts'.  I2CFinish indicates the
	// bus has been relinquished by an I2C 'stop'.
	I2CFinish
)
```



```go
const (
	// I2CModeController represents an I2C peripheral in controller mode.
	I2CModeController	I2CMode	= iota

	// I2CModeTarget represents an I2C peripheral in target mode.
	I2CModeTarget
)
```



```go
const Device = deviceName
```

Device is the running program's chip name, such as "ATSAMD51J19A" or
"nrf52840". It is not the same as the CPU name.

The constant is some hardcoded default value if the program does not target a
particular chip but instead runs in WebAssembly for example.


```go
const (
	KHz	= 1000
	MHz	= 1000_000
	GHz	= 1000_000_000
)
```

Generic constants.


```go
const NoPin = Pin(0xff)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	// GPIO
	PinInput	PinMode	= iota
	PinInputPullup
	PinInputPulldown
	PinOutput
	PinOutputOpenDrain
	PinDisable

	// ADC
	PinInputAnalog

	// UART
	PinModeUARTTX
	PinModeUARTRX

	// SPI
	PinModeSPISDI
	PinModeSPISDO
	PinModeSPICLK
	PinModeSPICS

	// I2C
	PinModeI2CSDA
	PinModeI2CSCL
)
```



```go
const (
	PinInputPullUp		= PinInputPullup
	PinInputPullDown	= PinInputPulldown
)
```

Deprecated: use PinInputPullup and PinInputPulldown instead.


```go
const (
	PinRising	PinChange	= iota + 2
	PinFalling
	PinToggle
)
```



```go
const (
	//                   [Pad]:      Alt Func 0       Alt Func 1       Alt Func 2           Alt Func 3            Alt Func 4            Alt Func 5   Alt Func 6            Alt Func 7            Alt Func 8             Alt Func 9
	//                   ----------  ---------------  ---------------  -------------------  --------------------  --------------------  -----------  --------------------  --------------------  ---------------------  ----------------
	PA0	= portA + 0	// [AD_B0_00]: FLEXPWM2_PWMA03  XBAR1_INOUT14    REF_CLK_32K          USB_OTG2_ID           LPI2C1_SCLS           GPIO1_IO00   USDHC1_RESET_B        LPSPI3_SCK             ~                      ~
	PA1	= portA + 1	// [AD_B0_01]: FLEXPWM2_PWMB03  XBAR1_INOUT15    REF_CLK_24M          USB_OTG1_ID           LPI2C1_SDAS           GPIO1_IO01   EWM_OUT_B             LPSPI3_SDO             ~                      ~
	PA2	= portA + 2	// [AD_B0_02]: FLEXCAN2_TX      XBAR1_INOUT16    LPUART6_TX           USB_OTG1_PWR          FLEXPWM1_PWMX00       GPIO1_IO02   LPI2C1_HREQ           LPSPI3_SDI             ~                      ~
	PA3	= portA + 3	// [AD_B0_03]: FLEXCAN2_RX      XBAR1_INOUT17    LPUART6_RX           USB_OTG1_OC           FLEXPWM1_PWMX01       GPIO1_IO03   REF_CLK_24M           LPSPI3_PCS0            ~                      ~
	PA4	= portA + 4	// [AD_B0_04]: SRC_BOOT_MODE00  MQS_RIGHT        ENET_TX_DATA03       SAI2_TX_SYNC          CSI_DATA09            GPIO1_IO04   PIT_TRIGGER00         LPSPI3_PCS1            ~                      ~
	PA5	= portA + 5	// [AD_B0_05]: SRC_BOOT_MODE01  MQS_LEFT         ENET_TX_DATA02       SAI2_TX_BCLK          CSI_DATA08            GPIO1_IO05   XBAR1_INOUT17         LPSPI3_PCS2            ~                      ~
	PA6	= portA + 6	// [AD_B0_06]: JTAG_TMS         GPT2_COMPARE1    ENET_RX_CLK          SAI2_RX_BCLK          CSI_DATA07            GPIO1_IO06   XBAR1_INOUT18         LPSPI3_PCS3            ~                      ~
	PA7	= portA + 7	// [AD_B0_07]: JTAG_TCK         GPT2_COMPARE2    ENET_TX_ER           SAI2_RX_SYNC          CSI_DATA06            GPIO1_IO07   XBAR1_INOUT19         ENET_1588_EVENT3_OUT   ~                      ~
	PA8	= portA + 8	// [AD_B0_08]: JTAG_MOD         GPT2_COMPARE3    ENET_RX_DATA03       SAI2_RX_DATA          CSI_DATA05            GPIO1_IO08   XBAR1_IN20            ENET_1588_EVENT3_IN    ~                      ~
	PA9	= portA + 9	// [AD_B0_09]: JTAG_TDI         FLEXPWM2_PWMA03  ENET_RX_DATA02       SAI2_TX_DATA          CSI_DATA04            GPIO1_IO09   XBAR1_IN21            GPT2_CLK              SEMC_DQS4               ~
	PA10	= portA + 10	// [AD_B0_10]: JTAG_TDO         FLEXPWM1_PWMA03  ENET_CRS             SAI2_MCLK             CSI_DATA03            GPIO1_IO10   XBAR1_IN22            ENET_1588_EVENT0_OUT  FLEXCAN3_TX            ARM_TRACE_SWO
	PA11	= portA + 11	// [AD_B0_11]: JTAG_TRSTB       FLEXPWM1_PWMB03  ENET_COL             WDOG1_WDOG_B          CSI_DATA02            GPIO1_IO11   XBAR1_IN23            ENET_1588_EVENT0_IN   FLEXCAN3_RX            SEMC_CLK6
	PA12	= portA + 12	// [AD_B0_12]: LPI2C4_SCL       CCM_PMIC_READY   LPUART1_TX           WDOG2_WDOG_B          FLEXPWM1_PWMX02       GPIO1_IO12   ENET_1588_EVENT1_OUT  NMI_GLUE_NMI           ~                      ~
	PA13	= portA + 13	// [AD_B0_13]: LPI2C4_SDA       GPT1_CLK         LPUART1_RX           EWM_OUT_B             FLEXPWM1_PWMX03       GPIO1_IO13   ENET_1588_EVENT1_IN   REF_CLK_24M            ~                      ~
	PA14	= portA + 14	// [AD_B0_14]: USB_OTG2_OC      XBAR1_IN24       LPUART1_CTS_B        ENET_1588_EVENT0_OUT  CSI_VSYNC             GPIO1_IO14   FLEXCAN2_TX           FLEXCAN3_TX            ~                      ~
	PA15	= portA + 15	// [AD_B0_15]: USB_OTG2_PWR     XBAR1_IN25       LPUART1_RTS_B        ENET_1588_EVENT0_IN   CSI_HSYNC             GPIO1_IO15   FLEXCAN2_RX           WDOG1_WDOG_RST_B_DEB  FLEXCAN3_RX             ~
	PA16	= portA + 16	// [AD_B1_00]: USB_OTG2_ID      QTIMER3_TIMER0   LPUART2_CTS_B        LPI2C1_SCL            WDOG1_B               GPIO1_IO16   USDHC1_WP             KPP_ROW07             ENET2_1588_EVENT0_OUT  FLEXIO3_FLEXIO00
	PA17	= portA + 17	// [AD_B1_01]: USB_OTG1_PWR     QTIMER3_TIMER1   LPUART2_RTS_B        LPI2C1_SDA            CCM_PMIC_READY        GPIO1_IO17   USDHC1_VSELECT        KPP_COL07             ENET2_1588_EVENT0_IN   FLEXIO3_FLEXIO01
	PA18	= portA + 18	// [AD_B1_02]: USB_OTG1_ID      QTIMER3_TIMER2   LPUART2_TX           SPDIF_OUT             ENET_1588_EVENT2_OUT  GPIO1_IO18   USDHC1_CD_B           KPP_ROW06             GPT2_CLK               FLEXIO3_FLEXIO02
	PA19	= portA + 19	// [AD_B1_03]: USB_OTG1_OC      QTIMER3_TIMER3   LPUART2_RX           SPDIF_IN              ENET_1588_EVENT2_IN   GPIO1_IO19   USDHC2_CD_B           KPP_COL06             GPT2_CAPTURE1          FLEXIO3_FLEXIO03
	PA20	= portA + 20	// [AD_B1_04]: FLEXSPIB_DATA03  ENET_MDC         LPUART3_CTS_B        SPDIF_SR_CLK          CSI_PIXCLK            GPIO1_IO20   USDHC2_DATA0          KPP_ROW05             GPT2_CAPTURE2          FLEXIO3_FLEXIO04
	PA21	= portA + 21	// [AD_B1_05]: FLEXSPIB_DATA02  ENET_MDIO        LPUART3_RTS_B        SPDIF_OUT             CSI_MCLK              GPIO1_IO21   USDHC2_DATA1          KPP_COL05             GPT2_COMPARE1          FLEXIO3_FLEXIO05
	PA22	= portA + 22	// [AD_B1_06]: FLEXSPIB_DATA01  LPI2C3_SDA       LPUART3_TX           SPDIF_LOCK            CSI_VSYNC             GPIO1_IO22   USDHC2_DATA2          KPP_ROW04             GPT2_COMPARE2          FLEXIO3_FLEXIO06
	PA23	= portA + 23	// [AD_B1_07]: FLEXSPIB_DATA00  LPI2C3_SCL       LPUART3_RX           SPDIF_EXT_CLK         CSI_HSYNC             GPIO1_IO23   USDHC2_DATA3          KPP_COL04             GPT2_COMPARE3          FLEXIO3_FLEXIO07
	PA24	= portA + 24	// [AD_B1_08]: FLEXSPIA_SS1_B   FLEXPWM4_PWMA00  FLEXCAN1_TX          CCM_PMIC_READY        CSI_DATA09            GPIO1_IO24   USDHC2_CMD            KPP_ROW03             FLEXIO3_FLEXIO08        ~
	PA25	= portA + 25	// [AD_B1_09]: FLEXSPIA_DQS     FLEXPWM4_PWMA01  FLEXCAN1_RX          SAI1_MCLK             CSI_DATA08            GPIO1_IO25   USDHC2_CLK            KPP_COL03             FLEXIO3_FLEXIO09        ~
	PA26	= portA + 26	// [AD_B1_10]: FLEXSPIA_DATA03  WDOG1_B          LPUART8_TX           SAI1_RX_SYNC          CSI_DATA07            GPIO1_IO26   USDHC2_WP             KPP_ROW02             ENET2_1588_EVENT1_OUT  FLEXIO3_FLEXIO10
	PA27	= portA + 27	// [AD_B1_11]: FLEXSPIA_DATA02  EWM_OUT_B        LPUART8_RX           SAI1_RX_BCLK          CSI_DATA06            GPIO1_IO27   USDHC2_RESET_B        KPP_COL02             ENET2_1588_EVENT1_IN   FLEXIO3_FLEXIO11
	PA28	= portA + 28	// [AD_B1_12]: FLEXSPIA_DATA01  ACMP_OUT00       LPSPI3_PCS0          SAI1_RX_DATA00        CSI_DATA05            GPIO1_IO28   USDHC2_DATA4          KPP_ROW01             ENET2_1588_EVENT2_OUT  FLEXIO3_FLEXIO12
	PA29	= portA + 29	// [AD_B1_13]: FLEXSPIA_DATA00  ACMP_OUT01       LPSPI3_SDI           SAI1_TX_DATA00        CSI_DATA04            GPIO1_IO29   USDHC2_DATA5          KPP_COL01             ENET2_1588_EVENT2_IN   FLEXIO3_FLEXIO13
	PA30	= portA + 30	// [AD_B1_14]: FLEXSPIA_SCLK    ACMP_OUT02       LPSPI3_SDO           SAI1_TX_BCLK          CSI_DATA03            GPIO1_IO30   USDHC2_DATA6          KPP_ROW00             ENET2_1588_EVENT3_OUT  FLEXIO3_FLEXIO14
	PA31	= portA + 31	// [AD_B1_15]: FLEXSPIA_SS0_B   ACMP_OUT03       LPSPI3_SCK           SAI1_TX_SYNC          CSI_DATA02            GPIO1_IO31   USDHC2_DATA7          KPP_COL00             ENET2_1588_EVENT3_IN   FLEXIO3_FLEXIO15

	PB0	= portB + 0	// [B0_00]:    LCD_CLK          QTIMER1_TIMER0   MQS_RIGHT            LPSPI4_PCS0           FLEXIO2_FLEXIO00      GPIO2_IO00   SEMC_CSX01            ENET2_MDC              ~                      ~
	PB1	= portB + 1	// [B0_01]:    LCD_ENABLE       QTIMER1_TIMER1   MQS_LEFT             LPSPI4_SDI            FLEXIO2_FLEXIO01      GPIO2_IO01   SEMC_CSX02            ENET2_MDIO             ~                      ~
	PB2	= portB + 2	// [B0_02]:    LCD_HSYNC        QTIMER1_TIMER2   FLEXCAN1_TX          LPSPI4_SDO            FLEXIO2_FLEXIO02      GPIO2_IO02   SEMC_CSX03            ENET2_1588_EVENT0_OUT  ~                      ~
	PB3	= portB + 3	// [B0_03]:    LCD_VSYNC        QTIMER2_TIMER0   FLEXCAN1_RX          LPSPI4_SCK            FLEXIO2_FLEXIO03      GPIO2_IO03   WDOG2_RESET_B_DEB     ENET2_1588_EVENT0_IN   ~                      ~
	PB4	= portB + 4	// [B0_04]:    LCD_DATA00       QTIMER2_TIMER1   LPI2C2_SCL           ARM_TRACE0            FLEXIO2_FLEXIO04      GPIO2_IO04   SRC_BOOT_CFG00        ENET2_TDATA03          ~                      ~
	PB5	= portB + 5	// [B0_05]:    LCD_DATA01       QTIMER2_TIMER2   LPI2C2_SDA           ARM_TRACE1            FLEXIO2_FLEXIO05      GPIO2_IO05   SRC_BOOT_CFG01        ENET2_TDATA02          ~                      ~
	PB6	= portB + 6	// [B0_06]:    LCD_DATA02       QTIMER3_TIMER0   FLEXPWM2_PWMA00      ARM_TRACE2            FLEXIO2_FLEXIO06      GPIO2_IO06   SRC_BOOT_CFG02        ENET2_RX_CLK           ~                      ~
	PB7	= portB + 7	// [B0_07]:    LCD_DATA03       QTIMER3_TIMER1   FLEXPWM2_PWMB00      ARM_TRACE3            FLEXIO2_FLEXIO07      GPIO2_IO07   SRC_BOOT_CFG03        ENET2_TX_ER            ~                      ~
	PB8	= portB + 8	// [B0_08]:    LCD_DATA04       QTIMER3_TIMER2   FLEXPWM2_PWMA01      LPUART3_TX            FLEXIO2_FLEXIO08      GPIO2_IO08   SRC_BOOT_CFG04        ENET2_RDATA03          ~                      ~
	PB9	= portB + 9	// [B0_09]:    LCD_DATA05       QTIMER4_TIMER0   FLEXPWM2_PWMB01      LPUART3_RX            FLEXIO2_FLEXIO09      GPIO2_IO09   SRC_BOOT_CFG05        ENET2_RDATA02          ~                      ~
	PB10	= portB + 10	// [B0_10]:    LCD_DATA06       QTIMER4_TIMER1   FLEXPWM2_PWMA02      SAI1_TX_DATA03        FLEXIO2_FLEXIO10      GPIO2_IO10   SRC_BOOT_CFG06        ENET2_CRS              ~                      ~
	PB11	= portB + 11	// [B0_11]:    LCD_DATA07       QTIMER4_TIMER2   FLEXPWM2_PWMB02      SAI1_TX_DATA02        FLEXIO2_FLEXIO11      GPIO2_IO11   SRC_BOOT_CFG07        ENET2_COL              ~                      ~
	PB12	= portB + 12	// [B0_12]:    LCD_DATA08       XBAR1_INOUT10    ARM_TRACE_CLK        SAI1_TX_DATA01        FLEXIO2_FLEXIO12      GPIO2_IO12   SRC_BOOT_CFG08        ENET2_TDATA00          ~                      ~
	PB13	= portB + 13	// [B0_13]:    LCD_DATA09       XBAR1_INOUT11    ARM_TRACE_SWO        SAI1_MCLK             FLEXIO2_FLEXIO13      GPIO2_IO13   SRC_BOOT_CFG09        ENET2_TDATA01          ~                      ~
	PB14	= portB + 14	// [B0_14]:    LCD_DATA10       XBAR1_INOUT12    ARM_TXEV             SAI1_RX_SYNC          FLEXIO2_FLEXIO14      GPIO2_IO14   SRC_BOOT_CFG10        ENET2_TX_EN            ~                      ~
	PB15	= portB + 15	// [B0_15]:    LCD_DATA11       XBAR1_INOUT13    ARM_RXEV             SAI1_RX_BCLK          FLEXIO2_FLEXIO15      GPIO2_IO15   SRC_BOOT_CFG11        ENET2_TX_CLK          ENET2_REF_CLK2          ~
	PB16	= portB + 16	// [B1_00]:    LCD_DATA12       XBAR1_INOUT14    LPUART4_TX           SAI1_RX_DATA00        FLEXIO2_FLEXIO16      GPIO2_IO16   FLEXPWM1_PWMA03       ENET2_RX_ER           FLEXIO3_FLEXIO16        ~
	PB17	= portB + 17	// [B1_01]:    LCD_DATA13       XBAR1_INOUT15    LPUART4_RX           SAI1_TX_DATA00        FLEXIO2_FLEXIO17      GPIO2_IO17   FLEXPWM1_PWMB03       ENET2_RDATA00         FLEXIO3_FLEXIO17        ~
	PB18	= portB + 18	// [B1_02]:    LCD_DATA14       XBAR1_INOUT16    LPSPI4_PCS2          SAI1_TX_BCLK          FLEXIO2_FLEXIO18      GPIO2_IO18   FLEXPWM2_PWMA03       ENET2_RDATA01         FLEXIO3_FLEXIO18        ~
	PB19	= portB + 19	// [B1_03]:    LCD_DATA15       XBAR1_INOUT17    LPSPI4_PCS1          SAI1_TX_SYNC          FLEXIO2_FLEXIO19      GPIO2_IO19   FLEXPWM2_PWMB03       ENET2_RX_EN           FLEXIO3_FLEXIO19        ~
	PB20	= portB + 20	// [B1_04]:    LCD_DATA16       LPSPI4_PCS0      CSI_DATA15           ENET_RX_DATA00        FLEXIO2_FLEXIO20      GPIO2_IO20   GPT1_CLK              FLEXIO3_FLEXIO20       ~                      ~
	PB21	= portB + 21	// [B1_05]:    LCD_DATA17       LPSPI4_SDI       CSI_DATA14           ENET_RX_DATA01        FLEXIO2_FLEXIO21      GPIO2_IO21   GPT1_CAPTURE1         FLEXIO3_FLEXIO21       ~                      ~
	PB22	= portB + 22	// [B1_06]:    LCD_DATA18       LPSPI4_SDO       CSI_DATA13           ENET_RX_EN            FLEXIO2_FLEXIO22      GPIO2_IO22   GPT1_CAPTURE2         FLEXIO3_FLEXIO22       ~                      ~
	PB23	= portB + 23	// [B1_07]:    LCD_DATA19       LPSPI4_SCK       CSI_DATA12           ENET_TX_DATA00        FLEXIO2_FLEXIO23      GPIO2_IO23   GPT1_COMPARE1         FLEXIO3_FLEXIO23       ~                      ~
	PB24	= portB + 24	// [B1_08]:    LCD_DATA20       QTIMER1_TIMER3   CSI_DATA11           ENET_TX_DATA01        FLEXIO2_FLEXIO24      GPIO2_IO24   FLEXCAN2_TX           GPT1_COMPARE2         FLEXIO3_FLEXIO24        ~
	PB25	= portB + 25	// [B1_09]:    LCD_DATA21       QTIMER2_TIMER3   CSI_DATA10           ENET_TX_EN            FLEXIO2_FLEXIO25      GPIO2_IO25   FLEXCAN2_RX           GPT1_COMPARE3         FLEXIO3_FLEXIO25        ~
	PB26	= portB + 26	// [B1_10]:    LCD_DATA22       QTIMER3_TIMER3   CSI_DATA00           ENET_TX_CLK           FLEXIO2_FLEXIO26      GPIO2_IO26   ENET_REF_CLK          FLEXIO3_FLEXIO26       ~                      ~
	PB27	= portB + 27	// [B1_11]:    LCD_DATA23       QTIMER4_TIMER3   CSI_DATA01           ENET_RX_ER            FLEXIO2_FLEXIO27      GPIO2_IO27   LPSPI4_PCS3           FLEXIO3_FLEXIO27       ~                      ~
	PB28	= portB + 28	// [B1_12]:    LPUART5_TX       CSI_PIXCLK       ENET_1588_EVENT0_IN  FLEXIO2_FLEXIO28      GPIO2_IO28            USDHC1_CD_B  FLEXIO3_FLEXIO28       ~                     ~                      ~
	PB29	= portB + 29	// [B1_13]:    WDOG1_B          LPUART5_RX       CSI_VSYNC            ENET_1588_EVENT0_OUT  FLEXIO2_FLEXIO29      GPIO2_IO29   USDHC1_WP             SEMC_DQS4             FLEXIO3_FLEXIO29        ~
	PB30	= portB + 30	// [B1_14]:    ENET_MDC         FLEXPWM4_PWMA02  CSI_HSYNC            XBAR1_IN02            FLEXIO2_FLEXIO30      GPIO2_IO30   USDHC1_VSELECT        ENET2_TDATA00         FLEXIO3_FLEXIO30        ~
	PB31	= portB + 31	// [B1_15]:    ENET_MDIO        FLEXPWM4_PWMA03  CSI_MCLK             XBAR1_IN03            FLEXIO2_FLEXIO31      GPIO2_IO31   USDHC1_RESET_B        ENET2_TDATA01         FLEXIO3_FLEXIO31        ~

	PC0	= portC + 0	// [SD_B1_00]: USDHC2_DATA3     FLEXSPIB_DATA03  FLEXPWM1_PWMA03      SAI1_TX_DATA03        LPUART4_TX            GPIO3_IO00   SAI3_RX_DATA           ~                     ~                      ~
	PC1	= portC + 1	// [SD_B1_01]: USDHC2_DATA2     FLEXSPIB_DATA02  FLEXPWM1_PWMB03      SAI1_TX_DATA02        LPUART4_RX            GPIO3_IO01   SAI3_TX_DATA           ~                     ~                      ~
	PC2	= portC + 2	// [SD_B1_02]: USDHC2_DATA1     FLEXSPIB_DATA01  FLEXPWM2_PWMA03      SAI1_TX_DATA01        FLEXCAN1_TX           GPIO3_IO02   CCM_WAIT              SAI3_TX_SYNC           ~                      ~
	PC3	= portC + 3	// [SD_B1_03]: USDHC2_DATA0     FLEXSPIB_DATA00  FLEXPWM2_PWMB03      SAI1_MCLK             FLEXCAN1_RX           GPIO3_IO03   CCM_PMIC_READY        SAI3_TX_BCLK           ~                      ~
	PC4	= portC + 4	// [SD_B1_04]: USDHC2_CLK       FLEXSPIB_SCLK    LPI2C1_SCL           SAI1_RX_SYNC          FLEXSPIA_SS1_B        GPIO3_IO04   CCM_STOP              SAI3_MCLK              ~                      ~
	PC5	= portC + 5	// [SD_B1_05]: USDHC2_CMD       FLEXSPIA_DQS     LPI2C1_SDA           SAI1_RX_BCLK          FLEXSPIB_SS0_B        GPIO3_IO05   SAI3_RX_SYNC           ~                     ~                      ~
	PC6	= portC + 6	// [SD_B1_06]: USDHC2_RESET_B   FLEXSPIA_SS0_B   LPUART7_CTS_B        SAI1_RX_DATA00        LPSPI2_PCS0           GPIO3_IO06   SAI3_RX_BCLK           ~                     ~                      ~
	PC7	= portC + 7	// [SD_B1_07]: SEMC_CSX01       FLEXSPIA_SCLK    LPUART7_RTS_B        SAI1_TX_DATA00        LPSPI2_SCK            GPIO3_IO07    ~                     ~                     ~                      ~
	PC8	= portC + 8	// [SD_B1_08]: USDHC2_DATA4     FLEXSPIA_DATA00  LPUART7_TX           SAI1_TX_BCLK          LPSPI2_SD0            GPIO3_IO08   SEMC_CSX02             ~                     ~                      ~
	PC9	= portC + 9	// [SD_B1_09]: USDHC2_DATA5     FLEXSPIA_DATA01  LPUART7_RX           SAI1_TX_SYNC          LPSPI2_SDI            GPIO3_IO09    ~                     ~                     ~                      ~
	PC10	= portC + 10	// [SD_B1_10]: USDHC2_DATA6     FLEXSPIA_DATA02  LPUART2_RX           LPI2C2_SDA            LPSPI2_PCS2           GPIO3_IO10    ~                     ~                     ~                      ~
	PC11	= portC + 11	// [SD_B1_11]: USDHC2_DATA7     FLEXSPIA_DATA03  LPUART2_TX           LPI2C2_SCL            LPSPI2_PCS3           GPIO3_IO11    ~                     ~                     ~                      ~
	PC12	= portC + 12	// [SD_B0_00]: USDHC1_CMD       FLEXPWM1_PWMA00  LPI2C3_SCL           XBAR1_INOUT04         LPSPI1_SCK            GPIO3_IO12   FLEXSPIA_SS1_B        ENET2_TX_EN           SEMC_DQS4               ~
	PC13	= portC + 13	// [SD_B0_01]: USDHC1_CLK       FLEXPWM1_PWMB00  LPI2C3_SDA           XBAR1_INOUT05         LPSPI1_PCS0           GPIO3_IO13   FLEXSPIB_SS1_B        ENET2_TX_CLK          ENET2_REF_CLK2          ~
	PC14	= portC + 14	// [SD_B0_02]: USDHC1_DATA0     FLEXPWM1_PWMA01  LPUART8_CTS_B        XBAR1_INOUT06         LPSPI1_SDO            GPIO3_IO14   ENET2_RX_ER           SEMC_CLK5              ~                      ~
	PC15	= portC + 15	// [SD_B0_03]: USDHC1_DATA1     FLEXPWM1_PWMB01  LPUART8_RTS_B        XBAR1_INOUT07         LPSPI1_SDI            GPIO3_IO15   ENET2_RDATA00         SEMC_CLK6              ~                      ~
	PC16	= portC + 16	// [SD_B0_04]: USDHC1_DATA2     FLEXPWM1_PWMA02  LPUART8_TX           XBAR1_INOUT08         FLEXSPIB_SS0_B        GPIO3_IO16   CCM_CLKO1             ENET2_RDATA01          ~                      ~
	PC17	= portC + 17	// [SD_B0_05]: USDHC1_DATA3     FLEXPWM1_PWMB02  LPUART8_RX           XBAR1_INOUT09         FLEXSPIB_DQS          GPIO3_IO17   CCM_CLKO2             ENET2_RX_EN            ~                      ~
	PC18	= portC + 18	// [EMC_32]:   SEMC_DATA10      FLEXPWM3_PWMB01  LPUART7_RX           CCM_PMIC_RDY          CSI_DATA21            GPIO3_IO18   ENET2_TX_EN            ~                     ~                      ~
	PC19	= portC + 19	// [EMC_33]:   SEMC_DATA11      FLEXPWM3_PWMA02  USDHC1_RESET_B       SAI3_RX_DATA          CSI_DATA20            GPIO3_IO19   ENET2_TX_CLK          ENET2_REF_CLK2         ~                      ~
	PC20	= portC + 20	// [EMC_34]:   SEMC_DATA12      FLEXPWM3_PWMB02  USDHC1_VSELECT       SAI3_RX_SYNC          CSI_DATA19            GPIO3_IO20   ENET2_RX_ER            ~                     ~                      ~
	PC21	= portC + 21	// [EMC_35]:   SEMC_DATA13      XBAR1_INOUT18    GPT1_COMPARE1        SAI3_RX_BCLK          CSI_DATA18            GPIO3_IO21   USDHC1_CD_B           ENET2_RDATA00          ~                      ~
	PC22	= portC + 22	// [EMC_36]:   SEMC_DATA14      XBAR1_IN22       GPT1_COMPARE2        SAI3_TX_DATA          CSI_DATA17            GPIO3_IO22   USDHC1_WP             ENET2_RDATA01         FLEXCAN3_TX             ~
	PC23	= portC + 23	// [EMC_37]:   SEMC_DATA15      XBAR1_IN23       GPT1_COMPARE3        SAI3_MCLK             CSI_DATA16            GPIO3_IO23   USDHC2_WP             ENET2_RX_EN           FLEXCAN3_RX             ~
	PC24	= portC + 24	// [EMC_38]:   SEMC_DM01        FLEXPWM1_PWMA03  LPUART8_TX           SAI3_TX_BCLK          CSI_FIELD             GPIO3_IO24   USDHC2_VSELECT        ENET2_MDC              ~                      ~
	PC25	= portC + 25	// [EMC_39]:   SEMC_DQS         FLEXPWM1_PWMB03  LPUART8_RX           SAI3_TX_SYNC          WDOG1_WDOG_B          GPIO3_IO25   USDHC2_CD_B           ENET2_MDIO            SEMC_DQS4               ~
	PC26	= portC + 26	// [EMC_40]:   SEMC_RDY         GPT2_CAPTURE2    LPSPI1_PCS2          USB_OTG2_OC           ENET_MDC              GPIO3_IO26   USDHC2_RESET_B        SEMC_CLK5              ~                      ~
	PC27	= portC + 27	// [EMC_41]:   SEMC_CSX00       GPT2_CAPTURE1    LPSPI1_PCS3          USB_OTG2_PWR          ENET_MDIO             GPIO3_IO27   USDHC1_VSELECT         ~                     ~                      ~
	_	= portC + 28	//
	_	= portC + 29	//
	_	= portC + 30	//
	_	= portC + 31	//

	PD0	= portD + 0	// [EMC_00]:   SEMC_DATA00      FLEXPWM4_PWMA00  LPSPI2_SCK           XBAR1_XBAR_IN02       FLEXIO1_FLEXIO00      GPIO4_IO00    ~                     ~                     ~                      ~
	PD1	= portD + 1	// [EMC_01]:   SEMC_DATA01      FLEXPWM4_PWMB00  LPSPI2_PCS0          XBAR1_IN03            FLEXIO1_FLEXIO01      GPIO4_IO01    ~                     ~                     ~                      ~
	PD2	= portD + 2	// [EMC_02]:   SEMC_DATA02      FLEXPWM4_PWMA01  LPSPI2_SDO           XBAR1_INOUT04         FLEXIO1_FLEXIO02      GPIO4_IO02    ~                     ~                     ~                      ~
	PD3	= portD + 3	// [EMC_03]:   SEMC_DATA03      FLEXPWM4_PWMB01  LPSPI2_SDI           XBAR1_INOUT05         FLEXIO1_FLEXIO03      GPIO4_IO03    ~                     ~                     ~                      ~
	PD4	= portD + 4	// [EMC_04]:   SEMC_DATA04      FLEXPWM4_PWMA02  SAI2_TX_DATA         XBAR1_INOUT06         FLEXIO1_FLEXIO04      GPIO4_IO04    ~                     ~                     ~                      ~
	PD5	= portD + 5	// [EMC_05]:   SEMC_DATA05      FLEXPWM4_PWMB02  SAI2_TX_SYNC         XBAR1_INOUT07         FLEXIO1_FLEXIO05      GPIO4_IO05    ~                     ~                     ~                      ~
	PD6	= portD + 6	// [EMC_06]:   SEMC_DATA06      FLEXPWM2_PWMA00  SAI2_TX_BCLK         XBAR1_INOUT08         FLEXIO1_FLEXIO06      GPIO4_IO06    ~                     ~                     ~                      ~
	PD7	= portD + 7	// [EMC_07]:   SEMC_DATA07      FLEXPWM2_PWMB00  SAI2_MCLK            XBAR1_INOUT09         FLEXIO1_FLEXIO07      GPIO4_IO07    ~                     ~                     ~                      ~
	PD8	= portD + 8	// [EMC_08]:   SEMC_DM00        FLEXPWM2_PWMA01  SAI2_RX_DATA         XBAR1_INOUT17         FLEXIO1_FLEXIO08      GPIO4_IO08    ~                     ~                     ~                      ~
	PD9	= portD + 9	// [EMC_09]:   SEMC_ADDR00      FLEXPWM2_PWMB01  SAI2_RX_SYNC         FLEXCAN2_TX           FLEXIO1_FLEXIO09      GPIO4_IO09   FLEXSPI2_B_SS1_B       ~                     ~                      ~
	PD10	= portD + 10	// [EMC_10]:   SEMC_ADDR01      FLEXPWM2_PWMA02  SAI2_RX_BCLK         FLEXCAN2_RX           FLEXIO1_FLEXIO10      GPIO4_IO10   FLEXSPI2_B_SS0_B       ~                     ~                      ~
	PD11	= portD + 11	// [EMC_11]:   SEMC_ADDR02      FLEXPWM2_PWMB02  LPI2C4_SDA           USDHC2_RESET_B        FLEXIO1_FLEXIO11      GPIO4_IO11   FLEXSPI2_B_DQS         ~                     ~                      ~
	PD12	= portD + 12	// [EMC_12]:   SEMC_ADDR03      XBAR1_IN24       LPI2C4_SCL           USDHC1_WP             FLEXPWM1_PWMA03       GPIO4_IO12   FLEXSPI2_B_SCLK        ~                     ~                      ~
	PD13	= portD + 13	// [EMC_13]:   SEMC_ADDR04      XBAR1_IN25       LPUART3_TX           MQS_RIGHT             FLEXPWM1_PWMB03       GPIO4_IO13   FLEXSPI2_B_DATA00      ~                     ~                      ~
	PD14	= portD + 14	// [EMC_14]:   SEMC_ADDR05      XBAR1_INOUT19    LPUART3_RX           MQS_LEFT              LPSPI2_PCS1           GPIO4_IO14   FLEXSPI2_B_DATA01      ~                     ~                      ~
	PD15	= portD + 15	// [EMC_15]:   SEMC_ADDR06      XBAR1_IN20       LPUART3_CTS_B        SPDIF_OUT             QTIMER3_TIMER0        GPIO4_IO15   FLEXSPI2_B_DATA02      ~                     ~                      ~
	PD16	= portD + 16	// [EMC_16]:   SEMC_ADDR07      XBAR1_IN21       LPUART3_RTS_B        SPDIF_IN              QTIMER3_TIMER1        GPIO4_IO16   FLEXSPI2_B_DATA03      ~                     ~                      ~
	PD17	= portD + 17	// [EMC_17]:   SEMC_ADDR08      FLEXPWM4_PWMA03  LPUART4_CTS_B        FLEXCAN1_TX           QTIMER3_TIMER2        GPIO4_IO17    ~                     ~                     ~                      ~
	PD18	= portD + 18	// [EMC_18]:   SEMC_ADDR09      FLEXPWM4_PWMB03  LPUART4_RTS_B        FLEXCAN1_RX           QTIMER3_TIMER3        GPIO4_IO18   SNVS_VIO_5_CTL         ~                     ~                      ~
	PD19	= portD + 19	// [EMC_19]:   SEMC_ADDR11      FLEXPWM2_PWMA03  LPUART4_TX           ENET_RDATA01          QTIMER2_TIMER0        GPIO4_IO19   SNVS_VIO_5             ~                     ~                      ~
	PD20	= portD + 20	// [EMC_20]:   SEMC_ADDR12      FLEXPWM2_PWMB03  LPUART4_RX           ENET_RDATA00          QTIMER2_TIMER1        GPIO4_IO20    ~                     ~                     ~                      ~
	PD21	= portD + 21	// [EMC_21]:   SEMC_BA0         FLEXPWM3_PWMA03  LPI2C3_SDA           ENET_TDATA01          QTIMER2_TIMER2        GPIO4_IO21    ~                     ~                     ~                      ~
	PD22	= portD + 22	// [EMC_22]:   SEMC_BA1         FLEXPWM3_PWMB03  LPI2C3_SCL           ENET_TDATA00          QTIMER2_TIMER3        GPIO4_IO22   FLEXSPI2_A_SS1_B       ~                     ~                      ~
	PD23	= portD + 23	// [EMC_23]:   SEMC_ADDR10      FLEXPWM1_PWMA00  LPUART5_TX           ENET_RX_EN            GPT1_CAPTURE2         GPIO4_IO23   FLEXSPI2_A_DQS         ~                     ~                      ~
	PD24	= portD + 24	// [EMC_24]:   SEMC_CAS         FLEXPWM1_PWMB00  LPUART5_RX           ENET_TX_EN            GPT1_CAPTURE1         GPIO4_IO24   FLEXSPI2_A_SS0_B       ~                     ~                      ~
	PD25	= portD + 25	// [EMC_25]:   SEMC_RAS         FLEXPWM1_PWMA01  LPUART6_TX           ENET_TX_CLK           ENET_REF_CLK          GPIO4_IO25   FLEXSPI2_A_SCLK        ~                     ~                      ~
	PD26	= portD + 26	// [EMC_26]:   SEMC_CLK         FLEXPWM1_PWMB01  LPUART6_RX           ENET_RX_ER            FLEXIO1_FLEXIO12      GPIO4_IO26   FLEXSPI2_A_DATA00      ~                     ~                      ~
	PD27	= portD + 27	// [EMC_27]:   SEMC_CKE         FLEXPWM1_PWMA02  LPUART5_RTS_B        LPSPI1_SCK            FLEXIO1_FLEXIO13      GPIO4_IO27   FLEXSPI2_A_DATA01      ~                     ~                      ~
	PD28	= portD + 28	// [EMC_28]:   SEMC_WE          FLEXPWM1_PWMB02  LPUART5_CTS_B        LPSPI1_SDO            FLEXIO1_FLEXIO14      GPIO4_IO28   FLEXSPI2_A_DATA02      ~                     ~                      ~
	PD29	= portD + 29	// [EMC_29]:   SEMC_CS0         FLEXPWM3_PWMA00  LPUART6_RTS_B        LPSPI1_SDI            FLEXIO1_FLEXIO15      GPIO4_IO29   FLEXSPI2_A_DATA03      ~                     ~                      ~
	PD30	= portD + 30	// [EMC_30]:   SEMC_DATA08      FLEXPWM3_PWMB00  LPUART6_CTS_B        LPSPI1_PCS0           CSI_DATA23            GPIO4_IO30   ENET2_TDATA00          ~                     ~                      ~
	PD31	= portD + 31	// [EMC_31]:   SEMC_DATA09      FLEXPWM3_PWMA01  LPUART7_TX           LPSPI1_PCS1           CSI_DATA22            GPIO4_IO31   ENET2_TDATA01          ~                     ~                      ~
)
```



```go
const (
	Mode0	= 0
	Mode1	= 1
	Mode2	= 2
	Mode3	= 3
)
```

SPI phase and polarity configs CPOL and CPHA


```go
const (
	// ParityNone means to not use any parity checking. This is
	// the most common setting.
	ParityNone	UARTParity	= iota

	// ParityEven means to expect that the total number of 1 bits sent
	// should be an even number.
	ParityEven

	// ParityOdd means to expect that the total number of 1 bits sent
	// should be an odd number.
	ParityOdd
)
```






## Variables

```go
var (
	DefaultUART = UART1
)
```

Default peripherals


```go
var (
	UART1	= &_UART1
	_UART1	= UART{
		Bus:		nxp.LPUART6,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART6_RX_SELECT_INPUT_DAISY_GPIO_AD_B0_03_ALT2,
			sel:	&nxp.IOMUXC.LPUART6_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART6_TX_SELECT_INPUT_DAISY_GPIO_AD_B0_02_ALT2,
			sel:	&nxp.IOMUXC.LPUART6_TX_SELECT_INPUT,
		},
	}
	UART2	= &_UART2
	_UART2	= UART{
		Bus:		nxp.LPUART4,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART4_RX_SELECT_INPUT_DAISY_GPIO_B1_01_ALT2,
			sel:	&nxp.IOMUXC.LPUART4_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART4_TX_SELECT_INPUT_DAISY_GPIO_B1_00_ALT2,
			sel:	&nxp.IOMUXC.LPUART4_TX_SELECT_INPUT,
		},
	}
	UART3	= &_UART3
	_UART3	= UART{
		Bus:		nxp.LPUART2,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART2_RX_SELECT_INPUT_DAISY_GPIO_AD_B1_03_ALT2,
			sel:	&nxp.IOMUXC.LPUART2_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART2_TX_SELECT_INPUT_DAISY_GPIO_AD_B1_02_ALT2,
			sel:	&nxp.IOMUXC.LPUART2_TX_SELECT_INPUT,
		},
	}
	UART4	= &_UART4
	_UART4	= UART{
		Bus:		nxp.LPUART3,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART3_RX_SELECT_INPUT_DAISY_GPIO_AD_B1_07_ALT2,
			sel:	&nxp.IOMUXC.LPUART3_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART3_TX_SELECT_INPUT_DAISY_GPIO_AD_B1_06_ALT2,
			sel:	&nxp.IOMUXC.LPUART3_TX_SELECT_INPUT,
		},
	}
	UART5	= &_UART5
	_UART5	= UART{
		Bus:		nxp.LPUART8,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART8_RX_SELECT_INPUT_DAISY_GPIO_AD_B1_11_ALT2,
			sel:	&nxp.IOMUXC.LPUART8_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART8_TX_SELECT_INPUT_DAISY_GPIO_AD_B1_10_ALT2,
			sel:	&nxp.IOMUXC.LPUART8_TX_SELECT_INPUT,
		},
	}
	UART6	= &_UART6
	_UART6	= UART{
		Bus:		nxp.LPUART1,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
	}
	UART7	= &_UART7
	_UART7	= UART{
		Bus:		nxp.LPUART7,
		Buffer:		NewRingBuffer(),
		txBuffer:	NewRingBuffer(),
		muxRX: muxSelect{
			mux:	nxp.IOMUXC_LPUART7_RX_SELECT_INPUT_DAISY_GPIO_EMC_32_ALT2,
			sel:	&nxp.IOMUXC.LPUART7_RX_SELECT_INPUT,
		},
		muxTX: muxSelect{
			mux:	nxp.IOMUXC_LPUART7_TX_SELECT_INPUT_DAISY_GPIO_EMC_31_ALT2,
			sel:	&nxp.IOMUXC.LPUART7_TX_SELECT_INPUT,
		},
	}
)
```



```go
var (
	SPI0	= SPI1	// SPI0 is an alias of SPI1 (LPSPI4)
	SPI1	= &_SPI1
	_SPI1	= SPI{
		Bus:	nxp.LPSPI4,
		muxSDI: muxSelect{
			mux:	nxp.IOMUXC_LPSPI4_SDI_SELECT_INPUT_DAISY_GPIO_B0_01_ALT3,
			sel:	&nxp.IOMUXC.LPSPI4_SDI_SELECT_INPUT,
		},
		muxSDO: muxSelect{
			mux:	nxp.IOMUXC_LPSPI4_SDO_SELECT_INPUT_DAISY_GPIO_B0_02_ALT3,
			sel:	&nxp.IOMUXC.LPSPI4_SDO_SELECT_INPUT,
		},
		muxSCK: muxSelect{
			mux:	nxp.IOMUXC_LPSPI4_SCK_SELECT_INPUT_DAISY_GPIO_B0_03_ALT3,
			sel:	&nxp.IOMUXC.LPSPI4_SCK_SELECT_INPUT,
		},
		muxCS: muxSelect{
			mux:	nxp.IOMUXC_LPSPI4_PCS0_SELECT_INPUT_DAISY_GPIO_B0_00_ALT3,
			sel:	&nxp.IOMUXC.LPSPI4_PCS0_SELECT_INPUT,
		},
	}
	SPI2	= &_SPI2
	_SPI2	= SPI{
		Bus:	nxp.LPSPI3,
		muxSDI: muxSelect{
			mux:	nxp.IOMUXC_LPSPI3_SDI_SELECT_INPUT_DAISY_GPIO_AD_B0_02_ALT7,
			sel:	&nxp.IOMUXC.LPSPI3_SDI_SELECT_INPUT,
		},
		muxSDO: muxSelect{
			mux:	nxp.IOMUXC_LPSPI3_SDO_SELECT_INPUT_DAISY_GPIO_AD_B1_14_ALT2,
			sel:	&nxp.IOMUXC.LPSPI3_SDO_SELECT_INPUT,
		},
		muxSCK: muxSelect{
			mux:	nxp.IOMUXC_LPSPI3_SCK_SELECT_INPUT_DAISY_GPIO_AD_B1_15,
			sel:	&nxp.IOMUXC.LPSPI3_SCK_SELECT_INPUT,
		},
		muxCS: muxSelect{
			mux:	nxp.IOMUXC_LPSPI3_PCS0_SELECT_INPUT_DAISY_GPIO_AD_B0_03_ALT7,
			sel:	&nxp.IOMUXC.LPSPI3_PCS0_SELECT_INPUT,
		},
	}
	SPI3	= &_SPI3
	_SPI3	= SPI{
		Bus:	nxp.LPSPI1,
		muxSDI: muxSelect{
			mux:	nxp.IOMUXC_LPSPI1_SDI_SELECT_INPUT_DAISY_GPIO_SD_B0_03_ALT4,
			sel:	&nxp.IOMUXC.LPSPI1_SDI_SELECT_INPUT,
		},
		muxSDO: muxSelect{
			mux:	nxp.IOMUXC_LPSPI1_SDO_SELECT_INPUT_DAISY_GPIO_SD_B0_02_ALT4,
			sel:	&nxp.IOMUXC.LPSPI1_SDO_SELECT_INPUT,
		},
		muxSCK: muxSelect{
			mux:	nxp.IOMUXC_LPSPI1_SCK_SELECT_INPUT_DAISY_GPIO_SD_B0_00_ALT4,
			sel:	&nxp.IOMUXC.LPSPI1_SCK_SELECT_INPUT,
		},
		muxCS: muxSelect{
			mux:	nxp.IOMUXC_LPSPI1_PCS0_SELECT_INPUT_DAISY_GPIO_SD_B0_01_ALT4,
			sel:	&nxp.IOMUXC.LPSPI1_PCS0_SELECT_INPUT,
		},
	}
)
```



```go
var (
	I2C0	= I2C1	// I2C0 is an alias for I2C1 (LPI2C1)
	I2C1	= &_I2C1
	_I2C1	= I2C{
		Bus:	nxp.LPI2C1,
		sda:	I2C1_SDA_PIN,
		scl:	I2C1_SCL_PIN,
		muxSDA: muxSelect{
			mux:	nxp.IOMUXC_LPI2C1_SDA_SELECT_INPUT_DAISY_GPIO_AD_B1_01_ALT3,
			sel:	&nxp.IOMUXC.LPI2C1_SDA_SELECT_INPUT,
		},
		muxSCL: muxSelect{
			mux:	nxp.IOMUXC_LPI2C1_SCL_SELECT_INPUT_DAISY_GPIO_AD_B1_00_ALT3,
			sel:	&nxp.IOMUXC.LPI2C1_SCL_SELECT_INPUT,
		},
	}
	I2C2	= &_I2C2
	_I2C2	= I2C{
		Bus:	nxp.LPI2C3,
		sda:	I2C2_SDA_PIN,
		scl:	I2C2_SCL_PIN,
		muxSDA: muxSelect{
			mux:	nxp.IOMUXC_LPI2C3_SDA_SELECT_INPUT_DAISY_GPIO_AD_B1_06_ALT1,
			sel:	&nxp.IOMUXC.LPI2C3_SDA_SELECT_INPUT,
		},
		muxSCL: muxSelect{
			mux:	nxp.IOMUXC_LPI2C3_SCL_SELECT_INPUT_DAISY_GPIO_AD_B1_07_ALT1,
			sel:	&nxp.IOMUXC.LPI2C3_SCL_SELECT_INPUT,
		},
	}
	I2C3	= &_I2C3
	_I2C3	= I2C{
		Bus:	nxp.LPI2C4,
		sda:	I2C3_SDA_PIN,
		scl:	I2C3_SCL_PIN,
		muxSDA: muxSelect{
			mux:	nxp.IOMUXC_LPI2C4_SDA_SELECT_INPUT_DAISY_GPIO_AD_B0_13_ALT0,
			sel:	&nxp.IOMUXC.LPI2C4_SDA_SELECT_INPUT,
		},
		muxSCL: muxSelect{
			mux:	nxp.IOMUXC_LPI2C4_SCL_SELECT_INPUT_DAISY_GPIO_AD_B0_12_ALT0,
			sel:	&nxp.IOMUXC.LPI2C4_SCL_SELECT_INPUT,
		},
	}
)
```



```go
var (
	ErrTimeoutRNG		= errors.New("machine: RNG Timeout")
	ErrClockRNG		= errors.New("machine: RNG Clock Error")
	ErrSeedRNG		= errors.New("machine: RNG Seed Error")
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel	= errors.New("machine: no channel available for pin interrupt")
)
```



```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.


```go
var (
	ErrTxInvalidSliceSize		= errors.New("SPI write and read slices must be same size")
	errSPIInvalidMachineConfig	= errors.New("SPI port was not configured properly by the machine")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func CPUReset

```go
func CPUReset()
```

CPUReset performs a hard system reset.


### func InitADC

```go
func InitADC()
```

InitADC is not used by this machine. Use `(ADC).Configure()`.


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.




## type ADC

```go
type ADC struct {
	Pin Pin
}
```




### func (ADC) Configure

```go
func (a ADC) Configure(config ADCConfig)
```

Configure initializes the receiver's ADC peripheral and pin for analog input.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get performs a single ADC conversion, returning a 16-bit unsigned integer.
The value returned will be scaled (uniformly distributed) if necessary so
that it is always in the range [0..65535], regardless of the ADC's configured
bit size (resolution).




## type ADCConfig

```go
type ADCConfig struct {
	Reference	uint32	// analog reference voltage (AREF) in millivolts
	Resolution	uint32	// number of bits for a single conversion (e.g., 8, 10, 12)
	Samples		uint32	// number of samples for a single conversion (e.g., 4, 8, 16, 32)
	SampleTime	uint32	// sample time, in microseconds (µs)
}
```

ADCConfig holds ADC configuration parameters. If left unspecified, the zero
value of each parameter will use the peripheral's default settings.





## type I2C

```go
type I2C struct {
	Bus	*nxp.LPI2C_Type

	// these pins are initialized by each global I2C variable declared in the
	// board_teensy4x.go file according to the board manufacturer's default pin
	// mapping. they can be overridden with the I2CConfig argument given to
	// (*I2C) Configure(I2CConfig).
	sda, scl	Pin

	// these hold the input selector ("daisy chain") values that select which pins
	// are connected to the LPI2C device, and should be defined where the I2C
	// instance is declared (e.g., in the board definition). see the godoc
	// comments on type muxSelect for more details.
	muxSDA, muxSCL	muxSelect
}
```




### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure is intended to setup an I2C interface for transmit/receive.


### func (*I2C) ReadRegister

```go
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (I2C) ReadRegisterEx

```go
func (i2c I2C) ReadRegisterEx(address uint8, register uint8, data []byte) error
```

ReadRegisterEx transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (I2C) SetBaudRate

```go
func (i2c I2C) SetBaudRate(br uint32) error
```

SetBaudRate sets the communication speed for I2C.


### func (I2C) Tx

```go
func (i2c I2C) Tx(addr uint16, w, r []byte) error
```



### func (*I2C) WriteRegister

```go
func (i2c *I2C) WriteRegister(address uint8, register uint8, data []byte) error
```

WriteRegister transmits first the register and then the data to the
peripheral device.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily write to such registers. Also, it only works for
devices with 7-bit addresses, which is the vast majority.


### func (I2C) WriteRegisterEx

```go
func (i2c I2C) WriteRegisterEx(address uint8, register uint8, data []byte) error
```

WriteRegisterEx transmits first the register and then the data to the
peripheral device.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily write to such registers. Also, it only works for
devices with 7-bit addresses, which is the vast majority.




## type I2CConfig

```go
type I2CConfig struct {
	Frequency	uint32
	SDA		Pin
	SCL		Pin
}
```

I2CConfig is used to store config info for I2C.





## type I2CMode

```go
type I2CMode int
```

I2CMode determines if an I2C peripheral is in Controller or Target mode.





## type I2CTargetEvent

```go
type I2CTargetEvent uint8
```

I2CTargetEvent reflects events on the I2C bus





## type NullSerial

```go
type NullSerial struct {
}
```

NullSerial is a serial version of /dev/null (or null router): it drops
everything that is written to it.



### func (NullSerial) Buffered

```go
func (ns NullSerial) Buffered() int
```

Buffered returns how many bytes are buffered in the UART. It always returns 0
as there are no bytes to read.


### func (NullSerial) Configure

```go
func (ns NullSerial) Configure(config UARTConfig) error
```

Configure does nothing: the null serial has no configuration.


### func (NullSerial) ReadByte

```go
func (ns NullSerial) ReadByte() (byte, error)
```

ReadByte always returns an error because there aren't any bytes to read.


### func (NullSerial) Write

```go
func (ns NullSerial) Write(p []byte) (n int, err error)
```

Write is a no-op: none of the data is being written and it will not return an
error.


### func (NullSerial) WriteByte

```go
func (ns NullSerial) WriteByte(b byte) error
```

WriteByte is a no-op: the null serial doesn't write bytes.




## type PDMConfig

```go
type PDMConfig struct {
	Stereo	bool
	DIN	Pin
	CLK	Pin
}
```






## type PWMConfig

```go
type PWMConfig struct {
	// PWM period in nanosecond. Leaving this zero will pick a reasonable period
	// value for use with LEDs.
	// If you want to configure a frequency instead of a period, you can use the
	// following formula to calculate a period from a frequency:
	//
	//     period = 1e9 / frequency
	//
	Period uint64
}
```

PWMConfig allows setting some configuration while configuring a PWM
peripheral. A zero PWMConfig is ready to use for simple applications such as
dimming LEDs.





## type Pin

```go
type Pin uint8
```

Pin is a single pin on a chip, which may be connected to other hardware
devices. It can either be used directly as GPIO pin or it can be used in
other peripherals like ADC, I2C, etc.



### func (Pin) Configure

```go
func (p Pin) Configure(config PinConfig)
```

Configure sets the GPIO pad and pin properties, and selects the appropriate
alternate function, for a given Pin and PinConfig.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get returns the current value of a GPIO pin.


### func (Pin) High

```go
func (p Pin) High()
```

High sets this GPIO pin to high, assuming it has been configured as an output
pin. It is hardware dependent (and often undefined) what happens if you set a
pin to high that is not configured as an output pin.


### func (Pin) Low

```go
func (p Pin) Low()
```

Low sets this GPIO pin to low, assuming it has been configured as an output
pin. It is hardware dependent (and often undefined) what happens if you set a
pin to low that is not configured as an output pin.


### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set changes the value of the GPIO pin. The pin must be configured as output.


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

This call will replace a previously set callback on this pin. You can pass a
nil func to unset the pin change interrupt. If you do so, the change
parameter is ignored and can be set to any value (such as 0).


### func (Pin) Toggle

```go
func (p Pin) Toggle()
```

Toggle switches an output pin from low to high or from high to low.




## type PinChange

```go
type PinChange uint8
```






## type PinConfig

```go
type PinConfig struct {
	Mode PinMode
}
```






## type PinMode

```go
type PinMode uint8
```

PinMode sets the direction and pull mode of the pin. For example, PinOutput
sets the pin as an output and PinInputPullup sets the pin as an input with a
pull-up.





## type RingBuffer

```go
type RingBuffer struct {
	rxbuffer	[bufferSize]volatile.Register8
	head		volatile.Register8
	tail		volatile.Register8
}
```

RingBuffer is ring buffer implementation inspired by post at
https://www.embeddedrelated.com/showthread/comp.arch.embedded/77084-1.php



### func (*RingBuffer) Clear

```go
func (rb *RingBuffer) Clear()
```

Clear resets the head and tail pointer to zero.


### func (*RingBuffer) Get

```go
func (rb *RingBuffer) Get() (byte, bool)
```

Get returns a byte from the buffer. If the buffer is empty,
the method will return a false as the second value.


### func (*RingBuffer) Put

```go
func (rb *RingBuffer) Put(val byte) bool
```

Put stores a byte in the buffer. If the buffer is already
full, the method will return false.


### func (*RingBuffer) Used

```go
func (rb *RingBuffer) Used() uint8
```

Used returns how many bytes in buffer have been used.




## type SPI

```go
type SPI struct {
	Bus	*nxp.LPSPI_Type

	// these hold the input selector ("daisy chain") values that select which pins
	// are connected to the LPSPI device, and should be defined where the SPI
	// instance is declared (e.g., in the board definition). see the godoc
	// comments on type muxSelect for more details.
	muxSDI, muxSDO, muxSCK, muxCS	muxSelect

	// these are copied from SPIConfig, during (*SPI).Configure(SPIConfig), and
	// should be considered read-only for internal reference (i.e., modifying them
	// will have no desirable effect).
	sdi, sdo, sck, cs	Pin
	frequency		uint32

	// auxiliary state data used internally
	configured	bool
}
```




### func (*SPI) Configure

```go
func (spi *SPI) Configure(config SPIConfig) error
```

Configure is intended to setup an SPI interface for transmit/receive.


### func (*SPI) Transfer

```go
func (spi *SPI) Transfer(w byte) (byte, error)
```

Transfer writes/reads a single byte using the SPI interface.


### func (SPI) Tx

```go
func (spi SPI) Tx(w, r []byte) error
```

Tx handles read/write operation for SPI interface. Since SPI is a syncronous write/read
interface, there must always be the same number of bytes written as bytes read.
The Tx method knows about this, and offers a few different ways of calling it.

This form sends the bytes in tx buffer, putting the resulting bytes read into the rx buffer.
Note that the tx and rx buffers must be the same size:

	spi.Tx(tx, rx)

This form sends the tx buffer, ignoring the result. Useful for sending "commands" that return zeros
until all the bytes in the command packet have been received:

	spi.Tx(tx, nil)

This form sends zeros, putting the result into the rx buffer. Good for reading a "result packet":

	spi.Tx(nil, rx)




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	SDI		Pin
	SDO		Pin
	SCK		Pin
	CS		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





## type UART

```go
type UART struct {
	Bus		*nxp.LPUART_Type
	Buffer		*RingBuffer
	Interrupt	interrupt.Interrupt

	// txBuffer should be allocated globally (such as when UART is created) to
	// prevent it being reclaimed or cleaned up prematurely.
	txBuffer	*RingBuffer

	// these hold the input selector ("daisy chain") values that select which pins
	// are connected to the LPUART device, and should be defined where the UART
	// instance is declared. see the godoc comments on type muxSelect for more
	// details.
	muxRX, muxTX	muxSelect

	// these are copied from UARTConfig, during (*UART).Configure(UARTConfig), and
	// should be considered read-only for internal reference (i.e., modifying them
	// will have no desirable effect).
	rx, tx	Pin
	baud	uint32

	// auxiliary state data used internally
	configured	bool
	transmitting	volatile.Register32
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
```

Configure initializes a UART with the given UARTConfig and other default
settings.


### func (*UART) Disable

```go
func (uart *UART) Disable()
```

Disable disables the UART interface.

If any buffered data has not yet been transmitted, Disable waits until
transmission completes before disabling the interface. The receiver UART's
interrupt is also disabled, and the RX/TX pins are reconfigured for GPIO
input (pull-up).


### func (*UART) Read

```go
func (uart *UART) Read(data []byte) (n int, err error)
```

Read from the RX buffer.


### func (*UART) ReadByte

```go
func (uart *UART) ReadByte() (byte, error)
```

ReadByte reads a single byte from the RX buffer.
If there is no data in the buffer, returns an error.


### func (*UART) Receive

```go
func (uart *UART) Receive(data byte)
```

Receive handles adding data to the UART's data buffer.
Usually called by the IRQ handler for a machine.


### func (*UART) Sync

```go
func (uart *UART) Sync() error
```

Sync blocks the calling goroutine until all data in the output buffer has
been transmitted.


### func (*UART) Write

```go
func (uart *UART) Write(data []byte) (n int, err error)
```

Write data over the UART's Tx.
This function blocks until the data is finished being sent.


### func (*UART) WriteByte

```go
func (uart *UART) WriteByte(c byte) error
```

WriteByte writes a byte of data over the UART's Tx.
This function blocks until the data is finished being sent.




## type UARTConfig

```go
type UARTConfig struct {
	BaudRate	uint32
	TX		Pin
	RX		Pin
	RTS		Pin
	CTS		Pin
}
```

UARTConfig is a struct with which a UART (or similar object) can be
configured. The baud rate is usually respected, but TX and RX may be ignored
depending on the chip and the type of object.





## type UARTParity

```go
type UARTParity uint8
```

UARTParity is the parity setting to be used for UART communication.





