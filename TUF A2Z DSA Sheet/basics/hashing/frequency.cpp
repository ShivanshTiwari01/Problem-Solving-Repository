// Program: Find frequency of each element in a given array

#include <bits/stdc++.h>
using namespace std;

int frequency(int arr[], int n) {
    
    unordered_map<int, int> map;
    for(int i=0; i<n; i++) {
        map[arr[i]]++;
    }

    for(auto x : map) {
        cout << x.first << " " << x.second << endl;
    }
    
}

int main() {
    
    int n=10;
    int arr[] = {2, 2, 1, 44, 44, 5, 33, 3, 5, 2};
    frequency(arr, n);

    return 0;
}