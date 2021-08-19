/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if (k <= 1) return head;
        ListNode dummy(0);
        dummy.next = head;
        ListNode* p0 = &dummy;
        while (p0->next) {
            ListNode *pcurr = p0->next;
            int cnt = 0;
            while (cnt < k && pcurr) {
                ++cnt;
                pcurr = pcurr->next;
            }
            if (cnt < k) break;
            ListNode *plast = p0->next;
            pcurr = plast->next;
            for (int i = 0; i < k - 1; ++i) {
                ListNode *pnext = pcurr->next;
                pcurr->next= plast;
                plast = pcurr;
                pcurr = pnext;
            }
            p0->next->next = pcurr;
            pcurr = p0->next;
            p0->next = plast;
            p0 = pcurr;
        }
        return dummy.next;
    }

    int main(int argc, const char *argv[]) {
        int n, k;
        while (cin >> k >> n) {
            ListNode dummy(0);
            ListNode *pn = &dummy;
            for (int i = 0; i < n; ++i) {
                int a;
                cin >> a;
                pn->next = new ListNode(a);
                pn = pn->next;
            }
            ListNode *ph = reverseKGroup(dummy.next, k);
            while (ph) {
                ListNode *pn = ph;
                cout << pn->val << " ";
                ph = pn->next;
                delete pn;
            }
            cout << endl;
        }
        return 0;
    }
};