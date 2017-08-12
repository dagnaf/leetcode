TARGET=main
CPPFLAGS+=-std=c++11 -I${problem_name}

${TARGET}: ${TARGET}.o
	${CXX} $^ -o $@

.PHONY: clean
clean:
	rm -rf *.o ${TARGET}
