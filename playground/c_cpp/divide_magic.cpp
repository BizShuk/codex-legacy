#include <iostream>
#include <iomanip>
using namespace std;

/*
 * This is a number analogy to a famous card trick. Ask the user to enter a three-digit number. Think of the number as ABC (where A, B, C are the three digits of the number). Now, find the remainders of the numbers formed by ABC, BCA, and CAB when divided by 11. We will call these remainders X, Y, Z. Add them up as X+Y, Y+Z, Z+X. If any of the sums are odd, increase or decrease it by 11 (whichever operation results in a positive number less than 20; note if the sum is 9, just report this and stop the process). Finally, divide each of the sums in half. The resulting digits are A, B, C. Write a program that implements this algorithm.
 *
 */



int main(int argc, char *argv[]) {
    int d3[3] = {9,5,3};

    cout << "Three digits: " 
        << setw(2) << d3[0] << "," 
        << setw(2) << d3[1] << ","
        << setw(2) << d3[2] << endl;

    int newnum[3],i;

    for (i = 0; i < 3; ++i) {
        newnum[i] = d3[i]*100 + d3[(i+1)%3]*10 + d3[(i+2)%3];
        newnum[i] %= 11;
        cout << newnum[i] << endl;
    }

    for (i = 0; i < 3; ++i) {
        int _tmp = newnum[i] + newnum[(i+1)%3];

        if (_tmp%2==1) {
            if (_tmp+11<20) {
                _tmp += 11;
            }else{
                _tmp -= 11;
            }
        }

        cout << _tmp/2 << endl;
    }


    return 0;
}



