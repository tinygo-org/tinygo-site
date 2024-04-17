
---
title: maixbit
---


## Constants

```go
const (
	P00	Pin	= 0
	P01	Pin	= 1
	P02	Pin	= 2
	P03	Pin	= 3
	P04	Pin	= 4
	P05	Pin	= 5
	P06	Pin	= 6
	P07	Pin	= 7
	P08	Pin	= 8
	P09	Pin	= 9
	P10	Pin	= 10
	P11	Pin	= 11
	P12	Pin	= 12
	P13	Pin	= 13
	P14	Pin	= 14
	P15	Pin	= 15
	P16	Pin	= 16
	P17	Pin	= 17
	P18	Pin	= 18
	P19	Pin	= 19
	P20	Pin	= 20
	P21	Pin	= 21
	P22	Pin	= 22
	P23	Pin	= 23
	P24	Pin	= 24
	P25	Pin	= 25
	P26	Pin	= 26
	P27	Pin	= 27
	P28	Pin	= 28
	P29	Pin	= 29
	P30	Pin	= 30
	P31	Pin	= 31
	P32	Pin	= 32
	P33	Pin	= 33
	P34	Pin	= 34
	P35	Pin	= 35
	P36	Pin	= 36
	P37	Pin	= 37
	P38	Pin	= 38
	P39	Pin	= 39
	P40	Pin	= 40
	P41	Pin	= 41
	P42	Pin	= 42
	P43	Pin	= 43
	P44	Pin	= 44
	P45	Pin	= 45
	P46	Pin	= 46
	P47	Pin	= 47
)
```

K210 IO pins.


