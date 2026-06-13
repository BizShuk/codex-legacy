/* book order prediction for next year
 *
 *
 */
#include <iostream>
#include <math.h>
using namespace std;

float show_prediction(int bcode, int enrollment, int inventory, int required, int newuse) {
    float newp = 0.9 , 
          usedp = 0.65;
    float total = 0;

    if (required==0) {
        newp = 0.4;
        usedp = 0.2;
    }

    if (newuse == 1) {
       total = floor( enrollment * newp ) - inventory;
    } else {
       total = floor( enrollment * usedp ) - inventory;
    }

    return total;
}

void order_book() {
    float bprice=0,
          btotal=0,
        all_order=0;
    int bcode,
        inventory=0,
        enrollment=0,
        required=0,
        newuse=0,
        done=1;


    do {
        cout << "Please enter the book code:"; cin >> bcode;
        cout << "Book price: "; cin >> bprice;
        cout << "Inventory: "; cin >> inventory;
        cout << "Enrollment: "; cin >> enrollment;
        cout << "1 for required , 0 for optional: "; cin >> required;
        required = required == 1 ? 1 : 0;
        cout << "1 for new , 0 for used: "; cin >> newuse;
        newuse = newuse == 1 ? 1 : 0;

        btotal = show_prediction(bcode, enrollment, inventory, required, newuse);

        cout << endl;
        cout << "*****************************" << endl;
        cout << "Book: " << bcode << endl;
        cout << "Price: " << bprice << endl;
        cout << "Inventory: " << inventory << endl;
        cout << "Enrollment: " << enrollment << endl << endl;
        cout << "*****************************" << endl;
        cout << "Need to order: " << btotal << endl;
        cout << "Total Cost: $" << btotal * bprice << endl;
        cout << "*****************************" << endl << endl;

        all_order += btotal * bprice;
        cout << "Enter 1 for predict another book, 0 for stop system: "; cin >> done;
    } while (done==1);

    
    cout << "*****************************" << endl;
    cout << "Total for all orders: $" << all_order << endl;
    cout << "*****************************" << endl;

}


int main(int argc, char *argv[]) {
    




    order_book();
    


    return 0;
}
