CC=gcc
TARGET=run
SRC=run.c

default : 32

32 : $(SRC)
	$(CC) -m32 -o $(TARGET) $(SRC)

64 : $(SRC)
	$(CC) -m64 -o $(TARGET) $(SRC)

clean:
	rm $(TARGET)

