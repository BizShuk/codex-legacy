#include <cppunit/extensions/HelperMacros.h>
#include "ComplexNumber.h"
/*
 * * A test case that is designed to produce
 * * example errors and failures
 * *
 * */
class TestComplexNumber : public CPPUNIT_NS::TestFixture
{
    //establish the test suit of TestComplexNumber
    CPPUNIT_TEST_SUITE( TestComplexNumber);
    //add test method testInit
    CPPUNIT_TEST( testInit);
    //add test method testEquals
    CPPUNIT_TEST( testEquals );
    //add test method testAdd
    CPPUNIT_TEST( testAdd );
    //finish the process
    CPPUNIT_TEST_SUITE_END();
public:
    //overide setUp(), init data etc
    void setUp();
    //overide tearDown(), free allocated memory, etc
    void tearDown();
    //test method testsample
    void runTest();
protected:
    //test method testAdd
    void testAdd();
    //test method testEquals
    void testEquals();
    //test method testInit
    void testInit();
private:
    //three instances of CompleNumber for test
    ComplexNumber *a, *b, *c;
};
