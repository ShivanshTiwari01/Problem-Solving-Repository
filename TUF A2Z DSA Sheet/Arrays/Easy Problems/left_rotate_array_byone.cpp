// Program: Left Rotate the array by one place.

#include <iostream>
#include <vector>
using namespace std;

int left_rotate_by_one(vector<int> &arr, int n) {
    int temp = arr[0];
    for (int i = 0; i < n - 1; i++) {
        arr[i] = arr[i + 1];
    }
    arr[n - 1] = temp;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    left_rotate_by_one(arr, n);
    for (int i : arr) {
        cout << i << " ";
    }

    return 0;
}
