#include <stdio.h>
#include "lib.h"
#include "libgo.h"

int main()
{
    HelloWorld();
    int x = Add(1, 2);
    printf("%d",x);
    return 0;
}

// clang -o hello -L. -ladd -lgo main.c   // -lhello => -l (library) is libhello
// ./hello  // execute the hellow output