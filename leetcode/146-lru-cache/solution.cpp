template<class K, class V, V val_404 = V()>
class LRUCache_t {
private:
    list<pair<K, V>> l;
    unordered_map<K, typename decltype(l)::iterator> h;
    size_t capacity;
public:
    LRUCache_t(size_t capacity) : capacity(capacity) {}

    V get(K key) {
        auto hit = h.find(key);
        if (hit != h.end()) {
            auto lit = hit->second;
            l.splice(l.begin(), l, lit);
            return l.begin()->second;
        }
        return val_404;
    }

    void put(K key, V val) {
        auto hit = h.find(key);
        if (hit == h.end()) {
            l.push_front(make_pair(key, val));
            h.insert(make_pair(key, l.begin()));
            if (l.size() > capacity) {
                h.erase(l.back().first);
                l.pop_back();
            }
        } else {
            hit->second->second = val;
            l.splice(l.begin(), l, hit->second);
        }
    }
};

class LRUCache : public LRUCache_t<int, int, -1> {
    using LRUCache_t<int, int, -1>::LRUCache_t;
};

class Solution {
public:
    int main(int argc, char const *argv[]) {
        int cap; cin >> cap;
        LRUCache cache(cap);
        string cmd;
        int key;
        int val;
        while (cin >> cmd) {
            if (cmd == "put") {
                cin >> key >> val;
                cache.put(key, val);
            } else {
                cin >> key;
                cout << cache.get(key) << endl;
            }
        }
        return 0;
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */