#include <iostream>
#include <array>

using namespace std;



array<int,3> sec_format(int secs) {
    array<int, 3> time_f {0,0,0};

    time_f[2] = secs/3600;
    time_f[1] = secs%3600/60;
    time_f[0] = secs%60;

    return time_f;
}

int main(int argc, char *argv[]) {
    int secs = 80000;

    cout << "Enter seconds:"; cin >> secs;

    array<int, 3> time_f ;
    time_f = sec_format(secs);


    cout << "Hours:" << time_f[2] << endl;
    cout << "Minutes:" << time_f[1] << endl;
    cout << "Seconds:" << time_f[0] << endl;

    return 0;
}
