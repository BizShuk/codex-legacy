#include <iostream>
#include <vector>
#include <string>
using namespace std;
using std::vector;
using std::string;



vector<vector <int> > subsetsWithDup(vector<int> &S) {
    sort(S.begin(), S.end());
    vector<vector<int> > ret = {{}};
    int size = 0, startIndex = 0;
    for (int i = 0; i < S.size(); i++) {
        cout << "size level:" << i << "," << size << "," << startIndex <<endl;
        startIndex = i >= 1 && S[i] == S[i - 1] ? size : 0;
        size = ret.size();
        for (int j = startIndex; j < size; j++) {
            cout << "j:" << j << ",size:"<< size<<endl;
            vector<int> temp = ret[j];
            temp.push_back(S[i]);
            ret.push_back(temp);
        }
    }
    return ret;
}
int main(){
        vector<int> a;
        a.push_back(4);
        a.push_back(1);
        a.push_back(0);
        vector<vector <int> > b;
        b = subsetsWithDup(a);
        
        
        vector<int> example {1,2,3,4,5,6, 7};

        for (int x = 0; x != b.size(); ++x){
            cout << "test" << x <<endl;
            for (int y = 0; y != b[x].size(); ++y){
                cout << b[x][y]  << endl;
            }
        }

        return 0;
}
