---
title: PWM
weight: 1
---

PWM, short for Pulse Width Modulation, is a trick to dim a digital output by switching the output on and off quickly enough so that the average remains at a value somewhere between on and off. Using this trick, a digital output starts to behave like an analog output. Some applications of PWM:

  * Dimming LEDs. By turning a LED on and off a few hundred times a second (or more), it appears as if the LED has a constant brightness. The duty cycle (how much percent of the time the output is high) determines how bright the LED appears to be.
  * Creating square waves, for example in simple speakers. By controlling the PWM frequency and leaving the duty cycle at 50%, a speaker can produce a square wave of a particular frequency.
  * Controlling servo motors. Most servos are controlled by sending a pulse between 1ms and 2ms to control the position, with a full PWM period of 20ms. That means that for the middle position, the signal should be 1.5ms high and then 18.5ms low (together 20ms).

The PWM API in TinyGo is a bit different from some other PWM APIs, such as the ones in [Arduino](https://www.arduino.cc/reference/en/language/functions/analog-io/analogwrite/), [Mbed](https://os.mbed.com/handbook/PwmOut), [CircuitPython](https://learn.adafruit.com/circuitpython-essentials/circuitpython-pwm), and some others. This might make the API a bit harder to use, but it also makes some things explicit that would otherwise be hidden and cause problems.

If you just want to fade an LED, feel free to skip the background section. But a bit of background might help to understand PWMs a bit more.

## Background

In the example below you can control a simulated LED using PWM. You can change the duty cycle and the speed and see how that affects the LED. Notice how, if the frequency is high enough, the output appears continuous.

<label>
  Duty cycle: <span id="pwm-dutycycle">50%</span>
  <input id="pwm-dutycycle-input" type="range" value="50" min="0" max="100"/>
</label>

<label>
  Speed: <span id="pwm-period">1000ms</span> or <span id="pwm-frequency">1.0Hz</span>
  <input id="pwm-period-input" type="range" value="0" min="0" max="30"/>
</label>

<svg viewBox="0 0 340 30" width="340" height="30">
  <rect x="0.5" y="0.5" width="302" height="28" fill="transparent" stroke="gray"/>
  <path id="pwm-path" fill="transparent" stroke="black" stroke-linecap="square" d="M 1.5 27.5 H 1.5 V 1.5 H 151.5 V 27.5 H 301.5"/>
  <circle id="pwm-led" cx="325" cy="14.5" r="10" fill="gray" stroke-width="1" stroke="gray"/>
</svg>

<style id="pwm-styles">
/* Note: these properties will be modified from JavaScript using CSSOM. */
#pwm-led {
  animation-duration: 1s;
  animation-name: pwm;
  animation-iteration-count: infinite;
}
@keyframes pwm {
  from {
    fill: transparent;
  }
  10% {
    fill: red;
  }
  50% {
    fill: red;
  }
  60% {
    fill: transparent;
  }
  to {
    fill: transparent;
  }
}
</style>

<script>
'use strict';
var period, frequency, dutyCycle;

