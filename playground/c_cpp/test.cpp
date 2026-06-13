#include <iostream>
#include <string.h>
#include <list>
using namespace std;

void test1()
{
    char b = 'a';
    char *a = &b;
    char w[] = "aaaaaa";
    char *x = w;
    cout << "x:" << strlen(x) << endl;

    cout << sizeof(*a) << endl;

    char d[] = "eeee";

    char *c[] = {d};

    cout << c[0] << strlen(c[0]) << endl;
}


void test2()
{
    int a;
    a++;
    cout << a << endl;
}


void continue_memory_array()
{
    int a[] = {1, 2, 3, 4, 5, 6, 7, 8};
    /*
        int* tmp;
        int* tmp1 = &a[5];
        int* tmp2 = &a[3];

        *tmp = *tmp1;
        *tmp1 = *tmp2;
        *tmp2 = *tmp;
    */


    int tmp;
    tmp = a[5];
    a[3] = a[5];
    a[5] = tmp;


    for (int i = 0; i < 8; ++i)
    {
        cout << a[i] << endl;
    }

}

class A
{
public:
    A ();
    ~A ();
    int num;
};

A::A()
{
}
A::~A()
{
}

A test_scope()
{
    A a;
    a.num = 4;
    return a;
}

int main(int argc, char *argv[])
{
    // test1();
    // test2();
    //continue_memory_array();

    A *b;
    if (!b)
    {
        cout << "123" << endl;
    }
    else
    {
        cout << "456";
    }


    A c = test_scope();

    cout << "t" << c.num << endl;


    list<int> _list;
    list<int>::iterator _list_it = _list.begin();

    cout << "begin:" << *_list_it << endl;
    _list_it++;
    cout << "begin++:" << *_list_it << endl;
    _list_it++;
    cout << "begin++:" << *_list_it << endl;


    return 0;
}
