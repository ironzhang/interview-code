# Makefile

CC = gcc
PROGS = string

all: $(PROGS)

string: string.c
	$(CC) ${CFLAGS} $^ -o $@

clean:
	rm -f ${PROGS}
