#include <iostream>
#include <array>
using namespace std;


array<array<int,5>, 5> bigchar(char a) {
    array<array<int,5>, 5> ret;
    switch (a) {
        case 'A': ret =  {{{0,0,1,0,0},{0,1,0,1,0},{1,0,0,0,1},{1,1,1,1,1},{1,0,0,0,1}}}; break;
        case 'B': ret =  {{{1,1,1,1,0},{1,0,0,0,1},{1,1,1,1,0},{1,0,0,0,1},{1,1,1,1,0}}}; break;
        case 'C': ret =  {{{},{},{},{},{}}}; break;
        case 'D': ret =  {{{},{},{},{},{}}}; break;
        case 'E': ret =  {{{},{},{},{},{}}}; break;
        case 'F': ret =  {{{},{},{},{},{}}}; break;
        case 'G': ret =  {{{},{},{},{},{}}}; break;
        case 'H': ret =  {{{},{},{},{},{}}}; break;
        case 'I': ret =  {{{},{},{},{},{}}}; break;
        case 'J': ret =  {{{},{},{},{},{}}}; break;
        case 'K': ret =  {{{},{},{},{},{}}}; break;
        case 'L': ret =  {{{},{},{},{},{}}}; break;
        case 'M': ret =  {{{},{},{},{},{}}}; break;
        case 'N': ret =  {{{},{},{},{},{}}}; break;
        case 'O': ret =  {{{},{},{},{},{}}}; break;
        case 'P': ret =  {{{},{},{},{},{}}}; break;
        case 'Q': ret =  {{{},{},{},{},{}}}; break;
        case 'R': ret =  {{{},{},{},{},{}}}; break;
        case 'S': ret =  {{{},{},{},{},{}}}; break;
        case 'T': ret =  {{{},{},{},{},{}}}; break;
        case 'U': ret =  {{{},{},{},{},{}}}; break;
        case 'V': ret =  {{{},{},{},{},{}}}; break;
        case 'W': ret =  {{{},{},{},{},{}}}; break;
        case 'X': ret =  {{{},{},{},{},{}}}; break;
        case 'Y': ret =  {{{},{},{},{},{}}}; break;
        case 'Z': ret =  {{{},{},{},{},{}}}; break;
        case '\'': ret =  {{{},{},{},{},{}}}; break;
    }
    return ret;
}

void print_char(array<array<int,5>,5> bigchar) {

    for (int i = 0; i < 5; i++) {
        for (int j = 0;  j < 5; j++) {
            if (bigchar[i][j]==1) {
                cout << "*";
            }else{
                cout << " ";
            }
        }
        cout << endl;
    }
    
}

int main(int argc, char *argv[]) {
    array<array<int,5>, 5> a; 
    a = bigchar('B');
    print_char(a);





    return 0;
}
