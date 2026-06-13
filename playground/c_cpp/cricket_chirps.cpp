#include <iostream>

using namespace std;

float cricket_temperature(int chirps) {
    float temp = ( (float) (chirps+40) / 4 );
    return temp;
}

int main(int argc, char *argv[]) {
    
    int chirps = 134;
    float temp;

    temp = cricket_temperature(chirps);
    cout << temp << endl;
    return 0;
}


