// Program: Implement the given selection sort.

#include <iostream>
#include <vector>
using namespace std;

int selection_sort(vector<int>& arr, int n) {
    
    for(int i=0; i<n-1; i++) {
        int min = i;
        for(int j=i; j<n-1; j++) {
            if(arr[j] < arr[min]) min = j;
        }
        swap(arr[min], arr[i]);
    }
    
}

int main() {
    
    int n = 10;
    vector<int> arr = {10, 3, 44, 5, 66, 4, 0, 999, 2, 1};
    selection_sort(arr, n);
    for(int i : arr) {
        cout << i << endl;
    }

    return 0;
}