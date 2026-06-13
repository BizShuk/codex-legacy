#include <iostream>
#include <vector>
#include <string>

using namespace std;

// class way
// swap way


void combine(vector<string> o, int s,int &p_count, string v,vector<string> &a){

    if (s>=0) {
        v = v + o.at(s);
        o.erase(o.begin()+s);
        if (o.size()==0) {
            a.push_back(v);
            p_count++;
            return ;
        }
    }

    for ( int i = 0;  i < o.size(); i++) {
        combine(o,i,p_count,v,a);
    }

    return;
}

int main(int argc, char *argv[])
{
    int book_num = 6;
    std::vector<string> books;
    std::vector<string> answer;
    int permutation_count = 0 ;


    for (int i = 0; i < book_num; ++i) {
        books.push_back(to_string(i+1));
    }

    vector<string>::iterator i = books.begin();
    while ( i != books.end()) {
        std::cout << "book: " << *i << std::endl;
        i++;
    }


    combine(books,-1,permutation_count,"",answer);


    for (int i = 0;i < answer.size();  i++) {
        std::cout << "Permutation : " << answer[i] << std::endl;
    }

    std::cout << "total permutation:" << permutation_count << std::endl;
    

    return 0;
}