```go
const (
	FUNC_JTAG_TCLK			FPIOAFunction	= 0	// JTAG Test Clock
	FUNC_JTAG_TDI			FPIOAFunction	= 1	// JTAG Test Data In
	FUNC_JTAG_TMS			FPIOAFunction	= 2	// JTAG Test Mode Select
	FUNC_JTAG_TDO			FPIOAFunction	= 3	// JTAG Test Data Out
	FUNC_SPI0_D0			FPIOAFunction	= 4	// SPI0 Data 0
	FUNC_SPI0_D1			FPIOAFunction	= 5	// SPI0 Data 1
	FUNC_SPI0_D2			FPIOAFunction	= 6	// SPI0 Data 2
	FUNC_SPI0_D3			FPIOAFunction	= 7	// SPI0 Data 3
	FUNC_SPI0_D4			FPIOAFunction	= 8	// SPI0 Data 4
	FUNC_SPI0_D5			FPIOAFunction	= 9	// SPI0 Data 5
	FUNC_SPI0_D6			FPIOAFunction	= 10	// SPI0 Data 6
	FUNC_SPI0_D7			FPIOAFunction	= 11	// SPI0 Data 7
	FUNC_SPI0_SS0			FPIOAFunction	= 12	// SPI0 Chip Select 0
	FUNC_SPI0_SS1			FPIOAFunction	= 13	// SPI0 Chip Select 1
	FUNC_SPI0_SS2			FPIOAFunction	= 14	// SPI0 Chip Select 2
	FUNC_SPI0_SS3			FPIOAFunction	= 15	// SPI0 Chip Select 3
	FUNC_SPI0_ARB			FPIOAFunction	= 16	// SPI0 Arbitration
	FUNC_SPI0_SCLK			FPIOAFunction	= 17	// SPI0 Serial Clock
	FUNC_UARTHS_RX			FPIOAFunction	= 18	// UART High speed Receiver
	FUNC_UARTHS_TX			FPIOAFunction	= 19	// UART High speed Transmitter
	FUNC_RESV6			FPIOAFunction	= 20	// Reserved function
	FUNC_RESV7			FPIOAFunction	= 21	// Reserved function
	FUNC_CLK_SPI1			FPIOAFunction	= 22	// Clock SPI1
	FUNC_CLK_I2C1			FPIOAFunction	= 23	// Clock I2C1
	FUNC_GPIOHS0			FPIOAFunction	= 24	// GPIO High speed 0
	FUNC_GPIOHS1			FPIOAFunction	= 25	// GPIO High speed 1
	FUNC_GPIOHS2			FPIOAFunction	= 26	// GPIO High speed 2
	FUNC_GPIOHS3			FPIOAFunction	= 27	// GPIO High speed 3
	FUNC_GPIOHS4			FPIOAFunction	= 28	// GPIO High speed 4
	FUNC_GPIOHS5			FPIOAFunction	= 29	// GPIO High speed 5
	FUNC_GPIOHS6			FPIOAFunction	= 30	// GPIO High speed 6
	FUNC_GPIOHS7			FPIOAFunction	= 31	// GPIO High speed 7
	FUNC_GPIOHS8			FPIOAFunction	= 32	// GPIO High speed 8
	FUNC_GPIOHS9			FPIOAFunction	= 33	// GPIO High speed 9
	FUNC_GPIOHS10			FPIOAFunction	= 34	// GPIO High speed 10
	FUNC_GPIOHS11			FPIOAFunction	= 35	// GPIO High speed 11
	FUNC_GPIOHS12			FPIOAFunction	= 36	// GPIO High speed 12
	FUNC_GPIOHS13			FPIOAFunction	= 37	// GPIO High speed 13
	FUNC_GPIOHS14			FPIOAFunction	= 38	// GPIO High speed 14
	FUNC_GPIOHS15			FPIOAFunction	= 39	// GPIO High speed 15
	FUNC_GPIOHS16			FPIOAFunction	= 40	// GPIO High speed 16
	FUNC_GPIOHS17			FPIOAFunction	= 41	// GPIO High speed 17
	FUNC_GPIOHS18			FPIOAFunction	= 42	// GPIO High speed 18
	FUNC_GPIOHS19			FPIOAFunction	= 43	// GPIO High speed 19
	FUNC_GPIOHS20			FPIOAFunction	= 44	// GPIO High speed 20
	FUNC_GPIOHS21			FPIOAFunction	= 45	// GPIO High speed 21
	FUNC_GPIOHS22			FPIOAFunction	= 46	// GPIO High speed 22
	FUNC_GPIOHS23			FPIOAFunction	= 47	// GPIO High speed 23
	FUNC_GPIOHS24			FPIOAFunction	= 48	// GPIO High speed 24
	FUNC_GPIOHS25			FPIOAFunction	= 49	// GPIO High speed 25
	FUNC_GPIOHS26			FPIOAFunction	= 50	// GPIO High speed 26
	FUNC_GPIOHS27			FPIOAFunction	= 51	// GPIO High speed 27
	FUNC_GPIOHS28			FPIOAFunction	= 52	// GPIO High speed 28
	FUNC_GPIOHS29			FPIOAFunction	= 53	// GPIO High speed 29
	FUNC_GPIOHS30			FPIOAFunction	= 54	// GPIO High speed 30
	FUNC_GPIOHS31			FPIOAFunction	= 55	// GPIO High speed 31
	FUNC_GPIO0			FPIOAFunction	= 56	// GPIO pin 0
	FUNC_GPIO1			FPIOAFunction	= 57	// GPIO pin 1
	FUNC_GPIO2			FPIOAFunction	= 58	// GPIO pin 2
	FUNC_GPIO3			FPIOAFunction	= 59	// GPIO pin 3
	FUNC_GPIO4			FPIOAFunction	= 60	// GPIO pin 4
	FUNC_GPIO5			FPIOAFunction	= 61	// GPIO pin 5
	FUNC_GPIO6			FPIOAFunction	= 62	// GPIO pin 6
	FUNC_GPIO7			FPIOAFunction	= 63	// GPIO pin 7
	FUNC_UART1_RX			FPIOAFunction	= 64	// UART1 Receiver
	FUNC_UART1_TX			FPIOAFunction	= 65	// UART1 Transmitter
	FUNC_UART2_RX			FPIOAFunction	= 66	// UART2 Receiver
	FUNC_UART2_TX			FPIOAFunction	= 67	// UART2 Transmitter
	FUNC_UART3_RX			FPIOAFunction	= 68	// UART3 Receiver
	FUNC_UART3_TX			FPIOAFunction	= 69	// UART3 Transmitter
	FUNC_SPI1_D0			FPIOAFunction	= 70	// SPI1 Data 0
	FUNC_SPI1_D1			FPIOAFunction	= 71	// SPI1 Data 1
	FUNC_SPI1_D2			FPIOAFunction	= 72	// SPI1 Data 2
	FUNC_SPI1_D3			FPIOAFunction	= 73	// SPI1 Data 3
	FUNC_SPI1_D4			FPIOAFunction	= 74	// SPI1 Data 4
	FUNC_SPI1_D5			FPIOAFunction	= 75	// SPI1 Data 5
	FUNC_SPI1_D6			FPIOAFunction	= 76	// SPI1 Data 6
	FUNC_SPI1_D7			FPIOAFunction	= 77	// SPI1 Data 7
	FUNC_SPI1_SS0			FPIOAFunction	= 78	// SPI1 Chip Select 0
	FUNC_SPI1_SS1			FPIOAFunction	= 79	// SPI1 Chip Select 1
	FUNC_SPI1_SS2			FPIOAFunction	= 80	// SPI1 Chip Select 2
	FUNC_SPI1_SS3			FPIOAFunction	= 81	// SPI1 Chip Select 3
	FUNC_SPI1_ARB			FPIOAFunction	= 82	// SPI1 Arbitration
	FUNC_SPI1_SCLK			FPIOAFunction	= 83	// SPI1 Serial Clock
	FUNC_SPI_PERIPHERAL_D0		FPIOAFunction	= 84	// SPI Peripheral Data 0
	FUNC_SPI_PERIPHERAL_SS		FPIOAFunction	= 85	// SPI Peripheral Select
	FUNC_SPI_PERIPHERAL_SCLK	FPIOAFunction	= 86	// SPI Peripheral Serial Clock
	FUNC_I2S0_MCLK			FPIOAFunction	= 87	// I2S0 Main Clock
	FUNC_I2S0_SCLK			FPIOAFunction	= 88	// I2S0 Serial Clock(BCLK)
	FUNC_I2S0_WS			FPIOAFunction	= 89	// I2S0 Word Select(LRCLK)
	FUNC_I2S0_IN_D0			FPIOAFunction	= 90	// I2S0 Serial Data Input 0
	FUNC_I2S0_IN_D1			FPIOAFunction	= 91	// I2S0 Serial Data Input 1
	FUNC_I2S0_IN_D2			FPIOAFunction	= 92	// I2S0 Serial Data Input 2
	FUNC_I2S0_IN_D3			FPIOAFunction	= 93	// I2S0 Serial Data Input 3
	FUNC_I2S0_OUT_D0		FPIOAFunction	= 94	// I2S0 Serial Data Output 0
	FUNC_I2S0_OUT_D1		FPIOAFunction	= 95	// I2S0 Serial Data Output 1
	FUNC_I2S0_OUT_D2		FPIOAFunction	= 96	// I2S0 Serial Data Output 2
	FUNC_I2S0_OUT_D3		FPIOAFunction	= 97	// I2S0 Serial Data Output 3
	FUNC_I2S1_MCLK			FPIOAFunction	= 98	// I2S1 Main Clock
	FUNC_I2S1_SCLK			FPIOAFunction	= 99	// I2S1 Serial Clock(BCLK)
	FUNC_I2S1_WS			FPIOAFunction	= 100	// I2S1 Word Select(LRCLK)
	FUNC_I2S1_IN_D0			FPIOAFunction	= 101	// I2S1 Serial Data Input 0
	FUNC_I2S1_IN_D1			FPIOAFunction	= 102	// I2S1 Serial Data Input 1
	FUNC_I2S1_IN_D2			FPIOAFunction	= 103	// I2S1 Serial Data Input 2
	FUNC_I2S1_IN_D3			FPIOAFunction	= 104	// I2S1 Serial Data Input 3
	FUNC_I2S1_OUT_D0		FPIOAFunction	= 105	// I2S1 Serial Data Output 0
	FUNC_I2S1_OUT_D1		FPIOAFunction	= 106	// I2S1 Serial Data Output 1
	FUNC_I2S1_OUT_D2		FPIOAFunction	= 107	// I2S1 Serial Data Output 2
	FUNC_I2S1_OUT_D3		FPIOAFunction	= 108	// I2S1 Serial Data Output 3
	FUNC_I2S2_MCLK			FPIOAFunction	= 109	// I2S2 Main Clock
	FUNC_I2S2_SCLK			FPIOAFunction	= 110	// I2S2 Serial Clock(BCLK)
	FUNC_I2S2_WS			FPIOAFunction	= 111	// I2S2 Word Select(LRCLK)
	FUNC_I2S2_IN_D0			FPIOAFunction	= 112	// I2S2 Serial Data Input 0
	FUNC_I2S2_IN_D1			FPIOAFunction	= 113	// I2S2 Serial Data Input 1
	FUNC_I2S2_IN_D2			FPIOAFunction	= 114	// I2S2 Serial Data Input 2
	FUNC_I2S2_IN_D3			FPIOAFunction	= 115	// I2S2 Serial Data Input 3
	FUNC_I2S2_OUT_D0		FPIOAFunction	= 116	// I2S2 Serial Data Output 0
	FUNC_I2S2_OUT_D1		FPIOAFunction	= 117	// I2S2 Serial Data Output 1
	FUNC_I2S2_OUT_D2		FPIOAFunction	= 118	// I2S2 Serial Data Output 2
	FUNC_I2S2_OUT_D3		FPIOAFunction	= 119	// I2S2 Serial Data Output 3
	FUNC_RESV0			FPIOAFunction	= 120	// Reserved function
	FUNC_RESV1			FPIOAFunction	= 121	// Reserved function
	FUNC_RESV2			FPIOAFunction	= 122	// Reserved function
	FUNC_RESV3			FPIOAFunction	= 123	// Reserved function
	FUNC_RESV4			FPIOAFunction	= 124	// Reserved function
	FUNC_RESV5			FPIOAFunction	= 125	// Reserved function
	FUNC_I2C0_SCLK			FPIOAFunction	= 126	// I2C0 Serial Clock
	FUNC_I2C0_SDA			FPIOAFunction	= 127	// I2C0 Serial Data
	FUNC_I2C1_SCLK			FPIOAFunction	= 128	// I2C1 Serial Clock
	FUNC_I2C1_SDA			FPIOAFunction	= 129	// I2C1 Serial Data
	FUNC_I2C2_SCLK			FPIOAFunction	= 130	// I2C2 Serial Clock
	FUNC_I2C2_SDA			FPIOAFunction	= 131	// I2C2 Serial Data
	FUNC_CMOS_XCLK			FPIOAFunction	= 132	// DVP System Clock
	FUNC_CMOS_RST			FPIOAFunction	= 133	// DVP System Reset
	FUNC_CMOS_PWDN			FPIOAFunction	= 134	// DVP Power Down Mode
	FUNC_CMOS_VSYNC			FPIOAFunction	= 135	// DVP Vertical Sync
	FUNC_CMOS_HREF			FPIOAFunction	= 136	// DVP Horizontal Reference output
	FUNC_CMOS_PCLK			FPIOAFunction	= 137	// Pixel Clock
	FUNC_CMOS_D0			FPIOAFunction	= 138	// Data Bit 0
	FUNC_CMOS_D1			FPIOAFunction	= 139	// Data Bit 1
	FUNC_CMOS_D2			FPIOAFunction	= 140	// Data Bit 2
	FUNC_CMOS_D3			FPIOAFunction	= 141	// Data Bit 3
	FUNC_CMOS_D4			FPIOAFunction	= 142	// Data Bit 4
	FUNC_CMOS_D5			FPIOAFunction	= 143	// Data Bit 5
	FUNC_CMOS_D6			FPIOAFunction	= 144	// Data Bit 6
	FUNC_CMOS_D7			FPIOAFunction	= 145	// Data Bit 7
	FUNC_SCCB_SCLK			FPIOAFunction	= 146	// SCCB Serial Clock
	FUNC_SCCB_SDA			FPIOAFunction	= 147	// SCCB Serial Data
	FUNC_UART1_CTS			FPIOAFunction	= 148	// UART1 Clear To Send
	FUNC_UART1_DSR			FPIOAFunction	= 149	// UART1 Data Set Ready
	FUNC_UART1_DCD			FPIOAFunction	= 150	// UART1 Data Carrier Detect
	FUNC_UART1_RI			FPIOAFunction	= 151	// UART1 Ring Indicator
	FUNC_UART1_SIR_IN		FPIOAFunction	= 152	// UART1 Serial Infrared Input
	FUNC_UART1_DTR			FPIOAFunction	= 153	// UART1 Data Terminal Ready
	FUNC_UART1_RTS			FPIOAFunction	= 154	// UART1 Request To Send
	FUNC_UART1_OUT2			FPIOAFunction	= 155	// UART1 User-designated Output 2
	FUNC_UART1_OUT1			FPIOAFunction	= 156	// UART1 User-designated Output 1
	FUNC_UART1_SIR_OUT		FPIOAFunction	= 157	// UART1 Serial Infrared Output
	FUNC_UART1_BAUD			FPIOAFunction	= 158	// UART1 Transmit Clock Output
	FUNC_UART1_RE			FPIOAFunction	= 159	// UART1 Receiver Output Enable
	FUNC_UART1_DE			FPIOAFunction	= 160	// UART1 Driver Output Enable
	FUNC_UART1_RS485_EN		FPIOAFunction	= 161	// UART1 RS485 Enable
	FUNC_UART2_CTS			FPIOAFunction	= 162	// UART2 Clear To Send
	FUNC_UART2_DSR			FPIOAFunction	= 163	// UART2 Data Set Ready
	FUNC_UART2_DCD			FPIOAFunction	= 164	// UART2 Data Carrier Detect
	FUNC_UART2_RI			FPIOAFunction	= 165	// UART2 Ring Indicator
	FUNC_UART2_SIR_IN		FPIOAFunction	= 166	// UART2 Serial Infrared Input
	FUNC_UART2_DTR			FPIOAFunction	= 167	// UART2 Data Terminal Ready
	FUNC_UART2_RTS			FPIOAFunction	= 168	// UART2 Request To Send
	FUNC_UART2_OUT2			FPIOAFunction	= 169	// UART2 User-designated Output 2
	FUNC_UART2_OUT1			FPIOAFunction	= 170	// UART2 User-designated Output 1
	FUNC_UART2_SIR_OUT		FPIOAFunction	= 171	// UART2 Serial Infrared Output
	FUNC_UART2_BAUD			FPIOAFunction	= 172	// UART2 Transmit Clock Output
	FUNC_UART2_RE			FPIOAFunction	= 173	// UART2 Receiver Output Enable
	FUNC_UART2_DE			FPIOAFunction	= 174	// UART2 Driver Output Enable
	FUNC_UART2_RS485_EN		FPIOAFunction	= 175	// UART2 RS485 Enable
	FUNC_UART3_CTS			FPIOAFunction	= 176	// UART3 Clear To Send
	FUNC_UART3_DSR			FPIOAFunction	= 177	// UART3 Data Set Ready
	FUNC_UART3_DCD			FPIOAFunction	= 178	// UART3 Data Carrier Detect
	FUNC_UART3_RI			FPIOAFunction	= 179	// UART3 Ring Indicator
	FUNC_UART3_SIR_IN		FPIOAFunction	= 180	// UART3 Serial Infrared Input
	FUNC_UART3_DTR			FPIOAFunction	= 181	// UART3 Data Terminal Ready
	FUNC_UART3_RTS			FPIOAFunction	= 182	// UART3 Request To Send
	FUNC_UART3_OUT2			FPIOAFunction	= 183	// UART3 User-designated Output 2
	FUNC_UART3_OUT1			FPIOAFunction	= 184	// UART3 User-designated Output 1
	FUNC_UART3_SIR_OUT		FPIOAFunction	= 185	// UART3 Serial Infrared Output
	FUNC_UART3_BAUD			FPIOAFunction	= 186	// UART3 Transmit Clock Output
	FUNC_UART3_RE			FPIOAFunction	= 187	// UART3 Receiver Output Enable
	FUNC_UART3_DE			FPIOAFunction	= 188	// UART3 Driver Output Enable
	FUNC_UART3_RS485_EN		FPIOAFunction	= 189	// UART3 RS485 Enable
	FUNC_TIMER0_TOGGLE1		FPIOAFunction	= 190	// TIMER0 Toggle Output 1
	FUNC_TIMER0_TOGGLE2		FPIOAFunction	= 191	// TIMER0 Toggle Output 2
	FUNC_TIMER0_TOGGLE3		FPIOAFunction	= 192	// TIMER0 Toggle Output 3
	FUNC_TIMER0_TOGGLE4		FPIOAFunction	= 193	// TIMER0 Toggle Output 4
	FUNC_TIMER1_TOGGLE1		FPIOAFunction	= 194	// TIMER1 Toggle Output 1
	FUNC_TIMER1_TOGGLE2		FPIOAFunction	= 195	// TIMER1 Toggle Output 2
	FUNC_TIMER1_TOGGLE3		FPIOAFunction	= 196	// TIMER1 Toggle Output 3
	FUNC_TIMER1_TOGGLE4		FPIOAFunction	= 197	// TIMER1 Toggle Output 4
	FUNC_TIMER2_TOGGLE1		FPIOAFunction	= 198	// TIMER2 Toggle Output 1
	FUNC_TIMER2_TOGGLE2		FPIOAFunction	= 199	// TIMER2 Toggle Output 2
	FUNC_TIMER2_TOGGLE3		FPIOAFunction	= 200	// TIMER2 Toggle Output 3
	FUNC_TIMER2_TOGGLE4		FPIOAFunction	= 201	// TIMER2 Toggle Output 4
	FUNC_CLK_SPI2			FPIOAFunction	= 202	// Clock SPI2
	FUNC_CLK_I2C2			FPIOAFunction	= 203	// Clock I2C2
	FUNC_INTERNAL0			FPIOAFunction	= 204	// Internal function signal 0
	FUNC_INTERNAL1			FPIOAFunction	= 205	// Internal function signal 1
	FUNC_INTERNAL2			FPIOAFunction	= 206	// Internal function signal 2
	FUNC_INTERNAL3			FPIOAFunction	= 207	// Internal function signal 3
	FUNC_INTERNAL4			FPIOAFunction	= 208	// Internal function signal 4
	FUNC_INTERNAL5			FPIOAFunction	= 209	// Internal function signal 5
	FUNC_INTERNAL6			FPIOAFunction	= 210	// Internal function signal 6
	FUNC_INTERNAL7			FPIOAFunction	= 211	// Internal function signal 7
	FUNC_INTERNAL8			FPIOAFunction	= 212	// Internal function signal 8
	FUNC_INTERNAL9			FPIOAFunction	= 213	// Internal function signal 9
	FUNC_INTERNAL10			FPIOAFunction	= 214	// Internal function signal 10
	FUNC_INTERNAL11			FPIOAFunction	= 215	// Internal function signal 11
	FUNC_INTERNAL12			FPIOAFunction	= 216	// Internal function signal 12
	FUNC_INTERNAL13			FPIOAFunction	= 217	// Internal function signal 13
	FUNC_INTERNAL14			FPIOAFunction	= 218	// Internal function signal 14
	FUNC_INTERNAL15			FPIOAFunction	= 219	// Internal function signal 15
	FUNC_INTERNAL16			FPIOAFunction	= 220	// Internal function signal 16
	FUNC_INTERNAL17			FPIOAFunction	= 221	// Internal function signal 17
	FUNC_CONSTANT			FPIOAFunction	= 222	// Constant function
	FUNC_INTERNAL18			FPIOAFunction	= 223	// Internal function signal 18
	FUNC_DEBUG0			FPIOAFunction	= 224	// Debug function 0
	FUNC_DEBUG1			FPIOAFunction	= 225	// Debug function 1
	FUNC_DEBUG2			FPIOAFunction	= 226	// Debug function 2
	FUNC_DEBUG3			FPIOAFunction	= 227	// Debug function 3
	FUNC_DEBUG4			FPIOAFunction	= 228	// Debug function 4
	FUNC_DEBUG5			FPIOAFunction	= 229	// Debug function 5
	FUNC_DEBUG6			FPIOAFunction	= 230	// Debug function 6
	FUNC_DEBUG7			FPIOAFunction	= 231	// Debug function 7
	FUNC_DEBUG8			FPIOAFunction	= 232	// Debug function 8
	FUNC_DEBUG9			FPIOAFunction	= 233	// Debug function 9
	FUNC_DEBUG10			FPIOAFunction	= 234	// Debug function 10
	FUNC_DEBUG11			FPIOAFunction	= 235	// Debug function 11
	FUNC_DEBUG12			FPIOAFunction	= 236	// Debug function 12
	FUNC_DEBUG13			FPIOAFunction	= 237	// Debug function 13
	FUNC_DEBUG14			FPIOAFunction	= 238	// Debug function 14
	FUNC_DEBUG15			FPIOAFunction	= 239	// Debug function 15
	FUNC_DEBUG16			FPIOAFunction	= 240	// Debug function 16
	FUNC_DEBUG17			FPIOAFunction	= 241	// Debug function 17
	FUNC_DEBUG18			FPIOAFunction	= 242	// Debug function 18
	FUNC_DEBUG19			FPIOAFunction	= 243	// Debug function 19
	FUNC_DEBUG20			FPIOAFunction	= 244	// Debug function 20
	FUNC_DEBUG21			FPIOAFunction	= 245	// Debug function 21
	FUNC_DEBUG22			FPIOAFunction	= 246	// Debug function 22
	FUNC_DEBUG23			FPIOAFunction	= 247	// Debug function 23
	FUNC_DEBUG24			FPIOAFunction	= 248	// Debug function 24
	FUNC_DEBUG25			FPIOAFunction	= 249	// Debug function 25
	FUNC_DEBUG26			FPIOAFunction	= 250	// Debug function 26
	FUNC_DEBUG27			FPIOAFunction	= 251	// Debug function 27
	FUNC_DEBUG28			FPIOAFunction	= 252	// Debug function 28
	FUNC_DEBUG29			FPIOAFunction	= 253	// Debug function 29
	FUNC_DEBUG30			FPIOAFunction	= 254	// Debug function 30
	FUNC_DEBUG31			FPIOAFunction	= 255	// Debug function 31
)
```

