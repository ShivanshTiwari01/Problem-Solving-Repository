// Program: Reverse the given array.

#include <bits/stdc++.h>
using namespace std;

int reverse_array(int arr[], int start, int end) {
    if(start < end) {
        swap(arr[start], arr[end]);
        reverse_array(arr, start+1, end-1);
    }
}

int main() {
    
    int n = 6;
    int arr[] = {33, 44, 21, 54, 11, 43};
    reverse_array(arr, 0, n-1);
    for(int i=0; i<n; i++){
        cout << arr[i] << " ";
    }
    return 0;
}
