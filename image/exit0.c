void _start() {
    asm(
        "mov $60, %%rax\n"   // syscall number for exit (60)
        "mov $0, %%rdi\n"    // exit code (0)
        "syscall\n"          // invoke syscall
        : : : "rax", "rdi"
    );
}
