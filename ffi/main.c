int Add(int a, int b){
    return a+b;
}

// clang -shared -fpic main.c -o libhello.dylib