// Update the graph and the LED animation.
function updateView() {
  if (frequency === undefined)
    return; // updateSpeed not yet called

  // Update the graph.
  let graphTop = 1.5;
  let graphBottom = 27.5;
  let graphStart = 1.5;
  let graphEnd = 301.5;
  let width = graphEnd - graphStart
  let widthTime = 1; // 1s
  let path;
  if (dutyCycle == 0) {
    // Duty cycle is 0%, draw a straight line at the bottom.
    path = 'M ' + graphStart + ' ' + graphBottom + ' H ' + graphEnd;
  } else if (dutyCycle == 1) {
    // Duty cycle is 100%, draw a straight line at the top.
    path = 'M ' + graphStart + ' ' + graphTop + ' H ' + graphEnd;
  } else {
    // Duty cycle is somewhere in between, draw a line starting at the start of
    // a pulse.
    path = 'M ' + graphStart + ' ' + graphBottom;
    for (let t=0; t<widthTime; t+=period) {
      let pulseStart = graphStart + t*width;
      path += ' H ' + pulseStart + ' V ' + graphTop;
      let pulseEnd = graphStart + (t+period*dutyCycle)*width;
      if (pulseEnd > graphEnd) {
        // End of the graph, don't move the line back.
        continue;
      }
      if (dutyCycle >= 1) {
        continue;
      }
      path += ' H ' + pulseEnd + ' V ' + graphBottom;
    }
  }
  path += ' H ' + graphEnd;
  document.querySelector('#pwm-path').setAttribute('d', path);

  // Update the LED animation CSS.
  let animation;
  let style;
  for (let rule of document.querySelector('#pwm-styles').sheet.cssRules) {
    if (rule.type == CSSRule.KEYFRAMES_RULE && rule.name == 'pwm') {
      animation = rule;
    } else if (rule.type == CSSRule.STYLE_RULE && rule.selectorText == '#pwm-led') {
      style = rule;
    }
  }
  if (dutyCycle == 1) {
    style.style.animationName = '';
    style.style.fill = 'red';
  } else if (dutyCycle == 0) {
    style.style.animationName = '';
    style.style.fill = 'transparent';
  } else if (period < 0.035) {
    // Speed is too high for most monitors/browsers to display properly.
    // This is a bit cheating, but there is no other way to make it show what
    // it is intended to show.
    style.style.animationName = '';
    style.style.fill = 'rgba(255, 0, 0, ' + dutyCycle + ')';
  } else {
    let delayPercent = (frequency/50) * 100; // soften transition in 20ms
    animation.cssRules[1].keyText = delayPercent + '%';
    animation.cssRules[2].keyText = (dutyCycle*100) + '%';
    animation.cssRules[3].keyText = (dutyCycle*100 + delayPercent) + '%';
    style.style.animationName = 'pwm';
    style.style.animationDuration = period + 's';
  }
}

// Called when the duty cycle slider is changed.
function updateDutyCycle() {
  let input = document.querySelector('#pwm-dutycycle-input');
  dutyCycle = input.value / 100;
  document.querySelector('#pwm-dutycycle').textContent = input.value + '%';
  updateView();
}

// Called when the speed slider is changed.
function updateSpeed() {
  let input = document.querySelector('#pwm-period-input');
  period = Math.pow(10, (30-input.value) / 10) / 1000;
  frequency = 1 / period;
  let periodString = (period * 1000).toFixed(0)+'ms';
  if (period < 0.01) {
    periodString = (period * 1e3).toFixed(1)+'ms';
  }
  if (period < 0.001) {
    periodString = (period * 1e6).toFixed(0)+'µs';
  }
  document.querySelector('#pwm-period').textContent = periodString;
  let frequencyString = frequency.toFixed(1) + 'Hz';
  if (frequency >= 100) {
    frequencyString = frequency.toFixed(0) + 'Hz';
  }
  document.querySelector('#pwm-frequency').textContent = frequencyString;
  updateView();
}

// Listen to slider changes.
document.querySelector('#pwm-dutycycle-input').addEventListener('input', updateDutyCycle);
document.querySelector('#pwm-period-input').addEventListener('input', updateSpeed);

// Make sure the graph and LED are initialized properly.
// The default values should work already, but Firefox keeps input state across
// refreshes so this needs to be updated anyway.
updateDutyCycle();
updateSpeed();
</script>

In most cases, a PWM is a combination of a counter that counts up or down and a few channels with a value that this counter is continuously compared against.

The below example is a very simple 3-bit PWM peripheral. It can only count up to 7 before it starts again at zero. Most PWM peripherals are 8, 16 or 24 bits and thus have a much larger range. You can control the 'top' value as in most real-world PWM peripherals. It also has just two channels: many real-world PWMs have four or sometimes more channels.

