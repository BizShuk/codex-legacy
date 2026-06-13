#include <iostream>
#include <array>

using namespace std;

int package1(int WeeklySales)
{
    return 600;
}

int package2(int WeeklySales)
{
    return 7*40+225*WeeklySales/10;
}

int package3(int WeeklySales)
{
    return (225/5+20)*WeeklySales;
}
int main(int argc, char *argv[])
{
    int WeeklySales=10;
    std::cout << "Enter week sales: ";
    while(!(cin>>WeeklySales)) {
        cin.clear();
        cin.ignore(100,'\n');
    }

    array<int,3> salary;
    salary[0] = package1(WeeklySales);
    salary[1] = package2(WeeklySales);
    salary[2] = package3(WeeklySales);
    
    std::cout << "P1:" << salary[0] << std::endl;
    std::cout << "P2:" << salary[1] << std::endl;
    std::cout << "P3:" << salary[2] << std::endl;


    int best=0,op=0;
    for (int i = 0; i < salary.size(); i++) {
        std::cout << i << ",op:" << op << ",best:" << best << std::endl;
        if (salary[i]>best) {
            best = salary[i];
            op = i ;
        }
    }
    std::cout << "P" << (op+1) << " is the best package: " << best << std::endl;


    return 0;
}



