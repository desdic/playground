#include <stdio.h>

int greet(const char *name, int year, char *out) {
    return sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);
}
