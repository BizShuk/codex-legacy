# C_CPP

### resource
- [Programming in C](http://www-ee.eng.hawaii.edu/Courses/EE150/Book/book.html)
- [CPPUnit](http://cppunit.sourceforge.net/doc/cvs/cppunit_cookbook.html)

### Practices and Exercises

Concept:
- [Pointer.cpp](pointer.cpp)
- [Sizeof](sizeof.cpp)
- [Class](class.cpp)
- [Reference](reference.cpp)
- [Garbage Collection](gc.cpp)
- [generic](generic.cpp)

Google for Education C++:
- Getting Started
    - [Example #1](helloworld.cpp)
    - [Example #2](get_input.cpp)
    - [Example #3](multiplication_table.cpp)
    - [Example #4](guess_game.cpp)
    - [Example #5](math_puzzles.cpp)
    - [Example #6](string_output.cpp)
    - [Example #7](best_sale.cpp)
    - [Example #8](scope_output.cpp)
    - [Example #9](file_io.cpp)
    - [Exercise #1](cricket_chirps.cpp)
    - [Exercise #2](exam.cpp)
    - [Exercise #3](sec_format.cpp)
    - [Exercise #4](char_bigsize.cpp)
    - [Exercise #5](reverse_num.cpp)
    - [Exercise #6](date_encode.cpp)
    - [Exercise #7](divide_magic.cpp)
- Next Steps
    - [Example #1](shooting_game.cpp)
    - [Example #2](pointer.cpp) function pointer1
    - [Example #3](vehicle_speed.cpp)
    - [Example #4](class.cpp) 
    - [Exercise #1](perfect_square.cpp)
    - [Exercise #2](order_predict.cpp)
    - [Project Database](Database.cpp) . Database.h Database.cpp Composer.h Composer.cpp Database_main.cpp Database_main.make
- C++ In Depth
    - [Exercise #1](pointer.cpp) function pointer2
    - Exercise #2 , check Next Steps **Exercises**
    - [Exercise #3](multi-dimensional_arrays.cpp) 
    - [Exercise #4](oo)
    - [CPPUnit Example](TestComplexNumber.cpp)
    - CPPUnit for Database Composer exercise
    - [Project article query](article_query.cpp)

- [reverse bits](reverseBits.cpp)





### build
	
##### build
`gcc`:  
- shared object , `gcc -shared -o libhello.so -fPIC hello.c`



g++:  
- `-c [-o <obj file>]`, compile only. Without `-c` , it'll compile and link together. 
- `-I <program path>` ,   --dynamic-linker 
- `-g` ,    generate debugging information
- `-Wall` , enable most warning messages
- `-shared -o xxxx.so `
- `-L <lib path>`
- shared object , `gcc -shared -o libhello.so -fPIC hello.c`
- `-l<package>` , after files which to be compiled


# C

### Sizeof
[Example](sizeof.cpp)
sizeof , size of bytes  check each type of bytes at sizeof.cpp
- int 4 bytes
- char 1 bytes
- float 4 bytes
- double 8 bytes

use sizeof to know memory size(byte) , but there is no way to know a pointer array size.



### memory
swap continuous memory => segmentation fault
```
    int a[10];
    swap_address(a[1],a[5]);
```

### pointer
[Example](pointer.cpp)  
[pointer mapping](http://stackoverflow.com/questions/3920729/in-c-c-is-char-arrayname-a-pointer-to-a-pointer-to-a-pointer-or-a-pointe?answertab=active#3925968)


# CPP

### CPP11

##### smart porinter
- unique_ptr , only one pointer.
- sharedd_ptr , after no references , it'll be delted
- weak_ptr , no reference count , maybe deleted by other.

##### Lambda expression
```cpp
int main()
{
    char s[]="Hello World!";
    int Uppercase = 0; //modified by the lambda
    for_each(s, s+sizeof(s), [&Uppercase] (char c) {
            if (isupper(c))
            Uppercase++;
            });
    cout<< Uppercase<<" uppercase letters in: "<< s<<endl;
}
```

##### rValue
move sematics

##### Auto type
- `auto i = f()`
- `for_each( auto const& i : ints)`

##### type long long
`auto national_debt=14400000000000LL;//long long`
`for(auto it=vi.begin(); it != vi.end(); ++it)`

##### Uniform initialization notation
- `int[] a(10,2)` , initial 10 elements with 2 value.
- `int[] a{10,2}` , initail 2 elements with 10 and 2 value.
- initial class data  
```
class C {
    private:
        int a = 7;
}
```

##### Deleted and Defaulted Functions
```
Class A
{
    A()=default; //C++11
    virtual ~A()=default; //C++11
};
```

##### nullptr

##### Delegating Constructors
```
class M //C++11 delegating constructors
{
    int x, y;
    char *p;
    public:
    M(int v) : x(v), y(0), p(new char [MAX]) {} //#1 target
    M(): M(0) {cout<<"delegating ctor"<<endl;} //#2 delegating
};
```

##### \_\_thread_local storage class

##### Control and query of object alignment

##### Static assertions

##### Variadic templates



### garbege collection
If it use **malloc** or **new** , it need to be **free()** or **delete**. And check [smart pointer](#smart-pointer)

##### new
If new is assigned to pointer or reference. After end of the scope , pointer and reference will be recycled, but not the new object.
In other situation , if you reassign pointer to a usual variable , it'll be recycled after end of the scope.

##### malloc
?



### Class
check example [here](class.cpp)

**this** is a pointer.

##### Different between new and no-new
- `Point p1 = Point(0,0)` , collect when scope is gone
- `Point p1 = new Point(0,0)` , need to delete p1

##### Differenet between . and ->
`*a.b` = `a->b`



### function parameter 

##### pass by value
```
void <func_name>(int p)
```

##### pass by reference
```
void <func_name>(int &p)
```
treat this like alias.


##### pass by address
```
void <func_name>(int *p)
```





##### virtual function
	ex: virtual void function_name(){};

	操作將延後到執行階段決定
		Wolf wolf = new Wolf();
		wolf.eat();		// wolf eat
		Animal* animal = wolf;
		animal.eat();	// animal eat

##### pure virtual void function
	ex: pure virtual void function_name()=0;

	child class 必定需要 overide

##### abstract class
	含有 pure virtual function 即稱為 , 如果用此class new instance會發生編譯錯誤


Casting example:
	ref. http://stackoverflow.com/questions/5313322/c-cast-to-derived-class
	class Animal { /* Some virtual members */ }
	class Dog: public Animal {};
	class Cat: public Animal {};

	Dog     dog;
	Cat     cat;
	// Notice no cast required. (Dogs and cats are animals).
	Animal& AnimalRef1 = dog;
	Animal& AnimalRef2 = cat;
	Animal* AnimalPtr1 = &dog;
	Animal* AnimlaPtr2 = &cat;

	// Throws an exception  AnimalRef1 is a dog
	Cat&    catRef1 = dynamic_cast<Cat&>(AnimalRef1);
	// Returns NULL         AnimalPtr1 is a dog
	Cat*    catPtr1 = dynamic_cast<Cat*>(AnimalPtr1);
	Cat&    catRef2 = dynamic_cast<Cat&>(AnimalRed2);  // Works
	Cat*    catPtr2 = dynamic_cast<Cat*>(AnimalPtr2);  // Works

	// This on the other hand makes no sense
	// An animal object is not a cat. Therefore it can not be treated like a Cat.
	Animal  a;
	// Throws an exception  Its not a CAT
	Cat&    catRef1 = dynamic_cast<Cat&>(a);
	// Returns NULL         Its not a CAT.	
	Cat*    catPtr1 = dynamic_cast<Cat*>(&a);


##### public private protected 差異

|作用域　 |當前類|同一package|子孫類|其他package|
|---------|------|-----------|------|-----------|
|public 　|　√　|　　√　　 |　√　|　　√     |
|protected|  √　|　　√　　 |　√　|　　×     |
|friendly |　√　|　　√　　 |　×　|　  ×     |
|private　|　√　|　　×　　 |　×　|　  ×     |
 


### CppUnit
- [Official site](http://cppunit.sourceforge.net/doc/cvs/cppunit_cookbook.html)  
- [Example](TestComplexNumber.cpp)

1. include
```
#include <cppunit/extensions/HelperMacros.h> 
```
2. create test class and extend public CPPUNIT_NS::TestFixture  
`class TestComplexNumber : public CPPUNIT_NS::TestFixture`
3. 




### Basic Assert
CPPUNIT_ASSERT:
```
    CPPUNIT_ASSERT( Complex (10, 1) == Complex (10, 1) );
```

##### Fixture
- setUp() ,  intiialize the variables
- tearDown() , release any permanent resources you allocated in setUp()

