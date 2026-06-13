#include <iostream>
using namespace std;

void Duplicate(int *a, int *b, int *c) {
    *a *= 2;
    *b *= 2;
    *c *= 2;
}

void Unknown(int *p, int num);
void HardToFollow(int *p, int q, int *num);

void pointer1() {
    int x = 1, y = 3, z = 7;
    Duplicate(&x, &y, &z);
    // The following outputs: x=2, y=6, z=14.
    cout << "x=" << x << ", y=" << y << ", z=" << z << endl;
}

void pointer2() {
    int *q;
    int trouble[3];

    trouble[0] = 1;
    q = &trouble[1];
    *q = 2;
    trouble[2] = 3;

    HardToFollow(q, trouble[0], &trouble[2]);
    Unknown(&trouble[0], *q);
    cout << *q << " " << trouble[0] << " " << trouble[2] << endl;
}


int main() {

    // 
    pointer1();

    // google c++ education "C++ In Depth"
    pointer2();

    return 0;
}




void Unknown(int *p, int num) {
    int *q;

    q = &num;
    *p = *q + 2;
    num = 7;
}


void HardToFollow(int *p, int q, int *num) {
    *p = q + *num;
    *num = q;
    num = p;
    p = &q;
    Unknown(num, *p);
}
