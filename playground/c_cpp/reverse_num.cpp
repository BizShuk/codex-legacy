#include <iostream>

using namespace std;


int reverse_num(int num) {
    int ret = 0;

    while (num>0) {
        ret = ret*10 + num%10;

        num = num / 10;
    }


    return ret;
}


int main(int argc, char *argv[]) {

    int input,rev,rev2;
    cout << "input number : "; cin >> input;
    rev = reverse_num(input);
    cout << "reverse it : " << rev << endl;
    cout << "subtract " << input << " - " << rev << " = " << input-rev << endl;
    rev2 = reverse_num(input-rev);
    cout << "reverse it : " << rev2 << endl;
    cout << "add: " << rev2 << "+" << rev-input << " = " << rev2 + input - rev << endl;



    return 0;
}



