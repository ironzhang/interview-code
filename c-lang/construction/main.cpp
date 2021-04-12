#include <stdio.h>

class test {
public:
	test()
		:m_a(0)
	{}

	void print() {
		printf("m_a=%d, m_b=%d, m_c=%p, m_d=%p, m_e=%p\n", m_a, m_b, m_c, m_d, m_e);
	}

private:
	int m_a;
	int m_b;
	int *m_c;
	int *m_d;
	int *m_e;
};

int main() {
	test t;
	t.print();
	return 0;
}
