// Reverse the given no and ignore the trailin zeros.
#include <bits/stdc++.h>
using namespace std;

int reverse_number(int no) {
    int rev = 0;
    while(no > 0) {

        if(no%10 != 0) rev = rev*10 + no%10;
        no = no/10;
        
    }
    return rev;
}

int main() {
    int n;
    cin >> n;
    cout << reverse_number(n);
}
