# MAX Falcon Programmer

A cross-platform, CLI firmware programmer for the [MAX Falcon-8 custom programmable keyboard](http://www.maxkeyboard.com/max-falcon-8-custom-programmable-mini-macropad-mechanical-keyboard-assembled.html), a handy little 8 key keyboard made by MAX.

The programmer supplied by the fine folks at MAX is Windows only, and while it
does work under Wine on linux and OSX, it's a bit... clunky. This project
emulates what the official programmer does.

[![Build Status](https://travis-ci.org/dshahbaz/myfalcon.svg?branch=master)](https://travis-ci.org/dshahbaz/myfalcon)

## WARNING

**THERE ARE NO CLAIMS BEING MADE AS TO WHETHER OR NOT THIS WILL DAMAGE YOUR
KEYBOARD. This tool is provided as-is.**

That said, I have been using this extensively during development/testing, and
have not damaged the keyboard. If you have a Windows machine with the [official
firmware programmer from MAX](http://www.maxkeyboard.com/download.html), it
should be enough of a safety net. Just follow the directions in the [offical
video](https://www.youtube.com/watch?v=IpYg5A78-hs) to reset your keyboard.

## Getting Started

IF YOU JUST WANT A BINARY TO RUN, see [releases](releases).

Otherwise, see [Building](#building).

### Defining a layout

Firmware definitions are written in protocol buffer *text format* (`.textpb`).
There is an example in
[firmware/examples/max_falcon8_test.textpb](firmware/examples/max_falcon8_test.textpb),
but basically:

To map a button to a key:

```
button1 {
  key: KEY_a_A
}
```

(A full list of USB HID key values is in
[firmware/proto/hid.proto](firmware/proto/hid.proto). See below for other
special key constants).

Alternatively, if a key can be represented as a string, you can also do:

```
button1 {
  string: 'a'
}
```

To assign a program to a button is a bit more involved:

```
button2 {
  # ctrl-d
  program {
    program_set {
      modifier: LCtr  # See below for possible modifiers.
      keys: [KEY_d_D]
    }
  }
}
```

A `program` can contain up to 100 `program_set`s. A `program_set` can contain up
to 6 keys. A more complicated example:

```
button4 {
  # Types "HELLO friend"
  program {
    program_set {
      modifier: LShi
      milliseconds_between_keys: 2  # Between 0 and 30 milliseconds.
      keys: [KEY_h_H]
    }

    program_set {
      modifier: LShi
      milliseconds_between_keys: 2
      keys: [KEY_e_E]
    }

    program_set {
      modifier: LShi
      milliseconds_between_keys: 2
      keys: [KEY_l_L]
    }

    program_set {
      modifier: LShi
      milliseconds_between_keys: 2
      keys: [KEY_l_L]
    }

    program_set {
      modifier: LShi
      milliseconds_between_keys: 2
      keys: [KEY_o_O]
    }

    program_set {
      milliseconds_between_keys: 2
      keys: [KEY_Spacebar]
    }

    program_set {
      milliseconds_between_keys: 2
      keys: [KEY_f_F, KEY_r_R, KEY_i_I, KEY_e_E, KEY_n_N, KEY_d_D]
      # Note: `string` values are not supported in program_sets.
    }
  }
}
```

#### Key constants

Available modifiers (defined in
[firmware/proto/firmware_keys.proto](firmware/proto/firmware_keys.proto)):

```
NoModifier
LCtr
LShi
LAlt
LWin
RCtr
RShi
RAlt
RWin
RWin_RSHi
RWin_RCtr
RWin_RAlt
RWin_RCtr_RShi
RCtr_RAlt
RCtr_RShi
RAlt_RShi
RAlt_RCtr_Rshi
```

Special keys (defined in
[firmware/proto/hid.proto](firmware/proto/hid.proto)):

```
Next_track
Previous_track
Stop
Play_Pause
Mute
Vol_Up
Vol_Down
Media_Select
Mail
Calculator
WWW_Search
WWW_Home
WWW_Back
WWW_Forward
WWW_Stop
WWW_Refresh
WWW_Favorites
```

### Running

Connect the keyboard and put it into programming mode (see the
[official video](https://www.youtube.com/watch?v=IpYg5A78-hs) for how this is
done). Mount it (or make sure it is automounted).

Assuming you've saved your layout from *Defining a layout* in
`my_new_layout.textpb`, you can verify the file without updating the firmware by
running:

```
myfalcon -text_proto my_new_layout.textpb -verify_only
```

Now to write the firmware:

```
myfalcon -text_proto my_new_layout.textpb -firmware_bin_path /path/to/firmware.bin
```

where `/path/to/firmware.bin` varies depending on OS and where the keyboard gets
mounted. For example, on linux, I usually see this file mounted under
`/media/$USER/CRP\ DISABLD/firmware.bin`. On OSX, this is
`/Volumes/CRP\ DISABLD/firmware.bin`. There's a lot of variability on this,
though, so don't just cut and paste. After plugging in the keyboard, you need to
find this yourself by running e.g. `mount` and seeing what looks promising.

**Important:** after running, don't forget to *unmount/eject* the keyboard!

Now put the keyboard back into normal (non-programmable) mode (again see the
video for how this is done). Try your new layout!

## [Building](#building)

This section is only for those wishing to hack on this. It's NOT required for
running (see *Getting Started*).

### Prerequisites for building

You need:

* https://golang.org/doc/install
* https://bazel.build/ - for building

### Building the firmware programmer

```
bazel build //firmware:myfalcon
```

Your binaries end up in `bazel-bin`.

### Running the tests

```
bazel test ...
```

## Built With

* https://github.com/bazelbuild/rules_go - golang build rules
* https://developers.google.com/protocol-buffers/docs/proto3 - for serializing firmware layouts

## Contributing

https://www.contributor-covenant.org/version/1/4/code-of-conduct.html

Pull requests welcome. 

## Versioning

http://semver.org/

## Deployment

Happens through travis-ci. See `.travis.yml`

https://travis-ci.org/dshahbaz/myfalcon

## Authors

* **Dimi Shahbaz** - *Initial work* - [dshahbaz](https://github.com/dshahbaz)

## License

[LICENSE](LICENSE)

## Acknowledgements

* Mateusz Perlak, initial efforts into this idea (http://blog.mateusz.perlak.com/index.php/2016/12/05/max-falcon-8-keyboard-hacking/)
* Ergodox Firmware, USB HID constants (https://github.com/benblazak/ergodox-firmware)
* The fine folks at http://www.maxkeyboard.com!

