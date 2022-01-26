#include  <stdio.h>
int Add(int a, int b){
    printf("Welcome from external C function\n");
    return a + b;
}

// clang -shared -fpic lib.c -o libadd.a
// or
// clang -shared -fpic -Wall -g lib.c -o libadd.so