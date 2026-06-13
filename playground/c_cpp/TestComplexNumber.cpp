#include "TestComplexNumber.h"
#include <cppunit/extensions/TestFactoryRegistry.h>
#include <cppunit/ui/text/TestRunner.h>
#include <cppunit/TestResult.h>


// Need to register test suites
CPPUNIT_TEST_SUITE_REGISTRATION( TestComplexNumber );

void TestComplexNumber::setUp()
{
    a = new ComplexNumber(1.0, 2.0);
    b = new ComplexNumber(1.0, 2.0);
    c = new ComplexNumber(2.0, 4.0);
}
void TestComplexNumber::tearDown()
{
    delete a;
    delete b;
    delete c;
}

void TestComplexNumber::runTest()
{
    CPPUNIT_ASSERT(*a == *b);
    CPPUNIT_ASSERT(!(*a == *c));
}

void TestComplexNumber::testInit()
{
    CPPUNIT_ASSERT(a->realPart == 1.0);
}
void TestComplexNumber::testAdd()
{
    CPPUNIT_ASSERT(*a + *b == *c);
}
void TestComplexNumber::testEquals()
{
    CPPUNIT_ASSERT(*a == *b);
    CPPUNIT_ASSERT(!(*a == *c));
}

int main( int argc, char **argv)
{
    CppUnit::TextUi::TestRunner runner;
    CppUnit::TestFactoryRegistry &registry
        = CppUnit::TestFactoryRegistry::getRegistry();
    // Auto test all test suites
    runner.addTest( registry.makeTest() );
    runner.run();


    // Test Suite with runner
    CppUnit::TestSuite* suite = new CppUnit::TestSuite("testsuite");
    CppUnit::TestResult result;
    CppUnit::TextUi::TestRunner runner1;

    suite->addTest( new CppUnit::TestCaller<TestComplexNumber>( "testEquality", &TestComplexNumber::runTest ) );
    suite->run( &result );

    runner1.addTest( suite );
    runner1.run();

    return 0;
}
