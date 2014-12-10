#CTFCODE

Collection of somewhat useful stuff for CTFs.

##Contents

###./ShellcodeTester

Simple tool to load and execute shellcode, either from stdin or from a file.
Supports x32 and x64.

Usage:

```bash
cd ShellcodeTester
make 32
./run /path/to/my/shellcode
```

> Obviously, be careful with this :)

###./ShellcodeBuilder

Basically just a Makefile and assembly templates to build custom shellcode.

Usage:

```bash
cd ShellcodeBuilder
# write your shellcode
vim shellcode32.asm
make 32
```

Afterwards go ahead and directly test your new code with the ShellcodeTester:

```bash
../ShellcodeTester/run shellcode
```

###./ExploitTemplates

These mostly take care of the networking stuff so your exploit doesn't have to, but they also provide some commonly needed funtionality, e.g. packing and unpacking of binary data.

See [here](http://kitctf.de/writeups/9447ctf2014/2014/12/01/europe-writeup/) for a usage example (the code uses an older version of the templates though) or just play around with them a bit on your own.
