# Process vs Thread

## Process

state: new, ready, running, waiting, terminated, suspended
heavier for creation, termination, context-switch
isolated memory
communication: IPC
switch between process: interupt to kernel

## Thread

state, running, ready, blocked
lighter for creation, termination, context-switch
shared memory
communication: share memory
switch between thread: no interupt
