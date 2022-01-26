#include <stdio.h>
#include "libadd.h"
#include "libgo.h"

int main()
{
    HelloWorld();
    int x = Add(1, 2);
    printf("%d",x);
    return 0;
}

// clang -o main -L. -ladd -lgo main.c   // -ladd => -l (library) is libadd
// ./hello  // execute the hellow output