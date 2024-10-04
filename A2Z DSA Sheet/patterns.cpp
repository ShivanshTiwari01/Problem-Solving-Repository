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

int pattern3(int n) {
    for (int i=1; i<=n; i++) {
        for (int j=1; j<=i; j++){
            cout << j;
        }
        cout << endl;
    }
}

int pattern4(int n) {
    for (int i=1; i<=n; i++) {
        for (int j=1; j<=i; j++){
            cout << i;
        }
        cout << endl;
    }
}

int pattern5(int n) {
    for (int i=n; i>0; i--) {
        for (int j=0; j<i; j++){
            cout << "*";
        }
        cout << endl;
    }
}


int main() {

    int n;
    cin >> n;
    pattern5(n);
    
    return 0;
}
