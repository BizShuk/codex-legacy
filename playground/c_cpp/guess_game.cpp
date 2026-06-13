#include <iostream>
#include <time.h>
using namespace std;

int main(int argc, char *argv[]) {
    
    int max = 100;
    int rand_num,guess;
    char getright = '<';
    char 
    srand (time(NULL));
    
    rand_num = rand() % max + 1;
    
    std::cout << "Start to guessing" << std::endl;
    while ( getright != '=' ){
        std::cout << "Enter a number:";
        if (!(cin>>guess)) {
            std::cout << "wrong guess type" << std::endl;
            cin.clear();
            cin.ignore(100,'\n');
        }else{
            
            if ( guess == rand_num ) {
                getright = '=';
            }else if ( guess < rand_num){
                getright = '<';
            }else if ( guess > rand_num){
                getright = '>';
            }

            std::cout << "Answer:" << guess << " " << getright << "real answer"  << std::endl;

        }
    
    }

    std::cout << "Mission completed" << std::endl;
    return 0;
}
