#include <bits/stdc++.h>
using namespace std;

int count_digits(int no) {
    int count=0;
    while(no%10)  {
        count++;
        no /= 10;
    }
    return count;
}

int main() {
    int n;
    cin >> n;
    cout << count_digits(n);
}
