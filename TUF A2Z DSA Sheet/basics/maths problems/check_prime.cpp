// Program: Check if the given no is prime or not.

#include <bits/stdc++.h>
using namespace std;

int check_prime(int n) {
    int flag = false;

    for(int i=2; i<n; i++) {
        if(n%i == 0) {
            flag = true;
            break;
        }
    }

    return flag;
}

int main() {
    int n;
    cin >> n;
    int result = check_prime(n);
    if(result == false) {
        cout << "no is prime" << endl;
    }
    else {
        cout << "no is not prime" << endl;
    }

    return 0;
}
