# stolen from https://rylev.github.io/DMG-01/public/book/
## this one seems cool to learn https://www.youtube.com/user/eaterbc
## cool docs (is incomplete)
    https://gekkio.fi/files/gb-docs/gbctr.pdf
## cool docs 2 (this ones is complete but uses times new roman)
    http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf
## operation table table z80 so take it with a grain of salt
    https://clrhome.org/table/
## z80 docs
    http://z80-heaven.wikidot.com/instructions-set:XX
## c++ impl
    https://github.com/jgilchrist/gbemu/

how to open windows of something in gdlv

window *name*

literally

window breakpoints

opens the window of breakpoints

The contents of B are rotated left one bit position. The contents of bit 7 are
copied to the carry flag and the previous contents of the carry flag are copied
to bit 0.

# idea for simply testing lang

```asm
SLA A 0b10000000
EXP A 110000000
EXP F.C true
```
