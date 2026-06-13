
繼承執行順序
    public class A {
            {
                System.out.println("Quote A");
            }
             
            static public void printStatic(){
                System.out.println("Static method A");
            }
            public void m1(){
                System.out.println("method A");
            }
            public A(){
                System.out.println("Constructor A");
            }
    }

    public class B extends A{
        {
            System.out.println("Quote B");
        }
        static public void printStatic(){
            System.out.println("Static method B");
        }
        public void m1(){
            System.out.println("method B");
        }
        public B(){
            System.out.println("Constructor B");
        }
    }

    public static void main(String[] args) {
        A test_inheritance1 = new B();
        test_inheritance1.printStatic(); 
        test_inheritance1.m1();
        
        B test_inheritance2 = new B();
        test_inheritance2.printStatic();
        test_inheritance2.m1();
    }

    Output => 
        "Quote A"
        "Constructor A"
        "Quote B"
        "Constructor B"
        "Static method A"
        "method B"

        "Quote A"
        "Constructor A"
        "Quote B"
        "Constructor B"
        "Static method B"       
        "method B"              

        method vs static method 執行順序
            method by new class
            static method by used class

        constructor 執行順序
            1. 父類別 in class  
            2. 父類別 Constructor
            3. 子類別 in class  
            4. 子類別 Constructor

多型 polymorphism
    A test= B();
    test.m1();

函式的多型,virtual,try and catch and finally的使用,

Thread
Runnable

Difference public, private , and protected

try catch finally的使用


Singleton 模式
    Singleton 模式可以保證一個類別只有一個實例，並提供一個訪問（visit）這個實例的方法。

    public class Singleton {
        private static Singleton instance = null;
        private Singleton(){}
        public static Singleton getInstance() {
            if (instance == null){
                synchronized(Singleton.class){          // 將 synchronized放在 getInstance() 上的話 會造成每次access都會效能低落  這樣的話 只有第一次init時會比較慢
                    if(instance == null) {
                         instance = new Singleton();
                    }
                }
            }
            return instance;
        }
    }
