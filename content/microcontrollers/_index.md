---
title: "Go on microcontrollers"
chapter: true
weight: 7
---

# Go on microcontrollers

TinyGo was designed to run on microcontrollers, but the Go language wasn't. This means there are a few challenges to writing Go code for microcontrollers.

Microcontrollers have very little RAM and execute code directly from flash. Also, constant globals are generally put in flash whenever possible. The Go language itself heavily relies on garbage collection so care must be taken to avoid dynamic memory allocation.
