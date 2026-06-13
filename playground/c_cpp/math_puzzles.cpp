#include <iostream>

int main(int argc, char *argv[])
{
    int count[3] = {0};
    float hourse_cost=10,pig_cost=3,rabbit_cost=0.5;
    int money=100,animals_num=100;
/*
    std::cout << "Enter total money:" << std::endl;
    std::cin >> money;
    std::cout << "Enter total animals:" << std::endl;
    std::cin >> animals_num;
*/

    int count_op =0 ;


    int remain_money = money;
    for ( int i = 0; i * hourse_cost <= remain_money ; i++) {

        remain_money -= i*hourse_cost;

        for ( int j = 0; j * pig_cost <= remain_money; j++) {

            remain_money -= j * pig_cost;

            for ( int k = 0; k * rabbit_cost <= remain_money; k++) {
                count_op++;
                if (remain_money-(k*rabbit_cost)==0 && (i+j+k)==100) {
                    count[0]=i;
                    count[1]=j;
                    count[2]=k;
                    break;
                }

            }
        }

        remain_money = money;
    }

    if (count[0]!=0 || count[1]!=0 || count[2]!=0) {
        std::cout << "You have a match: buy " << count[0] << " houses, " << count[1] << " pigs, " << count[2] << " rabbits." << std::endl;
    }else{
        std::cout << "Not find a matche." << std::endl;
    }
    std::cout << count_op <<"loop" << std::endl;
    return 0;
}
