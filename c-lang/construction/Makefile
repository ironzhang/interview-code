# Makefile

TARGET=construction

CC=g++
IDIR=-I.
CFLAGS=$(IDIR) -g -W -Wall
LFLAGS=

C_FILES=$(wildcard *.cpp)
S_FILES=$(C_FILES:.cpp=.s)
O_FILES=$(C_FILES:.cpp=.o)

.PHONY: clean assemble

all: $(TARGET) assemble

$(TARGET): $(O_FILES)
	$(CC) -o $(TARGET) $(LFLAGS) $^

assemble: $(S_FILES)

%.o: %.cpp
	$(CC) $(CFLAGS) -c $^ -o $@

%.s: %.cpp
	$(CC) $(CFLAGS) -S $^ -o $@

clean:
	rm -f $(O_FILES) $(S_FILES) $(TARGET)