<label>
  Counter top: <span id="timer-top">7</span>
  <input type="range" id="timer-top-input" value="7" min="0" max="7" oninput="updateTimerTop()"/>
</label>

<label>
  Channel compare A: <span class="timer-ch">5</span>
  <input type="range" class="timer-ch-input" value="5" min="0" max="7" oninput="updateChannelCompare()"/>
</label>

<label>
  Channel compare B: <span class="timer-ch">2</span>
  <input type="range" class="timer-ch-input" value="2" min="0" max="7" oninput="updateChannelCompare()"/>
</label>

<svg viewBox="0 0 250 50" width="250" height="50" style="background: #eee">
  <line x1="49.5" y1="0.5" x2="49.5" y2="49.5" fill="transparent" stroke="gray" stroke-linecap="square"/>
  <line x1="0.5" y1="49.5" x2="249.5" y2="49.5" fill="transparent" stroke="gray" stroke-linecap="square"/>
  <line id="timer-counter-bar-top" x1="28.5" y1="20.5" x2="39.5" y2="20.5" stroke="red" stroke-width="1" stroke-linecap="square"/>
  <line id="timer-counter-bar" x1="38.5" y1="20.5" x2="49.5" y2="20.5" stroke="black" stroke-width="1" stroke-linecap="square"/>
  <path id="timer-counter" stroke-width="1" fill="transparent" stroke="black"/>
</svg>

<script>
'use strict';

var timerTop;
var timerMax = 7;
var channels = [5, 2];
var counter = 7; // wrap around on first step
var counterHistory = [];
var counterAnimation;

function updateTimerTop() {
  timerTop = +document.querySelector('#timer-top-input').value;
  document.querySelector('#timer-top').textContent = timerTop;

  let counterBarTop = document.querySelector('#timer-counter-bar-top');
  counterBarTop.setAttribute('y1', 48-timerTop*6);
  counterBarTop.setAttribute('y2', 48-timerTop*6);
}

function updateChannelCompare() {
  let inputs = document.querySelectorAll('.timer-ch-input');
  let spans = document.querySelectorAll('.timer-ch');
  for (let i=0; i<inputs.length; i++) {
    channels[i] = inputs[i].value;
    spans[i].textContent = inputs[i].value;
  }
}

function doCounterStep() {
  if (counter == timerTop || counter == timerMax) {
    counter = 0;
  } else {
    counter += 1;
  }

  let counterBottom = 48;
  let counterDistance = 6;
  let counterStart = 50;

  let counterBar = document.querySelector('#timer-counter-bar');
  counterBar.setAttribute('y1', counterBottom-counter*counterDistance);
  counterBar.setAttribute('y2', counterBottom-counter*counterDistance);

  // Draw new path, including the history.
  counterHistory.unshift(counter);
  if (counterHistory.length > 22)
    counterHistory.pop();
  let path = '';
  for (let i=0; i<counterHistory.length; i++) {
    path += ' M ' + (counterStart + i*10-10) + ' ' + (counterBottom - counterHistory[i]*counterDistance) + ' h 10';
  }
  let counterEl = document.querySelector('#timer-counter');
  counterEl.setAttribute('d', path);
  if (counterAnimation) {
    counterAnimation.cancel();
  }
  counterAnimation = counterEl.animate([{transform: 'translateX(0px)'}, {transform: 'translateX(20px)'}], 666);

  setTimeout(function() {
    requestAnimationFrame(doCounterStep);
  }, 333);
}

// Update values on load.
updateChannelCompare();
updateTimerTop();

// Start animation.
doCounterStep();
</script>

While playing around with this, you may notice a few things:

  * The counter top determines the period of the output signal. It is easy to calculate the frequency from that, which is the inverse of the period.
  * The channel compare value determines the duty cycle of the signal. Every time the counter matches zero, all channel outputs are turned on and every time a channel matches the compare value, the channel is turned off again.
  * While you can change the period (and thus the frequency) of the output signal by changing the top value, the channel compare values will not automatically change with it and thus the duty cycle will change if you change the frequency: even though the 'on' period remains the same, the 'off' period changes.
  * While you can control the duty cycle of every channel independently, you cannot control the period (and thus the frequency) for individual channels. You can only change it for the entire PWM peripheral.

