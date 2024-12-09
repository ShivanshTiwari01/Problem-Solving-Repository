// Program: Implementation of Insertion Sort Algorithm.

#include <iostream>
#include <vector>
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

void insertion_sort_recursive(int arr[], int i, int n) {

    // Base Case: i == n.
    if (i == n) return;

    int j = i;
    while (j > 0 && arr[j - 1] > arr[j]) {
        swap(arr[j-1], arr[j]);
        j--;
    }

    insertion_sort_recursive(arr, i + 1, n);
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
