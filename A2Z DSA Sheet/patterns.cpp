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

int pattern6(int n) {
    for (int i=n; i>0; i--) {
        for (int j=1; j<=i; j++){
            cout << j;
        }
        cout << endl;
    }
}

int pattern7(int n) {
    for (int i=0; i<n; i++) {
        for(int j=0; j < n-i-1; j++) {
            cout << " ";
        }
        for(int k=0; k < 2*i+1; k++) {
            cout << "*";
        }
        for(int l=0; l < n-i-1; l++) {
            cout << " ";
        }
        cout << endl;
    }    
}

int pattern8(int n) {
    for(int i=0; i<n; i++) {
        for(int j=0; j<i; j++) {
            cout << " ";
        }
        for(int k=0; k < 2*n-(2*i+1); k++) {
            cout << "*";
        }
        for(int l=0; l<i; l++) {
            cout << " ";
        }
        cout << endl;
    }
}

int pattern9(int n) {
    for (int i=0; i<n; i++) {
        for(int j=0; j < n-i-1; j++) {
            cout << " ";
        }
        for(int k=0; k < 2*i+1; k++) {
            cout << "*";
        }
        for(int l=0; l < n-i-1; l++) {
            cout << " ";
        }
        cout << endl;
    }
    for(int i=0; i<n; i++) {
        for(int j=0; j<i; j++) {
            cout << " ";
        }
        for(int k=0; k < 2*n-(2*i+1); k++) {
            cout << "*";
        }
        for(int l=0; l<i; l++) {
            cout << " ";
        }
        cout << endl;
    }
}

int pattern10(int n) {
    for(int i=0; i<= 2*n-1; i++){
        int stars = i;
        if(i>n){
            stars = 2*n-i;
        }
        for(int j=1; j<=stars; j++){
            cout << "*";
        }
        cout << endl;
    }
}

int pattern11(int n){
    int start = 1;
    for(int i=0; i<n; i++){
        if(i%2 == 0) start = 1;
        else start = 0;
        for(int j=0; j<=i; j++){
            cout << start;
            start = 1-start;
        }
        cout << endl;
    }
}

int pattern12(int n){
    int spaces = 2*(n-1);
    for(int i=0; i<n; i++){
        for(int j=1; j<=i; j++){
            cout << j;
        }
        for(int j=1; j<=spaces; j++){
            cout << " ";
        }
        for(int j=i; j>=1; j--){
            cout << j;
        }
        cout << endl;
        spaces -= 2;
    }
}

int pattern13(int n){
    int t=1;
    for(int i=0; i<n; i++){
        for(int j=0; j<i; j++){
            cout << t << " ";
            t += 1;
        }
        cout << endl;
    }
}

int pattern14(int n){

}

int main() {

    int n;
    cin >> n;
    pattern14(n);
    
    return 0;
}
