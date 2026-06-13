.PHONY: clean, mrproper
CC = g++
CFLAGS = -g

all: clean Composer.o Database.o Database.out

Composer.o: Composer.cpp
	$(CC) $(CFLAGS) -c -o $@ $+

Database.o: Database.cpp
	$(CC) $(CFLAGS) -c -o $@ $+

Database.out: Composer.o Database.o Database_main.cpp
	$(CC) $(CFLAGS) -o $@ $+

clean:
	rm -f *.o 

mrproper: clean
	rm -f Database.out

