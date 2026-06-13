#include <iostream>
using namespace std;

int main(int argc, char *argv[])
{
    int input=0;
    do{
        std::cout << "Enter a number (-1=quit):" << std::endl;
        if (!(cin>>input)) {
            std::cout << "input is not a number..." << std::endl;
            cin.clear();
            cin.ignore(100,'\n');
            continue;
        }

        if (!cin.fail() && input!=-1 ) {
            std::cout << "input number:" << input << std::endl;
        }
    
    }while( input!=-1 );

    std::cout << "terminating..." << std::endl;



    return 0;
}
