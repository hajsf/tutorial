#include  <stdio.h>
int Add(int a, int b){
    printf("Welcome from external C function\n");
    return a+b;
}

// clang -shared -fpic main.c -o libhello.dylib