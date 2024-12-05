// Program: Implementation of Quick Sort Algorithm.

#include <bits/stdc++.h>
using namespace std;

int partition(vector<int> &arr, int low, int high) {
    
    int pivot = arr[low];
    int i = low;
    int j = high;

    while(i < j) {
        while(arr[i] <= pivot and i <= high-1) i++;
        while(arr[j] > pivot and j >= low+1) j--;
        if(i < j ) swap(arr[i], arr[j]);
    }
    swap(arr[low], arr[j]);
    return j;
    
}


int quick_sort(vector<int> &arr, int low, int high) {
    
    if(low < high) {
        int pIndex = partition(arr, low, high);
        quick_sort(arr, low, pIndex-1);
        quick_sort(arr, pIndex+1, high);
    }
    
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
