/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        if (root == nullptr) return;
        while (root->left != nullptr) {
            TreeLinkNode *p1 = root;
            TreeLinkNode *p2 = nullptr;
            while (p1 != nullptr) {
                if (p2 != nullptr) {
                    p2->next = p1->left;
                }
                if (p1->left != nullptr) {
                    p1->left->next = p1->right;
                    p2 = p1->right;
                    p1 = p1->next;
                } else {
                    break;
                }
            }
            root = root->left;
        }
    }

    int main(int argc, const char *argv[]) {
        int n = 10;
        vector<TreeLinkNode> nodes(n+1, TreeLinkNode(0));
        for (int i = 1; i <= n; ++i) {
            nodes[i].val = i;
            if ((i << 1) <= n) nodes[i].left = &nodes[i << 1];
            if ((i << 1 | 1) <= n) nodes[i].right = &nodes[i << 1 | 1];
        }
        connect(&nodes[1]);
        TreeLinkNode *p = &nodes[1];

        while (p != nullptr) {
            TreeLinkNode *q = p;
            while (q != nullptr) {
                cout << (q->val) << endl;
                q = q->next;
            }
            p = p->left;
        }
        return 0;
    }
};