#include <bits/stdc++.h>
using namespace std;

int pattern1(int n) {
    for(int i=0; i<n; i++) {
        for(int j=0; j<n; j++) {
            cout << "*";
        }
        cout << endl;
    }
}

int pattern2(int n) {
    for (int i=0; i<n; i++) {
        for (int j=0; j<=i; j++){
            cout << "*";
        }
        cout << endl;
    }
}

int main() {

    int n;
    cin >> n;
    pattern2(n);
    
    return 0;
}
