// Program: Move the zeroes in array to the last of the array.

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void move_zeroes(vector<int> &arr, int n) {
    int j = -1;
    for (int i = 0; i < n; i++) {
        if (arr[i] == 0) {
            j = i;
            break;
        }
    }
    if (j == -1)
        return;
    for (int i = j + 1; i < n; i++) {
        if (arr[i] != 0) {
            swap(arr[i], arr[j]);
            j++;
        }
    }
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    move_zeroes(arr, n);
    for (int i : arr) {
        cout << i << " ";
    }

    return 0;
}