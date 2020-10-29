#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

static const uint32_t MAGIC_NUMBER = 0x01234567;

typedef const char *string_t;

struct string_s {
	uint32_t magic;
	int32_t len;
	char *str;
};

string_t string_new(const char *str, int32_t len) {
	struct string_s *p;
	assert(str != NULL);
	assert(len >= 0);
	p = (struct string_s *)malloc(sizeof(*p) + len + 1);
	if (p == NULL) {
		return NULL;
	}
	p->magic = MAGIC_NUMBER;
	p->len = len;
	p->str = (char *)(p + 1);
	if (len > 0) {
		memcpy(p->str, str, len);
	}
	p->str[len] = '\0';
	return p->str;
}

void string_free(string_t s) {
}

int main() {
	const char str[] = "hello\0world";
	string_t s = string_new(str, sizeof(str));
	printf("%s\n", s);
	string_free(s);
}
