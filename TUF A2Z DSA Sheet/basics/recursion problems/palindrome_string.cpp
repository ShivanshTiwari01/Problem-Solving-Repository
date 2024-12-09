// Program: Check if the given string is palindrome or not.

#include <iostream>
using namespace std;

bool palindrome(int i, string& s) {

    if(i >= s.length()/2) return true;

    if(s[i] != s[s.length()-i-1]) return false;

    return palindrome(i+1, s);

}

int main() {
    
    string s;
    cin >> s;
    cout << palindrome(0 ,s);
    cout << endl;

    return 0;
}
