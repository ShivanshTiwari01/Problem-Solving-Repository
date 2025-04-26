// Program to traverse a linked list

#include <iostream>
using namespace std;

// Node structure
struct Node {
	int data;
	Node* next;
};

// Function to create a new node
Node* createNode(int data) {
	Node* newNode = new Node();
	newNode->data = data;
	newNode->next = nullptr;
	return newNode;
}

// Function to traverse the linked list
void traverseLinkedList(Node* head) {
	Node* current = head;
	while (current != nullptr) {
		cout << current->data << " ";
		current = current->next;
	}
	cout << endl;
}

// Main function
int main() {
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> 5
	Node* head = createNode(1);
	head->next = createNode(2);
	head->next->next = createNode(3);
	head->next->next->next = createNode(4);
	head->next->next->next->next = createNode(5);
	// Traversing the linked list
	cout << "Linked List: ";
	traverseLinkedList(head);
	// Freeing memory (not shown for simplicity)
	return 0;
}

