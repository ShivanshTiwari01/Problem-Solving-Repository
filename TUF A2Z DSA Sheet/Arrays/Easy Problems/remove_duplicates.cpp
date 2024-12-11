// Program: Remove duplicates from a sorted array.

#include <iostream>
#include <vector>
using namespace std;

int remove_duplicates(vector<int> &arr, int n) {
    int i = 0;
    for (int j = 1; j < n; j++) {
        if (arr[i] != arr[j]) {
            i++;
            arr[i] = arr[j];
        }
    }
    return i + 1;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    int k = remove_duplicates(arr, n);
    for (int i = 0; i < k; i++) {
        cout << arr[i] << " ";
    }

    return 0;
}
