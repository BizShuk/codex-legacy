.PHONY: clean, mrproper
CC = g++
CFLAGS = -g 

all: clean TestComplexNumber

TestComplexNumber: TestComplexNumber.h TestComplexNumber.cpp
	$(CC) $(CFLAGS) $+ -lcppunit -o $@

clean:
	rm -f *.o

mrproper: clean
	rm -f TestComplexNumber
