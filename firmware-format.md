# Keys

## Diagram

Looking at the top of the Max Falcon-8:

+----+----+----+----+
| 1  | 2  | 3  | 4  |
+-------------------+
| 5  | 6  | 7  | 8  |
+----+----+----+----+

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

## Values

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

## Firmware locations

### Prog1

Start location: 0x539c
set 1 location: 0x539c
set 2 location: 0x53a4
set 3 location: 0x53ac
...
set 100 location: 0x56b4  <- (0x539c + 800 - 8)

## Program Set format

prog start + set offset + 0x00: Modifier(s) (see below)
prog start + set offset + 0x01: Number of milliseconds (0.0 - 3.0 / 0x00 - 0x1e)
prog start + set offset + 0x02: First key
prog start + set offset + 0x03: Second key
prog start + set offset + 0x04: Third key
prog start + set offset + 0x05: Fourth key
prog start + set offset + 0x06: Fifth key
prog start + set offset + 0x07: Sixth key

### Modifiers

The list below is gleened from <http://blog.mateusz.perlak.com/index.php/2016/12/05/max-falcon-8-keyboard-hacking/>

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


