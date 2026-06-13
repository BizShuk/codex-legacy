#include <iostream>
using namespace std;


int main(int argc, char *argv[])
{
    std::cout << "1\t2\t3\t4\t5\t6\t7\t8\t9\t" << std::endl;

    for (int i = 1;  i < 10; ++i) {
        std::cout << i << "| ";
        for (int j = 1;  j< 10; ++j) {
            std::cout << i*j << "\t";
        }
        std::cout << std::endl;
    }







    return 0;
}
