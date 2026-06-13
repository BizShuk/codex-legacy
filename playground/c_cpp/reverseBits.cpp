#include <iostream>
#include <bitset>

using namespace std;


class Solution {
    public:
        uint32_t reverseBits(uint32_t n) {
            cout << "n:" <<n <<endl;
            n = (n >> 16) | (n << 16);
            n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8);
            n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4);
            n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2);
            n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1);
            cout << "n:" <<n <<endl;
            cout << "bits:" << std::bitset<32>(n) << endl;
            return n;
        }
};


int main(int argc, char *argv[])
{
    Solution a;
    a.reverseBits(1);
    a.reverseBits(2);


    return 0;
}
