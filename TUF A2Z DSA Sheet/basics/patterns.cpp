// Striver takeuforward patterns program in same order as given on website - https://takeuforward.org/strivers-a2z-dsa-course/must-do-pattern-problems-before-starting-dsa/

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
    for(int i=0; i<n; i++) {
        for(int j=0; j<=i; j++) {
            cout << "*";
        }
        cout << endl;
    }
}

int pattern3(int n) {
    for(int i=1; i<=n; i++) {
        for(int j=1; j<=i; j++) {
            cout << j;
        }
        cout << endl;
    }
}

int pattern4(int n) {
    for(int i=1; i<=n; i++) {
        for(int j=1; j<=i; j++) {
            cout << i;
        }
        cout << endl;
    }
}

int pattern5(int n) {
    for(int i=n; i>0; i--) {
        for (int j=0; j<i; j++) {
            cout << "*";
        }
        cout << endl;
    }
}

int pattern6(int n) {
    for(int i=n; i>0; i--) {
        for(int j=1; j<=i; j++){
            cout << j;
        }
        cout << endl;
    }
}

int pattern7(int n) {
    for(int i=0; i<n; i++) {
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
    for(int i=0; i<n; i++) {
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
    for(int i=0; i<= 2*n-1; i++) {
        int stars = i;
        if(i>n) stars = 2*n-i;
        for(int j=1; j<=stars; j++) {
            cout << "*";
        }
        cout << endl;
    }
}

int pattern11(int n) {
    int start = 1;
    for(int i=0; i<n; i++) {
        if(i%2 == 0) start = 1;
        else start = 0;
        for(int j=0; j<=i; j++) {
            cout << start;
            start = 1-start;
        }
        cout << endl;
    }
}

int pattern12(int n) {
    int spaces = 2*(n-1);
    for(int i=0; i<n; i++) {
        for(int j=1; j<=i; j++) {
            cout << j;
        }
        for(int j=1; j<=spaces; j++) {
            cout << " ";
        }
        for(int j=i; j>=1; j--) {
            cout << j;
        }
        cout << endl;
        spaces -= 2;
    }
}

int pattern13(int n) {
    int t = 1;
    for(int i=0; i<n; i++) {
        for(int j=0; j<i; j++) {
            cout << t << " ";
            t += 1;
        }
        cout << endl;
    }
}

int pattern14(int n) {
    for(int i=0; i<n; i++) {
        for(char ch='A'; ch<='A'+i; ch++) {
            cout << ch << " ";
        }
        cout << endl;
    }
}

int pattern15(int n) {
    for(int i=0; i<n; i++) {
        for(char ch='A'; ch<='A'+(n-i-1); ch++) {
            cout << ch << " ";
        }
        cout << endl;
    }
}

int pattern16(int n) {
    for(int i=0; i<n; i++){
        char ch = 'A'+i;
        for(int j=0; j<=i; j++) {
            cout << ch << " ";
        }
        cout << endl;
    }
}

int pattern17(int n) {
    for(int i=0; i<n; i++){
        for(int j=0; j<n-i-1; j++) {
            cout << " ";
        }
        char ch = 'A';
        int breakpoint = (2*i+1)/2;
        for(int j=1; j<=2*i+1; j++) {
            cout << ch;
            if(j <= breakpoint) ch++;
            else ch--;
        }
        for(int j=0; j<n-i-1; j++) {
            cout << " ";
        }
        cout << endl;
    }
}

int pattern18(int n) {
    for(int i=0; i<n; i++) {
        for(char ch='A'+n-1-i; ch<='A'+n-1; ch++) {
            cout << ch << " ";
        }
        cout << endl;
    }
}

int pattern19(int n) {
    int spaces = 0;
    for(int i=0; i<n; i++) {
        for(int j=1; j<=n-i; j++) {
            cout << "*";
        }
        for(int j=0; j<spaces; j++) {
            cout << " ";
        }
        for(int j=1; j<=n-i; j++) {
            cout << "*";
        }
        spaces += 2;
        cout << endl;
    }
    spaces = 2*n-2;
    for(int i=0; i<n; i++) {
        for(int j=0; j<=i; j++) {
            cout << "*";
        }
        for(int j=0; j<spaces; j++) {
            cout << " ";
        }
        for(int j=0; j<=i; j++) {
            cout << "*";
        }
        spaces -= 2;
        cout << endl;
    }
}

int pattern20(int n) {
    int spaces = 2*n-2;
    for(int i=1; i<=2*n-1; i++) {
        int stars = i;
        if(i>n) stars = 2*n-i;
        for(int j=1; j<=stars; j++) {
            cout << "*";
        }
        for(int j=1; j<=spaces; j++) {
            cout << " ";
        }
        for(int j=1; j<=stars; j++) {
            cout << "*";
        }
        cout << endl;
        if(i<n) spaces -= 2;
        else spaces +=2;
    }
}

int pattern21(int n) {
    for(int i=0; i<n; i++) {
        for(int j=0; j<n; j++) {
            if(i==0 or j==0 or i==n-1 or j==n-1) {
                cout << "*";
            }
            else{
                cout << " ";
            }
        }
        cout << endl;
    }
}

int pattern22(int n) {
    for(int i=0; i<2*n-1; i++) {
        for(int j=0; j<2*n-1; j++) {
            int top = i;
            int bottom = j;
            int right = (2*n-2) - j;
            int left = (2*n-2) - i;

            cout << (n-min(min(top, bottom), min(left, right))) << " ";
        }
        cout << endl;
    }
}


int main() {

    int n;
    cin >> n;
    pattern22(n);
    
    return 0;
}

