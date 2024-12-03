// Program: Impelementation of Bubble Sort Algorithm

#include <bits/stdc++.h>
using namespace std;

int bubble_sort(vector<int>& arr, int n) {
    
    for(int i=n-1; i>=0; i--) {
        int did_swap = 0;
        for(int j=0; j<=i-1; j++) {
            if(arr[j] > arr[j+1]) {
                swap(arr[j], arr[j+1]);
                did_swap = 1;
            }
        }
        if(did_swap == 0) {
            break;
        }
    }
    
}

int main() {
    
    vector<int> arr = {23,11,33,1,5,3,6,88,6,55};
    int n = 10;

    bubble_sort(arr, n);
    for(int i : arr) {
        cout << i << " ";
    }

    return 0;
}