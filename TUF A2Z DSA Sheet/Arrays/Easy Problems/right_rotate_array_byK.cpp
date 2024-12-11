// Program: Right Rotate the array by K elements.

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int right_rotate_array_byK(vector<int> &arr, int n, int k) {
    k = k % n;
    reverse(arr.begin(), arr.end() - k);
    reverse(arr.end() - k, arr.end());
    reverse(arr.begin(), arr.end());
}

int main() {

    int n, k;
    cin >> n >> k;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    right_rotate_array_byK(arr, n, k);
    for (int i : arr) {
        cout << i << " ";
    }

    return 0;
}
