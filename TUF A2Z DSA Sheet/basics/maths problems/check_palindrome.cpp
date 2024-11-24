// Program to check if a given number is palindrome or not.

#include <bits/stdc++.h>
using namespace std;

string check_palindrome(int n) {
    int temp = n;
    int rev = 0;
    while(temp) {
        rev = rev*10 + temp%10;
        temp /= 10;
    }
    if(rev == n) return "true";
    return "false";
}

int main() {
    int n;
    cin >> n;
    cout << check_palindrome(n);
    return 0;
}
