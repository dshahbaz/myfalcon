# Overview

The below details were gleaned from running the Windows-based MAX-Falcon
programmer (a bunch of times) and comparing the resulting `firmware.bin` with
the original unmodified `.bin`.

# Keys

## Diagram

Looking at the top of the Max Falcon-8:

| 1 | 2 | 3 | 4 |
|---|---|---|---|
| 5 | 6 | 7 | 8 |

## Firmware locations

These are the offsets into `firmware.bin` where the keys are programmed.

1: 0x5189
2: 0x5181
3: 0x5179
4: 0x5149
5: 0x518a
6: 0x5182
7: 0x517a
8: 0x514a

## Key values

For the authoritative key codes list that the Max Falcon-8 appears to use, see:

<https://github.com/benblazak/ergodox-firmware/blob/master/src/lib/usb/usage-page/keyboard.h>

Additional values, overridden from the above list, are:

Prog1: 0xd7
Prog2: 0xd8
Prog3: 0xd9
Prog4: 0xda
Prog5: 0xdb
Prog6: 0xdc
Prog7: 0xdd
Prog8: 0xde

Next track: 0xe9
Previous track: 0xea
Stop: 0xeb
Play/Pause: 0xec
Mute: 0xed
Vol Up: 0xee
Vol Down: 0xef
Media Select: 0xf0
Mail: 0xf1
Calculator: 0xf2
WWW Search: 0xf4
WWW Home: 0xf5
WWW Back: 0xf6
WWW Forward: 0xf7
WWW Stop: 0xf8
WWW Refresh: 0xf9
WWW Favorites: 0xfa

# Programs

Programs can contain up to 100 program sets. Since program sets are 8 bytes (see
below), this makes a total of 800 bytes per program. Programs are zero-padded?? XXX

## Program set format

Each program set occupies 8 bytes, which are:

program set offset + 0x00: Modifier(s) (see below)
program set offset + 0x01: Number of milliseconds between keys (0.0 - 3.0 / 0x00 - 0x1e)
program set offset + 0x02: First key
program set offset + 0x03: Second key
program set offset + 0x04: Third key
program set offset + 0x05: Fourth key
program set offset + 0x06: Fifth key
program set offset + 0x07: Sixth key

## Firmware locations

Each program set occupies 8 bytes, with a maximum of 100 program sets per
program, for a total of 800 bytes. The offsets below are for the start of the
program.

Program 1 location: 0x539c
set 1 location: 0x539c
set 2 location: 0x53a4
set 3 location: 0x53ac
...
set 100 location: 0x56b4  ( == 0x539c + 800 - 8)

Program 2 location: 0x56bc
set 1 location: 0x56bc
...
set 100 location: 0x59d4  ( == 0x56bc + 800 - 8)

Program 3 location: 0x59dc 
set 1 location: 0x59dc
...
set 100 location: 0x5cf4 ( == 0x59dc + 800 - 8)

Program 4 location: 0x5cfc 
set 1 location: 0x5cfc
...
set 100 location: 0x600c ( == 0x5cfc + 800 - 8)

Program 5 location: 0x601c 
set 1 location: 0x601c
...
set 100 location: 0x6334 ( == 0x601c + 800 - 8)

Program 6 location: 0x633c 
set 1 location: 0x633c
...
set 100 location: 0x6654 ( == 0x633c + 800 - 8)

Program 7 location: 0x665c 
set 1 location: 0x665c
...
set 100 location: 0x6974 ( == 0x665c + 800 - 8)

Program 8 location: 0x697c 
set 1 location: 0x697c
...
set 100 location: 0x6c94 ( == 0x697c + 800 - 8)

### Modifiers

The list below is gleaned from
<http://blog.mateusz.perlak.com/index.php/2016/12/05/max-falcon-8-keyboard-hacking/>

* LCtr = 0x01
* LShi = 0x02
* LAlt = 0x03
* LWin = 0x04
* RCtr = 0x05
* RShi = 0x06
* RAlt = 0x07
* RWin = 0x08
* RWin+RSHi = 0x09
* RWin+RCtr = 0x0A
* RWin+RAlt = 0x0B
* RWin+RCtr+RShi = 0x0C
* RCtr+RAlt = 0x0D
* RCtr+RShi = 0x0E
* RAlt+RShi = 0x0F
* RAlt+RCtr+Rshi = 0x10

