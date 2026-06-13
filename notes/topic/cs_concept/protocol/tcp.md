# TCP

### Three-way handshake

1. SYN, SEQ=[client random number]
2. SYN+ACK, SEQ=[server random number] ACK=[client random number] + 1
3. ACK, SEQ=[client random number]+1 ACK=[server random number] + 1
