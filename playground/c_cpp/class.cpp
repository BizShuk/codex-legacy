#include <iostream>
using namespace std;


class Time {
public:
    Time (int hour, int minute, int second);
    virtual ~Time ();
    void display();

private:
    int hour_;
    int minute_;
    int second_;
};


Time::~Time() {
    cout << "decontructor:" << endl;   
    this->display();
}

Time::Time(int hour, int minute, int second) {
    hour_ = hour;
    minute_ = minute;
    second_ = second;
}

void Time::display() {
    cout << (hour_ % 12 ? hour_ %12 : 12 ) << ":"
        << (minute_ < 10 ? "0" : "" ) << minute_ << ":"
        << (second_ < 10 ? "0" : "" ) << second_ << " "
        << (hour_ < 12  ? "AM" : "PM" ) 
        << endl;
}



Time returnTime1() {
    Time a(23,30,29);           // a default constructor , same as new
    return a;
}

Time& returnTime2(int h) {
    Time *b = new Time(h,10,20);
    return *b;
}

Time* returnTime3() {
    Time* b = new Time(2,11,22);
    return b;
}

Time* returnTime4() {
    Time c = Time(2,15,22);
    Time* b = &c;
    return b;
}


int main(int argc, char *argv[]) {
    // this can't assign to  pointer(*a) ( *a = &returnTime())  , it'll cause segmentation fault
    cout << endl;
    cout << "a:" << endl;
    Time a = returnTime1();
    a.display();

    // mem leak , will not delete after scope , pointer(*b) will be delete by end of scope , but not new
    cout << endl;
    cout << "b:" << endl;
    Time *b = &returnTime2(1);
    b->display();


    // no mem leak , will be delete , even use new as constructor
    cout << endl;
    cout << "b1:" << endl;
    Time b1 = returnTime2(2);
    b1.display();
    Time* b2 = &b1;
    b2->display();

    cout << endl;
    cout << "c:" << endl;
    Time *c = returnTime3();
    c->display();
    Time c1 = *c;
    c1.display();
    


    cout << endl;
    cout << "d:" << endl;
    Time* d = new Time(1,2,3);
    d->display();
    Time& d1 =  *d;
    d1.display();
    Time d2 = *d;
    d2.display();



    cout << endl;
    cout << "e:" << endl;
    Time* e = returnTime4();
    e->display();       // will get wrong address



    cout << endl << "start recycle:" << endl;
    return 0;
}

















