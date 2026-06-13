#include <iostream>

using namespace std;


int main(int argc, char *argv[])
{
    



    unsigned long int x=2,square_tmp,y=1,plus_tmp=0;
    char keepgo='y';


    while (keepgo!='n') {
        square_tmp  = x*x;

        
        while (plus_tmp < square_tmp) {
            plus_tmp += y;
            y++;
        }
    
    
    
        if (square_tmp < plus_tmp) {
            cout << square_tmp << " is not a perfect square number."<< endl;
        }
        if (square_tmp == plus_tmp) {
            cout << square_tmp << " is a perfect square number which is " << x << endl;
            cout << "Do you want to stop now(Y/N)? "; cin >> keepgo;
        }
        x++;
    
    }











    return 0;
}
