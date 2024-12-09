// Program: Print all divisors of a given number.

#include <iostream>
#include <vector>
using namespace std;

vector<int> find_divisors(int n) {
    vector<int> divisors;

    int sqrt_n = sqrt(n);
    for(int i=1; i<=sqrt_n; i++) {
        
        if(n%i == 0) divisors.push_back(i);

        if(i != n/i) divisors.push_back(n/i);
    } 
    return divisors;
}

int main() {

    int n;
    cin >> n;

    vector<int> divisors = find_divisors(n);
    cout << "divisors of n are: ";
    
    for(int i : divisors) {
        cout << i << " ";
    }
    cout << endl;

    return 0;
}
