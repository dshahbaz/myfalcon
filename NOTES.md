# Running

`bazel run :main -- -text_proto $PWD/test.textpb -firmware_bin_path /media/$USER/CRP\ DISABLD/firmware.bin`

# Comparing hex dumps

`dhex -cd 1  user_modify.bin  program_modify.bin`

F3 next
F4 prev
F10 quit

# Debugging keystrokes

`sudo showkey -a`

