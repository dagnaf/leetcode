class Solution {
public:
    string multiply(string num1, string num2) {
        int n1 = num1.size();
        int n2 = num2.size();
        vector<int> ret(n1 + n2 + 2, 0);
        for (int i = n1 - 1, k1 = 0; i >= 0; --i, ++k1) {
            for (int j = n2 - 1, k2 = 0; j >= 0; --j, ++k2) {
                ret[k1 + k2] += (num1[i] - '0') * (num2[j] - '0');
            }
        }
        for (int i = 0; i < n1 + n2 + 1; ++i) {
            ret[i + 1] += ret[i] / 10;
            ret[i] %= 10;
        }
        for_each(ret.begin(), ret.end(), [] (int &i) { i += '0'; });
        // cout << string(ret.rbegin(), ret.rend()) << endl;
        auto beg = ret.rbegin();
        while (beg != ret.rend() && *beg == '0') ++beg;
        if (beg == ret.rend()) return "0";
        return string(beg, ret.rend());
    }

    int main(int argc, const char *argv[]) {
        string a, b;
        while (cin >> a >> b) {
            cout << multiply(a, b) << endl;
        }
        return 0;
    }
};