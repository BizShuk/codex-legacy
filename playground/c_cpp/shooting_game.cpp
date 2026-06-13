#include <iostream>
#include <cmath>
using namespace std;

void startup() {
    cout << "Welcome to Artillery." << endl;
    cout << "You are in the middle of a war and being charged by thousands of enemies." << endl;
    cout << "You have one cannon, which you can shoot at any angle." << endl;
    cout << "You only have 10 cannonballs for this target.." << endl;
    cout << "Let's begin..." << endl;
}


int fly_distance(float angel) {
    float velocity = 200.0; 
    float gravity = 32.2;
    
    float time_in_air = (2.0 * velocity * sin(angel)) / gravity;
    float distance = round((velocity * cos(angel)) * time_in_air);

    return (int) round(distance);
}


int new_enemy_distance(){
    srand(time(NULL));

    int _dis = rand() % 900 + 200;
    cout << "the enemy is at " << _dis << endl;

    return _dis;
}


int fire() {
    float angel=0;
    int shot_count=0,hit=0,fly_dis=0;
    float enemy_distance = 0;
    
    enemy_distance = new_enemy_distance();

    do{
        cout << "What angle?"; cin >> angel;
        if (cin.fail()) {
            cin.clear();
            cin.ignore(100,'\n');
            continue;
        }
        angel = (angel * 3.1415)/180.0;
        fly_dis = fly_distance(angel);
        shot_count++;

        

        if ( abs(enemy_distance-fly_dis) > 1 ) {
            cout << "you over shot by " << abs(enemy_distance - fly_dis) << endl;
        }else{
            hit = 1;
        }
    
    } while ( hit == 0);
    cout << "You hit hime !!!" << endl;
    cout << "It tooke you " << shot_count << "shits" << endl;
    
    return 1;
}



int main(int argc, char *argv[]) {
    char done;
    startup();
    int killed = 0;
    do {
        killed += fire();
        cout << "You have killed " << killed << "enemies." << endl;
        cout << "I see another one, care to shoot again? (Y/N) " << endl;
        cin >> done;
    } while (done != 'n');

    return 0;
}




