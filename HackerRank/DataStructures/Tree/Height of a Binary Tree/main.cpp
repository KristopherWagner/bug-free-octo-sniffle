// https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree
#include <bits/stdc++.h>

using namespace std;

class Node {
    public:
        int data;
        Node *left;
        Node *right;
        Node(int d) {
            data = d;
            left = NULL;
            right = NULL;
        }
};

class Solution {
    public:
  		Node* insert(Node* root, int data) {
            if(root == NULL) {
                return new Node(data);
            } else {
                Node* cur;
                if(data <= root->data) {
                    cur = insert(root->left, data);
                    root->left = cur;
                } else {
                    cur = insert(root->right, data);
                    root->right = cur;
               }

               return root;
           }
        }
/*The tree node has data, left child and right child 
class Node {
    int data;
    Node* left;
    Node* right;
};

*/
    int height(Node* root) {
        int leftHeight = 0;
        int rightHeight = 0;
        if (root->left != NULL) {
            leftHeight = 1 + height(root->left);
        }
        if (root->right != NULL) {
            rightHeight = 1 + height(root->right);
        }
        if (leftHeight > rightHeight) {
            return leftHeight;
        }
        return rightHeight;
    }

};
