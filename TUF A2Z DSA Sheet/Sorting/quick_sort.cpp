// Program: Implementation of Quick Sort Algorithm.

#include <bits/stdc++.h>
using namespace std;

int quick_sort(vector<int> arr, int start, int end) {
    
    
    
}

int main() {
    
    int n;
    cin >> n;
    vector<int> arr(n); 
    for(int i=0; i<n; i++) {
        cin >> arr[i];
    }

    quick_sort(arr, 0, n-1);

    for(int i : arr) {
        cout << i << " ";
    }

    return 0;
}