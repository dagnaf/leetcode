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
    ListNode* swapPairs(ListNode* head) {
        ListNode dummy(0);
        dummy.next = head;
        ListNode *pcurr = head;
        ListNode *plast = &dummy;
        while (pcurr != nullptr && pcurr->next != nullptr) {
            ListNode *pcurr_next = pcurr->next;
            ListNode *pcurr_next_next = pcurr_next->next;
            pcurr->next = pcurr_next_next;
            pcurr_next->next = pcurr;
            plast->next = pcurr_next;
            plast = pcurr;
            pcurr = pcurr_next_next;
        }
        return dummy.next;
    }

    int main(int argc, const char *argv[]) {
        int n;
        while (cin >> n) {
            ListNode dummy(0);
            ListNode *pn = &dummy;
            for (int i = 0; i < n; ++i) {
                int a;
                cin >> a;
                pn->next = new ListNode(a);
                pn = pn->next;
            }
            ListNode *ph = swapPairs(dummy.next);
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