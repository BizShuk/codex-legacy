#include <iostream>
using namespace std;

int main(int argc, char *argv[])
{
    int a;
    int b[5];
    char c;
    char d[5];
    float e;
    float f[5];
    double g;
    double h[5];


    cout << "int      sizeof: " << sizeof(a) << endl;
    cout << "int[]    sizeof: " << sizeof(b) << endl;
    cout << "char     sizeof: " << sizeof(c) << endl;
    cout << "char[]   sizeof: " << sizeof(d) << endl;
    cout << "float    sizeof: " << sizeof(e) << endl;
    cout << "float[]  sizeof: " << sizeof(f) << endl;
    cout << "double   sizeof: " << sizeof(g) << endl;
    cout << "double[] sizeof: " << sizeof(h) << endl;


    return 0;
}
