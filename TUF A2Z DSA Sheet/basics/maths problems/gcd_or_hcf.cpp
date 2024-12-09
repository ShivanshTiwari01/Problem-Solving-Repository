// Program to find HCF of given numbers.

#include <iostream>
using namespace std;

int findGCD(int a, int b) {
    if(a == 0) return b;
    if(b == 0) return a;
    if(a == b) return a;
    if(a > b) return findGCD(a - b, b);
    return findGCD(a, b -a);
}

int main() {
    int a, b;
    cin >> a >> b;
    cout << "GCD of a and b is: " << findGCD(a, b);
    return 0;
}
