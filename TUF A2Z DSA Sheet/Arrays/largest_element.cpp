#include <iostream>
#include <vector>
using namespace std;

int largest_element(vector<int> &arr, int n) {
    int largest = -9999;
    for (int i : arr) {
        if (i > largest) {
            largest = i;
        }
    }
    return largest;
}

int main() {

    int n;
    cin >> n;
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }
    cout << largest_element(arr, n);

    return 0;
}
