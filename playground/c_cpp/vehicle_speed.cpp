#include <iostream>
using namespace std;


void accerate(int& speed, int amount) {
    speed += amount;
}


int main(int argc, char *argv[]) {
    int current_speed = 30;
    int accel_amount = 0;

    
    cout << "current speed: " << current_speed << endl;


    do{
        cout << "Accelerate by how much?";
        if ( !(cin >> accel_amount)) {
            cout << "Not a speed number." << endl;
            cin.clear();
            cin.ignore(100,'\n');
            continue;
        }
        accerate(current_speed,accel_amount);
        cout << "current speed :" << current_speed << endl;

    } while( current_speed != 0 );

    cout << "Current speed is " << current_speed << " , and we are freezed." << endl;

    return 0;
}
