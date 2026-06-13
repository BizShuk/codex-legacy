# Fork

Clone a process to execute

When fork is executed, it copies the caller's memory (code, globals, heap and stack), registers, and open files.
Linux implement Copy-on-Write strategy. As long as no write operation, the memory keeps only one copy.
