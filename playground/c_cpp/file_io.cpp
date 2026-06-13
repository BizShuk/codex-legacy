#include <iostream>
#include <fstream>
using namespace std;


int main(int argc, char *argv[])
{
    
    char first_name[30], last_name[30]; int age;
    char file_name[20] = "file_io.output";

    std::cout << "Enter First Name:"; cin >> first_name;
    std::cout << "Enter Last Name:"; cin >> last_name;
    std::cout << "Enter Age:"; cin >> age;
    std::cout << std::endl;


    // write file
    ofstream People;
    People.open(file_name, ios_base::app); // std::ios_base::app for append | std::ios_base::out

    People << first_name << endl << last_name << endl << age <<endl; 
    People.close();



    // read file
    ifstream People_in(file_name);

    while (1) {
        People_in >> first_name >> last_name >> age;    // >> get line
        if (People_in.eof()) {
            break;
        }
        cout << endl << "First Name: " << first_name;
        cout << endl << "Last Name:  " << last_name;
        cout << endl << "Enter Age:  " << age;
        cout << endl;
    }

    People_in.close();

    return 0;
}