Every pin on the Kendryte K210 is assigned to an FPIOA function.
Each pin can be configured with every function below.


```go
const (
	D0	= P00	// JTAG_TCLK
	D1	= P01	// JTAG_TDI
	D2	= P02	// JTAG_TMS
	D3	= P03	// JTAG_TDO
	D4	= P04	// UARTHS_RX
	D5	= P05	// UARTHS_TX
	D6	= P06	// RESV0
	D7	= P07	// RESV0
	D8	= P08	// GPIO1
	D9	= P09	// GPIO2
	D10	= P10	// GPIO3
	D11	= P11	// GPIO4
	D12	= P12	// GPIO5
	D13	= P13	// GPIO6
	D14	= P14	// GPIO7
	D15	= P15	// GPIO8
	D16	= P16	// GPIOHS0
	D17	= P17	// GPIOHS1
	D18	= P18	// GPIOHS2
	D19	= P19	// GPIOHS3
	D20	= P20	// GPIOHS4
	D21	= P21	// GPIOHS5
	D22	= P22	// GPIOHS6
	D23	= P23	// GPIOHS7
	D24	= P24	// GPIOHS8
	D25	= P25	// GPIOHS9
	D26	= P26	// GPIOHS10 / SPI0_SDI
	D27	= P27	// GPIOHS11 / SPI0_SCLK
	D28	= P28	// GPIOHS12 / SPI0_SDO
	D29	= P29	// GPIOHS13
	D30	= P30	// GPIOHS14
	D31	= P31	// GPIOHS15
	D32	= P32	// GPIOHS16
	D33	= P33	// GPIOHS17
	D34	= P34	// GPIOHS18
	D35	= P35	// GPIOHS19
)
```