There are also a few things that are supported by many real-world PWMs but not by this simplified example:

  * (Nearly?) all PWM peripherls support a prescaler to reduce the tick frequency, which allows for much longer periods and thus much lower frequencies than would otherwise be possible.
  * Most PWM peripherals are actually called timer/counter peripherals, because they can also count events (by incrementing the counter with some external trigger instead of from an internal clock source) and they allow timestamping events by loading the current counter into one of the channels. These channels are also called capture/compare channels because they cannot just compare the value to the current counter but can also be loaded with the current counter.
  * PWM peripherals support many more configurations than just counting up. Some count down by default (reloading with the top value when they hit zero) and some count both up and down in a kind of triangle waveform. You don't need all of that for just controlling LEDs, but some hardware needs very precise control over the generated waveform.
  * These peripherals are also called timers, because sometimes they can be used as just that: a timer, with no external connections. You can keep track of the current time by reading the counter value and registering an interrupt on overflow so the timer value keeps incrementing.
  * Some chips (such as many AVRs) may not have a dedicated top value. Instead, you can assign one of the capture/compare channels as the top value. Of course, that means you cannot use that channel for generating a PWM output.

Most of these features are either abstracted away by the TinyGo API or left to be implemented in the future.

## Fading an LED

Fading an LED is relatively simple. For example, this code sets the on-board LED of an [Arduino Nano 33 IoT](../../arduino-nano33-iot/) to 25% output:

```go
func main() {
    // Configure the PWM peripheral. The default PWMConfig is fine for LEDs.
    pwm := machine.TCC0
    err := pwm.Configure(machine.PWMConfig{})
    handleError(err, "failed to configure TCC0")

    // Obtain a PWM channel.
    ch, err := pwm.Channel(machine.LED)
    handleError(err, "failed to obtain PWM channel for LED")

    // Set the channel to a 25% duty cycle.
    pwm.Set(ch, pwm.Top() / 4)

    // Wait before exiting.
    time.Sleep(time.Hour)
}
```

There are several steps involved here:

 1. The PWM must be configured first. An empty (zero) `PWMConfig` is designed to work fine for LEDs. Which PWM to pick depends on the chip and the board. If you check the [board description for this board](../../arduino-nano33-iot/), you can see that the Arduino Nano 33 IoT has TCC2 and TCC0 connected to the LED pin - we are using TCC0 here.
 2. Next a channel is obtained for this PWM. In general, one channel is connected to a single pin although there are exceptions.
 3. Finally you can set the output to 25% of the "top" value of the PWM, which is the value for a duty cycle of 100%. The top value varies a lot by hardware and by configuration, so you can't rely on any particular values. `pwm.Set(ch, pwm.Top())` sets the duty cycle to 100% and `pwm.Set(ch, 0)` sets the duty cycle to 0%.

You can also control multiple LEDs from a single PWM peripheral, as long as they are supported by that PWM peripheral. For example, a [Particle Argon](../../particle-argon) has an RGB LED that you can control like this:

```go
func main() {
    // Configure the PWM peripheral. The default PWMConfig is fine for LEDs.
    pwm := machine.PWM0
    err := pwm.Configure(machine.PWMConfig{})
    handleError(err, "failed to configure PWM0")

    // Obtain the various channels.
    red, err := pwm.Channel(machine.LED_RED)
    handleError(err, "failed to obtain PWM channel for LED_RED")
    green, err := pwm.Channel(machine.LED_GREEN)
    handleError(err, "failed to obtain PWM channel for LED_GREEN")
    blue, err := pwm.Channel(machine.LED_BLUE)
    handleError(err, "failed to obtain PWM channel for LED_BLUE")

    // Invert the output, because this is a common anode RGB LED.
    // The same effect could be obtained by using the inverse value in the Set
    // method (pwm.Set(red, 0) for full output, for example).
    pwm.SetInverting(red, true)
    pwm.SetInverting(green, true)
    pwm.SetInverting(blue, true)

    // Set the RGB LED to orange.
    pwm.Set(red, pwm.Top())     // 100% duty cycle
    pwm.Set(green, pwm.Top()/4) // 25% duty cycle
    pwm.Set(blue, 0)            // 0% duty cycle

    // Wait before exiting.
    time.Sleep(time.Hour)
}
```

