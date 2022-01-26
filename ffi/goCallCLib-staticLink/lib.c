#include  <stdio.h>
int Add(int a, int b){
    printf("Welcome from external C function\n");
    return a + b;
}

// clang -shared -fpic lib.c -o ../libraries/libadd.a
// or
// clang -shared -fpic -Wall -g lib.c -o ../libraries/libadd.so