Pins on the MAix Bit.


```go
const (
	LED		= LED1
	LED1		= LED_RED
	LED2		= LED_GREEN
	LED3		= LED_BLUE
	LED_RED		= D13
	LED_GREEN	= D12
	LED_BLUE	= D14
)
```



```go
const (
	UART_TX_PIN	= D5
	UART_RX_PIN	= D4
)
```

Default pins for UARTHS.


```go
const (
	SPI0_SCK_PIN	= D27
	SPI0_SDO_PIN	= D28
	SPI0_SDI_PIN	= D26
)
```

SPI pins.


```go
const (
	I2C0_SDA_PIN	= D34
	I2C0_SCL_PIN	= D35
)
```

I2C pins.


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
	PinInput	PinMode	= iota
	PinInputPullup
	PinInputPulldown
	PinOutput
)
```

Pin modes.


```go
const (
	PinInputPullUp		= PinInputPullup
	PinInputPullDown	= PinInputPulldown
)
```

Deprecated: use PinInputPullup and PinInputPulldown instead.


```go
const (
	PinRising	PinChange	= 1 << iota
	PinFalling
	PinToggle	= PinRising | PinFalling
)
```

GPIOHS pin interrupt events.


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
var DefaultUART = UART0
```



```go
var (
	SPI0	= SPI{
		Bus: kendryte.SPI0,
	}
	SPI1	= SPI{
		Bus: kendryte.SPI1,
	}
)
```

