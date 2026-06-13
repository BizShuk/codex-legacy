#include <iostream>
#include <string>
#include <map>

using namespace std;

bool IsAnagram(string s, string t)
{
    if (t.length() != s.length()) {
        return false;
    }
    std::map<char, int> ori;
    for ( int i = 0; i < s.length(); i++) {
        ori[s[i]]++;
    }

    for (auto i = t.begin(); i != t.end(); i++) {
        if (ori.find(*i) != ori.end() && ori[*i] > 0 ) {
            ori[*i]--;
        } else {
            return false;
        }
    }

    return true;
}
int main(int argc, char *argv[])
{
    auto tmp = IsAnagram("test1234","4321test"); cout << tmp << endl;
    tmp = IsAnagram("acca","accc"); cout << tmp << endl;
    tmp = IsAnagram("ab","a"); cout << tmp << endl;
    tmp = IsAnagram("a","ab"); cout << tmp << endl;
    tmp = IsAnagram("wwww123","www123"); cout << tmp << endl;
    tmp = IsAnagram("x123","321x"); cout << tmp << endl;
    return 0;
}



