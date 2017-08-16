/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

class Solution {
private:
    enum order { ASC, DESC };

    int get_order(vector<ListNode*>& lists) {
        int ord = ASC;
        for (ListNode *&ph : lists) {
            bool rev = false;
            ListNode *pn = ph;
            while (pn != nullptr) {
                if (pn->next && pn->val != pn->next->val) {
                    if (ord == ASC && pn->val > pn->next->val) rev = true;
                    if (ord == DESC && pn->val < pn->next->val) rev = true;
                    if (ord == -1) ord = (pn->val < pn->next->val ? ASC : DESC);
                    break;
                }
                pn = pn->next;
            }
            if (rev) reverseList(ph);
        }
        return ord;
    }

    void reverseList(ListNode *&ptail) {
        if (ptail == nullptr) return;
        ListNode *plast = nullptr;
        while (ptail->next != nullptr) {
            ListNode *pnext = ptail->next;
            ptail->next = plast;
            plast = ptail;
            ptail = pnext;
        }
        ptail->next = plast;
    }

public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        int ord = get_order(lists);
        auto cmp = [ord] (const ListNode *lhs, const ListNode *rhs) -> bool {
            if (lhs == nullptr || rhs == nullptr) return lhs == nullptr;
            return ord == ASC ? lhs->val > rhs->val : lhs->val < rhs->val;
        };
        priority_queue<ListNode *, vector<ListNode *>, decltype(cmp)> q(cmp, lists);
        ListNode h(0);
        ListNode *pn = &h;
        while (!q.empty()) {
            ListNode *pt = q.top(); q.pop();
            if (!pt) continue;
            if (pt->next) q.push(pt->next);
            pn->next = pt;
            pn = pt;
        }
        return h.next;
    }

    int main(int argc, const char *argv[]) {
        int k;
        while (cin >> k) {
            // read k lists
            vector<ListNode *> lists(k);
            for (int i = 0; i < k; ++i) {
                int n;
                cin >> n;
                ListNode h(0);
                ListNode *pn = &h;
                for (int j = 0; j < n; ++j) {
                    int a;
                    cin >> a;
                    pn->next = new ListNode(a);
                    pn = pn->next;
                }
                lists[i] = h.next;
            }
            // merge
            ListNode *pn = mergeKLists(lists);
            // print and release
            while (pn != nullptr) {
                ListNode *pd = pn;
                cout << pn->val << " ";
                pn = pn->next;
                delete pd;
            }
            cout << endl;
        }
        return 0;
    }
};