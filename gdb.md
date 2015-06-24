GDB Basics
==========

If you have the source, compile with -g to enable debug information, gdb then has access to the source code.

### Show register content

    i r
    (info registers)

### Start execution

    r
    r < input_file
    r -x -y -z blabla

### Set a breakpoint

    b main          (needs symbols)
    b *0x08041234

### Continue execution

    c

### Single Step

    ni/nexti            (treats a 'call' as one instruction)
    si/stepi            (also steps into the callee)

    next                (needs -g, next source line)
    step                (needs -g, next source line)

### Display memory

    x/32wx $esp
    x/16gx 0x08049876
    x/f $eax            (float)
    x/20i main          (disassemble)

#### Sizes

    b   byte
    h   16bit word
    w   32bit word
    g   64bit word

### Show call stack

    bt
    (backtrace)

### Show process memory layout

    info proc mappings

### Execute shell commands

    !ls

### Show source code of function (needs -g)

    list main

### Display current instruction on each break/step

    display/i $eip
