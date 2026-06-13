cpp 物件導向概念

http://silverfoxkkk.pixnet.net/blog/post/43002735-cpp%3A%E4%B8%89%E3%80%81%E7%89%A9%E4%BB%B6%E5%B0%8E%E5%90%91
http://memo.cgu.edu.tw/shin-yan/1002_ComputerProgramming/Ch13.pdf

virtual function
	ex: virtual void function_name(){};

	操作將延後到執行階段決定
		Wolf wolf = new Wolf();
		wolf.eat();		// wolf eat
		Animal* animal = wolf;
		animal.eat();	// animal eat

pure virtual void function
	ex: pure virtual void function_name()=0;

	child class 必定需要 overide

abstract class
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


public private protected 差異






