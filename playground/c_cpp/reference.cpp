#include <iostream>

using namespace std;

void DoIt(int &foo, int goo);
void Duplicate(int &a,int &b,int &c);


int main(int argc, char *argv[]) {

    // DoIt
    int *foo, *goo;
    foo = new int;    *foo = 1;
    goo = new int;    *goo = 3;
    *foo = *goo + 3;            // *foo = 3 + 3;
    foo = goo;                  // foo = goo , point to *goo
    *goo = 5;                   // *goo = 5  => *foo =5
    *foo = *goo + *foo;         // *foo = 5 + 5

    DoIt(*foo, *goo);
    cout << *foo << endl;





    // Duplicate
    int x=1,y=3,z=7;
    Duplicate(x,y,z);
    cout << "x="<< x << ", y="<< y << ", z="<< z;



    return 0;
}

void Duplicate(int &a,int &b, int &c) {
    a *= 2;
    b *= 2;
    c *= 2;

}

void DoIt(int &foo, int goo) {
    foo = goo +3;
    goo = foo +4;
    foo = goo +3;
    goo = foo;
    goo = 3;
}
