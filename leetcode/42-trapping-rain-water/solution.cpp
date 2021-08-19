ostream &operator <<(ostream &os, const stack<int> &s) {
    auto t = s;
    vector<int> v;
    while (!t.empty()) { v.push_back(t.top()); t.pop(); }
    os << "[ ";
    for (int i : v) {
        os << i << " ";
    }
    os << "]";
    return os;
}
class Solution {
public:
    int trap(vector<int>& height) {
        int n = height.size();
        stack<int> shi;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            vector<int> tmpi;
            while (!shi.empty() && height[shi.top()] <= height[i]) {
// cout << "shi.top " << shi.top() << " out" << endl;
                tmpi.push_back(shi.top());
                shi.pop();
            }
            if (!shi.empty()) {
                tmpi.push_back(shi.top());
            }
            if (tmpi.size()) {
                int m = tmpi.size();
                int tmph = min(height[i], height[tmpi[m - 1]]);
                for (int j = 0; j < m - 1; ++j) {
// cout << "when tmph " << tmph << " inc " << ((tmpi[j] - tmpi[j + 1]) * (tmph - height[tmpi[j]])) << endl;
                    ans += (tmpi[j] - tmpi[j + 1]) * (tmph - height[tmpi[j]]);
                }
            }
            shi.push(i);
// cout << "stack after " << i << ": " << shi << endl;
        }
        return ans;
    }

    int main(int argc, const char *argv[]) {
        int n;
        while (cin >> n) {
            vector<int> height(n);
            for (int i = 0; i < n; ++i) {
                cin >> height[i];
            }
            cout << trap(height) << endl;
        }
        return 0;
    }
};