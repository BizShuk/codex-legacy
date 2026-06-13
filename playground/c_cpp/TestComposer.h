#include <cppunit/extensions/HelperMacros.h>
#include <Composer.h>

class TestComposer : public CPPUNIT_NS::TestFixture
{
    // establish the test suit of TestComposer
    CPPUNIT_TEST_SUITE( TestComposer);
    // add test method testInit
    CPPUNIT_TEST( testInit);
    // add test method testEquals
    CPPUNIT_TEST( testEquals );
    // add test method testAdd
    CPPUNIT_TEST( testAdd );
    // finish the process
    CPPUNIT_TEST_SUITE_END();
public:
    // overide setUp(), init data etc
    void setUp();
    // overide tearDown(), free allocated memory, etc
    void tearDown();
protected:
    // test method testAdd
    void testAdd();
    // test method testEquals
    void testEquals();
    // test method testInit
    void testInit();
private:
    // three instances of Composer for test
    Composer *a, *b, *c;
};
