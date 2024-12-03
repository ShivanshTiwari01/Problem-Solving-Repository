// Program: Implementation of Insertion Sort Algorithm.

#include <bits/stdc++.h>
using namespace std;

int insertion_sort(vector<int>& arr, int n) {
    
    for(int i=0; i<n; i++) {
        int j=i;
        while(j>0 and arr[j-1]>arr[j]) {
            swap(arr[j-1], arr[j]);
            j--;
        }
    }
    
}

int main() {
    
    int n;
    cin >> n;
    vector<int> arr(n); 
    for(int i=0; i<n; i++) {
        cin >> arr[i];
    }

    insertion_sort(arr, n);

    for(int i : arr) {
        cout << i << " ";
    }

    return 0;
}