SPI on the MAix Bit.


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
	UART0	= &_UART0
	_UART0	= UART{Bus: kendryte.UARTHS, Buffer: NewRingBuffer()}
)
```



```go
var (
	I2C0	= (*I2C)(unsafe.Pointer(kendryte.I2C0))
	I2C1	= (*I2C)(unsafe.Pointer(kendryte.I2C1))
	I2C2	= (*I2C)(unsafe.Pointer(kendryte.I2C2))
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





## type FPIOAFunction

```go
type FPIOAFunction uint8
```






## type I2C

```go
type I2C struct {
	Bus kendryte.I2C_Type
}
```

I2C on the K210.



### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure is intended to setup the I2C interface.


### func (*I2C) ReadRegister

```go
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (*I2C) SetBaudRate

```go
func (i2c *I2C) SetBaudRate(br uint32) error
```

SetBaudRate sets the communication speed for I2C.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


### func (*I2C) WriteRegister

```go
func (i2c *I2C) WriteRegister(address uint8, register uint8, data []byte) error
```

WriteRegister transmits first the register and then the data to the
peripheral device.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily write to such registers. Also, it only works for
devices with 7-bit addresses, which is the vast majority.




## type I2CConfig

```go
type I2CConfig struct {
	Frequency	uint32
	SCL		Pin
	SDA		Pin
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

Configure this pin with the given configuration.
The pin must already be set as GPIO or GPIOHS pin.


### func (Pin) FPIOAFunction

```go
func (p Pin) FPIOAFunction() FPIOAFunction
```

FPIOAFunction returns the current FPIOA function of the pin.


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
func (p Pin) Set(high bool)
```

Set the pin to high or low.


### func (Pin) SetFPIOAFunction

```go
func (p Pin) SetFPIOAFunction(f FPIOAFunction)
```

SetFPIOAFunction is used to configure the pin for one of the FPIOA functions.
Each pin on the Kendryte K210 can be configured with any of the available FPIOA functions.


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

You can pass a nil func to unset the pin change interrupt. If you do so,
the change parameter is ignored and can be set to any value (such as 0).
If the pin is already configured with a callback, you must first unset
this pins interrupt before you can set a new callback.




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
	Bus *kendryte.SPI_Type
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure is intended to setup the SPI interface.
Only SPI controller 0 and 1 can be used because SPI2 is a special
peripheral-mode controller and SPI3 is used for flashing.


### func (SPI) Transfer

```go
func (spi SPI) Transfer(w byte) (byte, error)
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
	SCK		Pin
	SDO		Pin
	SDI		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





## type UART

```go
type UART struct {
	Bus	*kendryte.UARTHS_Type
	Buffer	*RingBuffer
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