Finally, this mechanism doesn't just let you fade an LED, but also blink it. This is done by using a PWM frequency of 1Hz or period of 1 second, which can be set using the `Period` member of `machine.PWMConfig`. This code sample is again for the [Arduino Nano 33 IoT](../../arduino-nano33-iot/):

```go
func main() {
    // Configure the PWM peripheral. A period of 1e9 nanoseconds (or 1 second)
    // results in a frequency of 1Hz.
    pwm := machine.TCC0
    err := pwm.Configure(machine.PWMConfig{Period: 1e9})
    handleError(err, "failed to configure TCC0")

    // Obtain a PWM channel.
    ch, err := pwm.Channel(machine.LED)
    handleError(err, "failed to obtain PWM channel for LED")

    // Set the channel to a 50% duty cycle, for equal on and off time.
    pwm.Set(ch, pwm.Top()/2)

    // Wait before exiting, while the LED continues blinking.
    time.Sleep(time.Hour)
}
```

This can have some advantages over manually blinking using the `High()` and `Low()` methods, such as that the blinking will be very precisely timed (no interference from interrupts) and you don't need to dedicate a goroutine to blinking this LED. Not all hardware supports these long PWM periods, however.


## Creating sound

For a speaker it's important to get the frequency exactly right, otherwise it'll have the wrong tone. Apart from the frequency it's important that the duty cycle is always 50% to get a square wave.

The following example is written for an [Adafruit Circuit Playground Bluefruit](../../circuit-playground-bluefruit). The frequencies in this case are obtained [from Wikipedia](https://en.wikipedia.org/wiki/Piano_key_frequencies) and match emergency vehicle sounds in some places.

```go
const lowNote = uint64(1e9) / 880  // A5 = 880Hz
const highNote = uint64(1e9) / 988 // B5 = 987.7666Hz

func main() {
    // Configure the PWM peripheral with a specific frequency.
    pwm := machine.PWM0
    err := pwm.Configure(machine.PWMConfig{
        // Use the lowest frequency note (or longest period) in the Period
        // member below. This ensures the PWM can later easily be reconfigured
        // with a higher frequency or shorter period. The other way around is
        // unsupported and may or may not work.
        Period: lowNote,
    })
    handleError(err, "failed to configure PWM")

    // Obtain a PWM channel from a pin connected to the speaker.
    ch, err := pwm.Channel(machine.D12)
    handleError(err, "failed to obtain PWM channel for the audio output")

    // Sound a siren. Be aware that this can be quite loud!
    for {
        // Create the high note (B5).
        pwm.SetPeriod(highNote)
        pwm.Set(ch, pwm.Top()/2)
        time.Sleep(500 * time.Millisecond)

        // Create the low note (A5).
        pwm.SetPeriod(lowNote)
        pwm.Set(ch, pwm.Top()/2)
        time.Sleep(500 * time.Millisecond)
    }
}
```

In this example the PWM is initially configured with the lowest frequency (the longest period). Then it is possible to later reconfigure the frequency using `SetPeriod` to the same or a higher frequency.

`SetPeriod` has the side effect that the duty cycle changes when the frequency changes. This happens because `SetPeriod` does not change the "on" time of any of the channels, it only changes the period. Therefore, to remain at 50% duty cycle the duty cycle needs to be set to 50% again after the period is changed.
