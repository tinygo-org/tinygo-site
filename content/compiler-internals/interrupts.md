---
title: "Interrupts"
weight: 2
---

Interrupts are crucial if you want to do anything high-performance on microcontrollers. Unfortunately, the only close substitute in the Go language ([`os/signal.Notify`](https://godoc.org/os/signal#Notify)) is rather heavyweight. Therefore TinyGo uses a simpler way to work with interrupts.

Note: in most cases, you shouldn't need to work directly with interrupts. The machine package tries to abstract it away to provide a simple interface to work with. However, this page is here if you want to work on the machine package or need to override an interrupt.

## Basics

Interrupts are a bit like threads, but without any real concurrency. You could see them as a kind of callbacks at the hardware level.

Perhaps the most simple example of an interrupt is the UART (serial) interrupt on receiving a byte. Almost every microcontroller supports this. To use it, you generally will need to take the following steps:

 1. You need to define an interrupt handler. This is a function like any other, but is somehow marked specially to function as an interrupt handler. It depends on the compiler and programming language how this is done exactly.
 2. You need to configure the UART to enable receive interrupts. This is a setting in the UART peripheral itself, not in the interrupt handler.
 3. Depending on the chip, you need to set a priority for the interrupt. Not all microcontrollers support interrupt priorities, but many (especially more advanced) do. You set this not in the peripheral but in a centralized interrupt controller.
 4. Finally, you need to enable the interrupt. Some very simple chips (like the AVR) do not support this: you are expected to control the interrupt by enabling/disabling interrupt sources as in step 2.

TinyGo abstracts most of the complications away, but it's good to be aware of what is going on behind the scenes.

## Example

The following is an example of how interrupts work:

```go
func (uart UART) Configure(config UARTConfig) {
	// [...] other configuration

	// Enable the receive interrupt (step 2).
	// What the below line does is that it enables exactly one interrupt source
	// in the UART: the RXDRDY (RX ready) source. Enabling this makes sure the
	// interrupt is triggered whenever a byte is received in the UART.
	nrf.UART0.INTENSET.Set(nrf.UART_INTENSET_RXDRDY_Msk)
	
	// Register a new interrupt handler (step 1). This is the TinyGo way of
	// saying to the compiler that the UART0.handleInterrupt function is
	// special and should be called whenever the UART0 interrupt is triggered.
	intr := interrupt.New(nrf.IRQ_UART0, UART0.handleInterrupt)

	// Now we have a handle to the interrupt. The default on this chip is the
	// highest possible priority. We'd like to set the UART to a lower
	// priority, which we do here. The magic constant here will in a future
	// version be replaced with a regular constant for a low-priority
	// interrupt.
	intr.SetPriority(0xc0)

	// Finally, the interrupt must be enabled. Without this, the interrupt will
	// still be triggered but the handler will never be called.
	intr.Enable()
}

// This is the function that will be called. As a convenience, the interrupt
// handle is also provided as a parameter but you can usually ignore it.
func (uart *UART) handleInterrupt(intr interrupt.Interrupt) {
	// Multiple interrupt sources are often mapped to a single interrupt
	// handler. Therefore, we need to differentiate between various events.
	// In this case, only one interrupt source has been configured so
	// technically we could avoid this check, but it's good practice (for
	// future changes) to check for the event anyway.
	if nrf.UART0.EVENTS_RXDRDY.Get() != 0 {
		// This particular chip won't retrigger the interrupt when this event
		// is not cleared, but it won't clear the event either. So if we want
		// to differentiate between events in the next interrupt (when this
		// interrupt was triggered by a hypothetical other interrupt source) we
		// need to manually clear it first.
		nrf.UART0.EVENTS_RXDRDY.Set(0x0)

		// And finally, we can receive the byte from the UART. This is done by
		// reading the RXD register, which also has the side effect of
		// informing the hardware that this byte has been read.
		b := byte(nrf.UART0.RXD.Get())

		// Now do whatever you'd like with the just received byte.
		uart.Receive(b)
	}
}
```


## Troubleshooting

### The interrupt won't fire

  * Check that you have enabled the interrupt source in the peripheral (step 2).
  * Check that you have enabled the interrupt with the `intr.Enable()` call (step 4).
  * Check that you are listening for the correct interrupt and not for an interrupt for a different peripheral, for example (step 1).

### The interrupt keeps firing

Depending on the chip family, you may need to clear an interrupt source. This is for example the case in the Microchip SAM family of microcontrollers. If you don't do this, the interrupt will fire continuously thinking it still needs to be handled even though you have long since handled the interrupt.

### I get a compile error

You may get a compile error like the following:

```text
src/machine/machine_nrf.go:91:23: closures are not supported in interrupt.New
```

This could mean several things:

  * If you're passing an inline function closure to `interrupt.New`, you are not allowed to use variables from the outer scope. However, global variables are fine.
  * If you're passing a bound method, make sure the method has a function receiver. That is, define the interrupt handler like so:
    
    ```go
    func (uart *UART) handleInterrupt(intr interrupt.Interrupt) { ... }
    ```
    
    instead of like so:
    
    ```go
    // this won't compile
    func (uart UART) handleInterrupt(intr interrupt.Interrupt) { ... }
    ```
    
    The reason is that the latter will actually make a copy of the UART variable to be stored with the interrupt, which is currently not supported in the compiler and will lead to reduced performance even when it becomes supported.
    
    Additionally, make sure the variable you're binding to is a global variable instead of a local variable.
