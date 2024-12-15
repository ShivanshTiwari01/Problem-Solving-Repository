#include <iostream>
#include <string>
using namespace std;

string removeOuterParentheses(string s) {
    string result;
    int balance = 0;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] == '(') {
            if (balance > 0) {
                result += s[i];
            }
            balance++;
        } else {
            balance--;
            if (balance > 0) {
                result += s[i];
            }
        }
    }
    return result;
}

int main() {

    string str;
    getline(cin, str);
    string new_str = removeOuterParentheses(str);
    cout << new_str;

    return 0;
}