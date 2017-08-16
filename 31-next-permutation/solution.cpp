class Solution {
public:
    void nextPermutation(vector<int>& nums) {
        const auto ib = nums.begin();
        const auto ie = nums.end();
        if (ib == ie || ib + 1 == ie) return;
        auto i = ie - 1;
        for (; i != ib; --i) {
            if (*(i - 1) < *i) {
                auto j = ie;
                while (! (*(i - 1) < *--j)); // must add paren !()
                // cout << "iter_swap " << *(i - 1) << " " << *j  << endl;
                iter_swap(i - 1, j);
                break;
            }
        }
        reverse(i, ie);
    }

    int main(int argc, const char *argv[]) {
        int n;
        while (cin >> n) {
            vector<int> a(n);
            for (int i = 0; i < n; ++i) cin >> a[i];
            vector<int> b = a;
            nextPermutation(a);
            copy(a.begin(), a.end(), ostream_iterator<int>(cout, " "));
            cout << endl;
            next_permutation(b.begin(), b.end());
            copy(b.begin(), b.end(), ostream_iterator<int>(cout, " "));
            cout << endl;
        }
        return 0;
    }
};