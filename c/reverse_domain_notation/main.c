#include <stdio.h>
#include <ctype.h>
#include <string.h>

#define MAXFQDNLEN 254

void reverse_domain_notation_lower(char *fqdn, unsigned int len) {
  char tmp[MAXFQDNLEN] = {0};
  unsigned int offset = 0, numbytes = 0;
  unsigned int end = len;

  // If we see an attempt to overflow the buffer with a long fqdn
  // we truncate the length to the maximum allowed range and reverse
  if (len > MAXFQDNLEN) {
    len = MAXFQDNLEN;
  }

  for (unsigned int loopindex = len; loopindex > 0; --loopindex) {
    unsigned int index = loopindex - 1;

    switch (fqdn[index]) {
    case '.':

      numbytes = end - (index + 1);

      memcpy(tmp + offset, fqdn + index + 1, numbytes);

      end = index;
      offset += numbytes;

      tmp[offset++] = '.';

      break;
    default:
      fqdn[index] = tolower(fqdn[index]);
    }
  }

  // Top part of the fqdn might not contain a . so we
  // need to copy the rest
  if (len > offset) {
    memcpy(tmp + offset, fqdn, len - offset);
  }

  // Overwrite original content with the reverse one.
  memcpy(fqdn, tmp, len);
}

int main(int argc, char **argv) { 
  char test[] = "testing.example.com";
  reverse_domain_notation_lower(test, strlen(test));
  printf("got %s\n", test);

  return 0; }
