#ifndef ComplexNumber_h
#define ComplexNumber_h 1
#include <iostream>
using namespace std;
class TestComplexNumber;
class ComplexNumber
{
    friend bool operator==(const ComplexNumber& a, const ComplexNumber& b)
    {
        return (a.realPart == b.realPart) && (a.imagePart == b.imagePart);
    }
    friend ComplexNumber operator+(const ComplexNumber& a, const ComplexNumber &b)
    {
        ComplexNumber c;
        c.realPart = a.realPart + b.realPart;
        c.imagePart = a.imagePart + b.imagePart;
        return c;
    }
public:
    ComplexNumber(double r, double i): realPart(r), imagePart(i) {}
    ComplexNumber(): realPart(0.0), imagePart(0.0) {}
private:
    double realPart, imagePart;
    friend class TestComplexNumber;
};
#endif
