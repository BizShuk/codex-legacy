#include <iostream>
#include <iomanip>
#include <string>
#include <sstream>

using namespace std;

void setw(char *x,char w[3]) {

};

string month_encode(char w) {
    string s;
    switch (w) {
        case 'A': s = "Jan"; break;
        case 'B': s = "Feb"; break;
        case 'C': s = "Mat"; break;
        case 'D': s = "Apr"; break;
        case 'E': s = "May"; break;
        case 'F': s = "Jun"; break;
        case 'G': s = "Jul"; break;
        case 'H': s = "Aug"; break;
        case 'I': s = "Sep"; break;
        case 'J': s = "Oct"; break;
        case 'K': s = "Nov"; break;
        case 'L': s = "Dec"; break;
    }
    return s;
}


char digit_encode(char w) {
    char s;
    switch (w) {
        case 'Q': s='1'; break;
        case 'R': s='2'; break;
        case 'S': s='3'; break;
        case 'T': s='4'; break;
        case 'U': s='5'; break;
        case 'V': s='6'; break;
        case 'W': s='7'; break;
        case 'X': s='8'; break;
        case 'Y': s='9'; break;
        case 'Z': s='0'; break;
    }
    return s;
}

int year_encode(char w) {
    int ret;
    switch (w) {
        case 'A': ret = 1995+1;   break;
        case 'B': ret = 1995+2;   break;
        case 'C': ret = 1995+3;   break;
        case 'D': ret = 1995+4;   break;
        case 'E': ret = 1995+5;   break;
        case 'F': ret = 1995+6;   break;
        case 'G': ret = 1995+7;   break;
        case 'H': ret = 1995+8;   break;
        case 'I': ret = 1995+9;   break;
        case 'J': ret = 1995+10;  break;
        case 'K': ret = 1995+11;  break;
        case 'L': ret = 1995+12;  break;
        case 'M': ret = 1995+13;  break;
        case 'N': ret = 1995+14;  break;
        case 'O': ret = 1995+15;  break;
        case 'P': ret = 1995+16;  break;
        case 'Q': ret = 1995+17;  break;
        case 'R': ret = 1995+18;  break;
        case 'S': ret = 1995+19;  break;
        case 'T': ret = 1995+20;  break;
        case 'U': ret = 1995+21;  break;
        case 'V': ret = 1995+22;  break;
        case 'W': ret = 1995+23;  break;
        case 'X': ret = 1995+24;  break;
        case 'Y': ret = 1995+25;  break;
        case 'Z': ret = 1995+26;  break;
    }

    return ret;
}


string decode_date(char s[4]) {
    string month = month_encode(s[0]);
    char c1 = digit_encode(s[1]);
    char c2 = digit_encode(s[2]);
    int year = year_encode(s[3]);

    stringstream buffer;
    buffer << setw(4) << month << "," << setw(2) << c1 << c2  << setw(5) << year << endl;

    return buffer.str();
}

int main(int argc, char *argv[]) {
    char s[5] = "ARZM";

    string ret ;
    ret = decode_date(s);
    cout << ret ;
    return 0;
}
