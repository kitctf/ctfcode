/*
 * run.c - test shellcode.
 *
 * (c) 2014 Samuel Groß
 */
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/mman.h>

/* should be enough to hold your shellcode, if not just set this to a higher value */
#define BUFSIZE 4096
/* set to 1 to enable debugging, will break before executing the shellcode */
#define DEBUGGING 0

/* either paste your shellcode in here ... */
char shellcode[] = "\x31\xc0\x50\x68\x2f\x2f\x73\x68\x68\x2f\x62\x69"
                   "\x6e\x89\xe3\x50\x53\x89\xe1\xb0\x0b\xcd\x80";

int main(int argc, char* argv[])
{
    size_t len;
    char *buf, *ptr;

    printf("[*] allocating executable memory...\n");
    buf = mmap(NULL, BUFSIZE, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_ANONYMOUS | MAP_PRIVATE, 0, 0);
    ptr = buf;
    printf("[+] buffer @ %p\n", buf);

#if DEBUGGING
    ptr[0] = '\xcc';
    ptr++;
#endif

    /* ... or pass it as filename to the program */
    if (argc > 1) {
        printf("[*] reading shellcode from file...\n");
        FILE *f = fopen(argv[1], "r");
        if (!f) {
            fprintf(stderr, "[-] Cannot open %s: %s\n", argv[1], strerror(errno));
            exit(-1);
        }
        len = fread(ptr, 1, BUFSIZE, f);

        fclose(f);
    } else {
        len = sizeof(shellcode);
        printf("[*] copying shellcode...\n");
        memcpy(ptr, shellcode, len);
    }
    printf("[+] done, size of shellcode: %zi bytes\n", len);

    printf("[*] jumping into shellcode...\n\n");
    (*(void (*)()) buf)();

    return 0;
}
