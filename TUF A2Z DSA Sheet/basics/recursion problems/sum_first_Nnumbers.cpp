// Program: Find sum of first n numbers.

#include <bits/stdc++.h>
using namespace std;

int sumN(int sum, int n) {
    
    if(n < 1) return sum;

    sumN(sum+n, n-1);
}

int main() {
    
    int n;
    cin >> n;
    cout << "Sum of first n numbers: " << sumN(0, n);

    return 0;
}
