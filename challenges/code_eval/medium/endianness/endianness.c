#include "stdio.h"

// Endianness
// https://www.codeeval.com/open_challenges/15/
int main() {
    int n = 1;

    if (*(char *)&n == 1) {
        printf("LittleEndian\n");
    } else {
        printf("BigEndian\n");
    }
}
