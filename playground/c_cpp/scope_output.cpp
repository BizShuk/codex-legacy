#include <iostream>

using namespace std;

int a = 18;
int b = 6;

int function1(int a, int b) {
    return a - b;
}

int function2() {
    int c;
    c = a + b;      // 
    return c;
}

int main() {
    int b = 12;
    int c = 0;
    a = function1(b, a);    // override global a by func1
    c = function2();        // in func2 , because last statement , a=-6
    cout << "a: " << a << " b: " << b << " c: " << c << endl;
}
