#include "TestComposer.h"
#include <cppunit/extensions/TestFactoryRegistry.h>
#include <cppunit/ui/text/TestRunner.h>

CPPUNIT_TEST_SUITE_REGISTRATION( TestComposer );
void TestComposer::setUp()
{
    a = new Composer();



    b = new Composer();



    c = new Composer();




}
void TestComposer::tearDown()
{
    delete a;
    delete b;
    delete c;
}
void TestComposer::testInit()
{
    CPPUNIT_ASSERT(a->realPart == 1.0);
}
void TestComposer::testAdd()
{
    CPPUNIT_ASSERT(*a + *b == *c);
}
void TestComposer::testEquals()
{
    CPPUNIT_ASSERT(*a == *b);
    CPPUNIT_ASSERT(!(*a == *c));
}



int main( int argc, char **argv)
{
    CppUnit::TextUi::TestRunner runner;
    CppUnit::TestFactoryRegistry &registry
        = CppUnit::TestFactoryRegistry::getRegistry();
    runner.addTest( registry.makeTest() );
    runner.run();
    return 0;
}
