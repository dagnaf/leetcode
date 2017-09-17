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
        if (!root) return;
        while (root) {
            TreeLinkNode *p1 = root;
            TreeLinkNode *p2 = nullptr;
            TreeLinkNode *root2 = nullptr;
            // cout << "begin level " << root->val << endl;
            while (p1) {
                // cout << " at level head " << root->val << " traverse " << p1->val << endl;
                if (p1->left) p1->left->next = p1->right;
                if (p2) {
                    if (p1->left) p2->next = p1->left;
                    else p2->next = p1->right;
                } else {
                    if (p1->left) p2 = p1->left;
                    else p2 = p1->right;
                    // cout << " set p2 " << (p2 == nullptr ? -1 : p2->val) << endl;
                    root2 = p2;
                }
                if (p2 && p2->next) {
                    // cout << "  move forward " << p2->val << " to " << p2->next->val << endl;
                    if (p2->next->next) p2 = p2->next->next;
                    else p2 = p2->next;
                }
                p1 = p1->next;
            }
            root = root2;
        }
    }

    void print_by_level(TreeLinkNode *root) {
        queue<TreeLinkNode *> q;
        q.push(root);
        while (!q.empty()) {
            TreeLinkNode *p = q.front(); q.pop();
            if (p->left) q.push(p->left);
            if (p->right) q.push(p->right);
            cout << p->val << " ";
        }
        cout << endl;
    }

    int main(int argc, const char *argv[]) {
        string line;
        while (getline(cin, line)) {
            stringstream ss(line);
            string s;
            int cnt = 0;
            while (ss >> s) {
                if (s != "#") ++cnt;
            }
            vector<TreeLinkNode> nodes(cnt, TreeLinkNode(0));
            stringstream sss(line);
            cnt = 0;
            int pa = 0;
            while (sss >> s) {
                if (s != "#") {
                    nodes[cnt].val = atoi(s.c_str());
                    if (cnt > 0) {
                        if (pa & 1) nodes[pa/2].right = &nodes[cnt];
                        else nodes[pa/2].left = &nodes[cnt];
                    }
                    ++cnt;
                }
                if (cnt > 1) ++pa;
            }
            print_by_level(&nodes[0]);
            connect(&nodes[0]);
            TreeLinkNode *p = &nodes[0];
            while (p) {
                TreeLinkNode *q = p;
                TreeLinkNode *r = nullptr;
                while (q) {
                    cout << q->val << " ";
                    if (!r) {
                        if (q->left) r = q->left;
                        else r = q->right;
                    }
                    q = q->next;
                }
                cout << endl;
                p = r;
            }
        }
        return 0;
    }
};