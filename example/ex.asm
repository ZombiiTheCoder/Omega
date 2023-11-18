;; Asm Format, Windows, x86_64, Intel
;; Assember NASM_FELF64
;; Linker GCC_LD

section .data
    hello_msg db "Hello World", 0   ; Null-terminated string for printing

section .text
    global _start

    ; Function: Hello_World2
Hello_World2:
    ; Print "Hello World"
    mov eax, 4            ; System call number for sys_write
    mov ebx, 1            ; File descriptor 1 is stdout
    mov ecx, hello_msg    ; Pointer to the message
    mov edx, 11           ; Message length
    int 0x80              ; Interrupt to invoke the kernel

    ret                   ; Return from the function

    ; Function: Hello_World
Hello_World:
    ; Print "Hello World" from Hello_World2
    call Hello_World2

    ; Print "Hello World" again
    mov eax, 4
    mov ebx, 1
    mov ecx, hello_msg
    mov edx, 11
    int 0x80

    ; Exit the program
    mov eax, 1            ; System call number for sys_exit
    xor ebx, ebx          ; Exit code 0
    int 0x80              ; Interrupt to invoke the kernel

section .bss
    ; (Optional) Define a BSS section for uninitialized data