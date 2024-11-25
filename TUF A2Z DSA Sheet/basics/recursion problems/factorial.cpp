// Program: Find factorial of the given number.

#include <bits/stdc++.h>
using namespace std;

int fact(int n) {
    if(n == 0) return 1;

    return n*fact(n-1);
}

int main() {
    
    int n;
    cin >> n;
    cout << "Factorial of given no is: " << fact(n);

    return 0;
}
